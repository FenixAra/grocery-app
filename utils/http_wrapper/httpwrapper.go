package http_wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FenixAra/grocery-app/utils/log"
)

const (
	JSON_DATA = iota
	FORM_DATA
)

//go:generate mockgen -source=httpwrapper.go -destination=mock_httpwrapper.go -package=http_wrapper
type IHTTPWrapper interface {
	Init(l *log.Logger)
	MakeRequest(method string, pType int, u string, payload interface{}, auth string, pass string, v interface{}) (int, error)
}

type HTTPWrapper struct {
	l *log.Logger
}

func (h *HTTPWrapper) Init(l *log.Logger) {
	h.l = l
}

func (h *HTTPWrapper) GetRequest(url string, v interface{}) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, nil
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(content, v)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func (h *HTTPWrapper) MakeRequest(method string, pType int, u string, payload interface{}, auth string, pass string, v interface{}) (int, error) {
	if auth == "" && method == "GET" {
		return h.GetRequest(u, v)
	}
	client := new(http.Client)
	var p []byte
	var err error
	switch pType {
	case JSON_DATA:
		p, err = json.Marshal(payload)
		if err != nil {
			h.l.Error("Unable to marshal Request:", u, ", Err:", err)
			return 0, err
		}
	}

	req, err := http.NewRequest(method, u, bytes.NewBuffer(p))
	if err != nil {
		h.l.Error("Unable to create Request:", u, ", Err:", err)
		return 0, err
	}

	if auth != "" {
		req.SetBasicAuth(auth, pass)
	}

	response, err := client.Do(req)
	if err != nil {
		h.l.Error("Unable to make Request:", u, ", Err:", err)
		return 0, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return response.StatusCode, nil
	}

	if v == nil {
		return response.StatusCode, nil
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.l.Error("Unable to read response body, Err:", err)
		return 0, err
	}

	err = json.Unmarshal(content, &v)
	if err != nil {
		h.l.Error("Unable to unmarshal block response, Err:", err)
		return 0, err
	}

	return response.StatusCode, nil
}
