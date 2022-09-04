package main

import (
	"MovieAPI/internal/config"
	"MovieAPI/internal/core"
	"MovieAPI/internal/healthcheck"
	"MovieAPI/internal/movie"
	"MovieAPI/internal/user"
	"MovieAPI/pkg/graceful"
	"MovieAPI/pkg/log"
	"MovieAPI/pkg/mongoHelper"
	"MovieAPI/pkg/utils"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	port        = "5555"
	environment = "dev"
	debugMode   = false
	showHelp    = false
	instance    *echo.Echo

	AppConfig *config.Configuration
)

func main() {
	initFlags()
	initConfig()
	initApplication()
	initLogger()

	db, err := mongoHelper.ConnectDb(&AppConfig.MongoSettings)

	if err != nil {
		panic(err)
	}

	healthcheck.RegisterHandlers(instance)
	movie.RegisterHandlers(instance, movie.NewRepository(db))
	user.RegisterHandlers(instance, user.NewRepository(db))

	runAsService()
}

func initFlags() {

	//app parameter convention is most likely pain in teh ass.
	//so os env params will overwrite the application parameters, app parameters will overwrite to default parameters
	flag.StringVar(&port, "p", utils.EnvString("PORT", port), "HTTP listen address")
	flag.StringVar(&environment, "e", "", "Environment of apps")
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.BoolVar(&debugMode, "d", debugMode, "Debug Mode (default:false)")

	flag.Usage = flagUsage
	flag.Parse()
	if showHelp {
		showUsageAndExit(0)
	}

	if len(environment) == 0 {
		environment = core.GetEnvironmentVariables()
	}
	if err := core.SetEnvironment(environment); err != nil {
		panic(err.Error())
	}
}

func flagUsage() {
	fmt.Printf("Usage: api [-p 1234] [-e test] [-d 1]")
	flag.PrintDefaults()
}
func initConfig() {

	//Build configuration
	var vip = viper.GetViper()
	configFileName := environment
	var configReader = config.NewConfigReader("./config/", "config."+configFileName, vip)
	AppConfig = configReader.GetAllValues()
}
func showUsageAndExit(exitCode int) {
	flagUsage()
	os.Exit(exitCode)
}
func initLogger() {

	instance.Logger = log.SetupLogger()

	log.Logger.SetLevel(logrus.ErrorLevel)
	if debugMode {
		//enable echo debug mode
		log.Logger.SetLevel(logrus.DebugLevel)
	}
}

func initApplication() {

	instance = echo.New()
	instance.Debug = debugMode
	instance.HideBanner = !debugMode
	instance.HidePort = !debugMode

}
func runAsService() {

	log.Logger.Info("Service is starting with ", environment, " environment ", " and will serve on ", port, " port")

	//run barry run!
	go func() {
		if err := instance.Start(fmt.Sprintf(":%s", port)); err != nil {
			log.Logger.Fatalf("shutting down the server", err)
			panic(err)
		}
	}()
	graceful.Shutdown(instance, 2*time.Second)
}
