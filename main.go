package main

import (
	"MicroserviceGo/internal/database"
	"MicroserviceGo/internal/server"
	"log"
)

func main() {
	dbSettings := database.DBConnSettings{
		Hostname: "go-db-server.cluster-ro-c9sc9fsnusbu.us-east-1.rds.amazonaws.com",
		Username: "postgres",
		Password: "",
		Dbname:   "postgres",
		Port:     5432,
		Ssl:      "disable",
	}
	db, err := database.NewDatabaseClient(dbSettings)
	if err != nil {
		log.Fatalf("Failed to initialize DB client: %s", err)
	}
	serv := server.NewEchoServer(db)

	if err := serv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
