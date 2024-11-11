package main

import (
	"swagger/database"
	"swagger/routers"
)

func main() {
	database.InitDB()
	routers.StartRouter().Run(":8080")
}
