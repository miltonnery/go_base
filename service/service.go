package service

import (
	"miltonnery/go_base/configuration"
	"miltonnery/go_base/dto/service"
	"miltonnery/go_base/log"
)

type Service interface {
	Execute(serviceRequest *service.Request) (serviceResponse *service.Response, err error)
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

func (i Impl) Execute(serviceRequest *service.Request) (serviceResponse *service.Response, err error) {
	panic("implement me")
}

// Auxiliary functions -------------------------------------------------------------------------------------------------/
