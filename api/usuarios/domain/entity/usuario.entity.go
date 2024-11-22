package entity

type UsuarioEntity struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"user_name"`
	Habilitado bool   `json:"habilitado"`
	Email      string `json:"email"`
	Telefono   string `json:"telefono"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
