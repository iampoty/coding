package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiniMaxSum(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  []int32
			Output []int64
		}
		Mocks []Mock
	)

	mocks := Mocks{
		{Input: []int32{1, 2, 3, 4, 5}, Output: []int64{10, 14}},
	}

	for _, mock := range mocks {
		output := miniMaxSum(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
