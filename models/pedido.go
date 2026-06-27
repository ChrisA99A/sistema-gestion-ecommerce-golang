package models

import "errors"

type Pedido struct {
	id        int
	usuarioID int
	items     []ItemCarrito
	subtotal  float64
	impuesto  float64
	total     float64
	estado    string
}

func NewPedido(id int, usuarioID int, items []ItemCarrito, subtotal float64, impuesto float64, total float64) (*Pedido, error) {
	if len(items) == 0 {
		return nil, errors.New("el carrito está vacío")
	}

	if total <= 0 {
		return nil, errors.New("el total del pedido debe ser mayor a cero")
	}

	return &Pedido{
		id:        id,
		usuarioID: usuarioID,
		items:     items,
		subtotal:  subtotal,
		impuesto:  impuesto,
		total:     total,
		estado:    "Pendiente",
	}, nil
}

func (p Pedido) ID() int {
	return p.id
}

func (p Pedido) UsuarioID() int {
	return p.usuarioID
}

func (p Pedido) Items() []ItemCarrito {
	return p.items
}

func (p Pedido) Subtotal() float64 {
	return p.subtotal
}

func (p Pedido) Impuesto() float64 {
	return p.impuesto
}

func (p Pedido) Total() float64 {
	return p.total
}

func (p Pedido) Estado() string {
	return p.estado
}

func (p *Pedido) CambiarEstado(estado string) error {
	if estado == "" {
		return errors.New("el estado del pedido no puede estar vacío")
	}

	p.estado = estado
	return nil
}
