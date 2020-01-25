package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize is to set the Env variable to connect to DB
func (a *App) Initialize(user, password, hostname, dbname string) {
	db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Error("Connection Failed to Open")
	}
	log.Info("Connection Established")

	db.CreateTable(&Comment{})
	// Also some useful functions
	db.HasTable(&Comment{})          // =>;; true
	db.DropTableIfExists(&Comment{}) //Drops the table if already exists
}

func (a *App) Run(addr string) {}

func connectionString() {

}
