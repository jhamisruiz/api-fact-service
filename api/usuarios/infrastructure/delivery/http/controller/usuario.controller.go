package controller

import (
	"api-fact-service/api/usuarios/application/usecase"
	"api-fact-service/api/usuarios/domain/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	usecase *usecase.UsuarioUsecase
}

func NewUsuarioController(uc *usecase.UsuarioUsecase) *UsuarioController {
	return &UsuarioController{usecase: uc}
}

func (c *UsuarioController) CreateUsuario(ctx *gin.Context) {
	var usuario entity.UsuarioEntity
	if err := ctx.BindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.usecase.CreateUsuario(&usuario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, usuario)
}

func (c *UsuarioController) GetAllUsuarios(ctx *gin.Context) {
	usuarios, err := c.usecase.GetAllUsuarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usuarios)
}

func (c *UsuarioController) GetUsuario(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	usuario, err := c.usecase.GetUsuario(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) UpdateUsuario(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var usuario entity.UsuarioEntity
	if err := ctx.BindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.usecase.UpdateUsuario(id, &usuario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) DeleteUsuario(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.usecase.DeleteUsuario(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
