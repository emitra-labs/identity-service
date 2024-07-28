package controller

import (
	"crypto"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtPrivateKey crypto.PrivateKey

func init() {
	var err error

	jwtPrivateKey, err = jwt.ParseEdPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		panic(err)
	}
}
