package json

import (
	"git.lifemiles.net/lm-access/acc-gateway-svc/log"
	"reflect"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------/
func TestNewLifeMilesLogDetailsJSON(t *testing.T) {
	mockedLifeMilesLogDetailsJSON := mockedNewLifeMilesLogDetailsJSON()
	tests := []struct {
		name string
		want *LogDetailsJSON
	}{
		{
			name: "New LogDetailsJSON",
			want: mockedLifeMilesLogDetailsJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLifeMilesLogDetailsJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLifeMilesLogDetailsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockedNewLifeMilesLogDetailsJSON() *LogDetailsJSON {
	return &LogDetailsJSON{
		UUID:             "",
		IP:               "",
		MembershipNumber: "",
		Channel:          "",
		TimeStamp:        "",
		ServiceName:      "",
		Hostname:         "",
		RequestBody:      "",
		ResponseBody:     "",
		DestinationURL:   "",
		Step:             "",
		Level:            "",
		Product:          "",
		Application:      "",
		Class:            "",
		Method:           "",
		Language:         "",
		Thread:           "",
		LogMessage:       "",
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestSettersAndGetters(t *testing.T) {

	//Log Struct filling
	mockedJSONStruct := makeMockedJSONStruct()
	mockedEmptyJSONLogStruct := &LogDetailsJSON{}

	tests := []struct {
		name   string
		fields *LogDetailsJSON
		args   log.LifeMilesLogDetail
	}{
		{
			name:   "Test Setters",
			fields: mockedEmptyJSONLogStruct,
			args:   mockedJSONStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.SetUUID(tt.args.GetUUID())
			tt.fields.SetIP(tt.args.GetIP())
			tt.fields.SetMembershipNumber(tt.args.GetMembershipNumber())
			tt.fields.SetChannel(tt.args.GetChannel())
			tt.fields.SetTimeStamp(tt.args.GetTimeStamp())
			tt.fields.SetServiceName(tt.args.GetServiceName())
			tt.fields.SetHostname(tt.args.GetHostname())
			tt.fields.SetRequestBody(tt.args.GetRequestBody())
			tt.fields.SetResponseBody(tt.args.GetResponseBody())
			tt.fields.SetDestinationURL(tt.args.GetDestinationURL())
			tt.fields.SetStep(tt.args.GetStep())
			tt.fields.SetLevel(tt.args.GetLevel())
			tt.fields.SetProduct(tt.args.GetProduct())
			tt.fields.SetApplication(tt.args.GetApplication())
			tt.fields.SetClass(tt.args.GetClass())
			tt.fields.SetMethod(tt.args.GetMethod())
			tt.fields.SetLanguage(tt.args.GetLanguage())
			tt.fields.SetThread(tt.args.GetThread())
			tt.fields.SetLogMessage(tt.args.GetLogMessage())
		})
	}
}

func makeMockedJSONStruct() *LogDetailsJSON {
	return &LogDetailsJSON{
		UUID:             "123456789qwerty",
		IP:               "192.168.1.2",
		MembershipNumber: "04331604200",
		Channel:          "Test_Channel",
		TimeStamp:        time.Now().UTC().Format(time.Stamp),
		ServiceName:      "lifemiles-logging-svc",
		Hostname:         "12345ABCD",
		RequestBody:      "",
		ResponseBody:     "",
		DestinationURL:   "/mocked/path",
		Step:             "mockedStep",
		Level:            "mockedLevel",
		Product:          "Lifemiles logs",
		Application:      "Mocked Application",
		Class:            "mocked class",
		Method:           "mocked method",
		Language:         "GO",
		Thread:           "N/A",
		LogMessage:       "Mocked log message",
	}
}
