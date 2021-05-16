package main

import (
	"github.com/zengineDev/dojo"
	"main/internal/interfaces/web"
)

func main() {

	cfg := dojo.LoadConfigs(&dojo.DefaultConfiguration{})

	app := dojo.New(*cfg)
	// load the config
	// load the middleware
	web.RegisterWebMiddlewares(app)
	// load the routes
	web.ConfigureWebRoute(app)
	// load middlwares
	// load services
	// start the server
	app.Serve()

}

func bootstrap() {
	// register the middleware and the stacks

	// load the routes

}
