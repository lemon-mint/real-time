package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"time"
	_ "time/tzdata"

	jsoniter "github.com/json-iterator/go"

	"github.com/beevik/ntp"
)

var json = jsoniter.ConfigFastest

var Offset = time.Duration(0)

type Time struct {
	T1 float64 `json:"t1"`
	T2 float64 `json:"t2"`
}

func timeHandle(w http.ResponseWriter, r *http.Request) {
	t := Time{
		T1: (float64(time.Now().Add(Offset).UnixNano()) / float64(1000)) / float64(1000),
	}
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	t.T2 = (float64(time.Now().Add(Offset).UnixNano()) / float64(1000)) / float64(1000)
	err := e.Encode(t)
	if err != nil {
		log.Println(err)
	}
}

func getBadgeURL(label, message, color, style string) string {
	return fmt.Sprintf("https://img.shields.io/static/v1?label=%s&message=%s&color=%s&style=%s", url.QueryEscape(label), url.QueryEscape(message), url.QueryEscape(color), url.QueryEscape(style))
}

func badgeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	// No Cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	version := r.URL.Query().Get("version")
	if version == "" {
		version = "latest"
	}
	w.Header().Set("X-Version", version)

	tz := r.URL.Query().Get("tz")
	if tz == "Local" || tz == "" {
		// Default to UTC
		tz = "UTC"
	}

	label := r.URL.Query().Get("label")
	if label == "" {
		label = tz
	}

	color := r.URL.Query().Get("color")
	if color == "" {
		color = "4bc51d"
	} else {
		color = strings.ToLower(color)
	}

	style := r.URL.Query().Get("style")
	if style == "" {
		style = "for-the-badge"
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		c := http.Client{
			Timeout: time.Second * 5,
		}
		resp, err := c.Get(getBadgeURL("Error", "Invalid timezone", "eb4511", style))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		w.WriteHeader(http.StatusBadRequest)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			return
		}
		return
	}
	_, tzOffset := time.Now().In(loc).Zone()
	var offset time.Duration = Offset
	offset += time.Duration(tzOffset)
	now := time.Now().Add(offset)
	data := now.Format("2006-01-02 15:04:05")

	c := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := c.Get(getBadgeURL(label, data, color, style))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return
	}
}

func syncTime() {
	var offset time.Duration

	response, err := ntp.Query(
		"time1.google.com",
	)
	if err != nil {
		log.Println(err)
	}
	offset += response.ClockOffset
	log.Println("Offset: from time1.google.com", response.ClockOffset, "Stratum", response.Stratum)

	response, err = ntp.Query(
		"time2.google.com",
	)
	if err != nil {
		log.Println(err)
	}
	offset += response.ClockOffset
	log.Println("Offset: from time2.google.com", response.ClockOffset, "Stratum", response.Stratum)

	response, err = ntp.Query(
		"time3.google.com",
	)
	if err != nil {
		log.Println(err)
	}
	offset += response.ClockOffset
	log.Println("Offset: from time3.google.com", response.ClockOffset, "Stratum", response.Stratum)

	response, err = ntp.Query(
		"time4.google.com",
	)
	if err != nil {
		log.Println(err)
	}
	offset += response.ClockOffset
	log.Println("Offset: from time4.google.com", response.ClockOffset, "Stratum", response.Stratum)

	atomic.StoreInt64((*int64)(&Offset), int64(offset/4))
}

//go:embed dist/*
var DistFS embed.FS

func main() {
	go func() {
		for {
			func() {
				defer func() {
					if err := recover(); err != nil {
						log.Println("==============================================================")
						log.Println("Worker panic:", err)
						log.Println("Recovering...")
						log.Println("==============================================================")
					}
				}()
				syncTime()
				log.Println("Synchronized time Offset:", Offset)
				time.Sleep(time.Second * 60)
			}()
		}
	}()
	mux := http.NewServeMux()
	// ./dist
	var dist http.FileSystem
	if _, err := os.Stat("dist"); err == nil {
		dist = http.Dir("dist")
	} else {
		f, err := fs.Sub(DistFS, "dist")
		if err != nil {
			log.Fatal(err)
		}
		dist = http.FS(f)
	}
	mux.HandleFunc("/time", timeHandle)

	healthz := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/health", healthz)
	mux.HandleFunc("/api/badge", badgeHandle)

	mux.Handle("/", http.FileServer(dist))
	lnHost := ":8080"
	hostEnv := os.Getenv("HOST")
	if hostEnv != "" {
		lnHost = hostEnv
	}
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		lnHost = ":" + portEnv
	} else {
		portEnv = "8080"
	}
	ipEnv := os.Getenv("IP")
	if ipEnv != "" {
		ip := net.ParseIP(ipEnv)
		if ip != nil {
			lnHost = ipEnv + ":" + portEnv
		}
		if ip.To16() != nil {
			lnHost = "[" + ipEnv + "]:" + portEnv
		}
	}
	//#nosec
	ln, err := net.Listen("tcp", lnHost)
	if err != nil {
		log.Fatal(err)
	}

	err = http.Serve(ln, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		mux.ServeHTTP(rw, r)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
