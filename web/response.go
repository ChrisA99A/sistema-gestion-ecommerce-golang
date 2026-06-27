package web

import (
	"encoding/json"
	"net/http"
)

func ResponderJSON(w http.ResponseWriter, estado int, datos any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(estado)
	json.NewEncoder(w).Encode(datos)
}

func ResponderError(w http.ResponseWriter, estado int, mensaje string) {
	ResponderJSON(w, estado, map[string]string{
		"error": mensaje,
	})
}
