package handlers

import (
	"net/http"

	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/services"
	"github.com/julienschmidt/httprouter"
)

func setAdminRoutes(router *httprouter.Router) {
	router.POST("/admin/categories", SaveCategory)
	router.POST("/admin/items", SaveItem)
	router.POST("/admin/discounts", SaveDiscount)
}

func SaveCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := services.NewCategory(rd.l, rd.dbConn)
	req := &dtos.Category{}
	err := LoadJson(r, req)
	if err != nil {
		writeJSONMessage("Unable to unmarshal json. Err:"+err.Error(), MSG, http.StatusBadRequest, rd)
	}

	err = p.Save(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}

func SaveItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	i := services.NewItem(rd.l, rd.dbConn)
	req := &dtos.Item{}
	err := LoadJson(r, req)
	if err != nil {
		writeJSONMessage("Unable to unmarshal json. Err:"+err.Error(), MSG, http.StatusBadRequest, rd)
	}

	err = i.Save(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}

func SaveDiscount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	d := services.NewDiscount(rd.l, rd.dbConn)
	req := &dtos.Discount{}
	err := LoadJson(r, req)
	if err != nil {
		writeJSONMessage("Unable to unmarshal json. Err:"+err.Error(), MSG, http.StatusBadRequest, rd)
	}

	err = d.Save(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}