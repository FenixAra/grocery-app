package handlers

import (
	"net/http"

	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/services"
	"github.com/julienschmidt/httprouter"
)

func setClientRoutes(router *httprouter.Router) {
	router.POST("/accounts", SaveAccount)
	router.POST("/registers/status", OccupyRegister)
}

func OccupyRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	s := services.NewRegister(rd.l, rd.dbConn)
	req := &dtos.OccupyRegister{}
	err := LoadJson(r, req)
	if err != nil {
		writeJSONMessage("Unable to unmarshal json. Err:"+err.Error(), MSG, http.StatusBadRequest, rd)
	}

	err = s.OccupyRegister(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}

func SaveAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	a := services.NewAccount(rd.l, rd.dbConn)
	req := &dtos.Account{}
	err := LoadJson(r, req)
	if err != nil {
		writeJSONMessage("Unable to unmarshal json. Err:"+err.Error(), MSG, http.StatusBadRequest, rd)
	}

	err = a.Save(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}
