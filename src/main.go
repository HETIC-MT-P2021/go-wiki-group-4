package main

import (
	"packages.hetic.net/gomvc/router"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// dbCon := connectToDB

	env, _ := godotenv.Read(".env")

	router.StartRouter(env["API_PORT"], dbCon)
}
