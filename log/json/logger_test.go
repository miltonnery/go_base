package json

import (
	"context"
	svcLog "git.lifemiles.net/lm-access/acc-gateway-svc/log"
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	lmLog "git.lifemiles.net/lm-go-libraries/lifemiles-go/log"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestNewLifeMilesServiceLogJSON(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedLifeMilesserviceLogJSON := makeMockedNewLifeMilesServiceLogJSON(mockedConfiguration, mockedFactory, mockedLogger)

	type args struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	tests := []struct {
		name string
		args args
		want *ServiceLogJSON
	}{
		{
			name: "Create new ServiceLogJSON struct",
			args: args{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			want: mockedLifeMilesserviceLogJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLifeMilesServiceLogJSON(tt.args.environment, tt.args.factory, tt.args.loggerJSON); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLifeMilesServiceLogJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeMockedNewLifeMilesServiceLogJSON(
	environment configuration.Config,
	factory svcLog.LifemilesLogFactory,
	loggerJSON lmLog.Logger) *ServiceLogJSON {
	return &ServiceLogJSON{
		environment: environment,
		factory:     factory,
		loggerJSON:  loggerJSON,
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesServiceLogJSON_Debug(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				message:  message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmslJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				loggerJSON:  tt.fields.loggerJSON,
			}
			lmslJSON.Debug(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesServiceLogJSON_Info(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedConfiguration.Set("log.logging-level", "INFO")
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				message:  message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmslJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				loggerJSON:  tt.fields.loggerJSON,
			}
			lmslJSON.Info(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesServiceLogJSON_Warn(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedConfiguration.Set("log.logging-level", "WARN")
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				message:  message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmslJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				loggerJSON:  tt.fields.loggerJSON,
			}
			lmslJSON.Warn(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesServiceLogJSON_Error(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedConfiguration.Set("log.logging-level", "ERROR")
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				message:  message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmslJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				loggerJSON:  tt.fields.loggerJSON,
			}
			lmslJSON.Error(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestLifeMilesServiceLogJSON_Fatal(t *testing.T) {
	mockedConfiguration := makeMockedConfiguration()
	mockedConfiguration.Set("log.logging-level", "FATAL")
	mockedFactory := makeMockedNewLifeMilesJSONLogFactory(mockedConfiguration)
	mockedLogger := lmLog.NewLogger()

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Config
		factory     svcLog.LifemilesLogFactory
		loggerJSON  lmLog.Logger
	}
	type args struct {
		request  *http.Request
		response *http.Response
		step     string
		message  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedFactory,
				loggerJSON:  mockedLogger,
			},
			args: args{
				request:  mockedRequest,
				response: nil,
				step:     step,
				message:  message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmslJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				loggerJSON:  tt.fields.loggerJSON,
			}
			lmslJSON.Fatal(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func makeMockedRequestForLogging() *http.Request {

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
