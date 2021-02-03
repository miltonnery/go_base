package dto

//This response is intended to capture the incoming responses from the gateway.
type InternalServiceResponse struct {
	Response interface{} `json:"response,omitempty"`
	Err      error       `json:"error,omitempty"`
}
