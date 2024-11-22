package service

import (
	"api-fact-service/api/usuarios/application/usecase"
	"api-fact-service/api/usuarios/domain/entity"
)

type UsuarioService struct {
	usecase *usecase.UsuarioUsecase
}

func NewUsuarioService(uc *usecase.UsuarioUsecase) *UsuarioService {
	return &UsuarioService{usecase: uc}
}

func (s *UsuarioService) RegisterUsuario(usuario *entity.UsuarioEntity) error {
	// Puedes añadir validaciones adicionales aquí antes de crear el usuario
	return s.usecase.CreateUsuario(usuario)
}
