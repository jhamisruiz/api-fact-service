package service

import (
	"api-fact-service/api/auth/domain/model"
	"api-fact-service/api/auth/domain/repository"
	"api-fact-service/api/usuarios/domain/entity"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo      repository.AuthRepository
	secretKey string
}

func NewAuthService(repo repository.AuthRepository, secretKey string) *AuthService {
	return &AuthService{repo: repo, secretKey: secretKey}
}

// Login valida el nombre de usuario y la contraseña
func (s *AuthService) Login(userName, password string) (*entity.LoginResponse, error) {
	// Obtiene el usuario desde el repositorio
	usuario, err := s.repo.GetByUserName(userName)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	// Compara la contraseña con la almacenada en la base de datos
	/*err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("credenciales incorrectas")
	}*/
	if usuario.Password != password {
		return nil, fmt.Errorf("credenciales incorrectas")
	}

	// Genera un token JWT
	token, err := s.generateJWT(usuario)
	if err != nil {
		return nil, err
	}

	response := &entity.LoginResponse{
		Token:    token,
		Username: usuario.UserName,
		Email:    usuario.Email,
		Name:     usuario.Name,
	}

	return response, nil
}

// generateJWT genera un token JWT para el usuario
func (s *AuthService) generateJWT(usuario *model.UsuarioModel) (string, error) {
	// Define el encabezado y los claims del token
	claims := jwt.MapClaims{
		"sub": usuario.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiración 1 hora en segundos
		"iat": time.Now().Unix(),                     // Fecha de emisión
	}

	// Crea el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Println("Error al generar JWT:", err)
		return "", err
	}

	return signedToken, nil
}
