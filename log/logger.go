package log

import (
	"net/http"
)

// SERVICE LOG DEFINITION
type SeriviceLog interface {
	Debug(request *http.Request, response *http.Response, step string, message string)
	Info(request *http.Request, response *http.Response, step string, message string)
	Warn(request *http.Request, response *http.Response, step string, message string)
	Error(request *http.Request, response *http.Response, step string, message string)
	Fatal(request *http.Request, response *http.Response, step string, message string)
}
