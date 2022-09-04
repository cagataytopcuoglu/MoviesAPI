package core

import (
	"MovieAPI/pkg/log"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Item not found")
	ErrConflict            = errors.New("Your Item already exist")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	log.Logger.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func GetResponse(ctx echo.Context, err error) error {

	log.Logger.Error(err)
	switch err {
	case ErrInternalServerError:
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	case ErrNotFound:
		return ctx.JSON(http.StatusNotFound, err.Error())
	case ErrConflict:
		return ctx.JSON(http.StatusConflict, err.Error())
	default:
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
}
