package jwt_test

import (
	"flag"
	"streamer/configs"
	"streamer/utils/jwt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	flag.Parse()
	assert := require.New(t)

	var exp int64 = 30

	privateKey, err := jwt.ParseSSHPrivateKey(*configs.JWT_PRIVATE)
	assert.Nil(err)

	authorizer := jwt.NewJwt(privateKey)

	tests := []struct {
		exp time.Duration
		err error
	}{
		{exp: 30, err: nil},
		{exp: 30 * 60, err: nil},
		{exp: 30 * 60 * 60, err: nil},
	}

	for _, test := range tests {
		token, err := authorizer.Generate(&jwt.Session{}, exp)
		assert.Equal(test.err, err)
		if err != nil {
			assert.NotEmpty(token)
		}
	}
}
