package model

type UsuarioModel struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	LastName   string `db:"last_name"`
	UserName   string `db:"user_name"`
	Habilitado string `db:"habilitado"`
	Email      string `db:"email"`
	Telefono   string `db:"telefono"`
	Password   string `json:"password"`
}
