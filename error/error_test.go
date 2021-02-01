package errorhandling

import (
	"reflect"
	"testing"
)

func TestInternalError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get internal error message",
			fields: fields{
				Code:    GenericUnknownError,
				Message: ErrorDescription(GenericUnknownError),
			},
			want: "400 generic-errors: unknown error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ie := InternalError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := ie.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalError(t *testing.T) {
	type args struct {
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
		want InternalError
	}{
		{
			name: "Create new internal error",
			args: args{
				code:    PersistenceConnectionLost,
				message: ErrorDescription(PersistenceConnectionLost),
			},
			want: InternalError{
				Code:    PersistenceConnectionLost,
				Message: ErrorDescription(PersistenceConnectionLost),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalError(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalErrorWithCustomizedMessage(t *testing.T) {
	type args struct {
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
		want InternalError
	}{
		{
			name: "New interal error with a customized message",
			args: args{
				code:    GenericUnknownError,
				message: "This is a customized message",
			},
			want: InternalError{
				Code:    GenericUnknownError,
				Message: "This is a customized message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalErrorWithCustomizedMessage(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalErrorWithCustomizedMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
