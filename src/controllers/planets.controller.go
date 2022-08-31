package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	helpers "github.com/flpzow/starwars-api-go/src/helpers"
	"github.com/flpzow/starwars-api-go/src/models"
	"github.com/flpzow/starwars-api-go/src/repositories"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type planetController struct {
	Repository repositories.PlanetRepository
	Swapi      repositories.SwapiRepository
}

var _ PlanetController = &planetController{}

func NewPlanetController(planetRepo repositories.PlanetRepository, swapiRepo repositories.SwapiRepository) *planetController {
	return &planetController{
		Repository: planetRepo,
		Swapi:      swapiRepo,
	}
}

type PlanetController interface {
	GetPlanet(w http.ResponseWriter, r *http.Request)
	GetPlanets(w http.ResponseWriter, r *http.Request)
	CreatePlanet(w http.ResponseWriter, r *http.Request)
	DeletePlanet(w http.ResponseWriter, r *http.Request)
}

func (ctler *planetController) GetPlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planetId := vars["planetId"]

	if !primitive.IsValidObjectID(planetId) {
		helpers.Error(w, http.StatusBadRequest, errors.New("ID is not valid: "+planetId))
		return
	}

	planet, err := ctler.Repository.GetPlanetById(planetId)

	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.JsonRes(w, http.StatusOK, planet)
}

func (ctler *planetController) GetPlanets(w http.ResponseWriter, r *http.Request) {
	var err error
	var planets []models.Planet

	query := r.FormValue("search")

	if query != "" {
		planets, err = ctler.Repository.GetPlanetByName(query)
	} else {
		planets, err = ctler.Repository.GetAllPlanets()
	}

	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JsonRes(w, http.StatusOK, planets)
}

func (ctler *planetController) CreatePlanet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	jsonReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var planet models.Planet

	if err = json.Unmarshal(jsonReq, &planet); err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	swapiPlanet, err := ctler.Swapi.GetPlanetByName(planet.Name)
	if err != nil {
		planet.Appearances = 0
	} else {
		planet.Appearances = len(swapiPlanet.Films)
	}

	err = ctler.Repository.CreatePlanet(&planet)

	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	helpers.JsonRes(w, http.StatusCreated, planet)
}

func (ctler *planetController) DeletePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planetId := vars["planetId"]

	if !primitive.IsValidObjectID(planetId) {
		helpers.Error(w, http.StatusBadRequest, errors.New("Id is not valid: "+planetId))
		return
	}

	planetExists, err := ctler.Repository.GetPlanetById(planetId)

	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if planetExists == (models.Planet{}) {
		helpers.Error(w, http.StatusNotFound, nil)
		return
	}

	err = ctler.Repository.DeletePlanet(planetId)

	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.JsonRes(w, http.StatusOK, nil)
}
