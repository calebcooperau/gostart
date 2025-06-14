package main

import (
	"gostart/config"
	"gostart/internal/application"
	"gostart/internal/auth"
	"gostart/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	auth.NewAuthentication(cfg)
	app, err := application.NewApplication(cfg)
	if err != nil {
		panic(err)
	}
	defer app.Database.Close()
	routes.SetupRoutes(app.Gin, app.Repositories, app.Middlewares, app.Logger)
	app.Start()
}
