package web

import (
	"sistema-gestion-ecommerce-golang/models"
	"sistema-gestion-ecommerce-golang/services"
)

type UsuarioDTO struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"`
}

type ProductoDTO struct {
	ID         int     `json:"id"`
	Nombre     string  `json:"nombre"`
	Categoria  string  `json:"categoria"`
	Precio     float64 `json:"precio"`
	Stock      int     `json:"stock"`
	Disponible bool    `json:"disponible"`
}

type ItemCarritoDTO struct {
	ProductoID int     `json:"producto_id"`
	Nombre     string  `json:"nombre"`
	Precio     float64 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
	Total      float64 `json:"total"`
}

type CarritoDTO struct {
	UsuarioID int              `json:"usuario_id"`
	Items     []ItemCarritoDTO `json:"items"`
	Subtotal  float64          `json:"subtotal"`
	Impuesto  float64          `json:"impuesto"`
	Total     float64          `json:"total"`
}

type PedidoDTO struct {
	ID        int              `json:"id"`
	UsuarioID int              `json:"usuario_id"`
	Items     []ItemCarritoDTO `json:"items"`
	Subtotal  float64          `json:"subtotal"`
	Impuesto  float64          `json:"impuesto"`
	Total     float64          `json:"total"`
	Estado    string           `json:"estado"`
}

type PagoDTO struct {
	PedidoID int     `json:"pedido_id"`
	Metodo   string  `json:"metodo"`
	Monto    float64 `json:"monto"`
	Estado   string  `json:"estado"`
}

func UsuarioToDTO(usuario *models.Usuario) UsuarioDTO {
	return UsuarioDTO{
		ID:     usuario.ID(),
		Nombre: usuario.Nombre(),
		Correo: usuario.Correo(),
		Rol:    usuario.Rol(),
	}
}

func ProductoToDTO(producto *models.Producto) ProductoDTO {
	return ProductoDTO{
		ID:         producto.ID(),
		Nombre:     producto.Nombre(),
		Categoria:  producto.Categoria(),
		Precio:     producto.Precio(),
		Stock:      producto.Stock(),
		Disponible: producto.EstaDisponible(),
	}
}

func ProductosToDTO(productos []*models.Producto) []ProductoDTO {
	resultado := []ProductoDTO{}

	for _, producto := range productos {
		resultado = append(resultado, ProductoToDTO(producto))
	}

	return resultado
}

func ItemCarritoToDTO(item models.ItemCarrito) ItemCarritoDTO {
	return ItemCarritoDTO{
		ProductoID: item.ProductoID(),
		Nombre:     item.Nombre(),
		Precio:     item.Precio(),
		Cantidad:   item.Cantidad(),
		Total:      item.Total(),
	}
}

func ItemsCarritoToDTO(items []models.ItemCarrito) []ItemCarritoDTO {
	resultado := []ItemCarritoDTO{}

	for _, item := range items {
		resultado = append(resultado, ItemCarritoToDTO(item))
	}

	return resultado
}

func CarritoToDTO(carrito *models.Carrito) CarritoDTO {
	subtotal := services.CalcularSubtotal(carrito.Items())
	impuesto := services.CalcularImpuesto(subtotal, 0.15)
	total := services.CalcularTotal(subtotal, impuesto, 0)

	return CarritoDTO{
		UsuarioID: carrito.UsuarioID(),
		Items:     ItemsCarritoToDTO(carrito.Items()),
		Subtotal:  subtotal,
		Impuesto:  impuesto,
		Total:     total,
	}
}

func PedidoToDTO(pedido *models.Pedido) PedidoDTO {
	return PedidoDTO{
		ID:        pedido.ID(),
		UsuarioID: pedido.UsuarioID(),
		Items:     ItemsCarritoToDTO(pedido.Items()),
		Subtotal:  pedido.Subtotal(),
		Impuesto:  pedido.Impuesto(),
		Total:     pedido.Total(),
		Estado:    pedido.Estado(),
	}
}

func PedidosToDTO(pedidos []*models.Pedido) []PedidoDTO {
	resultado := []PedidoDTO{}

	for _, pedido := range pedidos {
		resultado = append(resultado, PedidoToDTO(pedido))
	}

	return resultado
}

func PagoToDTO(pago *models.Pago) PagoDTO {
	return PagoDTO{
		PedidoID: pago.PedidoID(),
		Metodo:   pago.Metodo(),
		Monto:    pago.Monto(),
		Estado:   pago.Estado(),
	}
}
