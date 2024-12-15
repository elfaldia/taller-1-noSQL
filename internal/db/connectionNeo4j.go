package db

import (
	"context"

	"github.com/elfaldia/taller-noSQL/internal/env"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectNeo4jDB() *neo4j.DriverWithContext {
	ctx := context.Background()
	
	dbUri := env.GetString("NEO4J_URI", "")
	dbUser := env.GetString("NEO4J_USERNAME", "")
	dbPassword := env.GetString("NEO4J_PASSWORD", "")


	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""),	
	)

	if err != nil {
		panic(err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	return &driver
}