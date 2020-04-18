package verification_test

import (
	cbConfig "cab/config"
	vf "cab/models/verification"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAuthentication(t *testing.T) {
	t.Parallel()

	var user vf.User
	pg := cbConfig.InitPG()
	c := cbConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()

	token := ""
	username := ""
	password := ""
	err := user.UserAuthentication(c, token, username, password)
	require.Error(t, err)

	token = ""
	username = "user1"
	password = "user1password"
	err = user.UserAuthentication(c, token, username, password)
	require.Error(t, err)

	token = "token"
	username = ""
	password = "user1password"
	err = user.UserAuthentication(c, token, username, password)
	require.Error(t, err)

	token = "token"
	username = "user1"
	password = ""
	err = user.UserAuthentication(c, token, username, password)
	require.Error(t, err)

	token = "token"
	username = "user1"
	password = "user1password"
	err = user.UserAuthentication(c, token, username, password)
	require.NoError(t, err)

}
