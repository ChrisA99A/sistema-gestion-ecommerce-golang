package models

import "errors"

type Producto struct {
	id        int
	nombre    string
	categoria string
	precio    float64
	stock     int
	activo    bool
}

func NewProducto(id int, nombre string, categoria string, precio float64, stock int) (*Producto, error) {
	if nombre == "" {
		return nil, errors.New("el nombre del producto no puede estar vacío")
	}

	if precio <= 0 {
		return nil, errors.New("el precio del producto debe ser mayor a cero")
	}

	if stock < 0 {
		return nil, errors.New("el stock no puede ser negativo")
	}

	return &Producto{
		id:        id,
		nombre:    nombre,
		categoria: categoria,
		precio:    precio,
		stock:     stock,
		activo:    true,
	}, nil
}

func (p Producto) ID() int {
	return p.id
}

func (p Producto) Nombre() string {
	return p.nombre
}

func (p Producto) Precio() float64 {
	return p.precio
}

func (p Producto) Stock() int {
	return p.stock
}

func (p Producto) EstaDisponible() bool {
	return p.activo && p.stock > 0
}

func (p *Producto) ReducirStock(cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	if cantidad > p.stock {
		return errors.New("stock insuficiente")
	}

	p.stock -= cantidad
	return nil
}
