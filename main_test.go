package main

import (
	"testing"
)

func TestPrompt_Validate(t *testing.T) {
	//Arrange
	var tests = []struct {
		name        string
		input       string
		shouldError bool
	}{
		{"name", "foobar", false},
		{"name with space", "foo bar", false},
		{"1 character name", "f", false},
		{"name with special characters", "foo-bar!", false},
		{"empty string", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Act
			err := validateProjectName(test.input)

			//Assert
			if test.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !test.shouldError && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
		})
	}
}
