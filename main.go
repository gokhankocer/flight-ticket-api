package main

import (
	"www.github.com/gokhankocer/ticket-api/api"
	"www.github.com/gokhankocer/ticket-api/database"
)

func main() {

	database.ConnectPostgres()
	api.Router()
}
