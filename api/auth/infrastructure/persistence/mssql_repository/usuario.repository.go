package mssql_repository

import (
	"api-fact-service/api/auth/domain/model"
	"api-fact-service/api/auth/domain/repository"
	"database/sql"
	"fmt"
	"log"
)

// usuarioRepository es el repositorio que maneja las consultas de usuarios
type usuarioRepository struct {
	db *sql.DB
}

// NewUsuarioRepository crea un nuevo repositorio de usuario
func NewUsuarioRepository(db *sql.DB) repository.AuthRepository {
	return &usuarioRepository{db: db}
}

// GetByUserName obtiene un usuario por su nombre de usuario
func (r *usuarioRepository) GetByUserName(userName string) (*model.UsuarioModel, error) {
	var usuario model.UsuarioModel
	err := r.db.QueryRow("SELECT id, name, last_name, user_name, email, password FROM usuarios WHERE user_name = ? or email=?", userName, userName).Scan(
		&usuario.ID, &usuario.Name, &usuario.LastName, &usuario.UserName, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		log.Println("Error al obtener usuario:", err)
		return nil, err
	}
	return &usuario, nil
}
