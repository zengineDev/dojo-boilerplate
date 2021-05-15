package main

import (
	"github.com/zengineDev/dojo"
	"main/internal/httpHandlers"
)

func ConfigureRoute(app *dojo.Application) {

	app.Route.Get("/", func(ctx dojo.Context, app *dojo.Application) error {
		return app.View(ctx, "welcome", make(map[string]interface{}))
	})

	app.Route.Get("/login", httpHandlers.ShowLoginForm)
	app.Route.Post("/login", httpHandlers.HandleLoginInAttempt)
}
