package service

import (
	"github.com/ReyLegar/vkTestProject/internal/models"
	"github.com/ReyLegar/vkTestProject/internal/repository"
)

type ActorService struct {
	actorRepo repository.ActorRepository
}

func NewActorService(actorRepo repository.ActorRepository) *ActorService {
	return &ActorService{actorRepo: actorRepo}
}

func (s *ActorService) AddActor(actor models.Actor) (int, error) {
	actorID, err := s.actorRepo.Create(actor)
	if err != nil {
		return 0, err
	}

	return actorID, nil
}

func (s *ActorService) UpdateActor(actorID int, actor models.Actor) error {
	err := s.actorRepo.Update(actorID, actor)
	if err != nil {
		return err
	}

	return nil
}

func (s *ActorService) DeleteActor(actorID int) error {
	err := s.actorRepo.Delete(actorID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ActorService) GetAllActorsAndMovies() (map[string][]models.Movie, error) {
	return s.actorRepo.GetAllActorsAndMovies()
}
