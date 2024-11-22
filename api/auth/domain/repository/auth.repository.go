package repository

import "api-fact-service/api/auth/domain/model"

type AuthRepository interface {
	GetByUserName(usuario string) (*model.UsuarioModel, error)
}
