package sinclair

import (
	"backend/enum"
	"backend/structs"
	"strings"
	"testing"
)

func TestCalcSinclair(t *testing.T) {
	type args struct {
		result *structs.Lift
	}
	tests := []struct {
		name             string
		args             args
		expectedSinclair float32
	}{
		{
			name:             "NormalSinclairMalePre2001",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "1993-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 261.68857,
		},
		{
			name:             "NormalSinclairMale2001",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "2001-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 261.68857,
		},
		{
			name:             "NormalSinclairMale2021",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 298.24963,
		},
		{
			name:             "NormalSinclairMale2017",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "2017-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 285.66986,
		},
		{
			name:             "NormalSinclairFemale2021",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: 270.42316,
		},
		{
			name:             "NormalSinclairFemale2017",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(81), Total: structs.NewWeightKg(235), Sinclair: 0, Event: &structs.Event{Date: "2020-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: 270.17587,
		},
		{
			name:             "Over-rangeSinclairMale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 0,
		},
		{
			name:             "Over-rangeSinclairFemale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(160), Total: structs.NewWeightKg(510), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: 0,
		},
		{
			name:             "SuperHeavySinclairMale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: 400,
		},
		{
			name:             "SuperHeavySinclairFemale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewWeightKg(200), Total: structs.NewWeightKg(400), Sinclair: 0, Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcSinclair(tt.args.result)
			gotSinclair := float32(tt.args.result.Sinclair)
			switch {
			case strings.Contains(tt.name, "NormalSinclair") || strings.Contains(tt.name, "SuperHeavy"):
				if gotSinclair != tt.expectedSinclair {
					t.Errorf("CalcSinclair(): Normal Sinclair - got = %v, want %v", gotSinclair, tt.expectedSinclair)
				}
			case strings.Contains(tt.name, "Over-rangeSinclair"):
				if gotSinclair > tt.expectedSinclair {
					t.Errorf("CalcSinclair(): Over-range Sinclair - got = %v, want %v", gotSinclair, tt.expectedSinclair)
				}
			}
		})
	}
}
