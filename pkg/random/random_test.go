package random

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomInt(t *testing.T) {
	num := Int(1, 2)
	require.NotEmpty(t, num)
}

func TestRandomString(t *testing.T) {
	str := String(6)
	require.NotEmpty(t, str)
}

func TestRandomOwner(t *testing.T) {
	owner := Owner()
	require.NotEmpty(t, owner)
}

func TestRandomMoney(t *testing.T) {
	num := Money()
	require.NotEmpty(t, num)
}

func TestRandomCurrency(t *testing.T) {
	currency := Currency()
	require.NotEmpty(t, currency)
}

func TestRandomEmail(t *testing.T) {
	email := Email()
	require.NotEmpty(t, email)
	require.Contains(t, email, "@")
}
