package valueobject

type UsuarioResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"user_name"`
	Habilitado string `json:"habilitado"`
	Email      string `json:"email"`
	Telefono   string `json:"telefono"`
}

func NewUsuarioResponse(id int, name string) UsuarioResponse {
	return UsuarioResponse{
		ID:   id,
		Name: name,
		/*LastName:   last_name,
		UserName:   user_name,
		Habilitado: habilitado,
		Email:      email,
		Telefono:   telefono,*/
	}
}
