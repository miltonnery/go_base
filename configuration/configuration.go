package configuration

import (
	"time"
)

// Defines the basic functions needed for a configuration retrieval across the microservice
type Configuration interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringMapString(key string) map[string]string
	Set(key string, value interface{})
}
