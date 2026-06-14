package main

import (
	"fmt"

	"sistema-gestion-ecommerce-golang/interfaces"
	"sistema-gestion-ecommerce-golang/models"
	"sistema-gestion-ecommerce-golang/services"
)

func main() {
	fmt.Println("Sistema de Gestión de E-commerce en Golang")
	fmt.Println("------------------------------------------")

	usuario, err := models.NewUsuario(1, "Christian Cevallos", "christian@email.com", "cliente")
	if err != nil {
		fmt.Println("Error al crear usuario:", err)
		return
	}

	laptop, err := models.NewProducto(1, "Laptop", "Tecnología", 850.00, 5)
	if err != nil {
		fmt.Println("Error al crear laptop:", err)
		return
	}

	mouse, err := models.NewProducto(2, "Mouse", "Accesorios", 15.50, 20)
	if err != nil {
		fmt.Println("Error al crear mouse:", err)
		return
	}

	productos := []*models.Producto{laptop, mouse}

	var productoRepo interfaces.ProductoRepository
	productoRepo = services.NewProductoMemoriaRepository(productos)

	disponibles := services.FiltrarProductosDisponibles(productoRepo.Listar())

	fmt.Println("Productos disponibles:")
	for _, producto := range disponibles {
		fmt.Printf("%d. %s - $%.2f - Stock: %d\n",
			producto.ID(),
			producto.Nombre(),
			producto.Precio(),
			producto.Stock(),
		)
	}

	productoBuscado, err := productoRepo.BuscarPorID(1)
	if err != nil {
		fmt.Println("Error al buscar producto:", err)
		return
	}

	fmt.Println()
	fmt.Println("Producto consultado mediante interfaz:")
	fmt.Println(productoBuscado.Nombre())

	carrito := models.NewCarrito(usuario.ID())

	itemLaptop, err := models.NewItemCarrito(*laptop, 1)
	if err != nil {
		fmt.Println("Error al agregar laptop:", err)
		return
	}

	itemMouse, err := models.NewItemCarrito(*mouse, 2)
	if err != nil {
		fmt.Println("Error al agregar mouse:", err)
		return
	}

	carrito.AgregarItem(*itemLaptop)
	carrito.AgregarItem(*itemMouse)

	subtotal := services.CalcularSubtotal(carrito.Items())
	impuesto := services.CalcularImpuesto(subtotal, 0.15)
	total := services.CalcularTotal(subtotal, impuesto, 0)

	pedido, err := models.NewPedido(1, usuario.ID(), carrito.Items(), subtotal, impuesto, total)
	if err != nil {
		fmt.Println("Error al generar pedido:", err)
		return
	}

	var procesadorPago interfaces.PagoProcessor
	procesadorPago = services.PagoSimuladoService{}

	pago, err := procesadorPago.Procesar(1, "Tarjeta simulada", pedido.Total())
	if err != nil {
		fmt.Println("Error al procesar pago:", err)
		return
	}

	err = pedido.CambiarEstado("Pagado")
	if err != nil {
		fmt.Println("Error al cambiar estado del pedido:", err)
		return
	}

	err = services.ActualizarInventario(productos, carrito.Items())
	if err != nil {
		fmt.Println("Error al actualizar inventario:", err)
		return
	}

	ventasTotales := services.CalcularVentasTotales([]*models.Pedido{pedido})

	fmt.Println()
	fmt.Println("Detalle del carrito:")
	for _, item := range carrito.Items() {
		fmt.Printf("- %s x%d - Precio unitario: $%.2f - Total: $%.2f\n",
			item.Nombre(),
			item.Cantidad(),
			item.Precio(),
			item.Total(),
		)
	}
	fmt.Println()
	fmt.Println("Resumen de compra:")
	fmt.Println("Cliente:", usuario.Nombre())
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Impuesto: $%.2f\n", impuesto)
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("Estado del pedido:", pedido.Estado())
	fmt.Println("Pago:", pago.Estado())
	fmt.Printf("Ventas acumuladas: $%.2f\n", ventasTotales)

	fmt.Println()
	fmt.Println("Inventario actualizado:")
	fmt.Printf("%s - Stock: %d\n", laptop.Nombre(), laptop.Stock())
	fmt.Printf("%s - Stock: %d\n", mouse.Nombre(), mouse.Stock())
}
