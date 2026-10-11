package main

import (
	_ "main/docs"
)

// @title           Swagger API
// @version         1.0
// @description     This is a api web server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":80"))
}
