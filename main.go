package main

import (
    "net/http"
    "os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello, World!</h1>"))
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"  
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)

    println("Launching", port)

    err := http.ListenAndServe(":"+port, mux)
    if err != nil {
        println("Server error:", err.Error())
    }

    println("Session is end") 
}
