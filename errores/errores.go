package errores

import "errors"

var (
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrCarritoVacio         = errors.New("el carrito estÃ¡ vacÃ­o")
	ErrPagoInvalido         = errors.New("pago invÃ¡lido")
	ErrMontoInvalido        = errors.New("el monto debe ser mayor a cero")
)
