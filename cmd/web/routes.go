package main

import "github.com/zengineDev/dojo"

func ConfigureRoute(app *dojo.Application) {

	app.GET("/", func(ctx dojo.Context) error {
		app.View(ctx, "welcome", make(map[string]interface{}))
		return nil
	})
}
