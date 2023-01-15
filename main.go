package main

import (
	"log"
	"net/http"
	"time"
	"Movies/controller"
	_"github.com/lib/pq"
)

func main() {
    // Connect to database
    db, err := connectToDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize database
    if err = initDB(); err != nil {
        log.Fatal(err)
    }
    
    // Create movie controller
    movieCtrl := controller.NewMovieController(db)

    // Create router
    mux := http.NewServeMux()
    mux.HandleFunc("/movies/create", movieCtrl.CreateMovieHandler)
	mux.HandleFunc("/movies/get", movieCtrl.GetMovieHandler)
    mux.HandleFunc("/movies/view", movieCtrl.ViewMovieHandler)
    mux.HandleFunc("/movies/update", movieCtrl.UpdateMovieHandler)
    mux.HandleFunc("/movies/delete", movieCtrl.DeleteMovieHandler)

    // Create server
    server := &http.Server{
        Addr:         ":8000",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    // Start server
    log.Println("Starting server on :8000")
    log.Fatal(server.ListenAndServe())
}
