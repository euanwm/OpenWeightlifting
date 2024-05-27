// A Go port of OpenPowerlifting's WeightKg type.

package structs

import (
	"fmt"
	"math"
	"strconv"
)

// A weight in kilograms represented as a fixed-point integer.
// The integer representation holds two decimal places, such that
// the floating-point value "123.45" is stored as `12345`. Values
// that cannot be exactly represented round toward zero.
type WeightKg struct {
	value int32
}

// Returns a new WeightKg from a floating-point value.
// Values that cannot be exactly represented round toward zero.
// Infinite or NaN inputs are treated as zero.
func NewWeightKg(v float64) WeightKg {
	if math.IsInf(v, 0) || math.IsNaN(v) {
		return WeightKg{value: 0}
	}

	is_signed := v < 0                  // -0 is treated identically to 0.
	v = math.Floor(math.Abs(v) * 100.0) // Shift two decimal places left and truncate.

	i := int32(v)
	if is_signed {
		i = -i
	}
	return WeightKg{value: i}
}

// Returns a new WeightKg from a string value.
// Values that cannot be parsed return zero.
func NewWeightKgFromString(s string) WeightKg {
	// Explicitly allow writing the empty string instead of zero.
	if len(s) == 0 {
		return WeightKg{0}
	}

	// Otherwise, expect a floating-point value.
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return WeightKg{0}
	}
	return NewWeightKg(float)
}

// Returns a new WeightKg from an integer weight.
// This is mostly useful for values that are derived from enums.
func NewWeightKgFromInt32(i int32) WeightKg {
	return WeightKg{i * 100}
}

// Returns whether both weights are equal.
func (kg WeightKg) Equal(other WeightKg) bool {
	return kg.value == other.value
}

// Returns whether kg > other.
func (kg WeightKg) GreaterThan(other WeightKg) bool {
	return kg.value > other.value
}

// Returns whether kg >= other.
func (kg WeightKg) GreaterThanOrEqual(other WeightKg) bool {
	return kg.value >= other.value
}

// Returns whether kg < other.
func (kg WeightKg) LessThan(other WeightKg) bool {
	return kg.value < other.value
}

// Returns whether kg <= other.
func (kg WeightKg) LessThanOrEqual(other WeightKg) bool {
	return kg.value <= other.value
}

// Returns -1 if negative, 0 if zero, +1 if positive.
func (kg WeightKg) Sign() int {
	if kg.value > 0 {
		return 1
	}
	if kg.value < 0 {
		return -1
	}
	return 0
}

func (kg WeightKg) IsPositive() bool {
	return kg.value > 0
}

func (kg WeightKg) IsNegative() bool {
	return kg.value < 0
}

// Returns whether the weight is zero.
func (kg WeightKg) IsZero() bool {
	return kg.value == 0
}

// Returns the minimum of the two WeightKgs.
func (kg WeightKg) Min(other WeightKg) WeightKg {
	if kg.LessThan(other) {
		return kg
	}
	return other
}

// Returns the maximum of the two WeightKgs.
func (kg WeightKg) Max(other WeightKg) WeightKg {
	if kg.GreaterThan(other) {
		return kg
	}
	return other
}

// Returns the nearest float32 value.
func (kg WeightKg) Float32() float32 {
	return float32(kg.value) / 100
}

// Returns the nearest float64 value.
func (kg WeightKg) Float64() float64 {
	return float64(kg.value) / 100
}

// Renders the WeightKg as a string, looking like a floating-point number.
// Decimal places are rendered with as few zeros as possible.
//
// Examples:
// - input 123.00 returns "123".
// - input 123.40 returns "123.4".
// - input 123.45 returns "123.45"
func (kg WeightKg) String() string {
	// Fast path for the common zero value.
	if kg.value == 0 {
		return "0"
	}

	// For purposes of the later modulo, store a non-negative representation.
	non_negative := kg.value
	if non_negative < 0 {
		non_negative = -non_negative
	}

	integer := kg.value / 100
	fraction := non_negative % 100

	// Render the integer component, which can include a negative sign.
	acc := strconv.Itoa(int(integer))

	// Inspect the remaining fractional component.
	if fraction == 0 {
		return acc // No fractional component, so return the rendered integer.
	}
	if fraction%10 == 0 {
		return acc + "." + strconv.Itoa(int(fraction/10)) // Render "50" as ".5".
	}
	return acc + "." + fmt.Sprintf("%02d", fraction) // Render left-padded with '0' to two places.
}

// JSON deserialization.
func (kg *WeightKg) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}
	*kg = NewWeightKgFromString(string(bytes))
	return nil
}

// JSON serialization.
func (kg WeightKg) MarshalJSON() ([]byte, error) {
	return []byte(kg.String()), nil
}
