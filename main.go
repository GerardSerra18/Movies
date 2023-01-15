package main

import (
	//"database/sql"
	"log"
	"net/http"
	//"os"
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
    
    // Create actor controller
    movieCtrl := controller.NewMovieController(db)

    // Create router
    mux := http.NewServeMux()
    mux.HandleFunc("/actors/create", movieCtrl.CreateMovieHandler)
	mux.HandleFunc("/actors/get", movieCtrl.GetMovieHandler)
    mux.HandleFunc("/actors/view", movieCtrl.ViewMovieHandler)
    mux.HandleFunc("/actors/update", movieCtrl.UpdateMovieHandler)
    mux.HandleFunc("/actors/delete", movieCtrl.DeleteMovieHandler)
    mux.Handle("/actors/images/", http.StripPrefix("/movies/images/", http.FileServer(http.Dir("movies/images"))))

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
