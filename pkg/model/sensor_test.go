package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensor(t *testing.T) {
	s := Sensor{
		SensorID:       1,
		Symbol:         "exterior_temperature",
		Name:           "Exterior Temperature",
		Description:    "The Exterior temperature.",
		Obsolete:       false,
		Type:           "int16",
		Unit:           "C",
		UnitConversion: "10 /",
		DatasheetURL:   "http://example.com/",
	}
	assert.NoError(t, s.Validate())

	// Test for symbol invalid names
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "_ack"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "ack_"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "Ack"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: " ack"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "ack "}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "hackHack"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: "æøå"}.Validate(), ErrInvalidSymbolName)
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "int8", Symbol: ""}.Validate(), ErrInvalidSymbolName)

	// Test for zero SensorID
	assert.ErrorIs(t, Sensor{Type: "string", Symbol: "valid"}.Validate(), ErrSensorIDZero)

	// Make sure we only use MaxLength for variable length data
	assert.ErrorIs(t, Sensor{SensorID: 1, Type: "uint8", MaxLength: 1, Symbol: "valid"}.Validate(), ErrMaxLengthForNonVariableData)
	assert.NoError(t, Sensor{SensorID: 1, Type: "string", MaxLength: 1, Symbol: "valid"}.Validate())
	assert.NoError(t, Sensor{SensorID: 1, Type: "bytes", MaxLength: 1, Symbol: "valid"}.Validate())

}
