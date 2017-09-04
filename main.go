package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "Welcome to home.redig.us\nThis page is a work in progress.")
    if (err != nil) {
        log.Fatal(err)
    }
}

func main() {
    log.Println("Starting server")
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":80", nil)
    if (err != nil) {
        log.Fatal(err)
    }
}
