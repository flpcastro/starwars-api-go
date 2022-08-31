package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flpzow/starwars-api-go/src/controllers"
	"github.com/flpzow/starwars-api-go/src/db"
	"github.com/flpzow/starwars-api-go/src/repositories"
	"github.com/flpzow/starwars-api-go/src/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.NewDbClient()

	repositories := repositories.NewRepositories(db)

	controllers := controllers.NewController(repositories)

	route := routers.NewRouter(controllers)

	fmt.Println("running on port " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), route))
}
