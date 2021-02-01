package viper

// file name and extension
const (
	FileName = "application"
	FileType = "yaml"
)

type Setting interface {
	GetBasePath() string
	GetEnvironmentPath() string
	GetName() string
	GetType() string
	WithVault() bool
}

// SETTING Interface implementation ------------------------------------------------------------------------------------/

type ConfigSetting struct {
	basePath        string
	environmentPath string
	name            string
	ftype           string
	withVault       bool
}

func NewSetting(basePath, environmentPath, name, ftype string, withVault bool) *ConfigSetting {

	s := ConfigSetting{
		basePath:        basePath,
		environmentPath: environmentPath,
		name:            name,
		ftype:           ftype,
		withVault:       withVault,
	}
	s.CheckDefault()
	return &s
}

func NewDefaultSetting() *ConfigSetting {
	s := ConfigSetting{withVault: true}
	s.CheckDefault()
	return &s
}

func NewDefaultSettingWithoutVault() *ConfigSetting {
	s := ConfigSetting{withVault: false}
	s.CheckDefault()
	return &s
}

func NewSettingWithSamePath(path string) *ConfigSetting {
	s := ConfigSetting{withVault: true, basePath: path, environmentPath: path}
	s.CheckDefault()
	return &s
}

func NewSettingWithDifferentPath(basePath, environmentPath string) *ConfigSetting {
	s := ConfigSetting{withVault: true, basePath: basePath, environmentPath: environmentPath}
	s.CheckDefault()
	return &s
}

func (s *ConfigSetting) CheckDefault() {
	if s.basePath == "" {
		s.basePath = "."
	}

	if s.environmentPath == "" {
		s.environmentPath = "."
	}

	if s.name == "" {
		s.name = FileName
	}

	if s.ftype == "" {
		s.ftype = FileType
	}
}

func (s *ConfigSetting) GetBasePath() string {
	return s.basePath
}

func (s *ConfigSetting) GetEnvironmentPath() string {
	return s.environmentPath
}

func (s *ConfigSetting) GetName() string {
	return s.name
}

func (s *ConfigSetting) GetType() string {
	return s.ftype
}

func (s *ConfigSetting) WithVault() bool {
	return s.withVault
}
