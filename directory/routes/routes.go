package routes

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/zero-os/0-directory/directory/api"
	"github.com/zero-os/0-directory/directory/db"
	"github.com/zero-os/0-directory/directory/db/pool"
	"github.com/zero-os/0-directory/directory/db/provider"
)

func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(log.StandardLogger().Out, h)
}

func GetRouter(p api.ProvidersInterface, s api.SearchInterface) http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	api.ProvidersInterfaceRoutes(r, p)
	api.SearchInterfaceRoutes(r, s)

	pool.InitModels()
	provider.InitModels()

	// home page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// apidocs
	r.PathPrefix("/apidocs/").Handler(http.StripPrefix("/apidocs/", http.FileServer(http.Dir("./apidocs/"))))

	// Add middlewares
	router := NewRouter(r)

	dbmw := db.DBMiddleware()
	recovery := handlers.RecoveryHandler()

	router.Use(recovery, loggingMiddleware, dbmw)

	return router.Handler()
}
