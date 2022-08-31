package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/flpzow/starwars-api-go/src/models"
)

type swapiRepository struct{}

var _ SwapiRepository = &swapiRepository{}

type SwapiRepository interface {
	GetPlanetByName(name string) (models.PlanetApi, error)
}

func NewSwapiRepository() *swapiRepository {
	return &swapiRepository{}
}

func (swapi *swapiRepository) GetPlanetByName(name string) (models.PlanetApi, error) {
	http := http.Client{Timeout: time.Duration(60) * time.Second}
	res, err := http.Get(os.Getenv("PLANETS_API_BASE_URL") + name)

	if err != nil {
		fmt.Printf("Error %s", err)
		return models.PlanetApi{}, err
	}

	defer res.Body.Close()

	var planet models.PlanetApi

	jsonRes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return planet, err
	}

	var resultsPaged models.Paged
	if err = json.Unmarshal(jsonRes, &resultsPaged); err != nil {
		return planet, err
	}

	planet = resultsPaged.Results[0]

	return planet, nil
}
