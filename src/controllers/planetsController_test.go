package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flpzow/starwars-api-go/src/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwapiRepoMock struct {
	GetPlanetByNameMock func(name string) (models.PlanetApi, error)
}

func (m *SwapiRepoMock) GetPlanetByName(name string) (models.PlanetApi, error) {
	return m.GetPlanetByNameMock(name)
}

type PlanetRepoMock struct {
	GetAllPlanetsMock   func() ([]models.Planet, error)
	GetPlanetByIdMock   func(id string) (models.Planet, error)
	GetPlanetByNameMock func(name string) ([]models.Planet, error)
	CreatePlanetMock    func(planet *models.Planet) error
	DeletePlanetMock    func(id string) error
}

func (m *PlanetRepoMock) GetAllPlanets() ([]models.Planet, error) {
	return m.GetAllPlanetsMock()
}
func (m *PlanetRepoMock) GetPlanetById(id string) (models.Planet, error) {
	return m.GetPlanetByIdMock(id)
}
func (m *PlanetRepoMock) GetPlanetByName(name string) ([]models.Planet, error) {
	return m.GetPlanetByNameMock(name)
}
func (m *PlanetRepoMock) CreatePlanet(planet *models.Planet) error {
	return m.CreatePlanetMock(planet)
}
func (m *PlanetRepoMock) DeletePlanet(id string) error {
	return m.DeletePlanetMock(id)
}

func TestGetAllPlanet(t *testing.T) {

	request, _ := http.NewRequest(http.MethodGet, "/planets", nil)
	response := httptest.NewRecorder()

	planets := []models.Planet{
		{ID: primitive.NewObjectID(), Name: "Tatooine", Climate: "arid", Terrain: "desert"},
		{ID: primitive.NewObjectID(), Name: "Alderaan", Climate: "temperate", Terrain: "grasslands, mountains"},
	}

	SwapiRepositoryMock := &SwapiRepoMock{}

	PlanetsRepositoryMock := &PlanetRepoMock{}
	PlanetsRepositoryMock.GetAllPlanetsMock = func() ([]models.Planet, error) {
		return planets, nil
	}

	planetCtler := NewPlanetController(PlanetsRepositoryMock, SwapiRepositoryMock)

	planetCtler.GetPlanets(response, request)

	planetsR := []models.Planet{}

	_ = json.Unmarshal(response.Body.Bytes(), &planetsR)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.EqualValues(t, planets, planetsR)
	assert.EqualValues(t, len(planetsR), 2)
}

func TestGetAllPlanetByName(t *testing.T) {

	planet := "Tatooine"

	request, _ := http.NewRequest(http.MethodGet, "/planets?search="+planet, nil)
	response := httptest.NewRecorder()

	planets := []models.Planet{
		{ID: primitive.NewObjectID(), Name: "Tatooine", Climate: "arid", Terrain: "desert"},
	}

	SwapiRepositoryMock := &SwapiRepoMock{}

	PlanetsRepositoryMock := &PlanetRepoMock{}
	PlanetsRepositoryMock.GetPlanetByNameMock = func(name string) ([]models.Planet, error) {
		return planets, nil
	}

	planetCtler := NewPlanetController(PlanetsRepositoryMock, SwapiRepositoryMock)

	planetCtler.GetPlanets(response, request)

	planetsR := []models.Planet{}

	_ = json.Unmarshal(response.Body.Bytes(), &planetsR)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.EqualValues(t, planets, planetsR)
	assert.EqualValues(t, 1, len(planetsR))
}

func TestCreatePlanet(t *testing.T) {

	var jsonReq = []byte(`
		{
			"name": "Tatooine",
			"Climate": "Primavera",
			"Terrain": "Cerrado"
		}
	`)

	request, _ := http.NewRequest(http.MethodPost, "/planets", bytes.NewBuffer(jsonReq))
	response := httptest.NewRecorder()

	expectedPlanet := models.Planet{
		ID:          primitive.NewObjectID(),
		Name:        "Tatooine",
		Climate:     "Primavera",
		Terrain:     "Cerrado",
		Appearances: 5,
	}

	SwapiRepositoryMock := &SwapiRepoMock{}
	SwapiRepositoryMock.GetPlanetByNameMock = func(name string) (models.PlanetApi, error) {
		var planet = models.PlanetApi{
			Films: []string{
				"Filme 1",
				"Filme 2",
				"Filme 3",
				"Filme 5",
				"Filme 5",
			},
		}

		return planet, nil
	}

	PlanetsRepositoryMock := &PlanetRepoMock{}
	PlanetsRepositoryMock.CreatePlanetMock = func(planet *models.Planet) error {
		planet.ID = expectedPlanet.ID
		return nil
	}

	planetCtler := NewPlanetController(PlanetsRepositoryMock, SwapiRepositoryMock)

	planetCtler.CreatePlanet(response, request)

	planetR := models.Planet{}

	_ = json.Unmarshal(response.Body.Bytes(), &planetR)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, expectedPlanet, planetR)
}

func TestDeletePlanet(t *testing.T) {

	id, _ := primitive.ObjectIDFromHex("630cec980a61b5107a8f03a2")

	request, _ := http.NewRequest(http.MethodDelete, "/planets", nil)
	response := httptest.NewRecorder()

	request = mux.SetURLVars(request, map[string]string{
		"planetId": "630cec980a61b5107a8f03a2",
	})

	planet := models.Planet{
		ID:          id,
		Name:        "Tatooine",
		Climate:     "primavera",
		Terrain:     "caatinga",
		Appearances: 3,
	}

	SwapiRepositoryMock := &SwapiRepoMock{}
	PlanetsRepositoryMock := &PlanetRepoMock{}
	PlanetsRepositoryMock.GetPlanetByIdMock = func(id string) (models.Planet, error) {
		return planet, nil
	}

	PlanetsRepositoryMock.DeletePlanetMock = func(id string) error {
		return nil
	}

	planetCtler := NewPlanetController(PlanetsRepositoryMock, SwapiRepositoryMock)

	planetCtler.DeletePlanet(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}
