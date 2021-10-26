package testdata

import (
	"github.com/borud/registry/pkg/reg"
)

var Registry = []reg.Sensor{
	reg.Sensor{
		SensorID:          1,
		Symbol:            "abc123_temperature",
		Model:             "Bosch BME680",
		Name:              "Ambient Temperature",
		Description:       "Ambient temperature sensor for ABC123 device",
		Obsolete:          false,
		Type:              "int16",
		Unit:              "C",
		UnitConversion:    "v * 10",
		DatasheetURL:      "http://example.com/datasheet-1.pdf",
		ExampleValues:     []string{"100", "200", "305"},
		ExamplesConverted: []string{"10", "20", "30.5"},
	},

	reg.Sensor{
		SensorID:          2,
		Symbol:            "abc123_pressure",
		Model:             "Bosch BME680",
		Name:              "Barometric Pressure",
		Description:       "Barometric pressure sensor for ABC123 device",
		Obsolete:          false,
		Type:              "int16",
		Unit:              "hPA",
		UnitConversion:    "v / 100",
		DatasheetURL:      "http://example.com/datasheet-1.pdf",
		ExampleValues:     []string{"30000", "50000", "110000"},
		ExamplesConverted: []string{"300", "500", "1100"},
	},

	reg.Sensor{
		SensorID:          3,
		Symbol:            "abc123_noise",
		Model:             "Unobtanium XYZ123",
		Name:              "Ambient noise",
		Description:       "Ambient noise as measured by totally non-creepy microphone",
		Obsolete:          false,
		Type:              "int16",
		Unit:              "dBFS",
		UnitConversion:    "(20 * log10(abs(v)+1)) / 32767",
		DatasheetURL:      "http://example.com/datasheet-2.pdf",
		ExampleValues:     []string{"1", "100", "1000", "10000"},
		ExamplesConverted: []string{"0.0001837397355046121", "0.0012233780167745859", "0.0018313755165131496", "0.0024415072647949844"},
	},
}
