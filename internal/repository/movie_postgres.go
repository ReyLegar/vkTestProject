package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

type MoviePostgres struct {
	db *sql.DB
}

func NewMoviePostgres(db *sql.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (m *MoviePostgres) Create(movie models.Movie) (int, error) {
	query := `
        INSERT INTO Movies (Title, Description, ReleaseDate, Rating)
        VALUES ($1, $2, $3, $4)
        RETURNING MovieID
    `

	var movieID int
	err := m.db.QueryRow(
		query,
		movie.Title,
		movie.Description,
		movie.ReleaseDate,
		movie.Rating,
	).Scan(&movieID)
	if err != nil {
		fmt.Println("Error creating movie:", err)
		return 0, err
	}

	for _, actorID := range movie.Actors {
		_, err := m.db.Exec(`
            INSERT INTO MovieActors (MovieID, ActorID)
            VALUES ($1, $2)
        `, movieID, actorID)
		if err != nil {
			fmt.Println("Error adding actor to movie:", err)
			return 0, err
		}
	}

	return movieID, nil
}

func (m *MoviePostgres) Update(movieID int, movie models.Movie) error {
	query := `
        UPDATE Movies
        SET Title = $1, Description = $2, ReleaseDate = $3, Rating = $4
        WHERE MovieID = $5
    `

	_, err := m.db.Exec(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, movieID)
	if err != nil {
		fmt.Println("Error updating movie:", err)
		return err
	}

	return nil
}

func (m *MoviePostgres) Delete(movieID int) error {
	queryMovieActors := "DELETE FROM MovieActors WHERE MovieID = $1"

	_, err := m.db.Exec(queryMovieActors, movieID)
	if err != nil {
		fmt.Println("Error deleting movie actors:", err)
		return fmt.Errorf("failed to delete movie actors: %v", err)
	}

	queryMovies := "DELETE FROM Movies WHERE MovieID = $1"

	_, err = m.db.Exec(queryMovies, movieID)
	if err != nil {
		fmt.Println("Error deleting movie:", err)
		return fmt.Errorf("failed to delete movie: %v", err)
	}

	return nil
}

func (m *MoviePostgres) GetAll(sortBy string) ([]models.Movie, error) {
	var movies []models.Movie

	sortBy = strings.ToLower(sortBy)
	query := "SELECT * FROM Movies"
	switch sortBy {
	case "rating":
		query += " ORDER BY Rating DESC"
	case "title":
		query += " ORDER BY Title"
	case "release_date":
		query += " ORDER BY ReleaseDate"
	default:
		query += " ORDER BY Rating DESC"
	}

	rows, err := m.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching movies:", err)
		return nil, fmt.Errorf("failed to fetch movies: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.Movie
		var releaseDate time.Time
		err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Description, &releaseDate, &movie.Rating)
		if err != nil {
			fmt.Println("Error scanning movie row:", err)
			return nil, fmt.Errorf("failed to scan movie row: %v", err)
		}
		movie.ReleaseDate = releaseDate.Format("2006-01-02") // Преобразуем дату выпуска в строковый формат
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over movie rows:", err)
		return nil, fmt.Errorf("error while iterating over movie rows: %v", err)
	}

	return movies, nil
}

func (m *MoviePostgres) SearchByTitle(title string) ([]models.Movie, error) {
	query := `SELECT * FROM Movies WHERE LOWER(Title) LIKE LOWER($1)`
	rows, err := m.db.Query(query, fmt.Sprintf("%%%s%%", title))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("Error searching by title:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.MovieID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			fmt.Println("Error scanning movie row:", err)
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over movie rows:", err)
		return nil, err
	}

	return movies, nil
}

func (m *MoviePostgres) SearchByActorName(actorName string) ([]models.Movie, error) {
	query := `
		SELECT DISTINCT m.*
		FROM Movies m
		INNER JOIN MovieActors ma ON m.MovieID = ma.MovieID
		INNER JOIN Actors a ON ma.ActorID = a.ActorID
		WHERE a.Name LIKE $1
	`

	rows, err := m.db.Query(query, "%"+actorName+"%")
	if err != nil {
		fmt.Println("Error searching by actor name:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.MovieID,
			&movie.Title,
			&movie.Description,
			&movie.ReleaseDate,
			&movie.Rating,
		)
		if err != nil {
			fmt.Println("Error scanning movie row:", err)
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over movie rows:", err)
		return nil, err
	}

	return movies, nil
}
