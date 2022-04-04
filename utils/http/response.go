package http

import "github.com/labstack/echo/v4"

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse() *response {
	return nil
}

func (resp *response) bytes() []byte {
	return nil
}

// func (resp *response) string() string {
// 	return nil
// }

func (resp *response) sendResponse(c echo.Context) {

}

// StatusNoContent 200
func StatusNoContent() {

}

// StatusBadRequest 400
func StatusBadRequest() {

}

// StatusNotFound 404
func StatusNotFound() {

}

// StatusMethodNotAllowed 405
func StatusMethodNotAllowed() {

}

func StatusConflict() {

}

// StatusInternalServerError 500
func StatusInternalServerError() {

}
