package main

import (
	"www.github.com/gokhankocer/ticket-api/api"
	"www.github.com/gokhankocer/ticket-api/database"
	"www.github.com/gokhankocer/ticket-api/migrations"
)

func main() {

	database.ConnectPostgres()
	migrations.Migrate()
	api.Router()
}
