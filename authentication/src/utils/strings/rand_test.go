package strings_test

import (
	"streamer/utils/strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandString(t *testing.T) {
	assert := require.New(t)

	assert.NotPanics(func() {
		_ = strings.RandString(30)
	})

	assert.Len(strings.RandString(30), 30)
}
