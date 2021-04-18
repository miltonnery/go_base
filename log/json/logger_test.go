package json

import (
	"github.com/miltonnery/go_base/configuration"
	"github.com/miltonnery/go_base/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"reflect"
	"testing"
)

func TestNewServiceLogJSON(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	wantedServiceLogger := newMockedServiceLogJSON(mockedConfiguration, mockedLogFactory, zapLogger)

	type args struct {
		environment configuration.Configuration
		factory     log.LogFactory
		loggerJSON  *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *ServiceLogJSON
	}{
		{
			name: "Create a service logger in JSON format",
			args: args{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				loggerJSON:  zapLogger,
			},
			want: wantedServiceLogger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceLogJSON(tt.args.environment, tt.args.factory, tt.args.loggerJSON); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceLogJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newMockedServiceLogJSON(
	environment configuration.Configuration,
	factory log.LogFactory,
	loggerJSON *zap.Logger) *ServiceLogJSON {
	return &ServiceLogJSON{
		environment: environment,
		factory:     factory,
		zapLogger:   loggerJSON,
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_Debug(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
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
			name: "Create DEBUG Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
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
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}
			sLJSON.Debug(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_Info(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
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
			name: "Create INFO Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
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
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.Info(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_Warn(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
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
			name: "Create WARN Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
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
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.Warn(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_Error(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
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
			name: "Create ERROR Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
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
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.Error(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_Fatal(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	mockedRequest := makeMockedRequest()
	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
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
			name: "Create FATAL Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
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
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.Fatal(tt.args.request, tt.args.response, tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_DebugLite(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		step    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create DEBUG LITE Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				step:    step,
				message: message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.DebugLite(tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_InfoLite(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		step    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create INFO LITE Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				step:    step,
				message: message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.InfoLite(tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_WarnLite(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		step    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create WARN LITE Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				step:    step,
				message: message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.WarnLite(tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_ErrorLite(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		step    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create ERROR LITE Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				step:    step,
				message: message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.ErrorLite(tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_FatalLite(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	step := "TEST STEP"
	message := "TEST MESSAGE"

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		step    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create FATAL LITE Log",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				step:    step,
				message: message,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}

			sLJSON.FatalLite(tt.args.step, tt.args.message)
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestServiceLogJSON_checkLoggingLevel(t *testing.T) {

	mockedConfiguration := makeMockedConfiguration()
	mockedLogFactory := makeMockedNewJSONLogFactory(mockedConfiguration)
	zapLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	type fields struct {
		environment configuration.Configuration
		factory     log.LogFactory
		zapLogger   *zap.Logger
	}
	type args struct {
		loggingLevel string
		loggingType  string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantPermission bool
	}{
		{
			name: "Logging level INFO and logging type INFO",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				loggingLevel: "info",
				loggingType:  "info",
			},
			wantPermission: true,
		},
		{
			name: "Logging level WARN and logging type WARN",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				loggingLevel: "warn",
				loggingType:  "warn",
			},
			wantPermission: true,
		},
		{
			name: "Logging level ERROR and logging type ERROR",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				loggingLevel: "error",
				loggingType:  "error",
			},
			wantPermission: true,
		},
		{
			name: "Logging level FATAL and logging type FATAL",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				loggingLevel: "fatal",
				loggingType:  "fatal",
			},
			wantPermission: true,
		},
		{
			name: "Logging level incorrect and logging type FATAL",
			fields: fields{
				environment: mockedConfiguration,
				factory:     mockedLogFactory,
				zapLogger:   zapLogger,
			},
			args: args{
				loggingLevel: "something",
				loggingType:  "fatal",
			},
			wantPermission: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sLJSON := ServiceLogJSON{
				environment: tt.fields.environment,
				factory:     tt.fields.factory,
				zapLogger:   tt.fields.zapLogger,
			}
			if gotPermission := sLJSON.checkLoggingLevel(tt.args.loggingLevel, tt.args.loggingType); gotPermission != tt.wantPermission {
				t.Errorf("checkLoggingLevel() = %v, want %v", gotPermission, tt.wantPermission)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/
