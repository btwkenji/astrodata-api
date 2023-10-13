package core

import (
	"github.com/kenjitheman/astrodata-api/vars"
	"testing"
)

func TestCalculateDst(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0}
	expected := 1.0 * vars.CONSTANT

	result := CalculateDst(data)

	if result != expected {
		t.Errorf("[!] expected CalculateDst(%v) to be %f, but got %f", data, expected, result)
	}
}

func TestSplitData(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	expected := [][]float64{}

	result := SplitData(data)

	if !slicesAreEqualTwo(result, expected) {
		t.Errorf("[!] expected SplitData(%v) to be %v, but got %v", data, expected, result)
	}
}

func TestCalculateDstForParts(t *testing.T) {
	parts := [][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
		{10.0, 11.0, 12.0},
	}
	expected := []float64{1.0, 4.0, 7.0, 10.0}
	for i, v := range expected {
		expected[i] = v * vars.CONSTANT
	}

	result := CalculateDstForParts(parts)

	if !slicesAreEqualOne(result, expected) {
		t.Errorf("[!] expected CalculateDstForParts(%v) to be %v, but got %v", parts, expected, result)
	}
}

func slicesAreEqualTwo(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, slice := range a {
		if !slicesAreEqualOne(slice, b[i]) {
			return false
		}
	}
	return true
}

func slicesAreEqualOne(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
