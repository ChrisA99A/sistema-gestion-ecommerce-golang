package store

import "sistema-gestion-ecommerce-golang/models"

var UsuarioPrincipal *models.Usuario
var Productos []*models.Producto
var CarritoActual *models.Carrito
var Pedidos []*models.Pedido
var Pagos []*models.Pago

// InicializarDatos carga informaciÃ³n inicial para probar los servicios web.
// En esta etapa se utiliza memoria temporal porque el proyecto aÃºn no integra una base de datos.
func InicializarDatos() error {
	usuario, err := models.NewUsuario(1, "Christian Cevallos", "christian@email.com", "cliente")
	if err != nil {
		return err
	}

	laptop, err := models.NewProducto(1, "Laptop", "Tecnologia", 850.00, 5)
	if err != nil {
		return err
	}

	mouse, err := models.NewProducto(2, "Mouse", "Accesorios", 15.50, 20)
	if err != nil {
		return err
	}

	teclado, err := models.NewProducto(3, "Teclado", "Accesorios", 35.00, 10)
	if err != nil {
		return err
	}

	audifonos, err := models.NewProducto(4, "Audifonos", "Accesorios", 25.00, 0)
	if err != nil {
		return err
	}

	UsuarioPrincipal = usuario
	Productos = []*models.Producto{laptop, mouse, teclado, audifonos}
	CarritoActual = models.NewCarrito(usuario.ID())
	Pedidos = []*models.Pedido{}
	Pagos = []*models.Pago{}

	return nil
}

func SiguientePedidoID() int {
	return len(Pedidos) + 1
}
