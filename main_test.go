package main

import "testing"

func TestGetPostcode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "SE153RP", output: "SE153RP"},
		{input: "SE153", output: "SE153"},
	}

	for _, test := range tests {
		postcode := GetPostcode(test.input)
		if postcode != test.output {
			t.Errorf("Postcode: %s was incorrect, got: %s, want: %s.", test.input, postcode, test.output)
		}
	}
}
