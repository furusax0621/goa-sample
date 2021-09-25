// Code generated by goagen v1.5.10, DO NOT EDIT.
//
// API "goa-sample": user TestHelpers
//
// Command:
// $ main

package test

import (
	"context"
	"fmt"
	"github.com/furusax0621/goa-sample/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// GetUserOK runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetUserOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserController, userID string) (http.ResponseWriter, *app.UserMedia) {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/users/%v", userID),
	}
	req := httptest.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}
	prms["user_id"] = []string{fmt.Sprintf("%v", userID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, prms)
	getCtx, err := app.NewGetUserContext(goaCtx, req, service)
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	err = ctrl.Get(getCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.UserMedia
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.UserMedia)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.UserMedia", resp, resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
