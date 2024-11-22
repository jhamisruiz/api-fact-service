package main

import (
	"api-fact-service/api/auth/infrastructure/delivery/http/routes"
	usersRoutes "api-fact-service/api/usuarios/infrastructure/delivery/http/routes"
	"api-fact-service/config"
	"api-fact-service/utils/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Inicia la conexión a la base de datos
	// Obtiene la configuración de la base de datos
	dbConfig := config.GetDatabaseConfig()

	// Inicializa la conexión a la base de datos
	db, err := database.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	// Configura Gin y las rutas
	r := gin.Default()
	routeGroup := r.Group("/v1")

	// Registra las rutas
	routes.RegisterAuthRoutes(routeGroup, db, "secretkey123")
	usersRoutes.RegisterRoutes(routeGroup, db)

	// Inicia el servidor
	r.Run(":8080")
}
