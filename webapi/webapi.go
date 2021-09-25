package webapi

import (
	"net/http"

	"github.com/furusax0621/goa-sample/webapi/app"
	"github.com/shogo82148/goa-v1"
)

//go:generate go run ./gen/main.go

// WebAPI provides Web API.
type WebAPI struct{}

// API provides HTTP handler.
func (w *WebAPI) API() http.Handler {
	service := goa.New("sample")

	// mount controllers
	app.MountUserController(service, NewUserController(service))

	return service.Mux
}
