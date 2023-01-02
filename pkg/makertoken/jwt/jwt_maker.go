package jwt

import (
	"errors"
	"fmt"
	"template-go/pkg/makertoken"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func New(secretKey string) (makertoken.Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *makertoken.Payload, error) {
	payload, err := makertoken.NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenSigned, err := jwtToken.SignedString([]byte(maker.secretKey))

	return tokenSigned, payload, err
}

func (maker *JWTMaker) VerifyToken(tokenStr string) (*makertoken.Payload, error) {
	keyFunc := func(tokenJwt *jwt.Token) (interface{}, error) {
		_, ok := tokenJwt.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, makertoken.ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenStr, &makertoken.Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, makertoken.ErrExpiredToken) {
			return nil, makertoken.ErrExpiredToken
		}
		return nil, makertoken.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*makertoken.Payload)
	if !ok {
		return nil, makertoken.ErrInvalidToken
	}

	return payload, nil
}
