package main

import (
	"ady-trans-jaya/config"
	"ady-trans-jaya/routes"
)

func main() {

	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
