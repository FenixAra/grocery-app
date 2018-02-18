package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/utils/log"
)

const (
	ERR_MSG = "ERROR_MESSAGE"
	MSG     = "MESSAGE"
)

type RequestData struct {
	l      *log.Logger
	dbConn *db.DBConn
	s      time.Time
	w      http.ResponseWriter
	r      *http.Request
}

func logAndGetContext(w http.ResponseWriter, r *http.Request) *RequestData {
	var ref string
	w.Header().Add("X-Content-Type-Options", "nosniff")

	l := log.NewLogger(ref)
	dbConn := db.NewDBConn(l)
	l.Info("Serving Request: ", r.RequestURI, ", Method: ", r.Method)
	s := time.Now()

	return &RequestData{
		l:      l,
		dbConn: dbConn,
		s:      s,
		r:      r,
		w:      w,
	}
}

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case ERR_MSG:
		Obj.Status = "FAILED"

	case MSG:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONMessage(msg string, msgType string, httpCode int, rd *RequestData) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponse(d, code, rd)
}

func writeJSONStruct(v interface{}, code int, rd *RequestData) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessage("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONResponse(d, code, rd)
}

func writeJSONResponse(d []byte, code int, rd *RequestData) {
	if code == http.StatusInternalServerError {
		rd.l.Info("Status Code: ", code, ", Response time: ", time.Since(rd.s), " Response: ", string(d))
	} else {
		rd.l.Info("Status Code: ", code, ", Response time: ", time.Since(rd.s))
	}

	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
}
