package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (resp *response) bytes() []byte {
	data, _ := json.Marshal(resp)
	return data
}

func (resp *response) string() string {
	return string(resp.bytes())
}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)
	_, _ = w.Write(resp.bytes())
	log.Println(resp.string())
}

func StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) { //200
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

func StatusNoContent(w http.ResponseWriter, r *http.Request) { //204
	newResponse(nil, http.StatusNoContent).sendResponse(w, r)
}

func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) { //400
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) { //404
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) { //405
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
}

func StatusConflict(w http.ResponseWriter, r *http.Request, err error) { //409
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) { //500
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
