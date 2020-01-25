package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	TC     *oauth2.Client
}

//Initialize is to set the Env variable to connect to DB
func (a *App) Initialize(user, password, hostname, dbname string) {

	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s  port=5432 dbname=%s sslmode=disable", user, password, hostname, dbname))
	defer db.Close()
	if err != nil {
		log.Error("Connection Failed to Open")
		panic(err)
	}
	a.DB = db
	log.Info("Connection Established")

	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {}

func connectionString() {

}
