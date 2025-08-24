package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const (
	AccessTokenSubject  = "access-token"
	RefreshTokenSubject = "refresh-token"
)

type MyCustomClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

type JwtResponse struct {
	Status  bool
	Message string
	Data    interface{}
}

type JwtService struct {
	signedKey             []byte
	accessExpirationTime  *jwt.NumericDate
	refreshExpirationTime *jwt.NumericDate
}

func NewJwtToken() JwtService {

	signinKey := []byte("@@##AAAtt##$@#@%23432424asdsad345345SFD")

	return JwtService{
		signedKey:             signinKey,
		accessExpirationTime:  jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		refreshExpirationTime: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 10)),
	}
}

func (jwtToken JwtService) CreateAccessToken(phoneNumber string) string {
	return jwtToken.createToken(phoneNumber, AccessTokenSubject)
}

func (jwtToken JwtService) CreateRefreshToken(phoneNumber string) string {
	return jwtToken.createToken(phoneNumber, RefreshTokenSubject)
}

func (jwtToken JwtService) createToken(phoneNumber string, subject string) string {

	anotherClaim := jwt.MapClaims{
		"phone_number": phoneNumber,
		"exp":          jwtToken.accessExpirationTime,
		"sub":          subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, anotherClaim)

	signedToken, err := token.SignedString(jwtToken.signedKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	return signedToken
}

func (jwtToken JwtService) Get(tokenString string) (bool, string, *MyCustomClaims) {

	jwtResponse := jwtToken.ParseToken(tokenString)

	if !jwtResponse.Status {
		return false, jwtResponse.Message, nil
	}

	if claims, ok := jwtResponse.Data.(*jwt.Token).Claims.(*MyCustomClaims); ok {
		return true, "", claims
	} else {
		return false, "unknown claims type, cannot proceed", nil
	}
}

func (jwtToken JwtService) ParseToken(tokenString string) JwtResponse {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtToken.signedKey), nil
	})

	if err != nil {
		return JwtResponse{
			Status:  false,
			Message: "token signature is invalid",
		}
	}

	expirationTime, expirationTimeErr := token.Claims.GetExpirationTime()

	if expirationTimeErr != nil {
		return JwtResponse{
			Status:  false,
			Message: "Error in calculating expiration time",
		}
	}

	if expirationTime.Time.Before(time.Now()) {
		return JwtResponse{
			Status:  false,
			Message: "Token is expired",
		}
	}

	return JwtResponse{
		Status: true,
		Data:   token,
	}
}
