package main

import (

	_ "main/docs"
)

func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":80"))
}
