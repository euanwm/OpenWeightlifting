package sinclair

import (
	"backend/structs"
	"strings"
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
			name:             "NormalSinclairMalePre2001",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "1993-04-01"}, male: true},
			expectedSinclair: 261.68857,
		},
		{
			name:             "NormalSinclairMale2001",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "2001-04-01"}, male: true},
			expectedSinclair: 261.68857,
		},
		{
			name:             "NormalSinclairMale2021",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "2021-04-01"}, male: true},
			expectedSinclair: 298.24963,
		},
		{
			name:             "NormalSinclairMale2017",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "2017-04-01"}, male: true},
			expectedSinclair: 285.66986,
		},
		{
			name:             "NormalSinclairFemale2021",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "2021-04-01"}, male: false},
			expectedSinclair: 270.42316,
		},
		{
			name:             "NormalSinclairFemale2017",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Date: "2020-04-01"}, male: false},
			expectedSinclair: 270.17587,
		},
		{
			name:             "Over-rangeSinclairMale",
      args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0, Date: "2021-04-01"}, male: true},
			expectedSinclair: 0,
		},
		{
			name:             "Over-rangeSinclairFemale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0, Date: "2021-04-01"}, male: false},
			expectedSinclair: 0,
		},
		{
			name:             "SuperHeavySinclairMale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0, Date: "2021-04-01"}, male: true},
			expectedSinclair: 400,
		},
		{
			name:             "SuperHeavySinclairFemale",
			args:             args{result: &structs.Entry{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0, Date: "2021-04-01"}, male: false},
			expectedSinclair: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcSinclair(tt.args.result, tt.args.male)
			switch {
			case strings.Contains(tt.name, "NormalSinclair") || strings.Contains(tt.name, "SuperHeavy"):
				if tt.args.result.Sinclair != tt.expectedSinclair {
					t.Errorf("CalcSinclair(male=%t): Normal Sinclair - got = %v, want %v", tt.args.male, tt.args.result.Sinclair, tt.expectedSinclair)
				}
			case strings.Contains(tt.name, "Over-rangeSinclair"):
				if tt.args.result.Sinclair > tt.expectedSinclair {
					t.Errorf("CalcSinclair(male=%t): Over-range Sinclair - got = %v, want %v", tt.args.male, tt.args.result.Sinclair, tt.expectedSinclair)
				}
			}
		})
	}
}
