package httpserv

import (
	"github.com/labstack/echo/v4"
)

func (a *HTTPServer) apiSuccess(c echo.Context, data interface{}) error {
	return c.JSON(200, data)
}

type APIError struct {
	Error       string `json:"error"`
	Description string `json:"description,omitempty"`
	RequestID   string `json:"request_id,omitempty"`
}

func (a *HTTPServer) apiError(c echo.Context, httpCode int, messages ...string) error {
	e := APIError{}
	msgDebug := ""

	if len(messages) == 0 {
		e.Error = "Unknown error"
	} else if len(messages) == 1 {
		e.Error = messages[0]
	} else {
		e.Error = messages[0]
		e.Description = msgDebug + messages[1]
	}

	return c.JSON(httpCode, e)
}

func stringPtr(s string) *string {
	return &s
}
