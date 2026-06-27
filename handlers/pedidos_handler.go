package handlers

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/models"
	"sistema-gestion-ecommerce-golang/services"
	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

func GenerarPedidoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	subtotal := services.CalcularSubtotal(store.CarritoActual.Items())
	impuesto := services.CalcularImpuesto(subtotal, 0.15)
	total := services.CalcularTotal(subtotal, impuesto, 0)

	pedido, err := models.NewPedido(
		store.SiguientePedidoID(),
		store.CarritoActual.UsuarioID(),
		store.CarritoActual.Items(),
		subtotal,
		impuesto,
		total,
	)
	if err != nil {
		web.ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	store.Pedidos = append(store.Pedidos, pedido)

	// Después de generar el pedido, el carrito se vacía para evitar duplicar la misma compra.
	store.CarritoActual.Vaciar()

	web.ResponderJSON(w, http.StatusCreated, map[string]any{
		"mensaje": "pedido generado correctamente",
		"pedido":  web.PedidoToDTO(pedido),
	})
}

func ListarPedidosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	web.ResponderJSON(w, http.StatusOK, web.PedidosToDTO(store.Pedidos))
}
