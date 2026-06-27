package handlers

import (
	"encoding/json"
	"net/http"

	"sistema-gestion-ecommerce-golang/models"
	"sistema-gestion-ecommerce-golang/services"
	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

type AgregarCarritoRequest struct {
	ProductoID int `json:"producto_id"`
	Cantidad   int `json:"cantidad"`
}

func VerCarritoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	web.ResponderJSON(w, http.StatusOK, web.CarritoToDTO(store.CarritoActual))
}

func AgregarProductoCarritoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	var solicitud AgregarCarritoRequest
	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		web.ResponderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	repositorio := services.NewProductoMemoriaRepository(store.Productos)
	producto, err := repositorio.BuscarPorID(solicitud.ProductoID)
	if err != nil {
		web.ResponderError(w, http.StatusNotFound, err.Error())
		return
	}

	item, err := models.NewItemCarrito(*producto, solicitud.Cantidad)
	if err != nil {
		web.ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	store.CarritoActual.AgregarItem(*item)

	web.ResponderJSON(w, http.StatusCreated, map[string]any{
		"mensaje": "producto agregado al carrito correctamente",
		"carrito": web.CarritoToDTO(store.CarritoActual),
	})
}
