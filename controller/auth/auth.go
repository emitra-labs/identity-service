package auth

import (
	"crypto"
	"os"
	"time"

	"github.com/ukasyah-dev/common/auth"
)

var accessExpiresIn = 30 * time.Minute
var privateKey crypto.PrivateKey

func init() {
	var err error

	privateKey, err = auth.ParsePrivateKeyFromBase64(os.Getenv("BASE64_JWT_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}
}
