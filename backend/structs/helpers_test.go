package structs

import (
	"backend/enum"
	"reflect"
	"testing"
)

func TestIterateFloatSlice(t *testing.T) {
	type args struct {
		data []Entry
		item string
	}
	tests := []struct {
		name        string
		args        args
		wantFloatSl []float32
	}{
		{name: "IterateTotals", args: args{data: []Entry{{Total: NewWeightKg(1)}, {Total: NewWeightKg(2)}, {Total: NewWeightKg(3)}}, item: enum.Total}, wantFloatSl: []float32{1, 2, 3}},
		{name: "IterateBestSnatch", args: args{data: []Entry{{BestSn: NewWeightKg(1)}, {BestSn: NewWeightKg(2)}, {BestSn: NewWeightKg(3)}}, item: enum.BestSnatch}, wantFloatSl: []float32{1, 2, 3}},
		{name: "IterateBestCJ", args: args{data: []Entry{{BestCJ: NewWeightKg(1)}, {BestCJ: NewWeightKg(2)}, {BestCJ: NewWeightKg(3)}}, item: enum.BestCJ}, wantFloatSl: []float32{1, 2, 3}},
		{name: "IterateBodyweight", args: args{data: []Entry{{Bodyweight: NewWeightKg(1)}, {Bodyweight: NewWeightKg(2)}, {Bodyweight: NewWeightKg(3)}}, item: enum.Bodyweight}, wantFloatSl: []float32{1, 2, 3}},
		{name: "IterateNothing", args: args{data: []Entry{{Bodyweight: NewWeightKg(1)}, {Bodyweight: NewWeightKg(2)}, {Bodyweight: NewWeightKg(3)}}, item: ""}, wantFloatSl: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFloatSl := IterateFloatSlice(tt.args.data, tt.args.item); !reflect.DeepEqual(gotFloatSl, tt.wantFloatSl) {
				t.Errorf("IterateFloatSlice() = %v, want %v", gotFloatSl, tt.wantFloatSl)
			}
		})
	}
}
