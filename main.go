package main

import (
	crud "learngo/chapter3/CRUD"
	_ "learngo/docs"
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
	crud.ShowdataCRUD()

}
