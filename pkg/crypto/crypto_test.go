package crypto

import (
	"template-go/pkg/random"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var crypt Crypto

func init() {
	crypt = New()
}

func TestPassword(t *testing.T) {
	password := random.String(6)
	hashedPassword1, err := crypt.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = crypt.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := random.String(6)
	err = crypt.CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := crypt.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
