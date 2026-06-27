package services

import "sistema-gestion-ecommerce-golang/models"

// CalcularVentasTotales suma el valor de los pedidos que se encuentran pagados.
// Esta funciÃ³n permite generar un reporte bÃ¡sico de ventas dentro del sistema.
func CalcularVentasTotales(pedidos []*models.Pedido) float64 {
	var total float64

	for _, pedido := range pedidos {
		if pedido.Estado() == "Pagado" {
			total += pedido.Total()
		}
	}

	return total
}
