package connect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Postgres(opDocker bool) error {
	if opDocker != false {
		fmt.Print("docker ativado!")

		os.Exit(0)
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	fmt.Print("Connect to PostgreSQL!")

	defer db.Close()

	return nil
}
