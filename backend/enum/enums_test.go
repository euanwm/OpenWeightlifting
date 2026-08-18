package enum

import "testing"

func TestClassifyGender(t *testing.T) {
	tests := []struct {
		name       string
		raw        string
		wantGender string
	}{
		{name: "DirectMatchMale", raw: Male, wantGender: Male},
		{name: "DirectMatchFemale", raw: Female, wantGender: Female},
		{name: "ContainsMen", raw: "Men's", wantGender: Male},
		{name: "ContainsWomen", raw: "Women's", wantGender: Female},
		{name: "CatchUnknown", raw: "something else", wantGender: Unknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGender := ClassifyGender(tt.raw); gotGender != tt.wantGender {
				t.Errorf("ClassifyGender() = %v, want %v", gotGender, tt.wantGender)
			}
		})
	}
}
