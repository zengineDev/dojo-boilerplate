package handlers

import (
	"github.com/zengineDev/dojo"
	"github.com/zengineDev/dojo/errorsx"
	"main/internal/domain/services"
)

type AuthHandler struct {
	UserService *services.UserSvc
}

func (h AuthHandler) ShowLoginForm(ctx dojo.Context, app *dojo.Application) error {
	return app.View(ctx, "auth/login", make(map[string]interface{}))
}

type LoginData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h AuthHandler) HandleLoginInAttempt(ctx dojo.Context, app *dojo.Application) error {
	var data LoginData
	err := ctx.Bind(&data)
	if err != nil {
		return err
	}

	user, err := h.UserService.FindByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	match, err := app.Auth.ComparePasswordAndHash(data.Password, user.Password)
	if err != nil {
		return err
	}

	if !match {
		// TODO return the view with same flash
		app.Logger.Warning("password dont match")
	}

	app.Logger.Info(user)

	err = app.Auth.Login(ctx, &user)
	if err != nil {
		return err
	}

	app.Route.Redirect(ctx, "/")

	return nil
}

func (h *AuthHandler) HandleLogoutAttempt(ctx dojo.Context, app *dojo.Application) error {
	err := app.Auth.Logout(ctx)
	if err != nil {
		return errorsx.BadRequest(err)
	}
	app.Route.Redirect(ctx, "/")
	return nil
}
