package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go-web-graph-app/internal/db"
    "go-web-graph-app/internal/handlers"
)

func main() {
    // Initialize the database
    database, err := db.InitDB("file:graph.db?cache=shared&mode=rwc")
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer database.Close()

    // Set up the router
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/api/graph", handlers.GetGraphData(database)).Methods("GET")

    // Start the server
    log.Println("Server is running on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}