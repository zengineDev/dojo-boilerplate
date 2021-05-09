package main

import "github.com/zengineDev/dojo"

func main() {

	cfg := dojo.LoadConfigs(&dojo.DefaultConfiguration{})

	app := dojo.New(*cfg)
	// load the config
	// create an app
	// load the routes
	ConfigureRoute(app)
	// load middlwares
	// load services
	// start the server
	app.Serve()

}
