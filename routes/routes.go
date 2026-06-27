package routes

import (
	"net/http"

	"sistema-gestion-ecommerce-golang/handlers"
)

func RegistrarRutas() {
	http.HandleFunc("/api/salud", handlers.SaludHandler)
	http.HandleFunc("/api/usuarios", handlers.ListarUsuariosHandler)
	http.HandleFunc("/api/productos", handlers.ListarProductosHandler)
	http.HandleFunc("/api/productos/disponibles", handlers.ListarProductosDisponiblesHandler)
	http.HandleFunc("/api/producto", handlers.BuscarProductoHandler)
	http.HandleFunc("/api/inventario", handlers.ConsultarInventarioHandler)
	http.HandleFunc("/api/carrito", handlers.VerCarritoHandler)
	http.HandleFunc("/api/carrito/agregar", handlers.AgregarProductoCarritoHandler)
	http.HandleFunc("/api/pedidos", handlers.ListarPedidosHandler)
	http.HandleFunc("/api/pedidos/generar", handlers.GenerarPedidoHandler)
	http.HandleFunc("/api/pagos/procesar", handlers.ProcesarPagoHandler)
	http.HandleFunc("/api/reportes/ventas", handlers.ReporteVentasHandler)
}
