package json

import (
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	lmLog "git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
	"miltonnery/go_base/log"
	"net/http"
	"strings"
)

// SERVICE LOG IMPLEMENTATION FOR JSON

type ServiceLogJSON struct {
	environment configuration.Config
	factory     log.LogFactory
	loggerJSON  lmLog.Logger
}

//Constant definition
const (
	debug = "debug"
	info  = "info"
	warn  = "warn"
	error = "error"
	fatal = "fatal"
)

func NewLifeMilesServiceLogJSON(
	environment configuration.Config,
	factory log.LogFactory,
	loggerJSON lmLog.Logger) *ServiceLogJSON {
	return &ServiceLogJSON{
		environment: environment,
		factory:     factory,
		loggerJSON:  loggerJSON,
	}
}

func (lmslJSON ServiceLogJSON) Debug(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := lmslJSON.environment.GetString("log.logging-level")
	if lmslJSON.checkLoggingLevel(loggingLevel, debug) {
		levelInfo := lmslJSON.environment.GetString("log.values.log-level.debug")
		log := lmslJSON.factory.Create(request, response, step, levelInfo, message)
		lmslJSON.loggerJSON.Debug(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (lmslJSON ServiceLogJSON) Info(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := lmslJSON.environment.GetString("log.logging-level")
	if lmslJSON.checkLoggingLevel(loggingLevel, info) {
		levelInfo := lmslJSON.environment.GetString("log.values.log-level.informative")
		log := lmslJSON.factory.Create(request, response, step, levelInfo, message)
		lmslJSON.loggerJSON.Info(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (lmslJSON ServiceLogJSON) Warn(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := lmslJSON.environment.GetString("log.logging-level")
	if lmslJSON.checkLoggingLevel(loggingLevel, warn) {
		levelInfo := lmslJSON.environment.GetString("log.values.log-level.warning")
		log := lmslJSON.factory.Create(request, response, step, levelInfo, message)
		lmslJSON.loggerJSON.Warn(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (lmslJSON ServiceLogJSON) Error(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := lmslJSON.environment.GetString("log.logging-level")
	if lmslJSON.checkLoggingLevel(loggingLevel, error) {
		levelInfo := lmslJSON.environment.GetString("log.values.log-level.error")
		log := lmslJSON.factory.Create(request, response, step, levelInfo, message)
		lmslJSON.loggerJSON.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (lmslJSON ServiceLogJSON) Fatal(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := lmslJSON.environment.GetString("log.logging-level")
	if lmslJSON.checkLoggingLevel(loggingLevel, fatal) {
		levelInfo := lmslJSON.environment.GetString("log.values.log-level.fatal")
		log := lmslJSON.factory.Create(request, response, step, levelInfo, message)
		lmslJSON.loggerJSON.Error(log.GetLogMessage(), "LM_LOG", log)
	}
}

func (lmslJSON ServiceLogJSON) checkLoggingLevel(loggingLevel string, loggingType string) (permission bool) {

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
