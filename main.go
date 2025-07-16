package main

import (
	"log"
	"task_manager/db"
	"task_manager/router"
)

func init() {
	db.ConnectToDB()
}

func main() {
	err := router.Run()
	if err != nil {
		log.Printf("Failed to start the server: %s", err.Error())
	}
}
