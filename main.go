package main

import (
	"MicroserviceGo/internal/database"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
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
	//db, err := database.NewDatabaseClient(dbSettings)
	//if err != nil {
	//	log.Fatalf("Failed to initialize DB client: %s", err)
	//}
	//serv := server.NewEchoServer(db)
	//
	//if err := serv.Start(); err != nil {
	//	log.Fatalf(err.Error())
	//}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	rdsService := rdsdata.NewFromConfig(cfg)
	resp, err := rdsService.ExecuteStatement(context.TODO(), &rdsdata.ExecuteStatementInput{
		ResourceArn:           nil,
		SecretArn:             nil,
		Sql:                   nil,
		ContinueAfterTimeout:  false,
		Database:              nil,
		FormatRecordsAs:       "",
		IncludeResultMetadata: false,
		Parameters:            nil,
		ResultSetOptions:      nil,
		Schema:                nil,
		TransactionId:         nil,
	})

}
