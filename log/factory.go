package log

import "net/http"

// LOG FACTORY DEFINITION

type LogFactory interface {
	Create(request *http.Request, response *http.Response, step string, level string, message string) Detail
	CreateLite(step string, level string, message string) Detail
}
