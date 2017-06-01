package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/cli"
	"github.com/zero-os/0-directory/directory/api"
	"github.com/zero-os/0-directory/directory/db"
	"github.com/zero-os/0-directory/directory/goraml"
	"github.com/zero-os/0-directory/directory/routes"

	"gopkg.in/validator.v2"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "0-Directory"
	app.Version = version

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	// Set log output to stdout so we can pipe it
	log.SetOutput(os.Stdout)

	var debugLogging, ignoreDevcert bool
	var bindAddress, dbConnectionString string
	var tlsCert, tlsKey string

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
		cli.StringFlag{
			Name:        "bind, b",
			Usage:       "Bind address",
			Value:       ":8443",
			Destination: &bindAddress,
		},
		cli.StringFlag{
			Name:        "connectionstring, c",
			Usage:       "Mongodb connection string",
			Value:       "127.0.0.1:27017",
			Destination: &dbConnectionString,
		},
		cli.StringFlag{
			Name:        "cert, s",
			Usage:       "TLS certificate path",
			Value:       "",
			Destination: &tlsCert,
		},
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "TLS private key path",
			Value:       "",
			Destination: &tlsKey,
		},
		cli.BoolFlag{
			Name:        "ignore-devcert, i",
			Usage:       "Ignore default devcert even if exists",
			Destination: &ignoreDevcert,
		},
	}

	app.Before = func(c *cli.Context) error {
		if debugLogging {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug logging enabled")
		}

		// input validator
		validator.SetValidationFunc("multipleOf", goraml.MultipleOf)

		return nil
	}

	app.Action = func(c *cli.Context) {

		log.Infoln(app.Name, "version", app.Version)
		// Connect to DB!
		go db.Connect(dbConnectionString)
		defer db.Close()

		r := routes.GetRouter(api.ProvidersAPI{}, api.SearchAPI{})

		// Go make magic over HTTPS
		log.Info("Listening on ", bindAddress)
		log.Fatal(http.ListenAndServe(bindAddress, r))
	}

	app.Run(os.Args)

}
