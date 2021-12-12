// Code generated by goagen v1.5.11, DO NOT EDIT.
//
// API "goa-sample": Application Controllers
//
// Command:
// $ main

package app

import (
	"context"
	goa "github.com/shogo82148/goa-v1"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Get(*GetUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/v1/users/:user_id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Get", "route", "GET /v1/users/:user_id")
}
