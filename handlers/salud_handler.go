package handlers

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/web"
)

func SaludHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	web.ResponderJSON(w, http.StatusOK, map[string]string{
		"estado":  "activo",
		"sistema": "Sistema de Gestion de E-commerce en Golang",
	})
}
