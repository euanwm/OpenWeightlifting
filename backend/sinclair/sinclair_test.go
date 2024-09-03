package sinclair

import (
	"backend/structs"
	"testing"
)

func TestCalcSinclair(t *testing.T) {
	type args struct {
		result *structs.Entry
		male   bool
	}
	tests := []struct {
		name             string
		args             args
		expectedSinclair float32
	}{
		{
			name:             "NormalSinclairMale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0}, male: true},
			expectedSinclair: 285.66986,
		},
		{
			name:             "Over-rangeSinclairMale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0}, male: true},
			expectedSinclair: 0,
		},
		{
			name:             "NormalSinclairFemale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0}, male: false},
			expectedSinclair: 270.17587,
		},
		{
			name:             "Over-rangeSinclairFemale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0}, male: false},
			expectedSinclair: 0,
		},
		{
			name:             "SuperHeavySinclairMale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0}, male: true},
			expectedSinclair: 400,
		},
		{
			name:             "SuperHeavySinclairFemale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0}, male: false},
			expectedSinclair: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcSinclair(tt.args.result, tt.args.male)
			switch tt.name {
			case "Normal Sinclair":
				if tt.args.result.Sinclair != tt.expectedSinclair {
					t.Errorf("CalcSinclair(male=%t): Normal Sinclair - got = %v, want %v", tt.args.male, tt.args.result.Sinclair, tt.expectedSinclair)
				}
			case "Over-range Sinclair":
				if tt.args.result.Sinclair > tt.expectedSinclair {
					t.Errorf("CalcSinclair(male=%t): Over-range Sinclair - got = %v, want %v", tt.args.male, tt.args.result.Sinclair, tt.expectedSinclair)
				}
			}
		})
	}
}
