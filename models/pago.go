package models

import "errors"

type Pago struct {
	pedidoID int
	metodo   string
	monto    float64
	estado   string
}

func NewPagoSimulado(pedidoID int, metodo string, monto float64) (*Pago, error) {
	if metodo == "" {
		return nil, errors.New("el método de pago no puede estar vacío")
	}

	if monto <= 0 {
		return nil, errors.New("el monto del pago debe ser mayor a cero")
	}

	return &Pago{
		pedidoID: pedidoID,
		metodo:   metodo,
		monto:    monto,
		estado:   "Aprobado",
	}, nil
}

func (p Pago) Estado() string {
	return p.estado
}
