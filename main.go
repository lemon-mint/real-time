package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/beevik/ntp"
)

var Offset = time.Duration(0)

type Time struct {
	T1 int64 `json:"t1"`
	T2 int64 `json:"t2"`
}

func timeHandle(w http.ResponseWriter, r *http.Request) {
	real := time.Now().Add(Offset)
	t := Time{
		T1: real.UnixMilli(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	t.T2 = time.Now().Add(Offset).UnixMilli()
	json.NewEncoder(w).Encode(t)
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
	http.ListenAndServe(":8080", nil)
}
