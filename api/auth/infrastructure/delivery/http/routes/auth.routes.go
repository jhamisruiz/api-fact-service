package routes

import (
	"api-fact-service/api/auth/application/service"
	"api-fact-service/api/auth/infrastructure/delivery/http/controller"
	"api-fact-service/api/auth/infrastructure/persistence/mssql_repository"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registra las rutas de autenticación
func RegisterAuthRoutes(route *gin.RouterGroup, db *sql.DB, secretKey string) {
	// Crea el repositorio y el servicio de autenticación
	repo := mssql_repository.NewUsuarioRepository(db)
	authService := service.NewAuthService(repo, secretKey)

	// Crea el controlador de autenticación
	authCtrl := controller.NewAuthController(authService)

	// Registra la ruta de login
	route.POST("/login", authCtrl.Login)
}
