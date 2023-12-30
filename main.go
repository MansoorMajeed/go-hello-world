package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, "Error getting the hostname", http.StatusInternalServerError)
        return
    }

    currentTime := time.Now().Format(time.RFC1123)

    fmt.Fprintf(w, "Hello from: %s\nCurrent Time: %s\n", hostname, currentTime)
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("starting http server at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("failed to start server at :8080")
    }
}

