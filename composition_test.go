package main

import (
	"errors"
	"testing"
)

func TestAddOne(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		err      error
	}{
		{1, 2, nil},
		{2, 3, nil},
		{0, 0, errors.New("x cannot be 0")},
	}

	for _, tt := range tests {
		result, err := AddOne(tt.input)
		if result != tt.expected || (err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("AddOne(%d) = %d, %v; want %d, %v", tt.input, result, err, tt.expected, tt.err)
		}
	}
}

func TestDouble(t *testing.T) {
	tests := []struct {
		input    int
		expected int
		err      error
	}{
		{1, 2, nil},
		{2, 4, nil},
		{0, 0, errors.New("x cannot be 0")},
	}

	for _, tt := range tests {
		result, err := Double(tt.input)
		if result != tt.expected || (err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("Double(%d) = %d, %v; want %d, %v", tt.input, result, err, tt.expected, tt.err)
		}
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		f        functionAsType
		g        functionAsType
		expected int
	}{
		{
			func() (int, error) { return AddOne(1) },
			func() (int, error) { return Double(2) },
			3, // 2 + 4 = 6
		},
		{
			func() (int, error) { return AddOne(1) },
			func() (int, error) { return Double(0) },
			0, // Error in second function
		},
		{
			func() (int, error) { return AddOne(0) },
			func() (int, error) { return Double(2) },
			0, // Error in first function
		},
	}

	for _, tt := range tests {
		resultFunc := Compose(tt.f, tt.g)
		result := resultFunc()
		if result != tt.expected {
			t.Errorf("Compose() = %d; want %d", result, tt.expected)
		}
	}
}
