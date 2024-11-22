package mssql_repository

import (
	"api-fact-service/api/usuarios/domain/model"
	"api-fact-service/api/usuarios/domain/repository"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type usuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) repository.UsuarioRepository {
	return &usuarioRepository{db: db}
}

func (r *usuarioRepository) Create(usuario *model.UsuarioModel) error {
	_, err := r.db.Exec(
		"INSERT INTO usuarios (name, last_name, user_name, email, telefono) VALUES (?, ?, ?, ?, ?)",
		usuario.Name, usuario.LastName, usuario.UserName, usuario.Email, usuario.Telefono)
	return err
}

func (r *usuarioRepository) GetAll() ([]*model.UsuarioModel, error) {
	rows, err := r.db.Query("SELECT id, name, last_name, user_name, email, habilitado, telefono FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*model.UsuarioModel
	for rows.Next() {
		var model model.UsuarioModel
		if err := rows.Scan(&model.ID, &model.Name, &model.LastName, &model.UserName, &model.Email, &model.Habilitado, &model.Telefono); err != nil {
			return nil, err
		}
		models = append(models, &model)
	}
	return models, nil
}

func (r *usuarioRepository) GetByID(id int) (*model.UsuarioModel, error) {
	row := r.db.QueryRow("SELECT id, name, last_name, user_name, email, habilitado, telefono FROM usuarios WHERE id = ?", id)
	var model model.UsuarioModel
	err := row.Scan(&model.ID, &model.Name, &model.LastName, &model.UserName, &model.Email, &model.Habilitado, &model.Telefono)
	return &model, err
}

func (r *usuarioRepository) Update(id int, usuario *model.UsuarioModel) error {
	_, err := r.db.Exec(
		"UPDATE usuarios SET name = ?, last_name = ?,user_name= ?, email= ?,telefono= ? WHERE id = ?",
		usuario.Name, usuario.LastName, usuario.UserName, usuario.Email, usuario.Telefono, id)
	return err
}

func (r *usuarioRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	return err
}
