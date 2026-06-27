# Sistema de Gestion de E-commerce en Golang

## Datos del proyecto

**Estudiante:** Christian Porfirio Cevallos Armijos
**Asignatura:** Programacion Orientada a Objetos
**Docente:** Palacios Morocho Milton Ricardo
**Proyecto:** Sistema de Gestion de E-commerce en Golang
**Fecha:** Junio de 2026

## Descripcion general

Este proyecto corresponde al desarrollo de un Sistema de Gestion de E-commerce implementado en Golang. El sistema permite representar el funcionamiento basico de una tienda digital mediante modulos de usuarios, productos, inventario, carrito de compras, pedidos, pagos simulados y reportes.

El proyecto inicio con una fase de planeacion, donde se selecciono el sistema de e-commerce por ser una opcion modular, aplicable y viable para un desarrollo academico individual. Posteriormente, se desarrollo una primera version funcional en consola y, en la etapa final, se incorporaron servicios web con respuestas en formato JSON.

## Objetivo del sistema

Desarrollar un sistema basico de gestion de e-commerce en Golang, integrando estructuras, metodos, encapsulacion, manejo de errores, interfaces, servicios web y serializacion JSON, con el proposito de representar un flujo funcional de compra y evidenciar la aplicacion de los contenidos de la asignatura.

## Modulos implementados

El sistema integra los siguientes modulos:

* Usuarios
* Productos
* Inventario
* Carrito de compras
* Pedidos
* Pagos simulados
* Reportes basicos

## Estructura del proyecto

```text
sistema-gestion-ecommerce-golang/
│
├── go.mod
├── main.go
├── README.md
├── .gitignore
│
├── models/
│   ├── usuario.go
│   ├── producto.go
│   ├── carrito.go
│   ├── pedido.go
│   └── pago.go
│
├── services/
│   ├── producto_service.go
│   ├── carrito_service.go
│   ├── inventario_service.go
│   ├── pago_service.go
│   ├── reporte_service.go
│   └── producto_repository_memoria.go
│
├── interfaces/
│   └── repositories.go
│
├── errores/
│   └── errores.go
│
├── store/
│   └── memoria.go
│
├── web/
│   ├── response.go
│   └── dtos.go
│
├── handlers/
│   ├── salud_handler.go
│   ├── usuarios_handler.go
│   ├── productos_handler.go
│   ├── inventario_handler.go
│   ├── carrito_handler.go
│   ├── pedidos_handler.go
│   ├── pagos_handler.go
│   └── reportes_handler.go
│
└── routes/
    └── routes.go
```

## Explicacion de carpetas

| Carpeta      | Funcion                                                                                              |
| ------------ | ---------------------------------------------------------------------------------------------------- |
| `models`     | Contiene las estructuras principales del sistema, como Usuario, Producto, Carrito, Pedido y Pago.    |
| `services`   | Contiene la logica del negocio, como calculos, inventario, pagos, reportes y repositorio en memoria. |
| `interfaces` | Define comportamientos esperados para productos y pagos.                                             |
| `errores`    | Centraliza errores reutilizables del sistema.                                                        |
| `store`      | Guarda datos temporales en memoria mientras el servidor esta activo.                                 |
| `web`        | Contiene estructuras DTO y funciones para responder en formato JSON.                                 |
| `handlers`   | Contiene los servicios web del sistema.                                                              |
| `routes`     | Registra las rutas HTTP disponibles.                                                                 |

## Tecnologias utilizadas

* Golang
* Visual Studio Code
* GitHub
* PowerShell
* Servicios web HTTP
* Serializacion JSON

## Ejecucion del proyecto

Para ejecutar el sistema, abrir una terminal en la carpeta principal del proyecto y utilizar:

```bash
go run .
```

Si el sistema inicia correctamente, se mostrara:

```text
Sistema de Gestion de E-commerce en Golang
Servidor web iniciado correctamente
URL base: http://localhost:8080
Servicio de prueba: http://localhost:8080/api/salud
```

## Servicios web disponibles

| Metodo | Endpoint                     | Funcion                                    |
| ------ | ---------------------------- | ------------------------------------------ |
| GET    | `/api/salud`                 | Verifica que el servidor este activo.      |
| GET    | `/api/usuarios`              | Lista los usuarios registrados.            |
| GET    | `/api/productos`             | Lista todos los productos.                 |
| GET    | `/api/productos/disponibles` | Lista productos con stock disponible.      |
| GET    | `/api/producto?id=1`         | Busca un producto por su ID.               |
| GET    | `/api/inventario`            | Consulta el stock y estado de inventario.  |
| GET    | `/api/carrito`               | Muestra el carrito actual.                 |
| POST   | `/api/carrito/agregar`       | Agrega productos al carrito mediante JSON. |
| GET    | `/api/pedidos`               | Lista los pedidos generados.               |
| POST   | `/api/pedidos/generar`       | Genera un pedido desde el carrito.         |
| POST   | `/api/pagos/procesar`        | Procesa un pago simulado.                  |
| GET    | `/api/reportes/ventas`       | Muestra el reporte de ventas acumuladas.   |

## Ejemplos de uso

### Consultar estado del servidor

```powershell
Invoke-RestMethod -Method Get -Uri "http://localhost:8080/api/salud"
```

### Listar productos

```powershell
Invoke-RestMethod -Method Get -Uri "http://localhost:8080/api/productos"
```

### Agregar producto al carrito

```powershell
$body = @{
    producto_id = 1
    cantidad = 1
} | ConvertTo-Json

Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/carrito/agregar" -ContentType "application/json" -Body $body
```

### Generar pedido

```powershell
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/pedidos/generar"
```

### Procesar pago simulado

```powershell
$body = @{
    pedido_id = 1
    metodo = "Tarjeta simulada"
} | ConvertTo-Json

Invoke-RestMethod -Method Post -Uri "http://localhost:8080/api/pagos/procesar" -ContentType "application/json" -Body $body
```

## Flujo principal del sistema

El flujo funcional del sistema es el siguiente:

1. Se inicia el servidor web.
2. Se consultan los productos disponibles.
3. Se agregan productos al carrito.
4. Se genera un pedido a partir del carrito.
5. Se procesa un pago simulado.
6. Se actualiza el inventario.
7. Se consulta el reporte de ventas.

## Relacion con las unidades de la asignatura

| Unidad   | Aplicacion en el proyecto                                                                    |
| -------- | -------------------------------------------------------------------------------------------- |
| Unidad 1 | Planeacion del sistema, seleccion del e-commerce, modulos, alcance y programacion funcional. |
| Unidad 2 | Desarrollo de estructuras, metodos, paquetes y flujo funcional inicial.                      |
| Unidad 3 | Encapsulacion, manejo de errores, interfaces y comentarios en funciones importantes.         |
| Unidad 4 | Servicios web, endpoints HTTP y serializacion JSON.                                          |

## Alcance actual

El sistema permite simular un flujo basico de compra en un e-commerce. Incluye usuarios, productos, carrito, pedidos, pagos simulados, inventario y reportes.

## Limitaciones

El sistema no incluye todavia:

* Base de datos real.
* Inicio de sesion con contrasena.
* Pasarela de pago bancaria.
* Facturacion electronica.
* Gestion real de envios.
* Interfaz grafica web.
* Aplicacion movil.

Estas funcionalidades se consideran posibles mejoras futuras, ya que aumentarian el alcance tecnico del proyecto.

## Estado del proyecto

El proyecto se encuentra en una version academica funcional. Actualmente permite demostrar la integracion de estructuras, servicios internos, manejo de errores, interfaces, servicios web y respuestas JSON.
