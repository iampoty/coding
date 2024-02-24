package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyinteger(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  []int32
			Output int32
		}
		Mocks []Mock
	)

	mocks := Mocks{
		// {Input: []int32{1, 2, 3, 4, 3, 2, 1}, Output: 4},
		// {Input: []int32{1, 2, 3, 4, 4, 2, 1}, Output: 3},
		{Input: []int32{59, 88, 14, 8, 85, 1, 94, 74, 57, 96, 39, 2, 47, 43, 35, 17, 53, 52, 92, 31, 99, 48, 94, 30, 92, 60, 32, 45, 88, 13, 39, 50, 22, 65, 89, 46, 65, 76, 57, 67, 99, 35, 76, 46, 85, 82, 45, 62, 53, 80, 74, 22, 31, 52, 82, 13, 41, 96, 2, 1, 80, 62, 4, 20, 50, 89, 59, 67, 60, 8, 41, 14, 47, 48, 17, 4, 43, 30, 32}, Output: 20},
	}

	for _, mock := range mocks {
		output := lonelyinteger(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
