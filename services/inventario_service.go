package services

import (
	"sistema-gestion-ecommerce-golang/errores"
	"sistema-gestion-ecommerce-golang/models"
)

// ActualizarInventario descuenta del stock los productos agregados al carrito.
// Antes de modificar el inventario, valida que cada producto exista y tenga stock suficiente.
// Esta función evita que el sistema registre ventas de productos no disponibles.
func ActualizarInventario(productos []*models.Producto, items []models.ItemCarrito) error {
	for _, item := range items {
		productoEncontrado := false

		for _, producto := range productos {
			if producto.ID() == item.ProductoID() {
				productoEncontrado = true

				if producto.Stock() < item.Cantidad() {
					return errores.ErrStockInsuficiente
				}

				err := producto.ReducirStock(item.Cantidad())
				if err != nil {
					return err
				}
			}
		}

		if !productoEncontrado {
			return errores.ErrProductoNoEncontrado
		}
	}

	return nil
}
