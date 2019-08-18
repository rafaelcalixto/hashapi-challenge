package main

import (
    // Core libraries
    "fmt"
    "net/http"
    // Proprietary libraries
    api "hash_api"
)

// This is a simple package used just to call the API
func main() {
    fmt.Println("starting the API...")
    http.HandleFunc("/", api.Index_handler)
    // Route for the first task
    http.HandleFunc("/api/create_hash", api.CreateHash)
    // Route for the second task
    http.HandleFunc("/api/return_hashs", api.ReturnHash)
    // Route for the third task
    http.HandleFunc("/api/return_text", api.ReturnText)
    http.ListenAndServe(":8000", nil)
}
