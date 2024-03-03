package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlippingBits(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  int64
			Output int64
		}
		Mocks []Mock
	)

	mocks := Mocks{
		{Input: 3, Output: 0},
		{Input: 2147483647, Output: 2147483648},
		{Input: 1, Output: 4294967294},
		{Input: 0, Output: 4294967295},
	}

	for _, mock := range mocks {
		output := flippingBits(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
