package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{input: "SW25RR", expectedError: ""},
		{input: "SE153", expectedError: ""},
		{input: "SE", expectedError: "invalid postcode: not enough characters"},
		{input: "SE153RPP", expectedError: "invalid postcode: too many characters"},
	}

	for _, test := range tests {
		err := Validate(test.input)
		if test.expectedError != "" {
			assert.ErrorContains(t, err, test.expectedError)
		}
	}
}
