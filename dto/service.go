package dto

//Request body definition
type Request struct {
	FirstAttribute string `json:"first_attribute"`
}

// Response body definition
type Response struct {
	Output string `json:"output,omitempty"`
}

type ErrorResponse struct {
	Description string `json:"description,omitempty"`
}
