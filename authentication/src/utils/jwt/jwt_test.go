package jwt_test

import (
	"flag"
	"testing"
	"time"

	"authentication/configs"
	"authentication/utils/jwt"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	flag.Parse()
	assert := require.New(t)

	var exp int64 = 30

	privateKey, err := jwt.ParseSSHPrivateKey(configs.JwtPrivate)
	assert.Nil(err)

	jwt.Init(privateKey)

	tests := []struct {
		exp time.Duration
		err error
	}{
		{exp: 30, err: nil},
		{exp: 30 * 60, err: nil},
		{exp: 30 * 60 * 60, err: nil},
	}

	for _, test := range tests {
		token, err := jwt.Signer.Generate(&jwt.Session{}, time.Duration(exp))
		assert.Equal(test.err, err)
		if err != nil {
			assert.NotEmpty(token)
		}
	}
}
