package json

import (
	"context"
	"git.lifemiles.net/lm-access/acc-gateway-svc/log"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------/
func TestNewLifeMilesJSONLogFactory(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	fakeFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	type args struct {
		environment configuration.Config
	}
	tests := []struct {
		name string
		args args
		want *LifeMilesJSONLogFactory
	}{
		{
			name: "Create new JSON log factory",
			args: args{environment: mockedConfiguration},
			want: fakeFactory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLifeMilesJSONLogFactory(tt.args.environment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLifeMilesJSONLogFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeMockedConfiguration() (environment configuration.Config) {
	environment = configuration.GetInstance(
		configuration.NewSetting("./../../", "application-test", "yaml", false))
	return
}

func makeMockedNewLifeMilesJSONLogFactory(environment configuration.Config) *LifeMilesJSONLogFactory {
	return &LifeMilesJSONLogFactory{
		environment: environment,
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesJSONLogFactory_Create(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	step := "TEST STEP"
	level := "TEST LEVEL"
	message := "TEST MESSAGE"

	mockedRequest := makeMockedRequest()
	mockedLog := makeMockedLog(mockedConfiguration, mockedRequest, step, level, message)

	type fields struct {
		environment configuration.Config
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		level    string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   log.LifeMilesLogDetail
	}{
		{
			name:   "Create a new JSON log",
			fields: fields{mockedConfiguration},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				level:    level,
				message:  message,
			},
			want: mockedLog,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmf := LifeMilesJSONLogFactory{
				environment: tt.fields.environment,
			}
			if got := lmf.Create(tt.args.request, tt.args.response, tt.args.step, tt.args.level, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeMockedLog(environment configuration.Config, request *http.Request, step, level, message string) *LogDetailsJSON {
	//Log data verification --------------------------------------------------------------------------------------------/
	//Filling nil values for avoiding errors during runtime
	if request == nil {
		fakeBody := httptest.NewRecorder().Body
		request = httptest.NewRequest(http.MethodPost, "/empty-request", fakeBody)
	}

	response := httptest.NewRecorder().Result()

	requestUUID := request.Context().Value("request-id").(string)
	timestamp := time.Now().UTC().Format(time.Stamp)
	hostName, _ := os.Hostname()

	//Reading request
	reqBodyBytes, _ := ioutil.ReadAll(request.Body)
	requestBody := string(reqBodyBytes)

	//Reading response
	resBodyBytes, _ := ioutil.ReadAll(response.Body)
	responseBody := string(resBodyBytes)

	_, file, line, _ := runtime.Caller(2)
	lineString := strconv.Itoa(line)
	class := file + ":" + lineString

	//Log data filling -------------------------------------------------------------------------------------------------/
	lmlJSON := NewLifeMilesLogDetailsJSON()
	lmlJSON.SetUUID(requestUUID)
	lmlJSON.SetIP(request.RemoteAddr)
	lmlJSON.SetMembershipNumber(environment.GetString("log.values.not-available"))
	lmlJSON.SetChannel(environment.GetString("log.values.channel"))
	lmlJSON.SetTimeStamp(timestamp)
	lmlJSON.SetServiceName(environment.GetString("log.values.service-name"))
	lmlJSON.SetHostname(hostName)
	lmlJSON.SetRequestBody(requestBody)
	lmlJSON.SetResponseBody(responseBody)
	//lmlJSON.SetDestinationURL(request.URL.Path)
	lmlJSON.SetStep(step)
	lmlJSON.SetLevel(level)
	lmlJSON.SetProduct(environment.GetString("log.values.product"))
	lmlJSON.SetApplication(environment.GetString("log.values.application"))
	lmlJSON.SetClass(class)
	lmlJSON.SetMethod(class)
	lmlJSON.SetLanguage(environment.GetString("log.values.language"))
	lmlJSON.SetThread(environment.GetString("log.values.thread"))
	lmlJSON.SetLogMessage(message)

	return lmlJSON
}

func makeMockedRequest() *http.Request {

	fakeBody := httptest.NewRecorder().Body
	request := httptest.NewRequest(http.MethodPost, "/empty-request", fakeBody)

	hostname, _ := os.Hostname()
	requestID := uuid.New().String()
	ctx := context.Background()
	cwv := context.WithValue(ctx, "pod", hostname)
	cwv = context.WithValue(cwv, "request-id", requestID)
	request = request.WithContext(cwv)

	return request
}
