package service

import (
	"github.com/miltonnery/go_base/configuration"
	"github.com/miltonnery/go_base/dto"
	"github.com/miltonnery/go_base/log"
)

type Service interface {
	Execute(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error)
}

// SERVICE INTERFACE IMPLEMENTATION ------------------------------------------------------------------------------------/

type Impl struct {
	configuration configuration.Configuration
	logger        log.SeriviceLog
}

func NewImpl(configuration configuration.Configuration, logger log.SeriviceLog) *Impl {
	return &Impl{
		configuration: configuration,
		logger:        logger,
	}
}

func (i Impl) Execute(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error) {
	i.logger.InfoLite("Service", "Starting the service layer")
	serviceResponse = &dto.Response{Output: serviceRequest.FirstAttribute}
	//err = errorHandling.NewInternalError(errorHandling.BasicEmptyParameter)
	return
}

// Auxiliary functions -------------------------------------------------------------------------------------------------/
