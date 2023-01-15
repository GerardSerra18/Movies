package model


import (
    "database/sql"
)

type Movie struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Year        int    `json:"year"`
    Genre       string `json:"genre"`
    Rating      float64 `json:"rating"`
}

type MovieModel struct {
    DB *sql.DB
}


// CreateMovie creates a new movie in the database
func (m *MovieModel) CreateMovie(movie *Movie) (int, error) {
	// Insert query and parameters
	query := "INSERT INTO movies (title, year, genre, rating) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute query
	err = stmt.QueryRow(movie.Title, movie.Year, movie.Genre, movie.Rating).Scan(&movie.ID)
	if err != nil {
		return 0, err
	}
	return movie.ID, nil
}

// GetMovie retrieves a movie from the database by ID
func (a *MovieModel) GetMovie(id int) (*Movie, error) {
	// Select query and parameters
	query := "SELECT id, title, year, genre, rating FROM movies WHERE id = $1"
	stmt, err := a.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	movie := &Movie{}
	err = stmt.QueryRow(id).Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Genre, &movie.Rating)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

// UpdateMovie updates a movie in the database
func (a *MovieModel) UpdateMovie(movie *Movie) error {
    // Update query and parameters
    query := "UPDATE movies SET first_name = $1, last_name = $2, gender = $3, age = $4 WHERE id = $5"
    stmt, err := a.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute query
    _, err = stmt.Exec(movie.Title, movie.Year, movie.Genre, movie.Rating, movie.ID)
    if err != nil {
        return err
    }
    return nil
}


// DeleteMovie deletes a movie from the database by ID
func (a *MovieModel) DeleteMovie(id int) error {
    // Delete query and parameters
    query := "DELETE FROM movies WHERE id = $1"
    stmt, err := a.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute query
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
