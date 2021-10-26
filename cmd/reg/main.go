package main

import (
	"log"
	"reflect"

	"github.com/Knetic/govaluate"
)

func main() {
	log.Printf("nothing here yet")

	expr, err := govaluate.NewEvaluableExpression("v / 10.0")
	if err != nil {
		log.Fatalf("error creating expression: %v", err)
	}

	result, err := expr.Evaluate(map[string]interface{}{"v": 10.00})
	if err != nil {
		log.Fatalf("error evaluating expression: %v", err)
	}

	log.Printf("result: %+v (%s)", result, reflect.TypeOf(result))
}
