package models

type Actor struct {
	ActorID   int    `json:"actorId`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}
