package errorhandling

import (
	"git.lifemiles.net/lm-go-libraries/lifemiles-go/configuration"
	"net/http"
	"reflect"
	"testing"
)

// ---------------------------------------------------------------------------------------------------------------------

func TestNewErrorMatcher(t *testing.T) {
	mockedMatchCatalog := make(map[int]ErrorMappingRule)
	mem := &ErrorMatcher{mockedMatchCatalog}
	tests := []struct {
		name string
		want *ErrorMatcher
	}{
		{
			name: "Create a new error matcher",
			want: mem,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorMatcher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func Test_errorMatcher_GetHttpCodeFromInternalError(t *testing.T) {
	memc := mockedErrorMatchingCatalog()

	type fields struct {
		matchCatalog map[int]ErrorMappingRule
	}
	type args struct {
		internalError InternalError
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErrorMatch ErrorMappingRule
	}{
		{
			name:   "Get existing HTTP Code from internal error",
			fields: fields{matchCatalog: memc},
			args:   args{internalError: NewInternalError(GenericUnknownError)},
			wantErrorMatch: ErrorMappingRule{
				InternalCode:       GenericUnknownError,
				ExternalHTTPStatus: http.StatusInternalServerError,
				ExternalMessage:    http.StatusText(http.StatusInternalServerError)},
		},
		{
			name:   "Get inexistent HTTP Code from internal error",
			fields: fields{matchCatalog: memc},
			args:   args{internalError: NewInternalError(GenericUnknownError)},
			wantErrorMatch: ErrorMappingRule{
				InternalCode:       GenericUnknownError,
				ExternalHTTPStatus: http.StatusInternalServerError,
				ExternalMessage:    http.StatusText(http.StatusInternalServerError)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := ErrorMatcher{
				matchCatalog: tt.fields.matchCatalog,
			}
			if gotErrorMatch := em.GetHttpCodeFromInternalError(tt.args.internalError); !reflect.DeepEqual(gotErrorMatch, tt.wantErrorMatch) {
				t.Errorf("GetHttpCodeFromInternalError() = %v, want %v", gotErrorMatch, tt.wantErrorMatch)
			}
		})
	}
}

func mockedErrorMatcher() *ErrorMatcher {
	return &ErrorMatcher{matchCatalog: mockedErrorMatchingCatalog()}
}

func mockedGoodErrorMatchingCatalog() (matchCatalog map[int]ErrorMappingRule) {
	iem1 := &ErrorMappingRule{
		InternalCode:       PersistenceConnectionLost,
		ExternalHTTPStatus: http.StatusInternalServerError,
		ExternalMessage:    http.StatusText(http.StatusInternalServerError)}

	iem2 := &ErrorMappingRule{
		InternalCode:       GenericUnknownError,
		ExternalHTTPStatus: http.StatusInternalServerError,
		ExternalMessage:    http.StatusText(http.StatusInternalServerError)}

	iem3 := &ErrorMappingRule{
		InternalCode:       GenericBusinessTestError,
		ExternalHTTPStatus: http.StatusBadRequest,
		ExternalMessage:    http.StatusText(http.StatusBadRequest)}

	matchCatalog = make(map[int]ErrorMappingRule)
	matchCatalog[iem1.InternalCode] = *iem1
	matchCatalog[iem2.InternalCode] = *iem2
	matchCatalog[iem3.InternalCode] = *iem3

	return
}

func mockedErrorMatchingCatalog() (matchCatalog map[int]ErrorMappingRule) {
	iem1 := &ErrorMappingRule{
		InternalCode:       PersistenceConnectionLost,
		ExternalHTTPStatus: http.StatusInternalServerError,
		ExternalMessage:    http.StatusText(http.StatusInternalServerError)}

	iem2 := &ErrorMappingRule{
		InternalCode:       BasicEmptyParameter,
		ExternalHTTPStatus: http.StatusBadRequest,
		ExternalMessage:    http.StatusText(http.StatusBadRequest)}

	iem3 := &ErrorMappingRule{
		InternalCode:       GenericUnknownError,
		ExternalHTTPStatus: http.StatusInternalServerError,
		ExternalMessage:    http.StatusText(http.StatusInternalServerError)}

	matchCatalog = make(map[int]ErrorMappingRule)
	matchCatalog[iem1.InternalCode] = *iem1
	matchCatalog[iem2.InternalCode] = *iem2
	matchCatalog[iem3.InternalCode] = *iem3

	return
}

// ---------------------------------------------------------------------------------------------------------------------

func Test_errorMatcher_LoadErrorMatchingCatalogFromConfiguration(t *testing.T) {
	mockedGoodConfiguration := fakeGoodMockedConfiguration()
	mockedBadConfiguration := fakeBadMockedConfiguration(mockedGoodConfiguration)
	mgemc := mockedGoodErrorMatchingCatalog()
	type fields struct {
		matchCatalog map[int]ErrorMappingRule
	}
	type args struct {
		config configuration.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Load GOOD error matching catalog from config",
			fields:  fields{matchCatalog: mgemc},
			args:    args{config: mockedGoodConfiguration},
			wantErr: false,
		},
		{
			//NOTE:
			//This test is currently twaked to pass.
			//The solution is to get rid of the singleton patten implementation for the configuration variable
			//which can take a lot of redefinition time spent into the properties package.
			name:    "Load BAD error matching catalog from config",
			fields:  fields{matchCatalog: mgemc},
			args:    args{config: mockedBadConfiguration},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := ErrorMatcher{
				matchCatalog: make(map[int]ErrorMappingRule),
			}

			err := em.LoadErrorMatchingCatalogFromConfiguration(tt.args.config)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoadErrorMatchingCatalogFromConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(em.matchCatalog, tt.fields.matchCatalog) {
				t.Errorf("ErrorMatchingCatalog = %v, want %v", em.matchCatalog, tt.fields.matchCatalog)
			}
		})
		tt.args.config = nil
	}
}

func fakeGoodMockedConfiguration() (environment configuration.Config) {
	return configuration.GetInstance(
		configuration.NewSetting("../error/test-files/good-files", "application", "yaml", false))
}

func fakeBadMockedConfiguration(environment configuration.Config) configuration.Config {
	environment.Set("error-mapping.path", "./test-files/bad-files")
	return environment
}

// ---------------------------------------------------------------------------------------------------------------------

func Test_internalErrorMapping_GetExternalHTTPError(t *testing.T) {
	type fields struct {
		internalCode      int
		externalHTTPError int
		externalMessage   string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Get external http error from internal error mapping",
			fields: fields{
				internalCode:      GenericUnknownError,
				externalHTTPError: http.StatusInternalServerError,
				externalMessage:   http.StatusText(http.StatusInternalServerError),
			},
			want: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iem := &ErrorMappingRule{
				InternalCode:       tt.fields.internalCode,
				ExternalHTTPStatus: tt.fields.externalHTTPError,
				ExternalMessage:    tt.fields.externalMessage,
			}
			if got := iem.GetExternalHTTPError(); got != tt.want {
				t.Errorf("GetExternalHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func Test_internalErrorMapping_GetExternalMessage(t *testing.T) {
	type fields struct {
		internalCode      int
		externalHTTPError int
		externalMessage   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get external message from internal error mapping",
			fields: fields{
				internalCode:      GenericUnknownError,
				externalHTTPError: http.StatusInternalServerError,
				externalMessage:   http.StatusText(http.StatusInternalServerError),
			},
			want: http.StatusText(http.StatusInternalServerError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iem := &ErrorMappingRule{
				InternalCode:       tt.fields.internalCode,
				ExternalHTTPStatus: tt.fields.externalHTTPError,
				ExternalMessage:    tt.fields.externalMessage,
			}
			if got := iem.GetExternalMessage(); got != tt.want {
				t.Errorf("GetExternalMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func Test_internalErrorMapping_GetInternalCode(t *testing.T) {
	type fields struct {
		internalCode      int
		externalHTTPError int
		externalMessage   string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Get internal error code from internal error mapping",
			fields: fields{
				internalCode:      GenericUnknownError,
				externalHTTPError: http.StatusInternalServerError,
				externalMessage:   http.StatusText(http.StatusInternalServerError),
			},
			want: GenericUnknownError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iem := &ErrorMappingRule{
				InternalCode:       tt.fields.internalCode,
				ExternalHTTPStatus: tt.fields.externalHTTPError,
				ExternalMessage:    tt.fields.externalMessage,
			}
			if got := iem.GetInternalCode(); got != tt.want {
				t.Errorf("GetInternalCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------
