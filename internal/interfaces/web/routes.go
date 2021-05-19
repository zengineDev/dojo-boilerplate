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
	app.MiddlewareRegistry.Register("auth", middleware.Authentication())

}

func ConfigureWebRoute(app *dojo.Application) {
	app.Route.Use("csrf")

	app.Route.GetWithName("/","welcome", func(ctx dojo.Context, app *dojo.Application) error {
		return app.View(ctx, "welcome", make(map[string]interface{}))
	})

	app.Route.RouteGroup("/dashboard", func(router *dojo.Router) {
		router.Use("auth")
		router.GetWithName("/","dashboard", func(ctx dojo.Context, app *dojo.Application) error {
			return app.View(ctx, "dashboard", make(map[string]interface{}))
		})
	})

	httpErrors := middleware.HttpError()
	guestMiddleware := middleware.Guest()

	userStore := stores.UserStore{}
	userStore.Init()
	authHandler := handlers.AuthHandler{
		UserService: services.NewUserSvc(userStore),
	}

	app.Route.GetWithName("/login","auth.login.show", httpErrors(guestMiddleware(authHandler.ShowLoginForm)))
	app.Route.PostWithName("/login","auth.login.post", httpErrors(authHandler.HandleLoginInAttempt))
	app.Route.Post("/logout", httpErrors(authHandler.HandleLogoutAttempt))
}
