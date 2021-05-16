package web

import (
	"github.com/zengineDev/dojo"
	"github.com/zengineDev/dojo/middleware"
	"main/internal/domain/services"
	"main/internal/interfaces/web/handlers"
	"main/internal/stores"
)

func RegisterWebMiddlewares(app *dojo.Application) {
	csrf := middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper:      middleware.DefaultSkipper,
		TokenLength:  32,
		TokenLookup:  "form:" + "_csrf",
		ContextKey:   "csrf",
		CookieName:   "_csrf",
		CookieMaxAge: 86400,
	})
	app.MiddlewareRegistry.Register("csrf", csrf)

}

func ConfigureWebRoute(app *dojo.Application) {
	app.Route.Use("csrf")
	app.Route.Get("/", func(ctx dojo.Context, app *dojo.Application) error {
		return app.View(ctx, "welcome", make(map[string]interface{}))
	})

	httpErrors := middleware.HttpError()
	guestMiddleware := middleware.Guest()

	userStore := stores.UserStore{}
	userStore.Init()
	authHandler := handlers.AuthHandler{
		UserService: services.NewUserSvc(userStore),
	}

	app.Route.Get("/login", httpErrors(guestMiddleware(authHandler.ShowLoginForm)))
	app.Route.Post("/login", httpErrors(authHandler.HandleLoginInAttempt))
	app.Route.Post("/logout", httpErrors(authHandler.HandleLogoutAttempt))
}
