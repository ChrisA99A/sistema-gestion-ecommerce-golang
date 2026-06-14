package errores

import "errors"

var (
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrCarritoVacio         = errors.New("el carrito está vacío")
	ErrPagoInvalido         = errors.New("pago inválido")
	ErrMontoInvalido        = errors.New("el monto debe ser mayor a cero")
)
