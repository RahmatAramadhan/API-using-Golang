package main

import(
	"First-API-Golang/database"
	"First-API-Golang/models"
	"First-API-Golang/routes"

)

func main(){
	database.Connect()
	database.Migrate(&models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}