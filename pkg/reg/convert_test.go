package reg

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversion(t *testing.T) {
	con, err := NewConversion("v * 10", TypeInt8)
	assert.Nil(t, err)

	stringResult, err := con.ConvertString(int8(123))
	assert.Nil(t, err)
	assert.Equal(t, stringResult, "1230")

	floatResult, err := con.Convert(int8(123))
	assert.Nil(t, err)
	assert.Equal(t, floatResult, float64(123*10))

	// try with invalid types
	_, err = con.Convert(int(123))
	assert.Error(t, err)

	_, err = con.Convert("123")
	assert.Error(t, err)
}

// When the UnitConversion field is empty it means that we
// should just return the value unmodified.
func TestIdentityConversion(t *testing.T) {
	con, err := NewConversion("", TypeString)
	assert.Nil(t, err)

	stringResult, err := con.ConvertString("foo bar")
	assert.Nil(t, err)
	assert.Equal(t, "foo bar", stringResult)

	// Since this is a string it should return a string
	result, err := con.Convert("foo bar")
	assert.Nil(t, err)
	assert.Equal(t, "foo bar", result)
}

func TestDBFSConversion(t *testing.T) {
	con, err := NewConversion("(20 * log10(abs(v)+1)) / 32767", TypeInt16)
	assert.Nil(t, err)

	res, err := con.ConvertString(int16(10000))
	assert.Nil(t, err)

	log.Printf("dBFS: %s", res)
}
