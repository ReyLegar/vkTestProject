package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

type ActorPostgres struct {
	db *sql.DB
}

func NewActorPostgres(db *sql.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (a *ActorPostgres) Create(actor models.Actor) (int, error) {
	query := `INSERT INTO Actors (Name, Gender, BirthDate) VALUES ($1, $2, $3) RETURNING ActorID`
	var actorID int
	date, err := time.Parse("2006-01-02", actor.BirthDate)

	if err != nil {
		log.Println("Error parsing birth date:", err)
	}

	err = a.db.QueryRow(query, actor.Name, actor.Gender, date).Scan(&actorID)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return 0, err
	}

	return actorID, nil
}

func (a *ActorPostgres) Update(actorID int, actor models.Actor) error {
	query := `
		UPDATE Actors
		SET Name = $2, Gender = $3, BirthDate = $4
		WHERE ActorID = $1
	`

	date, err := time.Parse("2006-01-02", actor.BirthDate)
	if err != nil {
		log.Println("Error parsing birth date:", err)
		return err
	}

	_, err = a.db.Exec(query, actorID, actor.Name, actor.Gender, date)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return err
	}

	return nil
}

func (a *ActorPostgres) Delete(actorID int) error {
	query := `
		DELETE FROM Actors
		WHERE ActorID = $1
	`

	_, err := a.db.Exec(query, actorID)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return err
	}

	return nil
}

func (a *ActorPostgres) GetByName(name string) (models.Actor, error) {
	var actor models.Actor
	query := "SELECT * FROM Actors WHERE Name = $1"
	err := a.db.QueryRow(query, name).Scan(&actor.ActorID, &actor.Name, &actor.Gender, &actor.BirthDate)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return models.Actor{}, err
	}
	return actor, nil
}

func (a *ActorPostgres) GetByID(actorID int) (*models.Actor, error) {
	query := "SELECT Name, Gender, BirthDate FROM Actors WHERE ActorID = $1"

	row := a.db.QueryRow(query, actorID)

	var actor models.Actor
	var birthDate time.Time

	err := row.Scan(&actor.Name, &actor.Gender, &birthDate)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return nil, err
	}

	actor.BirthDate = birthDate.Format("2006-01-02")

	return &actor, nil
}

func (a *ActorPostgres) GetAllActorsAndMovies() (map[string][]models.Movie, error) {
	query := `
		SELECT a.Name, m.MovieID, m.Title, m.Description, m.ReleaseDate, m.Rating
		FROM Actors a
		INNER JOIN MovieActors ma ON a.ActorID = ma.ActorID
		INNER JOIN Movies m ON ma.MovieID = m.MovieID
	`

	rows, err := a.db.Query(query)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		return nil, err
	}
	defer rows.Close()

	actorsMovies := make(map[string][]models.Movie)
	for rows.Next() {
		var actorName string
		var movie models.Movie
		if err := rows.Scan(&actorName, &movie.MovieID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating); err != nil {
			log.Println("Error scanning rows:", err)
			return nil, err
		}

		actorsMovies[actorName] = append(actorsMovies[actorName], movie)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return actorsMovies, nil
}
