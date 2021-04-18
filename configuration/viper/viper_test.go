package viper

import (
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"testing"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------/

func TestNewConfiguration(t *testing.T) {
	basePath := "../../"
	environmentPath := "../../"
	activeEnvironment := "test"
	fileName := "application"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, true)
	mockedWantedConfiguration := newMockedConfiguration(mockedSetting)

	type args struct {
		setting Setting
	}
	tests := []struct {
		name string
		args args
		want *viperConfiguration
	}{
		{
			name: "Create NewConfiguration",
			args: args{setting: mockedSetting},
			want: mockedWantedConfiguration,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfiguration(tt.args.setting); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newMockedConfiguration(setting Setting) *viperConfiguration {
	v := viper.New()

	// Applying singleton pattern in order to handler only one instance
	// of viperConfiguration. For further information please follow the link.
	// https://progolang.com/how-to-implement-singleton-pattern-in-go/
	once.Do(func() {
		viperConfig = &viperConfiguration{
			viper:   v,
			setting: setting,
			flags:   readFlags(setting), // Read flags viperConfiguration
		}

		// Read viperConfiguration from file
		viperConfig.readConfig()

		// Read viperConfiguration from Vault and append to read viperConfiguration from file
		if viperConfig.flags.Env != DefaultEnv && setting.WithVault() {
			// TODO: Set vault reading here
		}
	})
	return viperConfig
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_Get(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	flags := &Flag{
		Env:      "test",
		BasePath: "../../",
		EnvPath:  "../../",
	}

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, true)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "Getting one key from configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   flags,
			},
			args: args{key: "log.startup-phrase.title.value"},
			want: ". GO BASE SERVICE .",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetString(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, true)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Getting one string key from configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "log.startup-phrase.title.value"},
			want: ". GO BASE SERVICE .",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetBool(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, true)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Getting one boolean value from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.boolean-value"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetBool(tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetInt(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Getting one int value from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.int-value"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetInt32(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		{
			name: "Getting one int32 value from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.int32-value"},
			want: 20000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetInt32(tt.args.key); got != tt.want {
				t.Errorf("GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetInt64(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "Getting one int64 value from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.int64-value"},
			want: 2000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetInt64(tt.args.key); got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetFloat64(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Getting one float64 value from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.float64-value"},
			want: 20000000.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetFloat64(tt.args.key); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetTime(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	wantedTime, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "Getting time from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.time-value"},
			want: wantedTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetTime(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetDuration(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	wantedDuration := 10 * time.Second

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name: "Getting duration from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "test.duration-value"},
			want: wantedDuration,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetDuration(tt.args.key); got != tt.want {
				t.Errorf("GetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_GetStringMapString(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	externalErrorMap := make(map[string]string, 0)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name: "Getting duration from key in configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{key: "error-map"},
			want: externalErrorMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}
			if got := vc.GetStringMapString(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_Set(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Setting key-value pair into environment configurations",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   nil,
			},
			args: args{
				key:   "mock.key",
				value: "mock value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}

			vc.Set(tt.args.key, tt.args.value)

			argValue := tt.args.value.(string)

			if !strings.EqualFold(argValue, "mock value") {
				t.Errorf("The set value is not the expected one. Expected: %v, Got: %v", tt.args.value, vc.GetString("mock.key"))
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func Test_viperConfiguration_readConfig(t *testing.T) {

	basePath := "../../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application-test"
	fileType := "yaml"

	mockedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, false)
	mockedConfiguration := newMockedConfiguration(mockedSetting)

	flags := &Flag{
		Env:      "test",
		BasePath: "../../",
		EnvPath:  "../../",
	}

	type fields struct {
		viper   *viper.Viper
		setting Setting
		flags   *Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Read config",
			fields: fields{
				viper:   mockedConfiguration.viper,
				setting: mockedSetting,
				flags:   flags,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := &viperConfiguration{
				viper:   tt.fields.viper,
				setting: tt.fields.setting,
				flags:   tt.fields.flags,
			}

			vc.readConfig()

		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/
