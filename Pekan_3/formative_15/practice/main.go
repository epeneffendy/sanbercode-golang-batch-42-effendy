package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "kucing"
	dbname   = "practicex"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=#{os.{"DB_HOST"}} port=#{os.{"DB_PORT"}} user=#{os.{"DB_USER"}} password=#{os.{"DB_PASSWORD"}} dbname=#{os.{"DB_NAME"}} sslmode=disable")

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	defer DB.Close()

	router := gin.Default()
	router.GET("/person", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
