package main

import (
	"log"
	"flag"
	"net/http"

	"../../communication"
)

var addr = flag.String("addr", ":2333", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "html/home.html")
}

func main() {
	flag.Parse()
	log.Println("http://localhost:2333")
	hub := communication.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.Handle("/ws", hub)
	http.ListenAndServe(*addr, nil)
}