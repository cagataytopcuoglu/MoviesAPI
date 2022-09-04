package healthcheck

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterHandlers(instance *echo.Echo) {
	instance.GET("status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "Status HealthCheck OK"})
	})
}
