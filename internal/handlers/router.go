package handlers

import (
	"net/http"

	"runtime/debug"

	"github.com/FenixAra/grocery-app/internal/config"
	"github.com/FenixAra/grocery-app/internal/config/globals"
	"github.com/FenixAra/grocery-app/utils/log"
	"github.com/julienschmidt/httprouter"
)

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	globals.Logger = log.NewLogger("")

	router := httprouter.New()
	router.PanicHandler = PanicHandler
	setPingRoutes(router)

	return router
}

func tokenAuthentication(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userName, _, ok := r.BasicAuth()
		if (ok && userName == config.SERVICE_TOKEN) || r.FormValue("token") == config.SERVICE_TOKEN {
			h(w, r, ps)
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	globals.Logger.Fatal("Recovering from panic, Reason: ", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
