package main

import (
	"fmt"
	"log"
	"net/http"

	"sistema-gestion-ecommerce-golang/routes"
	"sistema-gestion-ecommerce-golang/store"
)

func main() {
	if err := store.InicializarDatos(); err != nil {
		log.Fatal("Error al inicializar datos: ", err)
	}

	routes.RegistrarRutas()

	fmt.Println("Sistema de Gestion de E-commerce en Golang")
	fmt.Println("Servidor web iniciado correctamente")
	fmt.Println("URL base: http://localhost:8080")
	fmt.Println("Servicio de prueba: http://localhost:8080/api/salud")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error al iniciar servidor: ", err)
	}
}
