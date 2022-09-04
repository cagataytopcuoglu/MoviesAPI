package jwtHelper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SecretKey = "Movie-Api"
)

func GetJwtByUser(firstName, lastName string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["firstName"] = firstName
	claims["lastName"] = lastName
	claims["isAuthenticated"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "nil", err
	}
	return t, nil
}

func IsUserAuthenticated(soken string) bool {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(soken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}

	return false
}
