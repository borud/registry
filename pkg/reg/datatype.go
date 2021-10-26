package reg

// DataType describes the datatype of what a sensor emits.
type DataType string

const (
	// TypeUnknown invalid data type
	TypeUnknown = DataType("unknown")
	// TypeUint8 is a 8 bit unsigned int
	TypeUint8 = DataType("uint8")
	// TypeUint16 is a 16 bit unsigned int
	TypeUint16 = DataType("uint16")
	// TypeUint32 is a 32 bit unsigned int
	TypeUint32 = DataType("uint32")
	// TypeUint64 is a 64 bit unsigned int
	TypeUint64 = DataType("uint64")
	// TypeInt8 is a 8 bit signed int
	TypeInt8 = DataType("int8")
	// TypeInt16 is a 16 bit signed int
	TypeInt16 = DataType("int16")
	// TypeInt32 is a 32 bit signed int
	TypeInt32 = DataType("int32")
	// TypeInt64 is a 64 bit signed int
	TypeInt64 = DataType("int64")
	// TypeFloat32 is a 32 bit float
	TypeFloat32 = DataType("float32")
	// TypeFloat64 is a 64 bit float
	TypeFloat64 = DataType("float64")
	// TypeString is a string type
	TypeString = DataType("string")
	// TypeByteArray is a byte array
	TypeByteArray = DataType("bytes")
	// TypeBool is a boolean value
	TypeBool = DataType("bool")
)

var (
	typeNames = map[string]DataType{
		"unknown": TypeUnknown,
		"uint8":   TypeUint8,
		"uint16":  TypeUint16,
		"uint32":  TypeUint32,
		"uint64":  TypeUint64,
		"int8":    TypeInt8,
		"int16":   TypeInt16,
		"int32":   TypeInt32,
		"int64":   TypeInt64,
		"float32": TypeFloat32,
		"float64": TypeFloat64,
		"string":  TypeString,
		"bytes":   TypeByteArray,
		"bool":    TypeBool,
	}

	goDataTypes = map[DataType]string{
		TypeUint8:     "uint8",
		TypeUint16:    "uint16",
		TypeUint32:    "uint32",
		TypeUint64:    "uint64",
		TypeInt8:      "int8",
		TypeInt16:     "int16",
		TypeInt32:     "int32",
		TypeInt64:     "int64",
		TypeFloat32:   "float32",
		TypeFloat64:   "float64",
		TypeString:    "string",
		TypeByteArray: "[]byte",
		TypeBool:      "bool",
	}

	cDataTypes = map[DataType]string{
		TypeUint8:     "t_uint8",
		TypeUint16:    "t_uint16",
		TypeUint32:    "t_uint32",
		TypeUint64:    "t_uint64",
		TypeInt8:      "t_int8",
		TypeInt16:     "t_int16",
		TypeInt32:     "t_int32",
		TypeInt64:     "t_int64",
		TypeFloat32:   "float",
		TypeFloat64:   "double",
		TypeString:    "char*",
		TypeByteArray: "byte[]",
		TypeBool:      "bool",
	}
)

// DataTypeByName returns the datatype by name.
func DataTypeByName(s string) DataType {
	t, ok := typeNames[s]
	if !ok {
		return TypeUnknown
	}
	return t
}

// GoTypeName returns the Go type name for the data type.
func (d *DataType) GoTypeName() string {
	t, ok := goDataTypes[*d]
	if !ok {
		return string(TypeUnknown)
	}
	return t
}

// CTypeName returns the C type name for the data type.
func (d *DataType) CTypeName() string {
	t, ok := cDataTypes[*d]
	if !ok {
		return string(TypeUnknown)
	}
	return t
}

// DataTypeFromValue returns the data type name for value.
func DataTypeFromValue(d interface{}) DataType {
	switch d.(type) {
	case uint8:
		return TypeUint8
	case uint16:
		return TypeUint16
	case uint32:
		return TypeUint32
	case uint64:
		return TypeUint64
	case int8:
		return TypeInt8
	case int16:
		return TypeInt16
	case int32:
		return TypeInt32
	case int64:
		return TypeInt64
	case float32:
		return TypeFloat32
	case float64:
		return TypeFloat64
	case string:
		return TypeString
	case []byte:
		return TypeByteArray
	case bool:
		return TypeBool
	default:
		return TypeUnknown
	}
}
