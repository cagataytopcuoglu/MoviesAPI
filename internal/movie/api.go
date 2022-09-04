package movie

import (
	"MovieAPI/internal/core"
	"MovieAPI/pkg/jwtHelper"
	"MovieAPI/pkg/pagination"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	baseUrl = "api/movie"
)

func RegisterHandlers(instance *echo.Echo, repo Repository) {

	res := &resource{
		service: newService(repo),
	}

	instance.GET(fmt.Sprintf("%s/all", baseUrl), res.getAll)
	instance.GET(fmt.Sprintf("%s/id/:id", baseUrl), res.getByParam)
	instance.POST(fmt.Sprintf("%s", baseUrl), res.insertHandler)
	instance.PATCH(fmt.Sprintf("%s", baseUrl), res.updateHandler)
	instance.DELETE(fmt.Sprintf("%s/id/:id", baseUrl), res.deleteHandler)
}

type resource struct {
	service Service
}

func (r *resource) getByParam(ctx echo.Context) error {

	filter := map[string]string{}

	if id := ctx.Param("_id"); len(id) > 0 {
		filter["_id"] = id
	}
	result, err := r.service.GetByParam(filter)

	if err != nil {
		return core.GetResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, result)
}

func (r *resource) getAll(ctx echo.Context) error {

	page := pagination.NewFromRequest(ctx.Request(), -1)

	result, err := r.service.GetAll(page)

	if err != nil {
		return core.GetResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, result)
}

func (r *resource) insertHandler(ctx echo.Context) error {

	jwt := ctx.Request().Header.Get("X-Authorization")

	if !jwtHelper.IsUserAuthenticated(jwt) {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized User")
	}

	entity := new(Movie)
	err := ctx.Bind(&entity)
	if err != nil {
		return ctx.JSON(core.GetStatusCode(err), err.Error())
	}
	val := entity.Validate()
	if len(val) > 0 {
		return ctx.JSON(http.StatusBadRequest, val)
	}
	err = r.service.Add(entity)
	if err != nil {
		return core.GetResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"Success": "true"})
}
func (r *resource) updateHandler(ctx echo.Context) error {

	jwt := ctx.Request().Header.Get("X-Authorization")

	if !jwtHelper.IsUserAuthenticated(jwt) {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized User")
	}

	entity := new(Movie)
	err := ctx.Bind(&entity)
	if err != nil {
		return ctx.JSON(core.GetStatusCode(err), err.Error())
	}

	err = r.service.Update(entity)
	if err != nil {
		return core.GetResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"Success": "true"})
}

func (r *resource) deleteHandler(ctx echo.Context) error {
	jwt := ctx.Request().Header.Get("X-Authorization")

	if !jwtHelper.IsUserAuthenticated(jwt) {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized User")
	}

	id := ctx.Param("id")

	err := r.service.Delete(id)

	if err != nil {
		return core.GetResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"Success": "true"})
}
