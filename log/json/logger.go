package json

import (
	"go.uber.org/zap"
	"miltonnery/go_base/configuration"
	"miltonnery/go_base/log"
	"net/http"
	"strings"
)

// SERVICE LOG IMPLEMENTATION FOR JSON

type ServiceLogJSON struct {
	environment configuration.Configuration
	factory     log.LogFactory
	zapLogger   *zap.SugaredLogger
}

//Constant definition
const (
	debug = "debug"
	info  = "info"
	warn  = "warn"
	error = "error"
	fatal = "fatal"
)

func NewServiceLogJSON(
	environment configuration.Configuration,
	factory log.LogFactory,
	loggerJSON *zap.SugaredLogger) *ServiceLogJSON {
	return &ServiceLogJSON{
		environment: environment,
		factory:     factory,
		zapLogger:   loggerJSON,
	}
}

func (sLJSON ServiceLogJSON) Debug(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, debug) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.debug")
		log := sLJSON.factory.Create(request, response, step, levelInfo, message)
		sLJSON.zapLogger.Debug(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) Info(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, info) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.informative")
		log := sLJSON.factory.Create(request, response, step, levelInfo, message)
		sLJSON.zapLogger.Info(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) Warn(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, warn) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.warning")
		log := sLJSON.factory.Create(request, response, step, levelInfo, message)
		sLJSON.zapLogger.Warn(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) Error(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, error) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.error")
		log := sLJSON.factory.Create(request, response, step, levelInfo, message)
		sLJSON.zapLogger.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) Fatal(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, fatal) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.fatal")
		log := sLJSON.factory.Create(request, response, step, levelInfo, message)
		sLJSON.zapLogger.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

// Lite logging versions -----------------------------------------------------------------------------------------------/
func (sLJSON ServiceLogJSON) DebugLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, debug) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.debug")
		log := sLJSON.factory.CreateLite(step, levelInfo, message)
		sLJSON.zapLogger.Debug(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) InfoLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, info) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.informative")
		log := sLJSON.factory.CreateLite(step, levelInfo, message)
		sLJSON.zapLogger.Info(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) WarnLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, warn) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.warning")
		log := sLJSON.factory.CreateLite(step, levelInfo, message)
		sLJSON.zapLogger.Warn(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) ErrorLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, error) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.error")
		log := sLJSON.factory.CreateLite(step, levelInfo, message)
		sLJSON.zapLogger.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (sLJSON ServiceLogJSON) FatalLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, fatal) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.fatal")
		log := sLJSON.factory.CreateLite(step, levelInfo, message)
		sLJSON.zapLogger.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

// Auxiliary methods and functions -------------------------------------------------------------------------------------/

func (sLJSON ServiceLogJSON) checkLoggingLevel(loggingLevel string, loggingType string) (permission bool) {

	loggingLevel = strings.ToLower(loggingLevel)
	switch loggingLevel {
	case debug:
		{
			if strings.EqualFold(loggingType, debug) ||
				strings.EqualFold(loggingType, info) ||
				strings.EqualFold(loggingType, warn) ||
				strings.EqualFold(loggingType, error) ||
				strings.EqualFold(loggingType, fatal) {
				permission = true
				return
			}
		}
	case info:
		{
			if strings.EqualFold(loggingType, info) ||
				strings.EqualFold(loggingType, warn) ||
				strings.EqualFold(loggingType, error) ||
				strings.EqualFold(loggingType, fatal) {
				permission = true
				return
			}
		}
	case warn:
		{
			if strings.EqualFold(loggingType, warn) ||
				strings.EqualFold(loggingType, error) ||
				strings.EqualFold(loggingType, fatal) {
				permission = true
				return
			}
		}
	case error:
		{
			if strings.EqualFold(loggingType, error) ||
				strings.EqualFold(loggingType, fatal) {
				permission = true
				return
			}
		}

	case fatal:
		{
			if strings.EqualFold(loggingType, fatal) {
				permission = true
				return
			}
		}

	}

	return
}
