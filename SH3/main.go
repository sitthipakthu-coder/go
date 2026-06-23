package main

import (
	crud "SH3/CRUD"
	_ "SH3/docs" // load generated docs
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// httpserver.ShowHelloHTTP()
	// httpserver.ShowHelloFiber()
	crud.CRUD()
}
