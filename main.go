package main

import (
	"encoding/json"
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

func main() {
	response, err := ntp.Query("time1.google.com")
	if err != nil {
		panic(err)
	}
	atomic.StoreInt64((*int64)(&Offset), int64(response.ClockOffset))

	// ./dist
	dist := http.Dir("dist")
	http.Handle("/", http.FileServer(dist))
	http.HandleFunc("/time", timeHandle)
	http.ListenAndServe(":8080", nil)
}
