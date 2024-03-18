package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ReyLegar/vkTestProject/internal/models"
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

func TestActorPostgres_Create(t *testing.T) {
	setup()
	actorRepo := NewActorPostgres(db)
	actor := models.Actor{
		Name:      "John Doe",
		Gender:    "Male",
		BirthDate: time.Now().Format("2006-01-02"),
	}

	actorID, err := actorRepo.Create(actor)
	if err != nil {
		t.Errorf("Failed to create actor: %v", err)
	}

	if actorID == 0 {
		t.Error("Actor ID is zero")
	}
	tearDown()
}

func TestActorPostgres_Update(t *testing.T) {
	setup()
	actorRepo := NewActorPostgres(db)
	actor := models.Actor{
		Name:      "Jane Doe",
		Gender:    "Female",
		BirthDate: time.Now().Format("2006-01-02"),
	}

	actorID, err := actorRepo.Create(actor)
	if err != nil {
		t.Fatalf("Failed to create actor: %v", err)
	}

	actor.Name = "Jane Smith"
	err = actorRepo.Update(actorID, actor)
	if err != nil {
		t.Errorf("Failed to update actor: %v", err)
	}
	tearDown()
}

func TestActorPostgres_Delete(t *testing.T) {
	setup()
	actorRepo := NewActorPostgres(db)
	actor := models.Actor{
		Name:      "Bob Smith",
		Gender:    "Male",
		BirthDate: time.Now().Format("2006-01-02"),
	}

	actorID, err := actorRepo.Create(actor)
	if err != nil {
		t.Fatalf("Failed to create actor: %v", err)
	}

	err = actorRepo.Delete(actorID)
	if err != nil {
		t.Errorf("Failed to delete actor: %v", err)
	}
	tearDown()
}

func TestActorPostgres_GetByName(t *testing.T) {
	setup()
	actorRepo := NewActorPostgres(db)
	actor := models.Actor{
		Name:      "Alice",
		Gender:    "Female",
		BirthDate: time.Now().Format("2006-01-02"),
	}

	_, err := actorRepo.Create(actor)
	if err != nil {
		t.Fatalf("Failed to create actor: %v", err)
	}

	actorByName, err := actorRepo.GetByName("Alice")
	if err != nil {
		t.Errorf("Failed to get actor by name: %v", err)
	}

	if actorByName.Name != "Alice" {
		t.Errorf("Expected actor name to be 'Alice', got '%s'", actorByName.Name)
	}
	tearDown()
}

func TestActorPostgres_GetByID(t *testing.T) {
	setup()
	actorRepo := NewActorPostgres(db)
	actor := models.Actor{
		Name:      "Dave",
		Gender:    "Male",
		BirthDate: time.Now().Format("2006-01-02"),
	}

	actorID, err := actorRepo.Create(actor)
	if err != nil {
		t.Fatalf("Failed to create actor: %v", err)
	}

	actorByID, err := actorRepo.GetByID(actorID)
	if err != nil {
		t.Errorf("Failed to get actor by ID: %v", err)
	}

	if actorByID.Name != "Dave" {
		t.Errorf("Expected actor name to be 'Dave', got '%s'", actorByID.Name)
	}
	tearDown()
}

/* func TestActorPostgres_GetAllActorsAndMovies(t *testing.T) {
	actorRepo := NewActorPostgres(db)

	// Подготовка данных для теста предполагает наличие актеров и фильмов в базе данных

	actorsMovies, err := actorRepo.GetAllActorsAndMovies()
	if err != nil {
		t.Errorf("Failed to get all actors and movies: %v", err)
	}

	// Добавьте проверки здесь для убеждения, что данные возвращены корректно
} */
