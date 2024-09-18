package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

// ? belajar
var SECRET_KEY = []byte("patunganBaik_s3cr3t")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Implement JWT token generation logic here
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
