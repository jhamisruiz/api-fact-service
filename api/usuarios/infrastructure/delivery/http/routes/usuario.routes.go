package routes

import (
	"api-fact-service/api/usuarios/application/usecase"
	"api-fact-service/api/usuarios/infrastructure/delivery/http/controller"
	"api-fact-service/api/usuarios/infrastructure/persistence/mssql_repository"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.RouterGroup, db *sql.DB) {
	// Inicializa el repositorio, caso de uso y controlador
	repo := mssql_repository.NewUsuarioRepository(db)
	uc := usecase.NewUsuarioUsecase(repo)
	ctrl := controller.NewUsuarioController(uc)

	// Define las rutas para el recurso "usuarios"
	route.POST("/usuarios", ctrl.CreateUsuario)
	route.GET("/usuarios", ctrl.GetAllUsuarios)
	route.GET("/usuarios/:id", ctrl.GetUsuario)
	route.PUT("/usuarios/:id", ctrl.UpdateUsuario)
	route.DELETE("/usuarios/:id", ctrl.DeleteUsuario)
}
