package main

import (
	_ "main/docs"
)

// @title           Swagger API
// @version         1.0
// @description     This is a api web server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//os.Setenv("SQLCONNECTSTRING", "root:@tcp(20.99.156.107:3306)/godev")
	router := newRouter()
	router.Logger.Fatal(router.Start(":80"))
}
