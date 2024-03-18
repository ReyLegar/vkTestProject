package models

type Movie struct {
	MovieID     int     `json:"movieId"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float32 `json:"rating"`
	Actors      []int   `json:"actors"`
}
