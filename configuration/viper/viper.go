package viper

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

// Implementation initialization ---------------------------------------------------------------------------------------/

const (
	// Base path is defined as default
	EnvFlag      = "active_env"
	DefaultEnv   = "local"
	BasePathFlag = "base_path"
	EnvPathFlag  = "env_path"
)

type Flag struct {
	Env      string
	BasePath string
	EnvPath  string
}

type viperConfiguration struct {
	viper   *viper.Viper
	setting Setting
	flags   *Flag
}

var viperConfig *viperConfiguration
var once sync.Once

// Initialization viperConfiguration
func NewConfiguration(setting Setting) *viperConfiguration {
	viper := viper.New()

	// Applying singleton pattern in order to handler only one instance
	// of viperConfiguration. For further information please follow the link.
	// https://progolang.com/how-to-implement-singleton-pattern-in-go/
	once.Do(func() {
		viperConfig = &viperConfiguration{
			viper:   viper,
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

// Configuration interface implementation with viper -------------------------------------------------------------------/

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
func (vc *viperConfiguration) Get(key string) interface{} {
	return vc.viper.Get(key)
}

// ToString casts an interface to a string type.
func (vc *viperConfiguration) GetString(key string) string {
	return vc.viper.GetString(key)
}

// ToBool casts an interface to a bool type.
func (vc *viperConfiguration) GetBool(key string) bool {
	return vc.viper.GetBool(key)
}

// ToInt casts an interface to an int type.
func (vc *viperConfiguration) GetInt(key string) int {
	return vc.GetInt(key)
}

// ToInt32 casts an interface to an int32 type.
func (vc *viperConfiguration) GetInt32(key string) int32 {
	return vc.GetInt32(key)
}

// ToInt64 casts an interface to an int64 type
func (vc *viperConfiguration) GetInt64(key string) int64 {
	return vc.GetInt64(key)
}

// ToFloat64 casts an interface to a float64 type.
func (vc *viperConfiguration) GetFloat64(key string) float64 {
	return vc.GetFloat64(key)
}

// ToTime casts an interface to a time.Time type.
func (vc *viperConfiguration) GetTime(key string) time.Time {
	return vc.viper.GetTime(key)
}

// ToDuration casts an interface to a time.Duration type.
func (vc *viperConfiguration) GetDuration(key string) time.Duration {
	return vc.GetDuration(key)
}

// ToStringMapString casts an interface to a map[string]string type.
func (vc *viperConfiguration) GetStringMapString(key string) map[string]string {
	return vc.viper.GetStringMapString(key)
}

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
func (vc *viperConfiguration) Set(key string, value interface{}) {
	vc.viper.Set(key, value)
}

// Auxiliary functions ------------------------------------------------------------------------------------------------/

func (vc *viperConfiguration) readConfig() {
	// General viperConfiguration
	vc.viper.SetConfigName(vc.setting.GetName())
	vc.viper.SetConfigType(vc.setting.GetType())
	vc.viper.AddConfigPath(vc.flags.BasePath)

	// Read viperConfiguration from base files
	if err := vc.viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read viperConfiguration"+
			" "+vc.setting.GetName()+"."+vc.setting.GetType()+" file. Error: %s", err.Error())
	}

	// By Env viperConfiguration
	vc.viper.SetConfigName(vc.setting.GetName() + "-" + vc.flags.Env)
	vc.viper.SetConfigType(vc.setting.GetType())
	vc.viper.AddConfigPath(vc.flags.EnvPath)

	// Merge viperConfiguration between base files with env
	if err := vc.viper.MergeInConfig(); err != nil {
		log.Printf("Unable to to Merge viperConfiguration file"+
			" "+vc.setting.GetName()+"-"+vc.flags.Env+"."+vc.setting.GetType()+" Error: %s", err.Error())
	}
}

// readFlags() sets the configuration flags for configuring viper
func readFlags(setting Setting) *Flag {
	// base_path, env_path, active_env
	env := flag.String(EnvFlag, DefaultEnv, "a string representing the environment file to read")
	basePath := flag.String(BasePathFlag, setting.GetBasePath(), "a string to tell where is the base_path to read a yml")
	envPath := flag.String(EnvPathFlag, setting.GetBasePath(), "a string to tell where is the env_path to read a yml per environment")
	flag.Parse()

	return &Flag{
		Env:      *env,
		BasePath: *basePath,
		EnvPath:  *envPath,
	}
}
