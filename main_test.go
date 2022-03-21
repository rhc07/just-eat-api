package main

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		input string
		error error
	}{
		{input: "SW25RR", error: nil},
		{input: "SE153", error: nil},
		{input: "SE15", error: errors.New("invalid postcode: not enough characters")},
		{input: "SE153RPP", error: errors.New("invalid postcode: too many characters")},
	}

	for _, test := range tests {
		err := Validate(test.input)
		if err != test.error {
			t.Errorf("Postcode: %s was incorrect, got: %s, want: %s.", test.input, err, test.error)
		}
	}
}
