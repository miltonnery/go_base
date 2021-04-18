package viper

import (
	"reflect"
	"testing"
)

func TestNewSetting(t *testing.T) {

	basePath := "../"
	environmentPath := "."
	activeEnvironment := "test"
	fileName := "application"
	fileType := "yaml"
	withVault := false

	wantedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, withVault)

	type args struct {
		basePath          string
		environmentPath   string
		activeEnvironment string
		fileName          string
		fileType          string
		withVault         bool
	}
	tests := []struct {
		name string
		args args
		want *ConfigSetting
	}{
		{
			name: "Create NewSetting",
			args: args{
				basePath:          basePath,
				environmentPath:   environmentPath,
				activeEnvironment: activeEnvironment,
				fileName:          fileName,
				fileType:          fileType,
				withVault:         withVault,
			},
			want: wantedSetting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetting(tt.args.basePath, tt.args.environmentPath, tt.args.activeEnvironment, tt.args.fileName, tt.args.fileType, tt.args.withVault); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newMockedSetting(basePath, environmentPath, activeEnvironment string, fileName, fileType string, withVault bool) *ConfigSetting {
	s := ConfigSetting{
		basePath:          basePath,
		environmentPath:   environmentPath,
		activeEnvironment: activeEnvironment,
		fileName:          fileName,
		fileType:          fileType,
		withVault:         withVault,
	}
	return &s
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestNewDefaultSetting(t *testing.T) {

	basePath := "."
	environmentPath := "."
	activeEnvironment := ""
	fileName := "application"
	fileType := "yaml"
	withVault := true

	wantedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, withVault)

	tests := []struct {
		name string
		want *ConfigSetting
	}{
		{
			name: "Create NewDefaultSetting",
			want: wantedSetting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultSetting(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestNewDefaultSettingWithoutVault(t *testing.T) {

	basePath := "."
	environmentPath := "."
	activeEnvironment := ""
	fileName := "application"
	fileType := "yaml"
	withVault := false

	wantedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, withVault)

	tests := []struct {
		name string
		want *ConfigSetting
	}{
		{
			name: "Create NewDefaultSettingWithoutVault",
			want: wantedSetting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultSettingWithoutVault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultSettingWithoutVault() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestNewSettingWithSamePath(t *testing.T) {

	basePath := "."
	environmentPath := "."
	activeEnvironment := ""
	fileName := "application"
	fileType := "yaml"
	withVault := true

	wantedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, withVault)

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *ConfigSetting
	}{
		{
			name: "Create NewSettingWithSamePath",
			args: args{basePath},
			want: wantedSetting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSettingWithSamePath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSettingWithSamePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestNewSettingWithDifferentPath(t *testing.T) {

	basePath := "."
	environmentPath := "."
	activeEnvironment := ""
	fileName := "application"
	fileType := "yaml"
	withVault := true

	wantedSetting := newMockedSetting(basePath, environmentPath, activeEnvironment, fileName, fileType, withVault)

	type args struct {
		basePath        string
		environmentPath string
	}
	tests := []struct {
		name string
		args args
		want *ConfigSetting
	}{
		{
			name: "Test NewSettingWithDifferentPath",
			args: args{
				basePath:        basePath,
				environmentPath: environmentPath,
			},
			want: wantedSetting,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSettingWithDifferentPath(tt.args.basePath, tt.args.environmentPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSettingWithDifferentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_CheckDefault(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := false

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Check default settings",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			s.CheckDefault()
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_GetBasePath(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := true

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get base path from settings",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
			want: basePath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			if got := s.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_GetEnvironmentPath(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := true

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get environment path from settings",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
			want: environmentPath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			if got := s.GetEnvironmentPath(); got != tt.want {
				t.Errorf("GetEnvironmentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_GetName(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := true

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get file name from settings",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
			want: fileName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_GetType(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := true

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get file type from settings",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
			want: fileType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			if got := s.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/

func TestConfigSetting_WithVault(t *testing.T) {

	basePath := "."
	environmentPath := "."
	fileName := "application"
	fileType := "yaml"
	withVault := true

	type fields struct {
		basePath        string
		environmentPath string
		fileName        string
		fileType        string
		withVault       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Get with vault boolean flag",
			fields: fields{
				basePath:        basePath,
				environmentPath: environmentPath,
				fileName:        fileName,
				fileType:        fileType,
				withVault:       withVault,
			},
			want: withVault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ConfigSetting{
				basePath:        tt.fields.basePath,
				environmentPath: tt.fields.environmentPath,
				fileName:        tt.fields.fileName,
				fileType:        tt.fields.fileType,
				withVault:       tt.fields.withVault,
			}
			if got := s.WithVault(); got != tt.want {
				t.Errorf("WithVault() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------------------------------------------------/
