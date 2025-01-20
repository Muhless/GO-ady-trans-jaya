package main

import (
	"ady-trans-jaya/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
