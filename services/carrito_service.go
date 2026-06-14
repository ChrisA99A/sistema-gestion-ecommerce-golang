package services

import "sistema-gestion-ecommerce-golang/models"

func CalcularSubtotal(items []models.ItemCarrito) float64 {
	var subtotal float64

	for _, item := range items {
		subtotal += item.Total()
	}

	return subtotal
}

func CalcularImpuesto(subtotal float64, porcentaje float64) float64 {
	return subtotal * porcentaje
}

func CalcularTotal(subtotal float64, impuesto float64, descuento float64) float64 {
	return subtotal + impuesto - descuento
}
