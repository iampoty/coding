package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {

	assert := assert.New(t)
	type (
		Mock struct {
			Input  string
			Output string
		}
		Mocks []Mock
	)

	mocks := Mocks{
		// {Input: "12:40:22AM", Output: "00:40:22"},
		{Input: "12:45:54PM", Output: "12:45:54"},
	}

	for _, mock := range mocks {
		output := timeConversion(mock.Input)
		assert.Equal(output, mock.Output)

	}

}
