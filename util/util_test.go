package util

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	var randomStr string

	randomStr = RandomString(0)
	assert.Equal(t, "", randomStr)

	randomStr = RandomString(-1)
	assert.Equal(t, "", randomStr)

	// 测 5 次吧
	const randomTimes = 5
	for i := 0; i < randomTimes; i++ {
		length := rand.Intn(100)
		randomStr = RandomString(length)
		assert.Equal(t, length, len(randomStr))
	}
}
