package main

import (
	"assessment2/routes"
	"assessment2/utils/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	g := gin.Default()
	db := postgres.NewConnection(postgres.BaseConfig()).Database

	routes.InitHttpRoute(g, db)
	g.Run()
}
