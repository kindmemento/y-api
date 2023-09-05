package main

import "y-api/routes"

func main() {
	r := routes.SetupRouter()

	r.Run(":8080")
}
