package repository

import (
	"database/sql"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

type ActorRepository interface {
	Create(actor models.Actor) (int, error)
	Update(actorID int, actor models.Actor) error
	Delete(actorID int) error
	GetByName(name string) (models.Actor, error)
	GetByID(actorID int) (*models.Actor, error)
	GetAllActorsAndMovies() (map[string][]models.Movie, error)
}

type MovieRepository interface {
	Create(movie models.Movie) (int, error)
	Update(movieID int, movie models.Movie) error
	Delete(movieID int) error
	GetAll(string) ([]models.Movie, error)
	SearchByTitle(title string) ([]models.Movie, error)
	SearchByActorName(actorName string) ([]models.Movie, error)
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	ActorRepository
	MovieRepository
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		ActorRepository: NewActorPostgres(db),
		MovieRepository: NewMoviePostgres(db),
		Authorization:   NewAuthPostgres(db),
	}
}
