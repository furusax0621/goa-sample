package design

import (
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("goa-sample", func() {
	Title("Goa Sample")
	Scheme("https")
	Host("api.furusax.dev")
	Version("v1")
	BasePath("/v1")

	Consumes("application/json")
	Produces("application/json")
})
