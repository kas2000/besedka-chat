package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port     = "8080"
	logLevel = 5
	addr     = flag.String("addr", ":8080", "http service address")
)


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
	http.FileServer(http.Dir("./static"))
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

