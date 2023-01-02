package paseto

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	domainErrors "template-go/internal/core/domain/errors"
	"template-go/pkg/makertoken"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func New(symmetricKey string) (makertoken.Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *makertoken.Payload, error) {
	payload, err := makertoken.NewPayload(username, duration)
	if err != nil {
		appError := domainErrors.NewAppErrorWithType(domainErrors.UUIDGenError)
		return "", payload, appError
	}

	tokenEncrypt, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return tokenEncrypt, nil, domainErrors.NewAppErrorWithType(domainErrors.TokenGeneratorError)
	}
	return tokenEncrypt, payload, nil
}

func (maker *PasetoMaker) VerifyToken(tokenStr string) (*makertoken.Payload, error) {
	payload := &makertoken.Payload{}

	err := maker.paseto.Decrypt(tokenStr, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
	}
	err = payload.Valid()
	if err != nil {
		return nil, domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
	}

	return payload, nil
}
