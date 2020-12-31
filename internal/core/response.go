package core

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []APIError  `json:"errors,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string `json:"code,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}
