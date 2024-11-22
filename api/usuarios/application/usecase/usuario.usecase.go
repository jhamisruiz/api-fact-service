package usecase

import (
	"api-fact-service/api/usuarios/domain/entity"
	"api-fact-service/api/usuarios/domain/mapping"
	"api-fact-service/api/usuarios/domain/repository"
)

type UsuarioUsecase struct {
	repo repository.UsuarioRepository
}

func NewUsuarioUsecase(repo repository.UsuarioRepository) *UsuarioUsecase {
	return &UsuarioUsecase{repo: repo}
}

func (uc *UsuarioUsecase) CreateUsuario(usuario *entity.UsuarioEntity) error {
	model := mapping.ToUsuarioModel(usuario)
	// Pasar el modelo mapeado al repositorio
	return uc.repo.Create(model)
}

func (uc *UsuarioUsecase) GetAllUsuarios() ([]*entity.UsuarioEntity, error) {
	// Llama al repositorio para obtener todos los modelos desde la base de datos
	models, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Mapea los modelos a entidades del dominio
	var usuarios []*entity.UsuarioEntity
	for _, model := range models {
		usuarios = append(usuarios, mapping.ToUsuarioEntity(model))
	}
	return usuarios, nil
}

func (uc *UsuarioUsecase) GetUsuario(id int) (*entity.UsuarioEntity, error) {
	// Llama al repositorio para obtener el modelo desde la base de datos
	model, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Convierte el modelo a una entidad del dominio
	return mapping.ToUsuarioEntity(model), nil

}

func (uc *UsuarioUsecase) UpdateUsuario(id int, usuario *entity.UsuarioEntity) error {
	model := mapping.ToUsuarioModel(usuario)
	// Pasar el modelo mapeado al repositorio
	return uc.repo.Update(id, model)
}

func (uc *UsuarioUsecase) DeleteUsuario(id int) error {
	return uc.repo.Delete(id)
}
