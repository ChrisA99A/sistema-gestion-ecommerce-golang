package services

import (
	"sistema-gestion-ecommerce-golang/errores"
	"sistema-gestion-ecommerce-golang/models"
)

type PagoSimuladoService struct{}

// Procesar valida los datos del pago y genera un pago simulado.
// No se conecta a bancos reales, porque esta primera versiÃ³n solo representa el flujo bÃ¡sico de compra.
func (p PagoSimuladoService) Procesar(pedidoID int, metodo string, monto float64) (*models.Pago, error) {
	if metodo == "" {
		return nil, errores.ErrPagoInvalido
	}

	if monto <= 0 {
		return nil, errores.ErrMontoInvalido
	}

	return models.NewPagoSimulado(pedidoID, metodo, monto)
}
