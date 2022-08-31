package repositories

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct {
	Planets PlanetRepository
	Swapi   SwapiRepository
	db      *mongo.Client
}

func NewRepositories(db *mongo.Client) *Repositories {
	return &Repositories{
		Planets: NewPlanetsRepository(db),
		Swapi:   NewSwapiRepository(),
		db:      db,
	}
}
