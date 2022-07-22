package http_middleware

import (
	"context"
	"os"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/golang-jwt/jwt"
)

func JWTValidator () *jwtmiddleware.JWTMiddleware{

	keyFunc := func(ctx context.Context) (interface{}, error) {
        return []byte("secret"), nil
    }
	jwtValidator, _ := validator.New(keyFunc, validator.HS256, "http://localhost:8000", []string{"api:read"})
    jwtMiddware := jwtmiddleware.New(jwtValidator.ValidateToken)
	return jwtMiddware
}


func GenerateToken(user_id int8) (string, error) {
	token_lifespan, err := strconv.Atoi(string(5 * time.Minute))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}