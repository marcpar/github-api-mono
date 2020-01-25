package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	//TC     *oauth2.Client
}

//Initialize is to set the Env variable to connect to DB
func (a *App) Initialize(user, password, hostname, dbname string) {
	//Generates a connection string to DB
	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s  port=5432 dbname=%s sslmode=disable", user, password, hostname, dbname))
	defer db.Close()
	if err != nil {
		log.Error("Connection Failed to Open")
		panic(err)
	}
	a.DB = db
	log.Info("Connection Established")

	//
	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {}

// TEST: 2. Persisting Github comments against a given Github organization
// TEST: 3. Returning an array of all the comments that have been registered against a Github organization
// TEST: 4. Soft-deleting all comments associated with a particular organization. We define a "soft delete" to mean that deleted items should not be returned in GET calls, but should remain in the database for emergency retrieval and audit purposes.
// TEST: 5. Returning an array of members of an organization (with their login, avatar URL, the numbers of followers they have, and the number of people they're following), sorted in descending order by the number of followers.

func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/organization", a.createOrg).Methods("POST")
	// a.Router.HandleFunc("/org/{id:[0-9]+}/comments", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/org/{id:[0-9]+}/comments", a.writeCommentOrg).Methods("PUT")
	// a.Router.HandleFunc("/org/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}
