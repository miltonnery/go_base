package viper

// file name and extension
const (
	FileName = "application"
	FileType = "yaml"
)

type Setting interface {
	GetBasePath() string
	GetEnvironmentPath() string
	GetActiveEnvironment() string
	GetName() string
	GetType() string
	WithVault() bool
}

// SETTING Interface implementation ------------------------------------------------------------------------------------/

type ConfigSetting struct {
	basePath          string
	environmentPath   string
	activeEnvironment string
	fileName          string
	fileType          string
	withVault         bool
}

func NewSetting(basePath, environmentPath, activeEnvironment, fileName, fileType string, withVault bool) *ConfigSetting {

	s := ConfigSetting{
		basePath:          basePath,
		environmentPath:   environmentPath,
		activeEnvironment: activeEnvironment,
		fileName:          fileName,
		fileType:          fileType,
		withVault:         withVault,
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

	if s.fileName == "" {
		s.fileName = FileName
	}

	if s.fileType == "" {
		s.fileType = FileType
	}
}

func (s *ConfigSetting) GetBasePath() string {
	return s.basePath
}

func (s *ConfigSetting) GetEnvironmentPath() string {
	return s.environmentPath
}

func (s *ConfigSetting) GetActiveEnvironment() string {
	return s.activeEnvironment
}

func (s *ConfigSetting) SetActiveEnvironment(activeEnvironment string) {
	s.activeEnvironment = activeEnvironment
}

func (s *ConfigSetting) GetName() string {
	return s.fileName
}

func (s *ConfigSetting) GetType() string {
	return s.fileType
}

func (s *ConfigSetting) WithVault() bool {
	return s.withVault
}
