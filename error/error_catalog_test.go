package errorhandling

import "testing"

func TestErrorDescription(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get error description",
			args: args{code: GenericUnknownError},
			want: "generic-errors: unknown error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorDescription(tt.args.code); got != tt.want {
				t.Errorf("ErrorDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
