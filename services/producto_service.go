package services

import "sistema-gestion-ecommerce-golang/models"

func FiltrarProductosDisponibles(productos []*models.Producto) []*models.Producto {
	disponibles := []*models.Producto{}

	for _, producto := range productos {
		if producto.EstaDisponible() {
			disponibles = append(disponibles, producto)
		}
	}

	return disponibles
}
