package errorhandling

import (
	"github.com/spf13/viper"
	"miltonnery/go_base/configuration"
	"net/http"
)

//---------------------------------------------------------------------------------------------------------------------|

type ErrorMappingRule struct {
	InternalCode       int    `mapstructure:"internal-code"`
	ExternalHTTPStatus int    `mapstructure:"external-http-status"`
	ExternalMessage    string `mapstructure:"external-message"`
}

type ErrorMap struct {
	ErrorMappingRules []ErrorMappingRule `mapstructure:"error-map"`
}

func (emr *ErrorMappingRule) GetInternalCode() int {
	return emr.InternalCode
}

func (emr *ErrorMappingRule) GetExternalHTTPError() int {
	return emr.ExternalHTTPStatus
}

func (emr *ErrorMappingRule) GetExternalMessage() string {
	return emr.ExternalMessage
}

//---------------------------------------------------------------------------------------------------------------------|
type ErrorMatcher struct {
	matchCatalog map[int]ErrorMappingRule
}

func NewErrorMatcher() *ErrorMatcher {
	matchCatalog := make(map[int]ErrorMappingRule)
	return &ErrorMatcher{matchCatalog}
}

func (em ErrorMatcher) LoadErrorMatchingCatalogFromConfiguration(config configuration.Configuration) (err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(config.GetString("error-mapping.path"))
	v.SetConfigName(config.GetString("error-mapping.filename"))

	//Reading configuration
	if readErr := v.ReadInConfig(); readErr != nil {
		err = NewInternalErrorWithCustomizedMessage(IOFileEnvironmentConfigurationNotFound, readErr.Error())
		return
	}

	//Loading properties into mapping rules struct
	var errMap ErrorMap
	if unErr := v.Unmarshal(&errMap); err != nil {
		err = NewInternalErrorWithCustomizedMessage(IOViperUnmarshalProblem, unErr.Error())
		return
	}

	//Switching mapping rules to a map
	for _, iem := range errMap.ErrorMappingRules {
		em.matchCatalog[iem.InternalCode] = iem
	}
	return err
}

func (em ErrorMatcher) GetHttpCodeFromInternalError(internalError InternalError) (errorMatch ErrorMappingRule) {

	if v, found := em.matchCatalog[internalError.Code]; found {
		errorMatch = v
	} else {
		errorMatch = ErrorMappingRule{
			InternalCode:       GenericUnknownError,
			ExternalHTTPStatus: http.StatusInternalServerError,
			ExternalMessage:    http.StatusText(http.StatusInternalServerError),
		}
	}

	return
}

//---------------------------------------------------------------------------------------------------------------------|
