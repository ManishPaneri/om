package handlers

import (
	"encoding/json"
	"net/http"
	"om/controllers"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	data := controllers.MapProductUrl(w, r)
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
