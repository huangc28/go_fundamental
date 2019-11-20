package statrannum

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestStateRandomNumbers(t *testing.T) {
	num := 100

	num1, num2 := StatRandomNumbers(num)

	assert.Equal(t, num1, 46)
	assert.Equal(t, num2, 54)
}
