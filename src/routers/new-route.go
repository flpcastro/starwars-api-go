package routers

import (
	"github.com/flpzow/starwars-api-go/src/controllers"
	"github.com/flpzow/starwars-api-go/src/routers/route"
	"github.com/gorilla/mux"
)

func NewRouter(ctler *controllers.Controller) *mux.Router {
	router := mux.NewRouter()
	return route.Config(router, ctler)
}
