package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	testMan = Person{
		Name:   "bob",
		Age:    20,
		Weight: 88.1,
		Male:   true,
	}
	testWoman = Person{
		Name:   "jane",
		Age:    22,
		Weight: 90.9,
		Male:   false,
	}

	changes = map[string]interface{}{
		"Name":   "jane",
		"Age":    22,
		"Weight": 90.9,
		"Male":   false,
	}
)

func TestNulValue(t *testing.T) {
	exp := errors.New("impute is nil")
	rec := peopleChanger(nil, changes)
	assert.Equal(t, exp, rec, "Errors must be equal")
}

func TestNulInput(t *testing.T) {
	var exp error = nil
	rec := peopleChanger(&testMan, nil)
	assert.Equal(t, exp, rec, "Errors must be equal")
}

func TestNulBoth(t *testing.T) {
	exp := errors.New("impute is nil")
	rec := peopleChanger(nil, nil)
	assert.Equal(t, exp, rec, "Errors must be equal")
}

func TestPositive(t *testing.T) {
	remade := testMan
	peopleChanger(&remade, changes)

	assert.Equal(t, testWoman, remade)
	assert.Equal(t, true, reflect.DeepEqual(testWoman, remade))

}

func TestWrongStringType(t *testing.T) {
	strVal := map[string]interface{}{
		"Name": 10,
	}
	exp := errors.New("in value Name mismatched types, expected: string - received int\n")
	remade := testMan
	rec := peopleChanger(&remade, strVal)
	assert.Equal(t, exp, rec)
}

func TestWrongIntType(t *testing.T) {
	intVal := map[string]interface{}{
		"Age": "foo",
	}

	exp := errors.New("in value Age mismatched types, expected: int - received string\n")
	remade := testMan
	rec := peopleChanger(&remade, intVal)
	assert.Equal(t, exp, rec)
}

func TestWrongFloatType(t *testing.T) {
	floatVal := map[string]interface{}{
		"Weight": true,
	}

	exp := errors.New("in value Weight mismatched types, expected: float64 - received bool\n")
	remade := testMan
	rec := peopleChanger(&remade, floatVal)
	assert.Equal(t, exp, rec)
}

func TestWrongBoolType(t *testing.T) {
	boolVal := map[string]interface{}{
		"Male": 3.1415,
	}

	exp := errors.New("in value Male mismatched types, expected: bool - received float64\n")
	remade := testMan
	rec := peopleChanger(&remade, boolVal)
	assert.Equal(t, exp, rec)
}
