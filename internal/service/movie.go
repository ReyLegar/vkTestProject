package service

import (
	"github.com/ReyLegar/vkTestProject/internal/models"
	"github.com/ReyLegar/vkTestProject/internal/repository"
)

type MovieService struct {
	movieRepo repository.MovieRepository
	actorRepo repository.ActorRepository
}

func NewMovieService(movieRepo repository.MovieRepository, actorRepo repository.ActorRepository) *MovieService {
	return &MovieService{movieRepo: movieRepo, actorRepo: actorRepo}
}

func (s *MovieService) AddMovie(movie models.Movie) (int, error) {

	movieID, err := s.movieRepo.Create(movie)
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

func (s *MovieService) UpdateMovie(movieID int, movie models.Movie) error {

	err := s.movieRepo.Update(movieID, movie)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieService) DeleteMovie(movieID int) error {
	err := s.movieRepo.Delete(movieID)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieService) GetAllMovies(sortBy string) ([]models.Movie, error) {
	allMovies, err := s.movieRepo.GetAll(sortBy)
	if err != nil {
		return nil, err
	}

	return allMovies, nil

}

func (s *MovieService) SearchMoviesByTitle(title string) ([]models.Movie, error) {
	searchMovies, err := s.movieRepo.SearchByTitle(title)
	if err != nil {
		return nil, err
	}

	return searchMovies, nil
}

func (s *MovieService) SearchMoviesByActorName(actorName string) ([]models.Movie, error) {

	searchMovies, err := s.movieRepo.SearchByActorName(actorName)
	if err != nil {
		return nil, err
	}

	return searchMovies, nil
}
