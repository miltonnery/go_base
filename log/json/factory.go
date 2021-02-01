package json

import (
	"git.lifemiles.net/lm-access/acc-gateway-svc/log"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type LifeMilesJSONLogFactory struct {
	environment configuration.Config
}

func NewLifeMilesJSONLogFactory(environment configuration.Config) *LifeMilesJSONLogFactory {
	return &LifeMilesJSONLogFactory{
		environment: environment,
	}
}

func (lmf LifeMilesJSONLogFactory) Create(
	request *http.Request,
	response *http.Response,
	step string,
	level string,
	message string) log.LifeMilesLogDetail {

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
	lmlJSON := NewLifeMilesLogDetailsJSON()
	lmlJSON.SetUUID(requestUUID)
	lmlJSON.SetIP(logRequest.RemoteAddr)
	lmlJSON.SetMembershipNumber(lmf.environment.GetString("log.values.not-available"))
	lmlJSON.SetChannel(lmf.environment.GetString("log.values.channel"))
	lmlJSON.SetTimeStamp(timestamp)
	lmlJSON.SetServiceName(lmf.environment.GetString("log.values.service-name"))
	lmlJSON.SetHostname(hostName)
	//lmlJSON.SetRequestBody(requestBody)
	//lmlJSON.SetResponseBody(responseBody)
	//lmlJSON.SetDestinationURL(logRequest.URL.Path)
	lmlJSON.SetStep(step)
	lmlJSON.SetLevel(level)
	lmlJSON.SetProduct(lmf.environment.GetString("log.values.product"))
	lmlJSON.SetApplication(lmf.environment.GetString("log.values.application"))
	lmlJSON.SetClass(class)
	lmlJSON.SetMethod(class)
	lmlJSON.SetLanguage(lmf.environment.GetString("log.values.language"))
	lmlJSON.SetThread(lmf.environment.GetString("log.values.thread"))
	lmlJSON.SetLogMessage(message)

	return lmlJSON
}
