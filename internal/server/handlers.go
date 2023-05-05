package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bradford-hamilton/bt-data-server/internal/storage"
	"github.com/go-chi/render"
)

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("pong"))
}

type createDataDumpReq struct {
	Sensor     string `json:"sensor,omitempty"`
	DataValues string `json:"data_values,omitempty"`
}

func (a *API) createDataDump(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req createDataDumpReq
	if err := json.Unmarshal(b, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dd := storage.BTDataDump{
		Sensor:     req.Sensor,
		DataValues: req.DataValues,
	}

	if err := a.db.CreateDataDump(dd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, "{ \"status\": \"success\" }")
}
