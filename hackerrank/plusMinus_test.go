package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusMinus(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  []int32
			Output []string
		}
		Mocks []Mock
	)

	mocks := Mocks{
		{Input: []int32{-4, 3, -9, 0, 4, 1}, Output: []string{"0.500000", "0.333333", "0.166667"}},
	}

	for _, mock := range mocks {
		output := plusMinus(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
