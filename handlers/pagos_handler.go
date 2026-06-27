package handlers

import (
	"encoding/json"
	"net/http"

	"sistema-gestion-ecommerce-golang/interfaces"
	"sistema-gestion-ecommerce-golang/models"
	"sistema-gestion-ecommerce-golang/services"
	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

type ProcesarPagoRequest struct {
	PedidoID int    `json:"pedido_id"`
	Metodo   string `json:"metodo"`
}

func ProcesarPagoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	var solicitud ProcesarPagoRequest
	if err := json.NewDecoder(r.Body).Decode(&solicitud); err != nil {
		web.ResponderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	pedido := buscarPedidoPorID(solicitud.PedidoID)
	if pedido == nil {
		web.ResponderError(w, http.StatusNotFound, "pedido no encontrado")
		return
	}

	if pedido.Estado() == "Pagado" {
		web.ResponderError(w, http.StatusConflict, "el pedido ya fue pagado")
		return
	}

	var procesadorPago interfaces.PagoProcessor
	procesadorPago = services.PagoSimuladoService{}

	pago, err := procesadorPago.Procesar(pedido.ID(), solicitud.Metodo, pedido.Total())
	if err != nil {
		web.ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := pedido.CambiarEstado("Pagado"); err != nil {
		web.ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.ActualizarInventario(store.Productos, pedido.Items()); err != nil {
		web.ResponderError(w, http.StatusBadRequest, err.Error())
		return
	}

	store.Pagos = append(store.Pagos, pago)

	web.ResponderJSON(w, http.StatusOK, map[string]any{
		"mensaje":       "pago procesado correctamente",
		"pago":          web.PagoToDTO(pago),
		"estado_pedido": pedido.Estado(),
	})
}

func buscarPedidoPorID(id int) *models.Pedido {
	for _, pedido := range store.Pedidos {
		if pedido.ID() == id {
			return pedido
		}
	}

	return nil
}
