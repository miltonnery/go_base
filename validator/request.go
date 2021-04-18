package validator

import (
	"github.com/miltonnery/go_base/dto"
	errorHandling "github.com/miltonnery/go_base/error"
	"strings"
)

type RequestValidator struct {
	request *dto.Request
}

func NewRequestValidator(request *dto.Request) *RequestValidator {
	return &RequestValidator{request: request}
}

func (rv *RequestValidator) OK() error {

	if strings.EqualFold(rv.request.FirstAttribute, "") {
		err := errorHandling.NewInternalErrorWithCustomizedMessage(errorHandling.FieldMissing, "The email is empty")
		return err
	}

	if !strings.EqualFold(rv.request.FirstAttribute, "ABC") {
		err := errorHandling.NewInternalErrorWithCustomizedMessage(errorHandling.FieldDoesNotMatch, "The attributes don't match")
		return err
	}

	return nil
}
