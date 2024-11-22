package controller

import (
	"api-fact-service/api/auth/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login maneja la autenticación del usuario
func (c *AuthController) Login(ctx *gin.Context) {
	var credentials struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	// Bind los datos JSON a la estructura
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llama al servicio de autenticación
	logged, err := c.authService.Login(credentials.UserName, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Devuelve el token JWT
	ctx.JSON(http.StatusOK, logged)
}
