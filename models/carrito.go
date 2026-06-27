package models

import "errors"

type ItemCarrito struct {
	productoID int
	nombre     string
	precio     float64
	cantidad   int
}

func NewItemCarrito(producto Producto, cantidad int) (*ItemCarrito, error) {
	if cantidad <= 0 {
		return nil, errors.New("la cantidad debe ser mayor a cero")
	}

	if cantidad > producto.Stock() {
		return nil, errors.New("stock insuficiente")
	}

	return &ItemCarrito{
		productoID: producto.ID(),
		nombre:     producto.Nombre(),
		precio:     producto.Precio(),
		cantidad:   cantidad,
	}, nil
}

func (i ItemCarrito) ProductoID() int {
	return i.productoID
}

func (i ItemCarrito) Nombre() string {
	return i.nombre
}

func (i ItemCarrito) Precio() float64 {
	return i.precio
}

func (i ItemCarrito) Cantidad() int {
	return i.cantidad
}

func (i ItemCarrito) Total() float64 {
	return i.precio * float64(i.cantidad)
}

type Carrito struct {
	usuarioID int
	items     []ItemCarrito
}

func NewCarrito(usuarioID int) *Carrito {
	return &Carrito{
		usuarioID: usuarioID,
		items:     []ItemCarrito{},
	}
}

func (c Carrito) UsuarioID() int {
	return c.usuarioID
}

func (c *Carrito) AgregarItem(item ItemCarrito) {
	c.items = append(c.items, item)
}

func (c Carrito) Items() []ItemCarrito {
	return c.items
}

func (c *Carrito) Vaciar() {
	c.items = []ItemCarrito{}
}
