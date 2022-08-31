package route

import (
	"net/http"

	"github.com/flpzow/starwars-api-go/src/controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Config(router *mux.Router, ctler *controllers.Controller) *mux.Router {
	routers := CreatePlanetsRouters(ctler.PlanetsCtler)

	for _, route := range routers {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return router
}
