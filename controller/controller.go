package controller

import (
	"Movies/model"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
)

// MovieController struct represents the controller for movies
type MovieController struct {
	Model model.MovieModel
}

// NewMovieController creates a new MovieController
func NewMovieController(db *sql.DB) *MovieController {
	return &MovieController{
		Model: model.MovieModel{DB: db},
	}
}

// CreateMovieHandler handles the creation of a new movie
func (c *MovieController) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create movie
	id, err := c.Model.CreateMovie(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with movie ID
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

// GetMovieHandler handles the retrieval of an movie by ID
func (c *MovieController) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Get movie ID from URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Get movie
	movie, err := c.Model.GetMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with actor
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// ViewMovieHandler handles the display of a movie details
func (c *MovieController) ViewMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	movie, _ := c.Model.GetMovie(id)
	t, _ := template.ParseFiles("views/movie.html")
	t.Execute(w, movie)
}

// UpdateMovieHandler handles the update of a movie details
func (c *MovieController) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update movie
	err = c.Model.UpdateMovie(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Write([]byte("Movie updated successfully"))
}

// DeleteMovieHandler handles the deletion of a movie by ID
func (c *MovieController) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Get movie ID from URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Delete movie
	err = c.Model.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Write([]byte("Movie deleted successfully"))
}
