package service

import (
	"context"
	"encoding/json"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"miltonnery/go_base/dto"
	errorHandling "miltonnery/go_base/error"
	"net/http"
)

// ErrorDecoderEncoder is intended to add the error matching functionality to the encoders and decoders
//in the transport layer. There is no need to use this struct if you do not need the usage of error conversion
type ErrorDecoderEncoder struct {
	errorMatcher *errorHandling.ErrorMatcher
}

func NewErrorDecoderEncoder(errorMatcher *errorHandling.ErrorMatcher) (eed *ErrorDecoderEncoder) {
	return &ErrorDecoderEncoder{errorMatcher: errorMatcher}
}

func NewHTTPHandler(endpoints Endpoints, eed *ErrorDecoderEncoder) http.Handler {
	serviceHandler := httpTransport.NewServer(
		endpoints.Service,
		eed.DecodeRetrievePNRRequest,
		eed.EncodeRetrievePNRResponse,
	)

	r := mux.NewRouter()
	r.Handle("/path/to-endpoint", serviceHandler).Methods("POST")
	return r
}
func (eed ErrorDecoderEncoder) DecodeRetrievePNRRequest(_ context.Context, request *http.Request) (req interface{}, err error) {
	var serviceRequest dto.Request
	err = json.NewDecoder(request.Body).Decode(&serviceRequest)
	req = serviceRequest

	return serviceRequest, err
}
func (eed ErrorDecoderEncoder) EncodeRetrievePNRResponse(ctx context.Context, responseWriter http.ResponseWriter, response interface{}) error {
	//Error verification
	if svcResponse, ok := response.(dto.InternalServiceResponse); ok == true {
		if svcResponse.Err != nil {
			encodeError(ctx, svcResponse.Err, responseWriter, eed.errorMatcher)
			return nil
		}
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(responseWriter).Encode(response)
}

//Error conversion to HTTP codes
func encodeError(_ context.Context, err error, w http.ResponseWriter, matcher *errorHandling.ErrorMatcher) interface{} {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err.(type) {
	case errorHandling.InternalError:
		intErr := err.(errorHandling.InternalError)
		em := matcher.GetHttpCodeFromInternalError(intErr)
		em.GetInternalCode()
		w.WriteHeader(em.GetExternalHTTPError())
		return json.NewEncoder(w).Encode(dto.ErrorResponse{Description: em.GetExternalMessage()})
	}
	return nil
}
