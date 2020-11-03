package app

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type M map[string]interface{}

type MyClaims struct {
    jwt.StandardClaims
    Username string `json:"username"`
    Name 	 string `json:"name"`
    Job      string `json:"job"`
}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("signature of my APP")