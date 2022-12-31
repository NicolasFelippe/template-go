package uidgen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var uid UIDGen

func init() {
	uid = New()
}

func TestUuidValid(t *testing.T) {
	id := uid.New()
	require.NotEmpty(t, id)
	uuid, isValid := uid.IsValidUuid(id.String())
	require.True(t, isValid)
	require.NotEmpty(t, uuid)
}

func TestUuidInvalid(t *testing.T) {
	id := "00000000-0000-0000-0000-000000000000"
	require.NotEmpty(t, id)
	uuid, isValid := uid.IsValidUuid(id)
	require.False(t, isValid)
	require.Empty(t, uuid)
}
