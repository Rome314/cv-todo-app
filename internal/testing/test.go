package test

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"cv-todo-app/cmd/connections"
)

func GetPool() *pgxpool.Pool {

	pgHost := "localhost"
	pgPort := "5432"
	pgUser := "postgres"
	pgPassword := "simplepassword"
	pgDb := "postgres"

	dns := fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s", pgHost, pgPort, pgDb, pgUser, pgPassword)

	if d := os.Getenv("PG_TEST_DNS"); d != "" {
		dns = d
	}

	pool, err := connections.GetPostgresDatabase(dns)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}
