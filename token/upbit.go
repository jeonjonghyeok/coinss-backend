package upbit

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type customClaims struct {
	jwt.StandardClaims
	AccessKey string `json:"access_key"`
	Nonce     string `json:"nonce"`
}

func New(AccessKey, SecretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		AccessKey:      AccessKey,
		Nonce:          uuid.NewV4().String(),
		StandardClaims: jwt.StandardClaims{},
	})
	return token.SignedString([]byte(SecretKey))

}
