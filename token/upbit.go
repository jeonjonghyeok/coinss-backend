package upbit

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type customClaims struct {
	jwt.StandardClaims
	AccessKey string `json:"access_key"`
	Nonce     string `json:"nonce"`
	UserID    int    `json:"id"`
}

const mySigningKey = "coinss"

func NewUpbit(id int, AccessKey, SecretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		UserID:         id,
		AccessKey:      AccessKey,
		Nonce:          uuid.NewV4().String(),
		StandardClaims: jwt.StandardClaims{},
	})
	return token.SignedString([]byte(SecretKey))
}

func New(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			Issuer:    "jjh",
		},
	})
	return token.SignedString([]byte(mySigningKey))
}

func Parse(token string) (userID int, err error) {
	parsed, err := jwt.ParseWithClaims(token, &customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})

	if err != nil {
		return 0, err
	}

	if !parsed.Valid {
		return 0, errors.New("token is invalid")
	}

	if c, ok := parsed.Claims.(*customClaims); ok {
		return c.UserID, nil
	}

	return 0, errors.New("token is invalid")
}
