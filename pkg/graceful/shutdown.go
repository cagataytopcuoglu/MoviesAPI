package graceful

import (
	"MovieAPI/pkg/log"
	"context"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Shutdown shuts down the given HTTP server gracefully when receiving an os.Interrupt or syscall.SIGTERM signal.
// It will wait for the specified timeout to stop hanging HTTP handlers.
func Shutdown(instance *echo.Echo, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Logger.Infof("shutting down server with %s timeout", timeout)

	if err := instance.Shutdown(ctx); err != nil {
		log.Logger.Errorf("error while shutting down server: %v", err)
	} else {
		log.Logger.Info("server was shut down gracefully")
	}
}
