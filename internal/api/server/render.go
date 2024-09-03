package server

import (
	"net/http"

	"github.com/go-chi/render"
)

func ResponseJSON(w http.ResponseWriter, r *http.Request, code int, obj any) {
	if obj == nil {
		obj = struct{}{}
	}
	render.Status(r, code)
	render.JSON(w, r, obj)
}

func ErrorJSON(w http.ResponseWriter, r *http.Request, code int, err error) {
	resp := HttpError{
		ErrorMessage: err.Error(),
		ErrorCode:    code,
	}
	render.Status(r, code)
	render.JSON(w, r, &resp)
}
