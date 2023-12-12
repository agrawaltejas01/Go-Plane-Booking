package main

import (
	"Plane-Booking/app/routes"
	"Plane-Booking/db"
	"fmt"
)

func connectDb() {
	db.Connect()
	db.Migrate()
}

func main() {

	connectDb()
	router := routes.Routes()

	const PORT = ":8000"

	router.Run(PORT)
	fmt.Printf("Server Running on Port: %s", PORT)
}
