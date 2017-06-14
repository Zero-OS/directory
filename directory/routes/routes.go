package routes

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	assetfs "github.com/elazarl/go-bindata-assetfs"
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
	r := mux.NewRouter()

	api.ProvidersInterfaceRoutes(r, p)
	api.SearchInterfaceRoutes(r, s)

	pool.InitModels()
	provider.InitModels()

	// home page and apiconsole
	r.PathPrefix("/").Handler(http.FileServer(
		&assetfs.AssetFS{Asset: Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "",
		}))

	// Add middlewares
	router := NewRouter(r)

	dbmw := db.DBMiddleware()
	recovery := handlers.RecoveryHandler()

	router.Use(recovery, loggingMiddleware, dbmw)

	return router.Handler()
}
