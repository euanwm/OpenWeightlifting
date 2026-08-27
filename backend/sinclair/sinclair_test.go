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
		expectedSinclair structs.FixedFloat
	}{
		{
			name:             "NormalSinclairMalePre2001",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "1993-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(261.68),
		},
		{
			name:             "NormalSinclairMale2001",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2001-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(261.68857),
		},
		{
			name:             "NormalSinclairMale2021",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(298.24963),
		},
		{
			name:             "NormalSinclairMale2017",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2017-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(285.66986),
		},
		{
			name:             "NormalSinclairFemale2021",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: structs.NewFixedFloat(270.42316),
		},
		{
			name:             "NormalSinclairFemale2017",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(81), Total: structs.NewFixedFloat(235), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2020-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: structs.NewFixedFloat(270.17587),
		},
		{
			name:             "Over-rangeSinclairMale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(160), Total: structs.NewFixedFloat(510), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(0),
		},
		{
			name:             "Over-rangeSinclairFemale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(160), Total: structs.NewFixedFloat(510), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: structs.NewFixedFloat(0),
		},
		{
			name:             "SuperHeavySinclairMale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(200), Total: structs.NewFixedFloat(400), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Male}}},
			expectedSinclair: structs.NewFixedFloat(400.00),
		},
		{
			name:             "SuperHeavySinclairFemale",
			args:             args{result: &structs.Lift{Bodyweight: structs.NewFixedFloat(200), Total: structs.NewFixedFloat(400), Sinclair: structs.NewFixedFloat(0), Event: &structs.Event{Date: "2021-04-01"}, Lifter: &structs.Lifter{Gender: enum.Female}}},
			expectedSinclair: structs.NewFixedFloat(400.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcSinclair(tt.args.result)
			gotSinclair := tt.args.result.Sinclair
			switch {
			case strings.Contains(tt.name, "NormalSinclair") || strings.Contains(tt.name, "SuperHeavy"):
				if !tt.args.result.Sinclair.Equal(tt.expectedSinclair) {
					t.Errorf("CalcSinclair(): Normal Sinclair - got = %v, want %v", gotSinclair, tt.expectedSinclair)
				}
			case strings.Contains(tt.name, "Over-rangeSinclair"):
				if tt.args.result.Sinclair.GreaterThan(tt.expectedSinclair) {
					t.Errorf("CalcSinclair(): Over-range Sinclair - got = %v, want %v", gotSinclair, tt.expectedSinclair)
				}
			}
		})
	}
}
