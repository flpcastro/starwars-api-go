package route

import (
	"net/http"

	"github.com/flpzow/starwars-api-go/src/controllers"
)

func CreatePlanetsRouters(ctler controllers.PlanetController) []Route {
	return []Route{
		{
			URI:     "/planets",
			Method:  http.MethodPost,
			Handler: ctler.CreatePlanet,
		},
		{
			URI:     "/planets/{planetId}",
			Method:  http.MethodGet,
			Handler: ctler.GetPlanet,
		},
		{
			URI:     "/planets",
			Method:  http.MethodGet,
			Handler: ctler.GetPlanets,
		},
		{
			URI:     "/planets/{planetId}",
			Method:  http.MethodDelete,
			Handler: ctler.DeletePlanet,
		},
	}
}
