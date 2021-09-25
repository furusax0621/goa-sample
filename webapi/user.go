package webapi

import (
	"sync"
	"time"

	"github.com/furusax0621/goa-sample/webapi/app"
	"github.com/shogo82148/goa-v1"
)

// UserController implements controller of user resource.
type UserController struct {
	*goa.Controller
	userMap *sync.Map
}

// NewUserController returns UserController.
func NewUserController(service *goa.Service) *UserController {
	m := new(sync.Map)
	m.Store(1, &app.UserMedia{
		UserID:     1,
		GivenName:  "太郎",
		FamilyName: "田中",
		CreatedAt:  time.Date(2021, 9, 25, 11, 22, 33, 0, time.UTC),
	})

	return &UserController{
		Controller: service.NewController("UserController"),
		userMap:    m,
	}
}

// Get returns user.
func (c *UserController) Get(ctx *app.GetUserContext) error {
	data, ok := c.userMap.Load(ctx.UserID)
	if !ok {
		return ctx.NotFound()
	}

	media, ok := data.(*app.UserMedia)
	if !ok {
		return ctx.InternalServerError()
	}

	return ctx.OK(media)
}
