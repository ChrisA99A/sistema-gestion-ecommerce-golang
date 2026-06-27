package handlers

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

type InventarioDTO struct {
	ID          int    `json:"id"`
	Producto    string `json:"producto"`
	Stock       int    `json:"stock"`
	EstadoStock string `json:"estado_stock"`
}

func ConsultarInventarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "metodo no permitido")
		return
	}

	inventario := []InventarioDTO{}

	for _, producto := range store.Productos {
		estado := "Disponible"

		if producto.Stock() == 0 {
			estado = "Agotado"
		} else if producto.Stock() <= 5 {
			estado = "Stock bajo"
		}

		inventario = append(inventario, InventarioDTO{
			ID:          producto.ID(),
			Producto:    producto.Nombre(),
			Stock:       producto.Stock(),
			EstadoStock: estado,
		})
	}

	web.ResponderJSON(w, http.StatusOK, inventario)
}
