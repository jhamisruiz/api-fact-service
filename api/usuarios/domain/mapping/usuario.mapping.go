package mapping

import (
	"api-fact-service/api/usuarios/domain/entity"
	"api-fact-service/api/usuarios/domain/model"
)

func ToUsuarioModel(usuario *entity.UsuarioEntity) *model.UsuarioModel {
	return &model.UsuarioModel{
		ID:       usuario.ID,
		Name:     usuario.Name,
		LastName: usuario.LastName,
		UserName: usuario.UserName,
		//Habilitado: usuario.Habilitado,
		Email:    usuario.Email,
		Telefono: usuario.Telefono,
	}
}

func ToUsuarioEntity(model *model.UsuarioModel) *entity.UsuarioEntity {
	var habilitado bool
	if len(model.Habilitado) > 0 {
		habilitado = model.Habilitado[0] == 1
	}

	return &entity.UsuarioEntity{
		ID:         model.ID,
		Name:       model.Name,
		LastName:   model.LastName,
		UserName:   model.UserName,
		Habilitado: habilitado,
		Email:      model.Email,
		Telefono:   model.Telefono,
	}
}
