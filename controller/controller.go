package controller

import (
	"crypto"
	"os"

	"github.com/emitra-labs/common/auth"
)

var jwtPrivateKey crypto.PrivateKey

func init() {
	var err error

	jwtPrivateKey, err = auth.ParsePrivateKeyFromBase64(os.Getenv("BASE64_JWT_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}
}
