package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchingStrings(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Strings []string
			Queries []string
			Output  []int32
		}
		Mocks []Mock
	)

	mocks := Mocks{
		{Strings: []string{"aba", "baba", "aba", "xzxb"}, Queries: []string{"aba", "xzxb", "ab"}, Output: []int32{2, 1, 0}},
	}

	for _, mock := range mocks {
		output := matchingStrings(mock.Strings, mock.Queries)
		assert.Equal(output, mock.Output)

	}

}
