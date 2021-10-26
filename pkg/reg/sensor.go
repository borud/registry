package reg

import (
	"errors"
	"net/url"
	"regexp"
)

// ValidSymbolRegex is the regular expression for symbols.
// Symbols are all lower case.  A symbol name may only start with a letter.
// Words are separated by underscores (_).  A symbol name cannot start with
// a number or an underscore.  It cannot end with an underscore.  These rules
// are to aid in ensuring we can generate valid symbols for different languages.
var ValidSymbolRegex = regexp.MustCompile(`^[[:lower:]]+[[:lower:]0-9_]+[[:lower:]0-9]$`)

var (
	// ErrInvalidSymbolName means the symbol name does not conform to the Symbol name format.
	ErrInvalidSymbolName = errors.New("symbol name does not conform to naming convention")
	// ErrSymbolNameEmpty means the symbol name was empty.
	ErrSymbolNameEmpty = errors.New("symbol name was empty")
	// ErrSensorIDZero zero sensorID is is not allowed.
	ErrSensorIDZero = errors.New("SensorID was zero")
	// ErrDataTypeNotSet means there was no data type specified.
	ErrDataTypeNotSet = errors.New("data type missing")
	// ErrInvalidDataType means that the data type given was invalid
	ErrInvalidDataType = errors.New("data type invalid")
	// ErrMaxLengthForNonVariableData means we tried to specify a length for data that is not
	// of variable length.
	ErrMaxLengthForNonVariableData = errors.New("max length set for data that is not variable length")
	// ErrExampleValuesMismatch means that the ExampleValues and ExampleConverted arrays were of different
	// length, indicating that they are out of sync.
	ErrExampleValuesMismatch = errors.New("different number of entries in ExampleValues and ExampleConverted")
)

// Sensor is the structure of a sensor registry entry.
type Sensor struct {
	// Globally unique sensor ID.  You can only add new sensor IDs, not delete
	// or re-use sensor ID that have been defined before.  If a sensor ID is
	// not to be used make sure you set the `Obsolete` flag to `true`.
	SensorID uint32 `json:"sensorId" db:"sensor_id"`

	// Symbol is the name that is used when generating source code.  See the ValidSymbolRegex
	// for information on what symbols are valid.
	Symbol string `json:"symbol" db:"symbol"`

	// Model is the name of the sensor.  For instance "BME680".  This is not intended for
	// usage in end-user interfaces, but is more a convenience for technical uses.
	Model string `json:"model" db:"name"`

	// Name is a human readable name for the sensor meant for use in user interfaces.  Do not
	// use this for symbol names etc since this is free-form.
	Name string `json:"name" db:"name"`

	// Free form description of the sensor.
	Description string `json:"description" db:"description"`

	// Since we can't delete a sensor ID that is still in the wild we need some
	// way of signaling that we don't want to use this anymore.  The "Obsolete" flag
	// means that you should not use this sensor ID, but we can't delete it since
	// there may be devices that still use it so backend systems still have to
	// know about this sensor.
	Obsolete bool `json:"obsolete" db:"obsolete"`

	// Type is a (string) representation of what data type the sensor emits.
	Type DataType `json:"type" db:"type"`

	// MaxLength is only defined for TypeByteArray and TypeString values
	// and denotes their maximum length.
	MaxLength int `json:"maxLength" db:"maxLength"`

	// Unit is the unit of measurement emitted by the device. Whenever possible
	// use the SI symbol https://en.wikipedia.org/wiki/International_System_of_Units
	Unit string `json:"unit" db:"unit"`

	// UnitConversion is the arithmetical expression to convert the
	// raw value from the sensor to the Unit above.  An empty string means
	// the raw value from the sensor needs no conversion.
	UnitConversion string `json:"referenceUnit" db:"reference_unit"`

	// The URL where we can download the datasheet from.  Please make sure that the
	// URLs used here do not require login if possible (so that we can download and
	// cache datasheets).
	DatasheetURL string `json:"datasheetUrl" db:"datasheet_url"`

	// ExampleValues is an array of the string representation of example values.
	// This is both for documentation and test purposes.  For each example value
	// there must be a corresponding converted value (by applying the formula in UnitConversion)
	// in the ExampleConverted array.
	ExampleValues []string `json:"exampleValues" db:"example_values"`

	// ExamplesConverted represents the conversion of each of the values in ExampleValues
	// by applying the formula in UnitConversion.
	ExamplesConverted []string `json:"exampleConverted" db:"example_converted"`
}

// Validate ensures the Sensor entry is valid
func (s Sensor) Validate() error {
	// zero is a really poor choice for a sensor ID
	if s.SensorID == 0 {
		return ErrSensorIDZero
	}

	// data type is required
	if s.Type == "" {
		return ErrDataTypeNotSet
	}

	// ...and it must be a data type we support
	if DataTypeByName(string(s.Type)) == "unknown" {
		return ErrInvalidDataType
	}

	// make sure the symbol is valid
	if !ValidSymbolRegex.MatchString(s.Symbol) {
		return ErrInvalidSymbolName
	}

	// Make sure max length is only set for variable length data
	if (s.MaxLength > 0) && !(s.Type == TypeString || s.Type == TypeByteArray) {
		return ErrMaxLengthForNonVariableData
	}

	if len(s.ExampleValues) != len(s.ExamplesConverted) {
		return ErrExampleValuesMismatch
	}

	// if URL is present it must be correct
	if s.DatasheetURL != "" {
		_, err := url.Parse(s.DatasheetURL)
		if err != nil {
			return err
		}
	}
	return nil
}
