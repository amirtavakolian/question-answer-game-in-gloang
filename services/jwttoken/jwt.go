package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type JwtToken struct {
	SignedKey string
}

func NewJwtToken() JwtToken {
	return JwtToken{SignedKey: "a-random-key"}
	// todo => move signed key to env
}

func (jwtToken JwtToken) CreateToken(phoneNumber string) string {

	mySigningKey := []byte("golang-programming-language")

	anotherClaim := jwt.MapClaims{
		"phone_number": phoneNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, anotherClaim)

	signedToken, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	return signedToken
}
