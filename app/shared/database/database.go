package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

type PgInfo struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	Params   string
}

// Creating an postgresql URL
func GetPgUrl(pi PgInfo) string {
	return "postgresql://" +
		pi.User +
		":" +
		pi.Password +
		"@" +
		pi.Host +
		":" +
		fmt.Sprintf("%d", pi.Port) +
		"/" +
		pi.Name +
		pi.Params
}

// Connecting to the database
func Connect(pi PgInfo) {
	var err error
	if Db, err = sqlx.Connect("postgres", GetPgUrl(pi)); err != nil {
		log.Println("SQL Driver Error", err)
	}

	// Check if is alive
	if err = Db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	fmt.Println("Connected to Database")
}
