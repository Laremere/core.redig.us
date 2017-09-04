package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/ghthor/gowol"
)

func handler(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "Welcome to home.redig.us")
    if (err != nil) {
        log.Println(err)
    }
}

func wakeScott(w http.ResponseWriter, r *http.Request) {
    err := wol.MagicWake("30:5a:3a:05:da:20", "255.255.255.255")
    if (err != nil) {
	log.Println(err)
    }
}

func main() {
    log.Println("Starting server")
    http.HandleFunc("/", handler)
    http.HandleFunc("/wake/scott", wakeScott)
    err := http.ListenAndServe(":80", nil)
    if (err != nil) {
        log.Fatal(err)
    }
}
