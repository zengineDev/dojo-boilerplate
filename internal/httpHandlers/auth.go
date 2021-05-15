package httpHandlers

import (
	"github.com/zengineDev/dojo"
	"main/internal/entities"
)

func ShowLoginForm(ctx dojo.Context, app *dojo.Application) error {
	return app.View(ctx, "auth/login", make(map[string]interface{}))
}

type LoginData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func HandleLoginInAttempt(ctx dojo.Context, app *dojo.Application) error {
	var data LoginData
	var user entities.User
	err := ctx.Bind(&data)
	if err != nil {
		return err
	}

	err = user.FindByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	match, err := app.Auth.ComparePasswordAndHash(data.Password, user.Password)
	if err != nil {
		return err
	}

	if !match {
		// TODO return the view with same flash
	}

	err = app.Auth.Login(user)

	// TODO return an redirect

	app.Logger.Info(data)

	return nil
}
