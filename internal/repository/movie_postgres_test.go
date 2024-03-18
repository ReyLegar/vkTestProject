package repository_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ReyLegar/vkTestProject/internal/models"
	"github.com/ReyLegar/vkTestProject/internal/repository"
)

var db *sql.DB

func setup() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=2022 user=user dbname=mydb password=password sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
}

func tearDown() {
	db.Close()
}

func TestMoviePostgres_Create(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)
	movie := models.Movie{
		Title:       "Test Movie",
		Description: "Test Description",
		ReleaseDate: time.Now().Format("2006-01-02"),
		Rating:      5,
		Actors:      []int{1},
	}

	movieID, err := movieRepo.Create(movie)
	if err != nil {
		t.Fatalf("Failed to create movie: %v", err)
	}

	if movieID == 0 {
		t.Error("Movie ID is zero")
	}
	tearDown()
}

func TestMoviePostgres_Update(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)
	movie := models.Movie{
		Title:       "Test Movie",
		Description: "Updated Description",
		ReleaseDate: "2006-01-02",
		Rating:      4,
		Actors:      []int{1}, // Assuming actors with IDs 1 and 2 exist in the database
	}

	// Assuming there's a movie with ID 1 in the database
	err := movieRepo.Update(1, movie)
	if err != nil {
		t.Fatalf("Failed to update movie: %v", err)
	}
	tearDown()
}

func TestMoviePostgres_Delete(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)

	// Assuming there's a movie with ID 1 in the database
	err := movieRepo.Delete(1)
	if err != nil {
		t.Fatalf("Failed to delete movie: %v", err)
	}
	tearDown()
}

func TestMoviePostgres_GetAll(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)

	movies, err := movieRepo.GetAll("rating")
	if err != nil {
		t.Fatalf("Failed to get all movies: %v", err)
	}

	if len(movies) == 0 {
		t.Error("No movies returned")
	}
	tearDown()
}

func TestMoviePostgres_SearchByTitle(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)

	// Assuming there's a movie with title "Test Movie" in the database
	movies, err := movieRepo.SearchByTitle("Test Movie")
	if err != nil {
		t.Fatalf("Failed to search movies by title: %v", err)
	}

	if len(movies) == 0 {
		t.Error("No movies found by title")
	}
	tearDown()
}

func TestMoviePostgres_SearchByActorName(t *testing.T) {
	setup()
	movieRepo := repository.NewMoviePostgres(db)

	// Assuming there's an actor with name "John Doe" in the database
	movies, err := movieRepo.SearchByActorName("J")
	if err != nil {
		t.Fatalf("Failed to search movies by actor name: %v", err)
	}

	if len(movies) == 0 {
		t.Error("No movies found by actor name")
	}
	tearDown()
}
