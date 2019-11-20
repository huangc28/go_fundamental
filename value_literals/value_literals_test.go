package valueliteral

import "testing"

import "github.com/magiconair/properties/assert"

func TestValueLiterals(t *testing.T) {
	// hex(base16) from of 15
	assert.Equal(t, 0xF, 15)
	assert.Equal(t, 0XF, 15)

	// octal(base8) form of 15
	assert.Equal(t, 017, 15)
	assert.Equal(t, 0o17, 15)
	assert.Equal(t, 0O17, 15)
}

func TestTypeAlias(t *testing.T) {
	var v1 byte
	var v2 uint8

	// byte is an alias type for uint8
	assert.Equal(t, v1, v2)

	// rune is an alias type for int32
	var v3 rune
	var v4 int32
	assert.Equal(t, v3, v4)
}
