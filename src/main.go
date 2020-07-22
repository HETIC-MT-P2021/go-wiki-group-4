package main

import (
	"packages.hetic.net/gomvc/models"
	"packages.hetic.net/gomvc/router"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	env, _ := godotenv.Read(".env")

	if len(env["DB_PASSWORD"]) == 0 {
		panic("Ajouter les variables d'environnement au niveaux du .env.example")
	}

	dbCon := models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], env["DB_PORT"])

	router.StartRouter(env["API_PORT"], dbCon)
}
