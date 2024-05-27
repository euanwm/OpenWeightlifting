package structs

import (
	"encoding/json"
	"math"
	"testing"
)

type newCase struct {
	input    float64 // Testcase input as a float64.
	expected int32   // The expected internal representation after conversion.
}

// Tests that WeightKgNew properly converts float64 input values.
func TestWeightKg_New(t *testing.T) {
	cases := []newCase{
		// Sanity checking.
		{input: -0.0, expected: 0},
		{input: +0.0, expected: 0},
		{input: -1.0, expected: -100},
		{input: +1.0, expected: 100},
		{input: +123.45, expected: 12345},
		{input: -123.45, expected: -12345},

		// Rounding cases.
		{input: 1.0 / 3.0, expected: 33},
		{input: +123.4567, expected: 12345},
		{input: -123.4567, expected: -12345},

		// Invalid inputs that still need to be handled.
		{input: math.NaN(), expected: 0},
		{input: math.Inf(1), expected: 0},
		{input: math.Inf(-1), expected: 0},
	}
	for _, c := range cases {
		result := NewWeightKg(c.input).value
		if result != c.expected {
			t.Errorf("expected %d, got %d with input %f", c.expected, result, c.input)
		}
	}
}

type stringCase struct {
	input    int32  // The internal representation held in WeightKg.value.
	expected string // The expected output of WeightKg.String().
}

// Tests that WeightKg.String formats strings according to spec.
func TestWeightKg_String(t *testing.T) {
	cases := []stringCase{
		{input: 0, expected: "0"},

		// Decimal formatting with positive weights.
		{input: 12300, expected: "123"},
		{input: 12340, expected: "123.4"},
		{input: 12345, expected: "123.45"},
		{input: 12305, expected: "123.05"},
		{input: 12005, expected: "120.05"},

		// Decimal formatting with negative weights.
		{input: -12300, expected: "-123"},
		{input: -12340, expected: "-123.4"},
		{input: -12345, expected: "-123.45"},
		{input: -12305, expected: "-123.05"},
		{input: -12005, expected: "-120.05"},
	}
	for _, c := range cases {
		kg := WeightKg{c.input}
		result := kg.String()
		if result != c.expected {
			t.Errorf("expected %s, got %s with input %d", c.expected, result, c.input)
		}
	}
}

type jsonTest struct {
	MyKg WeightKg `json:"mykg"`
}

// Tests that WeightKg behaves like a float64 when serialized/deserialized to/from JSON.
func TestWeightKg_Json(t *testing.T) {
	data := jsonTest{NewWeightKg(123.45)}

	// Serialize to JSON. It should serialize as a float64.
	jsonData, err := json.Marshal(&data)
	if err != nil {
		t.Fatalf("Failed marshaling WeightKg to JSON: %v", err)
	}

	expected := `{"mykg":123.45}`
	if string(jsonData) != expected {
		t.Errorf("Error serializing: expected %s, got %s", expected, string(jsonData))
	}

	// Deserialize back from the serialized JSON. It should match the original.
	var parsed jsonTest
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		t.Fatalf("Failed unmarshaling WeightKg from JSON: %v", err)
	}

	if data != parsed {
		t.Errorf("Error deserializing: expected %+v, got %+v", data, parsed)
	}
}
