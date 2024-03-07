package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  [][]int32
			Output int32
		}
		Mocks []Mock
	)

	mocks := Mocks{
		{Input: [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, Output: 2},
		{Input: [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, Output: 15},
		// {Input: [][]int32{{}, {}, {}}, Output: 12},
	}

	for _, mock := range mocks {
		output := diagonalDifference(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
