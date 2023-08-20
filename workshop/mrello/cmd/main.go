package main

import (
	"database/sql"
	"fmt"
	"log"
	"mrello"
	"mrello/pgrepository"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	pgConnInfo := fmt.Sprintf("host='%s' port='%d' user='%s' password='%s' dbname='%s' sslmode=disable", os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	fmt.Println(pgConnInfo)
	db, err := sql.Open("postgres", pgConnInfo)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	defer db.Close()

	userRepo := pgrepository.NewUserRepository(db)
	handler := mrello.NewHandler(mrello.WithUserRepository(userRepo))

	e := gin.Default()
	mrello.Router(e, handler)

	e.Run(":8080")
}
