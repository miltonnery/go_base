package main

import (
	"fmt"
	"github.com/dimiro1/health"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	viperConfiguration "miltonnery/go_base/configuration/viper"
	errorHandling "miltonnery/go_base/error"
	logJSON "miltonnery/go_base/log/json"
	"miltonnery/go_base/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Setting up configuration
	cfgSetting := viperConfiguration.NewSettingWithSamePath("./")
	configuration := viperConfiguration.NewConfiguration(cfgSetting)

	//Service port
	port := configuration.GetString("server.port")

	// Setting logger
	logFactory := logJSON.NewLogFactory(configuration)
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer logger.Sync() // flushes buffer, if any

	serviceLogger := logJSON.NewServiceLogJSON(configuration, logFactory, logger)

	// Setting middlewares and service
	serviceImpl := service.NewImpl(configuration, serviceLogger)
	middleware := service.NewMiddlewareImpl(configuration, serviceLogger, serviceImpl)

	endpoint := service.Endpoints{
		Service: service.NewServiceEndpoint(middleware),
	}
	//Creating the server
	mux := http.NewServeMux()

	//Setting the health check
	handler := health.NewHandler()
	//handler.AddChecker("Repository checker", SET_REPO_IMPL_HERE)

	errorMatcher := errorHandling.NewErrorMatcher()
	errorMatcher.LoadErrorMatchingCatalogFromConfiguration(configuration)
	errEncDec := service.NewErrorDecoderEncoder(errorMatcher)

	//Setting endpoints
	mux.Handle("/health", handler)
	mux.Handle("/path/to-endpoint", service.NewHTTPHandler(endpoint, errEncDec))

	//Se crea el canal para el manejo de error
	errs := make(chan error, 2)
	go func() {
		portMsj := "Port: " + port + ". Listening..."
		serviceLogger.InfoLite("main", portMsj)
		errs <- http.ListenAndServe(":"+port, mux)
	}()
	//This go-routine detects when the service is shut down
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		//scheduler.Stop()
	}()
	terminationErr := <-errs
	terminatedMsj := "Details: " + terminationErr.Error()
	serviceLogger.InfoLite("Terminated: ", terminatedMsj)

}
