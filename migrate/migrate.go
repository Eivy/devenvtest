package migrate

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	conString := createConnectionString()
	m, err := migrate.New("file://./migrations", conString)
	if err != nil {
		log.Print(err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s, ok := q["step"]
	if !ok {
		log.Print("check migration version")
		v, d, err := m.Version()
		if err != nil {
			log.Print(err)
			w.Write([]byte(fmt.Sprintf("%v", err)))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("version: %d\ndirtry: %v", v, d)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("migration start")
	step, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		log.Print(err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = m.Steps(step)
	if err != nil {
		log.Print(err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("migration done")
	w.Write([]byte("migration done"))
}

func createConnectionString() string {
	host := getConnectionString("PGHOST")
	port := getConnectionString("PGPORT")
	user := getConnectionString("PGUSER")
	password := getConnectionString("PGPASSWORD")
	db := getConnectionString("PGDATABASE")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, db)
}

func getConnectionString(env string) string {
	str := os.Getenv(env)
	if str == "" {
		panic(fmt.Sprintf("you should set postgresql database with %s", env))
	}
	return str
}
