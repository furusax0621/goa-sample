package main

import (
	_ "github.com/furusax0621/goa-sample/webapi/design"

	"github.com/shogo82148/goa-v1/design"
	"github.com/shogo82148/goa-v1/goagen/codegen"
	genapp "github.com/shogo82148/goa-v1/goagen/gen_app"
	genswagger "github.com/shogo82148/goa-v1/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
		),
	)
}
