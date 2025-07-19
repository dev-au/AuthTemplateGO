package utils

import (
	"AuthTemplate/src"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GenerateJWT(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userID.String(),
			"exp": time.Now().Add(time.Hour * 24 * time.Duration(src.Config.JWTExpireDays)).Unix(),
		})
	jwtBinary := []byte(src.Config.JWTSecret)

	return token.SignedString(jwtBinary)
}

func VerifyJWT(tokenString string) (uuid.UUID, error) {
	jwtBinary := []byte(src.Config.JWTSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtBinary, nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["id"].(string)
		return uuid.MustParse(userID), nil
	}

	return uuid.Nil, fmt.Errorf("invalid token")
}
