package controller

import (
	"crypto"
	"os"

	"github.com/emitra-labs/pb/mail"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var jwtPrivateKey crypto.PrivateKey
var mailClientConn *grpc.ClientConn
var mailClient mail.MailClient

func init() {
	var err error

	jwtPrivateKey, err = jwt.ParseEdPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		panic(err)
	}

	mailClientConn, err = grpc.NewClient(
		os.Getenv("MAIL_GRPC_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	mailClient = mail.NewMailClient(mailClientConn)
}

func CloseClientConnections() error {
	mailClientConn.Close()
	return nil
}
