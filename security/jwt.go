package security

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// JWTInterface ...
type JWTInterface interface {
	Sign(authToken string) (string, error)
	Verify(jwtToken string) (bool, error)
}

// JWTService ...
type JWTService struct {
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
	expirationHours time.Duration
}

// NewJWTService ...
func NewJWTService(publicKey, privateKey string, expirationHours time.Duration) (JWTService, error) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return JWTService{}, errors.Wrap(err, "Failed to parse private key")
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return JWTService{}, errors.Wrap(err, "Failed to parse public key")
	}
	return JWTService{
		privateKey:      signKey,
		publicKey:       verifyKey,
		expirationHours: expirationHours,
	}, nil
}

// Sign ...
func (j *JWTService) Sign(authToken string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"token": authToken,
		"exp":   time.Now().Add(time.Hour * j.expirationHours).Unix(),
	})

	return jwtToken.SignedString(j.privateKey)
}

// Verify ...
func (j *JWTService) Verify(jwtToken string) (bool, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
