package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	errr := godotenv.Load("/home/andreas.timothy/go/src/demo-go-basic-backend/.env")
	if errr != nil {
		log.Fatalln(errr)
	}

	info := fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = sql.Open("postgres", info)
	if err != nil {
		fmt.Println(err)
	}
}

func Get() *sql.DB {
	return db
}