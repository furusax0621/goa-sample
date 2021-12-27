package webapi

import (
	"time"

	"github.com/furusax0621/goa-sample/webapi/app"
	"github.com/furusax0621/goa-sample/webapi/models"
	"github.com/shogo82148/goa-v1"
)

// UserController implements controller of user resource.
type UserController struct {
	*goa.Controller
	userMap *models.SyncUserMap
}

// NewUserController returns UserController.
func NewUserController(service *goa.Service) *UserController {
	m := models.NewSyncUserMap()
	m.Store("太郎", "田中", time.Date(2021, 9, 25, 11, 22, 33, 0, time.UTC))

	return &UserController{
		Controller: service.NewController("UserController"),
		userMap:    m,
	}
}

// Get returns user.
func (c *UserController) Get(ctx *app.GetUserContext) error {
	media, ok := c.userMap.Load(ctx.UserID)
	if !ok {
		return ctx.NotFound()
	}

	return ctx.OK(media)
}

// Post creates new user.
func (c *UserController) Post(ctx *app.PostUserContext) error {
	now := time.Now()
	media := c.userMap.Store(ctx.Payload.GivenName, ctx.Payload.FamilyName, now)

	return ctx.Created(media)
}
