package main

import (
	"MytheresaChallenge/database/connections"
	"MytheresaChallenge/router"
	"os"
)

func main()  {
	connections.Connect()

	r := router.GetRouter()
	port := os.Args[1]
	r.Run("0.0.0.0:" + port)
}
