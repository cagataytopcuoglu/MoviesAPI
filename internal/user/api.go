package user

import (
	"MovieAPI/internal/core"
	"MovieAPI/pkg/jwtHelper"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	baseUrl = "api/user"
)

func RegisterHandlers(instance *echo.Echo, repo Repository) {
	res := &resource{
		service: newService(repo),
	}

	instance.POST(fmt.Sprintf("%s/login", baseUrl), res.Login)
}

type resource struct {
	service Service
}

func (r *resource) Login(ctx echo.Context) error {
	entity := new(Login)
	err := ctx.Bind(&entity)
	if err != nil {
		return ctx.JSON(core.GetStatusCode(err), err.Error())
	}

	user, err := r.service.Login(entity)

	if err != nil {
		return ctx.JSON(core.GetStatusCode(err), err.Error())
	}

	if user == nil {
		return ctx.JSON(http.StatusAccepted, map[string]string{"message": "Invalid username or password"})
	}

	var token, error = jwtHelper.GetJwtByUser(user.Name, user.LastName)
	if error != nil {
		return ctx.JSON(core.GetStatusCode(err), err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
