package controllers

import "github.com/flpzow/starwars-api-go/src/repositories"

type Controller struct {
	PlanetsCtler PlanetController
}

func NewController(r *repositories.Repositories) *Controller {
	return &Controller{
		PlanetsCtler: NewPlanetController(r.Planets, r.Swapi),
	}
}
