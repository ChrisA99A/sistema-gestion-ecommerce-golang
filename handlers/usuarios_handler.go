package handlers

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

func ListarUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	usuarios := []web.UsuarioDTO{web.UsuarioToDTO(store.UsuarioPrincipal)}
	web.ResponderJSON(w, http.StatusOK, usuarios)
}
