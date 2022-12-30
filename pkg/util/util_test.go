package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomInt(t *testing.T) {
	num := RandomInt(1, 2)
	require.NotEmpty(t, num)
}

func TestRandomString(t *testing.T) {
	str := RandomString(6)
	require.NotEmpty(t, str)
}

func TestRandomOwner(t *testing.T) {
	owner := RandomOwner()
	require.NotEmpty(t, owner)
}

func TestRandomMoney(t *testing.T) {
	num := RandomMoney()
	require.NotEmpty(t, num)
}

func TestRandomCurrency(t *testing.T) {
	currency := RandomCurrency()
	require.NotEmpty(t, currency)
}

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	require.NotEmpty(t, email)
	require.Contains(t, email, "@")
}
