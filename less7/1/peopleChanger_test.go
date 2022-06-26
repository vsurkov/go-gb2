package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testMan = Person{
		Name:   "bob",
		Age:    22,
		Weight: 88.1,
		Male:   true,
	}

	changes = map[string]interface{}{
		"Name":   "jane",
		"Age":    22,
		"Weight": 90.9,
		"Male":   false,
	}
)

func TestNulValue(t *testing.T) {
	expectErr := errors.New("impute is nil")
	receivedErr := peopleChanger(nil, changes)
	assert.Equal(t, expectErr, receivedErr, "Errors must be equal")
}

func TestNulInput(t *testing.T) {
	var expectErr error = nil
	receivedErr := peopleChanger(&testMan, nil)
	assert.Equal(t, expectErr, receivedErr, "Errors must be equal")
}

func TestNulBoth(t *testing.T) {
	expectErr := errors.New("impute is nil")
	receivedErr := peopleChanger(nil, nil)
	assert.Equal(t, expectErr, receivedErr, "Errors must be equal")
}
