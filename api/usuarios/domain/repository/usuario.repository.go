package repository

import (
	"api-fact-service/api/usuarios/domain/model"
)

type UsuarioRepository interface {
	Create(usuario *model.UsuarioModel) error
	GetAll() ([]*model.UsuarioModel, error)
	GetByID(id int) (*model.UsuarioModel, error)
	Update(id int, usuario *model.UsuarioModel) error
	Delete(id int) error
}
