package service

import (
	"miltonnery/go_base/configuration"
	"miltonnery/go_base/dto"
	"miltonnery/go_base/log"
	"miltonnery/go_base/validator"
)

type middlewareImpl struct {
	configuration configuration.Configuration
	logger        log.SeriviceLog
	next          Service
}

func NewMiddlewareImpl(configuration configuration.Configuration, logger log.SeriviceLog, serviceImpl Service) *middlewareImpl {
	return &middlewareImpl{
		configuration: configuration,
		logger:        logger,
		next:          serviceImpl,
	}
}

func (m middlewareImpl) Execute(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error) {
	m.logger.InfoLite("Middleware", "Beginning of middleware layer")

	//Request validation
	m.logger.InfoLite("Middleware", "Request content validation")
	rv := validator.NewRequestValidator(serviceRequest)
	err = rv.OK()
	if err != nil {
		return
	}

	//TODO: Set the middleware logic here
	serviceResponse, err = m.next.Execute(serviceRequest)
	if err != nil {
		return nil, err
	}
	return
}
