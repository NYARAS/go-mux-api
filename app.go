package main

import (
	"database/sql"

	// tom: for Initialize

	"log"

	// tom: for route handlers
	"net/http"

	// tom: go get required
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
