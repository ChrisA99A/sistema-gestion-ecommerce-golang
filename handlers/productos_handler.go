package handlers

import (
	"net/http"
	"strconv"

	"sistema-gestion-ecommerce-golang/services"
	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

func ListarProductosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	web.ResponderJSON(w, http.StatusOK, web.ProductosToDTO(store.Productos))
}

func ListarProductosDisponiblesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	disponibles := services.FiltrarProductosDisponibles(store.Productos)
	web.ResponderJSON(w, http.StatusOK, web.ProductosToDTO(disponibles))
}

func BuscarProductoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	idTexto := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idTexto)
	if err != nil || id <= 0 {
		web.ResponderError(w, http.StatusBadRequest, "id de producto inválido")
		return
	}

	repositorio := services.NewProductoMemoriaRepository(store.Productos)
	producto, err := repositorio.BuscarPorID(id)
	if err != nil {
		web.ResponderError(w, http.StatusNotFound, err.Error())
		return
	}

	web.ResponderJSON(w, http.StatusOK, web.ProductoToDTO(producto))
}
