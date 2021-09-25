package main

import (
	"log"
	"net/http"

	"github.com/furusax0621/goa-sample/webapi"
)

func main() {
	api := new(webapi.WebAPI)
	log.Fatal(http.ListenAndServe(":8080", api.API()))
}
