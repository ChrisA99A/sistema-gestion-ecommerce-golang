package interfaces

import "sistema-gestion-ecommerce-golang/models"

type ProductoRepository interface {
	Listar() []*models.Producto
	BuscarPorID(id int) (*models.Producto, error)
}

type PagoProcessor interface {
	Procesar(pedidoID int, metodo string, monto float64) (*models.Pago, error)
}
