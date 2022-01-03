package main

import (
	"embed"
	"log"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"

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
			syncTime()
			log.Println("Synchronized time Offset:", Offset)
			time.Sleep(time.Second * 60)
		}
	}()
	mux := http.NewServeMux()
	// ./dist
	var dist http.FileSystem
	if _, err := os.Stat("dist"); err == nil {
		dist = http.Dir("dist")
	} else {
		dist = http.FS(DistFS)
	}
	mux.HandleFunc("/time", timeHandle)

	healthz := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/health", healthz)

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

	http.ListenAndServe(lnHost, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		mux.ServeHTTP(rw, r)
	}))
}
