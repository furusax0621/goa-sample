package webapi

import (
	"context"
	"testing"
	"time"

	"github.com/furusax0621/goa-sample/webapi/app"
	"github.com/furusax0621/goa-sample/webapi/app/test"
	"github.com/google/go-cmp/cmp"
	"github.com/shogo82148/goa-v1"
)

func TestUserController_Get(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// test data
	user1 := &app.UserMedia{
		UserID:     1,
		GivenName:  "花子",
		FamilyName: "山田",
		CreatedAt:  time.Date(2021, 12, 30, 13, 14, 15, 0, time.UTC),
	}

	service := goa.New("sample")
	c := NewUserController(service)

	// initialize userMap
	c.userMap.Store(1, user1)
	c.userMap.Store(2, "invalid data")

	t.Run("request 1, returns OK", func(t *testing.T) {
		_, got := test.GetUserOK(t, ctx, service, c, 1)
		if diff := cmp.Diff(user1, got); diff != "" {
			t.Errorf("response mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("request 2, returns Internal Server Error", func(t *testing.T) {
		test.GetUserInternalServerError(t, ctx, service, c, 2)
	})

	t.Run("request 3, returns Not Found", func(t *testing.T) {
		test.GetUserNotFound(t, ctx, service, c, 3)
	})
}
