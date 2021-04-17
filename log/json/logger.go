package json

import (
	"github.com/miltonnery/go_base/configuration"
	"github.com/miltonnery/go_base/log"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// SERVICE LOG IMPLEMENTATION FOR JSON

type ServiceLogJSON struct {
	environment configuration.Configuration
	factory     log.LogFactory
	zapLogger   *zap.Logger
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
	loggerJSON *zap.Logger) *ServiceLogJSON {
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
		logJSON := sLJSON.factory.Create(request, response, step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Debug("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) Info(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, info) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.informative")
		logJSON := sLJSON.factory.Create(request, response, step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Info("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) Warn(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, warn) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.warning")
		logJSON := sLJSON.factory.Create(request, response, step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Warn("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) Error(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, error) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.error")
		logJSON := sLJSON.factory.Create(request, response, step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Error("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) Fatal(request *http.Request, response *http.Response, step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, fatal) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.fatal")
		logJSON := sLJSON.factory.Create(request, response, step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Error("", zapFields...)
	}
}

// Lite logging versions -----------------------------------------------------------------------------------------------/
func (sLJSON ServiceLogJSON) DebugLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, debug) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.debug")
		logJSON := sLJSON.factory.CreateLite(step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Debug("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) InfoLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, info) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.informative")
		logJSON := sLJSON.factory.CreateLite(step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Info("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) WarnLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, warn) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.warning")
		logJSON := sLJSON.factory.CreateLite(step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Warn("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) ErrorLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, error) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.error")
		logJSON := sLJSON.factory.CreateLite(step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Error("", zapFields...)
	}
}

func (sLJSON ServiceLogJSON) FatalLite(step string, message string) {
	loggingLevel := sLJSON.environment.GetString("log.logging-level")
	if sLJSON.checkLoggingLevel(loggingLevel, fatal) {
		levelInfo := sLJSON.environment.GetString("log.values.log-level.fatal")
		logJSON := sLJSON.factory.CreateLite(step, levelInfo, message)
		zapFields := parseLogToZapFields(logJSON)
		sLJSON.zapLogger.Error("", zapFields...)
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

func parseLogToZapFields(logJSON log.Detail) (fields []zap.Field) {
	fields = append(fields, zap.String("level", logJSON.GetLevel()))
	fields = append(fields, zap.String("time", logJSON.GetTimeStamp()))
	fields = append(fields, zap.String("service-name", logJSON.GetServiceName()))
	fields = append(fields, zap.String("host", logJSON.GetHostname()))
	fields = append(fields, zap.String("ip", logJSON.GetIP()))
	fields = append(fields, zap.String("uuid", logJSON.GetUUID()))
	fields = append(fields, zap.String("step", logJSON.GetStep()))
	fields = append(fields, zap.String("message", logJSON.GetLogMessage()))
	fields = append(fields, zap.String("line", logJSON.GetMethod()))
	fields = append(fields, zap.String("request", logJSON.GetRequestBody()))
	fields = append(fields, zap.String("response", logJSON.GetResponseBody()))

	return
}
