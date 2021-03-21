package graph

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/eivy/aptitude_bulb/db"
)

var con *db.Queries

func New() (err error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	d, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password))
	if err != nil {
		return
	}
	con = db.New(d)
	return nil
}
