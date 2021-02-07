package helper

import "github.com/labstack/echo/v4"

// Error -
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewError custom error
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["message"] = v.Message
	default:
		e.Errors["messgae"] = v.Error()
	}
	return e
}
