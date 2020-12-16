package main

import (
	"rest_api/controllers"
	"rest_api/database"
)

func main() {
	db,err :=database.ConnectToDataBase()
	if err != nil {
		panic(err)
	}
	if err = controllers.InitialApiServer(db);err != nil{
		panic(err)
	}
}
