package database

import (
	"database/sql"
	"fmt"
)

var (
	DBConnection *sql.DB
)

func DBMigrate(dbParam *sql.DB) {
	migrations := &dbmigrate.PackrMigrationSourse{
		Box.packr.New("MIgration", "./sql_migrations"),
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	DBConnection = dbParam
	fmt.Println("Applied", n, "migration!")
}
