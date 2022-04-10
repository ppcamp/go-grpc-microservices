package jwt_test

import (
	"streamer/utils/jwt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSSHPrivateKey(t *testing.T) {
	assert := require.New(t)

	key := `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAArAAAABNlY2RzYS
1zaGEyLW5pc3RwNTIxAAAACG5pc3RwNTIxAAAAhQQBXkAw3Sr/erwBzorzxg4hJErdNhPK
hZhjc/mhb/4WdH5D90C3h8ZuxwjGYLjXAo9xxU4txxEY/IzdJx70120zOk8BBunr0SM5vK
mZc35dwrASEbYgBZnSihx5jvbIm3e4GWYn1PPnpmUXhRdWcmC9L+R7k/DNr7Ouydr0rMk8
Y5Mg+MMAAAEYbnyzrm58s64AAAATZWNkc2Etc2hhMi1uaXN0cDUyMQAAAAhuaXN0cDUyMQ
AAAIUEAV5AMN0q/3q8Ac6K88YOISRK3TYTyoWYY3P5oW/+FnR+Q/dAt4fGbscIxmC41wKP
ccVOLccRGPyM3Sce9NdtMzpPAQbp69EjObypmXN+XcKwEhG2IAWZ0ooceY72yJt3uBlmJ9
Tz56ZlF4UXVnJgvS/ke5Pwza+zrsna9KzJPGOTIPjDAAAAQgFbCxbLXNkxBcdk+46SXOwr
x8tIUfjKNd+LZoiu7vFTk+V2L8jvaKlj3anxhcyrvSf28D8Jna1LZ5Ru+AaXgFLJCgAAAB
ZwcGNhbXBAREVTS1RPUC0xNE9WNTVQAQIDBA==
-----END OPENSSH PRIVATE KEY-----`

	_, err := jwt.ParseSSHPrivateKey(key)
	assert.Nil(err)

	wrongKey := []rune(key)
	wrongKey[37] = rune('a')

	_, err = jwt.ParseSSHPrivateKey(string(wrongKey))
	assert.NotNil(err)
}
