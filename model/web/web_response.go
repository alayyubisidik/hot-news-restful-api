package web

type WebResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Errors []DetailError `json:"errors"`
}

type DetailError struct {
	Field   string `json:"field"`
	Message string    `json:"message"`
}
 