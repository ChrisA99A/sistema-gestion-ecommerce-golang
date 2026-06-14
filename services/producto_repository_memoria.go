package services

import (
	"sistema-gestion-ecommerce-golang/errores"
	"sistema-gestion-ecommerce-golang/models"
)

type ProductoMemoriaRepository struct {
	productos []*models.Producto
}

func NewProductoMemoriaRepository(productos []*models.Producto) *ProductoMemoriaRepository {
	return &ProductoMemoriaRepository{
		productos: productos,
	}
}

func (r *ProductoMemoriaRepository) Listar() []*models.Producto {
	return r.productos
}

func (r *ProductoMemoriaRepository) BuscarPorID(id int) (*models.Producto, error) {
	for _, producto := range r.productos {
		if producto.ID() == id {
			return producto, nil
		}
	}

	return nil, errores.ErrProductoNoEncontrado
}
