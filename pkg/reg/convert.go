package reg

// note that some functions were borrowed from
// https://github.com/jamillosantos/functions-for-govaluate

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/Knetic/govaluate"
)

// Unit represents from raw data to the target unit representation.
// The expression language is documented here: https://github.com/Knetic/govaluate/blob/master/MANUAL.md
type Unit struct {
	expression string
	expr       *govaluate.EvaluableExpression
	inputType  DataType
}

var (
	// ErrWrongParamCount indicates that the function got the wrong parameter count
	ErrWrongParamCount = errors.New("wrong parameter count")

	// ErrWrongParamType indecates that the function got the wrong parameter type
	ErrWrongParamType = errors.New("wrong parameter type")
)

var functionMap = map[string]govaluate.ExpressionFunction{
	"log10": functionLog10,
	"abs":   functionAbs,
}

// NewConversionFromSensor
func NewConversionFromSensor(sensor Sensor) (*Unit, error) {
	err := sensor.Validate()
	if err != nil {
		return nil, err
	}
	return NewConversion(sensor.UnitConversion, sensor.Type)
}

// NewConversion creates a new conversion.
func NewConversion(expression string, inputType DataType) (*Unit, error) {
	if DataTypeByName(string(inputType)) == TypeUnknown {
		return nil, fmt.Errorf("input type '%s' unknown", inputType)
	}

	// if the expression is blank it represents the identity expression so
	// we just modify expression to reflect that
	if expression == "" {
		expression = "v"
	}

	expr, err := govaluate.NewEvaluableExpressionWithFunctions(expression, functionMap)
	if err != nil {
		return nil, err
	}

	return &Unit{
		expression: expression,
		expr:       expr,
		inputType:  inputType,
	}, nil
}

// Convert value and return as interface{}, most likely float64 for arithmetic
func (c *Unit) Convert(v interface{}) (interface{}, error) {
	if DataTypeFromValue(v) != c.inputType {
		return float64(0), fmt.Errorf("expected input type '%s', got '%s'", c.inputType, reflect.TypeOf(v))
	}
	return c.expr.Evaluate(map[string]interface{}{"v": v})
}

// ConvertString value and return as string
func (c *Unit) ConvertString(v interface{}) (string, error) {
	if DataTypeFromValue(v) != c.inputType {
		return "", fmt.Errorf("expected input type '%s', got '%s'", c.inputType, reflect.TypeOf(v))
	}

	res, err := c.expr.Evaluate(map[string]interface{}{"v": v})
	return fmt.Sprintf("%v", res), err
}

func functionLog10(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := toFloat64(args[0])
		if !ok {
			return math.NaN(), ErrWrongParamType
		}
		return math.Log10(v), nil
	}
	return nil, ErrWrongParamCount
}

func functionAbs(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := toFloat64(args[0])
		if !ok {
			return math.NaN(), ErrWrongParamType
		}
		return math.Abs(v), nil
	}
	return nil, ErrWrongParamCount
}

func toFloat64(v interface{}) (float64, bool) {
	switch vv := v.(type) {
	case float64:
		return vv, true
	case float32:
		return float64(vv), true
	case int8:
		return float64(vv), true
	case int32:
		return float64(vv), true
	case int64:
		return float64(vv), true
	case uint8:
		return float64(vv), true
	case uint32:
		return float64(vv), true
	case uint64:
		return float64(vv), true
	case int:
		return float64(vv), true
	case uint:
		return float64(vv), true
	default:
		return 0, false
	}
}
