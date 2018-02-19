package handlers

import (
	"net/http"

	"github.com/FenixAra/grocery-app/internal/services"
	"github.com/julienschmidt/httprouter"
)

func setPingRoutes(router *httprouter.Router) {
	router.GET("/ping", Ping)
}

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := services.NewPing(rd.l, rd.dbConn)

	err := p.Ping()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("pong", MSG, http.StatusOK, rd)
}
