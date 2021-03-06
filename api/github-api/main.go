package main

import (
	"github.com/Netflix/go-env"
	log "github.com/sirupsen/logrus"
)

/*
Environment Variable PostgresUser
*/

type Environment struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabaseName     string `env:"POSTGRES_DB"`
	Prod             bool   `env:"PROD"`
	AccessToken      string `env:"GITHUB_ACCESS_TOKEN"`
}

// This is where I initialize the application
func main() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {

		log.Error("error in unmarshel ENV", err)
	}

	//FOR DEV only TODO: put conditoin if nil or not equal to PROD
	if environment.Prod == false || environment.Prod != true {
		log.Info("ENV credentials: ", environment)
	}
	a := App{}

	a.Initialize(
		environment.PostgresUser,
		environment.PostgresPassword,
		environment.DatabaseHost,
		environment.DatabaseName)

	IntializeModel(a.DB)
	a.Run(":8080")
}
