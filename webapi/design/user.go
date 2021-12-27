package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/users")
	Action("get", func() {
		Routing(GET("/:user_id"))
		Params(func() {
			Param("user_id", Integer)
		})
		Response(OK, UserMedia)
		Response(NotFound)
	})

	Action("post", func() {
		Routing(POST(""))
		Payload(func() {
			Member("given_name", String, func() {
				Description("名前")
				Example("太郎")
			})
			Member("family_name", String, func() {
				Description("姓")
				Example("田中")
			})

			Required("given_name", "family_name")
		})
		Response(Created, UserMedia)
	})
})

var UserMedia = MediaType("application/vnd.user_media+json", func() {
	Attributes(func() {
		Attribute("user_id", Integer, func() {
			Description("ユーザーID")
			Example(1)
		})
		Attribute("given_name", String, func() {
			Description("名前")
			Example("太郎")
		})
		Attribute("family_name", String, func() {
			Description("姓")
			Example("田中")
		})
		Attribute("created_at", DateTime, func() {
			Description("ユーザー作成日時")
			Example("2021-09-25T11:22:33Z")
		})

		Required("user_id")
		Required("given_name")
		Required("family_name")
		Required("created_at")
	})

	View("default", func() {
		Attribute("user_id")
		Attribute("given_name")
		Attribute("family_name")
		Attribute("created_at")
	})
})
