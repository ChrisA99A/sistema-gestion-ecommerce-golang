package handlers

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/services"
	"sistema-gestion-ecommerce-golang/store"
	"sistema-gestion-ecommerce-golang/web"
)

func ReporteVentasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		web.ResponderError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	ventasTotales := services.CalcularVentasTotales(store.Pedidos)

	web.ResponderJSON(w, http.StatusOK, map[string]any{
		"ventas_totales":      ventasTotales,
		"cantidad_pedidos":    len(store.Pedidos),
		"cantidad_pagos":      len(store.Pagos),
		"descripcion_reporte": "suma de pedidos con estado Pagado",
	})
}
