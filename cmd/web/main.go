package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
	"syedzayyan.com/gosrvpi/pkg/models/sqlite3"
)

type application struct {
	globalPath   *string
	signUpSecret *string
	session      *sessions.Session
	users        *sqlite3.UserModel
	authState    bool
}

func main() {
	addr := flag.String("addr", ":5000", "HTTP network address")
	var globalPath = flag.String("hostDir", "./files", "Directory to files")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	signUpSecret := flag.String("signUpSecret", "1234", "SignupPath")
	flag.Parse()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	db, errr := openDB()
	if errr != nil {
		log.Fatal(errr)
	}
	defer db.Close()
	app := &application{
		globalPath:   globalPath,
		signUpSecret: signUpSecret,
		session:      session,
		users:        &sqlite3.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
