package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func timeHandlerClosure(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}

	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler)

	// with closure
	th := timeHandlerClosure(time.RFC1123)
	mux.Handle("/closure", th)
	log.Println("Listening..")
	http.ListenAndServe(":3001", mux)
}
