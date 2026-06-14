# Sistema de Gestión de E-commerce en Golang 

## Descripción

Proyecto académico desarrollado en Golang para representar un Sistema de Gestión de E-commerce. El sistema permite simular un flujo básico de compra mediante usuarios, productos, carrito de compras, pedidos, pagos simulados, inventario y reportes básicos.

Este repositorio corresponde al desarrollo del Trabajo Autónomo de Programación Orientada a Objetos, tomando como base la planeación inicial del sistema y continuando con la codificación funcional.

## Objetivo

Desarrollar un avance funcional de un sistema de e-commerce en Golang, aplicando estructuras, métodos, encapsulación, manejo de errores, interfaces y comentarios en funcionalidades complejas.

## Módulos implementados

- Usuarios
- Productos
- Inventario
- Carrito de compras
- Pedidos
- Pagos simulados
- Reportes básicos

## Desarrollo del Paso 1

En el Paso 1 se inició la codificación del sistema. Se crearon las estructuras principales del proyecto:

- Usuario
- Producto
- ItemCarrito
- Carrito
- Pedido
- Pago

También se implementó un flujo básico que permite:

1. Crear un usuario.
2. Crear productos.
3. Listar productos disponibles.
4. Agregar productos al carrito.
5. Calcular subtotal, impuesto y total.
6. Generar un pedido.
7. Simular un pago.
8. Actualizar inventario.

## Desarrollo del Paso 2

En el Paso 2 se mejoró la estructura interna del sistema aplicando:

- Encapsulación mediante atributos privados y métodos públicos.
- Manejo de errores mediante el paquete `errores`.
- Interfaces mediante el paquete `interfaces`.
- Comentarios en funciones complejas como inventario, pago y reportes.

## Estructura del proyecto

```text
sistema-gestion-ecommerce-golang/
│
├── go.mod
├── main.go
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
└── errores/
    └── errores.go