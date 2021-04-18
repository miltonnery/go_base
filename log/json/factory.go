package json

import (
	"github.com/miltonnery/go_base/configuration"
	"github.com/miltonnery/go_base/log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type LogFactory struct {
	configuration configuration.Configuration
}

func NewLogFactory(configuration configuration.Configuration) *LogFactory {
	return &LogFactory{
		configuration: configuration,
	}
}

func (jlf LogFactory) Create(
	request *http.Request,
	response *http.Response,
	step string,
	level string,
	message string) log.Detail {

	//Log data verification --------------------------------------------------------------------------------------------/
	//Filling nil values for avoiding errors during runtime
	var logRequest http.Request
	if request == nil {
		fakeBody := strings.NewReader("{}")
		logRequest = *httptest.NewRequest(http.MethodPost, "/empty-request", fakeBody)
	} else {
		logRequest = *request
	}
	//var logResponse http.Response
	//if response == nil {
	//	logResponse = *httptest.NewRecorder().Result()
	//} else {
	//	logResponse = *response
	//}

	requestUUID, ok := logRequest.Context().Value("request-id").(string)
	if ok {
		requestUUID = logRequest.Context().Value("request-id").(string)
	} else {
		requestUUID = "No UUID"
	}

	timestamp := time.Now().UTC().Format(time.Stamp)
	hostName, _ := os.Hostname()

	//Reading request

	//reqBodyBytes, _ := ioutil.ReadAll(logRequest.Body)
	//requestBody := string(reqBodyBytes)

	//Reading response
	//resBodyBytes, _ := ioutil.ReadAll(logResponse.Body)
	//responseBody := string(resBodyBytes)

	_, file, line, _ := runtime.Caller(2)
	lineString := strconv.Itoa(line)
	class := file + ":" + lineString

	//Log data filling -------------------------------------------------------------------------------------------------/
	lmlJSON := NewLogDetailsJSON()
	lmlJSON.SetUUID(requestUUID)
	lmlJSON.SetIP(logRequest.RemoteAddr)
	lmlJSON.SetTimeStamp(timestamp)
	lmlJSON.SetServiceName(jlf.configuration.GetString("log.values.service-name"))
	lmlJSON.SetHostname(hostName)
	//lmlJSON.SetRequestBody(requestBody)
	//lmlJSON.SetResponseBody(responseBody)
	//lmlJSON.SetDestinationURL(logRequest.URL.Path)
	lmlJSON.SetStep(step)
	lmlJSON.SetLevel(level)
	lmlJSON.SetProduct(jlf.configuration.GetString("log.values.product"))
	lmlJSON.SetApplication(jlf.configuration.GetString("log.values.application"))
	lmlJSON.SetClass(class)
	lmlJSON.SetMethod(class)
	lmlJSON.SetLanguage(jlf.configuration.GetString("log.values.language"))
	lmlJSON.SetLogMessage(message)

	return lmlJSON
}

func (jlf LogFactory) CreateLite(step string, level string, message string) log.Detail {

	//Log data verification --------------------------------------------------------------------------------------------/
	//Filling nil values for avoiding errors during runtime

	timestamp := time.Now().UTC().Format(time.Stamp)
	hostName, _ := os.Hostname()

	_, file, line, _ := runtime.Caller(2)
	lineString := strconv.Itoa(line)
	class := file + ":" + lineString

	//Log data filling -------------------------------------------------------------------------------------------------/
	lmlJSON := NewLogDetailsJSON()
	lmlJSON.SetTimeStamp(timestamp)
	lmlJSON.SetServiceName(jlf.configuration.GetString("log.values.service-name"))
	lmlJSON.SetUUID("N/A")
	lmlJSON.SetIP("N/A")
	lmlJSON.SetHostname(hostName)
	lmlJSON.SetStep(step)
	lmlJSON.SetLevel(level)
	lmlJSON.SetProduct(jlf.configuration.GetString("log.values.product"))
	lmlJSON.SetApplication(jlf.configuration.GetString("log.values.application"))
	lmlJSON.SetClass(class)
	lmlJSON.SetMethod(class)
	lmlJSON.SetLanguage(jlf.configuration.GetString("log.values.language"))
	lmlJSON.SetLogMessage(message)

	return lmlJSON
}
