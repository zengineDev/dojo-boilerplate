package handlers

import (
	"github.com/zengineDev/dojo"
	"main/internal/domain/services"
	"net/http"
)

type AuthHandler struct {
	UserService *services.UserSvc
}

func (h AuthHandler) ShowLoginForm(ctx dojo.Context, app *dojo.Application) error {
	data := make(map[string]interface{})
	data["Flash"] = ctx.Session().GetFlash("message")
	old := ctx.Session().GetFlash(dojo.FlashOldKey)
	if old != nil {
		data["Old"] = old[0].(map[string]interface{})
	}
	return app.View(ctx, "auth/login", data)
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
		ctx.Session().Flash("message", "auth.failed")
		old := make(map[string]interface{})
		old["Email"] = data.Email
		ctx.Session().WithOld(old)
		app.Logger.Warning("email not found")
		app.Route.Redirect(ctx, "/login")
	}

	match, err := app.Auth.ComparePasswordAndHash(data.Password, user.Password)
	if err != nil {
		return err
	}

	if !match {
		ctx.Session().Flash("message", "auth.failed")
		old := make(map[string]interface{})
		old["Email"] = data.Email
		ctx.Session().WithOld(old)
		app.Logger.Warning("password dont match")
		app.Route.Redirect(ctx, "/login")
	}

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
		return dojo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	app.Route.Redirect(ctx, "/")
	return nil
}
