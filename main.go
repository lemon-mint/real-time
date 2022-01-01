package main

import (
	"log"
	"net/http"
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
	//Use GMT Date
	w.Header().Set("Date", time.Now().Add(Offset).UTC().Format(time.RFC1123))
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

func main() {
	go func() {
		for {
			syncTime()
			log.Println("Synchronized time Offset:", Offset)
			time.Sleep(time.Second * 60)
		}
	}()
	// ./dist
	dist := http.Dir("dist")
	http.Handle("/", http.FileServer(dist))
	http.HandleFunc("/time", timeHandle)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", nil)
}
