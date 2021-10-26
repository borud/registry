package reg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataType(t *testing.T) {
	// Make sure all the types we need are there.
	assert.Equal(t, DataTypeFromValue(uint8(1)), TypeUint8)
	assert.Equal(t, DataTypeFromValue(uint16(1)), TypeUint16)
	assert.Equal(t, DataTypeFromValue(uint32(1)), TypeUint32)
	assert.Equal(t, DataTypeFromValue(uint64(1)), TypeUint64)
	assert.Equal(t, DataTypeFromValue(int8(1)), TypeInt8)
	assert.Equal(t, DataTypeFromValue(int16(1)), TypeInt16)
	assert.Equal(t, DataTypeFromValue(int32(1)), TypeInt32)
	assert.Equal(t, DataTypeFromValue(int64(1)), TypeInt64)
	assert.Equal(t, DataTypeFromValue([]byte{}), TypeByteArray)
	assert.Equal(t, DataTypeFromValue("foo"), TypeString)
	assert.Equal(t, DataTypeFromValue(true), TypeBool)

	// Check that we have valid names for these types for C and Go
	for k, v := range typeNames {
		// make sure we get sensible type name for all types that are not "unkown"
		if k != "unknown" {
			assert.NotEqual(t, v.GoTypeName(), "unknown")
			assert.NotEqual(t, v.CTypeName(), "unknown")
		}
	}
}
