package repositories

import (
	"context"
	"os"

	"github.com/flpzow/starwars-api-go/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type planetsRepository struct {
	PlanetsCollection *mongo.Collection
}

type PlanetRepository interface {
	GetAllPlanets() ([]models.Planet, error)
	GetPlanetById(id string) (models.Planet, error)
	GetPlanetByName(name string) ([]models.Planet, error)
	CreatePlanet(planet *models.Planet) error
	DeletePlanet(id string) error
}

var _ PlanetRepository = &planetsRepository{}

var (
	collection = "planets"
)

func NewPlanetsRepository(db *mongo.Client) *planetsRepository {
	collection := db.Database(os.Getenv("MONGO_DB")).Collection(collection)

	return &planetsRepository{
		PlanetsCollection: collection,
	}
}

func (r *planetsRepository) GetAllPlanets() ([]models.Planet, error) {
	var planet models.Planet
	var planets []models.Planet

	filter := bson.D{}

	client, err := r.PlanetsCollection.Find(context.TODO(), filter)

	if err != nil {
		defer client.Close(context.TODO())
		return nil, err
	}

	for client.Next(context.TODO()) {
		err := client.Decode(&planet)

		if err != nil {
			return nil, err
		}

		planets = append(planets, planet)
	}

	return planets, nil
}

func (r *planetsRepository) GetPlanetById(id string) (models.Planet, error) {
	var planet models.Planet

	idPlanet, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return planet, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idPlanet}}

	err = r.PlanetsCollection.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		return planet, err
	}

	return planet, nil
}

func (r *planetsRepository) GetPlanetByName(name string) ([]models.Planet, error) {
	var planet models.Planet
	var planets []models.Planet

	filter := bson.M{
		"name": bson.M{
			"$regex": primitive.Regex{
				Pattern: name,
				Options: "i",
			},
		},
	}

	client, err := r.PlanetsCollection.Find(context.TODO(), filter)

	if err != nil {
		defer client.Close(context.TODO())
		return nil, err
	}

	for client.Next(context.TODO()) {
		err := client.Decode(&planet)

		if err != nil {
			return nil, err
		}

		planets = append(planets, planet)
	}

	return planets, nil
}

func (r *planetsRepository) CreatePlanet(planet *models.Planet) error {
	planet.ID = primitive.NewObjectID()

	_, err := r.PlanetsCollection.InsertOne(context.TODO(), planet)

	if err != nil {
		return err
	}

	return nil
}

func (r *planetsRepository) DeletePlanet(id string) error {
	idPlanet, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idPlanet}}

	_, err = r.PlanetsCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}
