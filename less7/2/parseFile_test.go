package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	exp := 2
	rec, err := parseFile("parseFile.go", "parseFile")
	assert.NoError(t, err)
	assert.Equal(t, exp, rec, "return must be 2")
}
