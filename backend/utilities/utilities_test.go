package utilities

import (
	"reflect"
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	type args struct {
		sl   []string
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Contains", args: args{sl: []string{"a", "b", "c"}, name: "b"}, want: true},
		{name: "DoesNotContain", args: args{sl: []string{"a", "b", "c"}, name: "d"}, want: false},
		{name: "EmptyList", args: args{sl: []string{}, name: "d"}, want: false},
		{name: "EmptyName", args: args{sl: []string{"a", "b", "c"}, name: ""}, want: false},
		{name: "EmptyListAndName", args: args{sl: []string{}, name: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.sl, tt.args.name); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat(t *testing.T) {
	type args struct {
		preFloatStr string
	}
	tests := []struct {
		name         string
		args         args
		wantRetFloat float32
	}{
		{name: "Float", args: args{preFloatStr: "1.0"}, wantRetFloat: 1.0},
		{name: "NotFloat", args: args{preFloatStr: "a"}, wantRetFloat: 0.0},
		{name: "EmptyString", args: args{preFloatStr: ""}, wantRetFloat: 0.0},
		{name: "FloatWithCommas", args: args{preFloatStr: "1,0"}, wantRetFloat: 0.0},
		{name: "Float64", args: args{preFloatStr: "0.12345678912121212"}, wantRetFloat: 0.12345679},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetFloat := Float(tt.args.preFloatStr); gotRetFloat != tt.wantRetFloat {
				t.Errorf("Float() = %v, want %v", gotRetFloat, tt.wantRetFloat)
			}
		})
	}
}

func TestMapContains(t *testing.T) {
	type args struct {
		StrQuery string
		mapData  map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "MapContains", args: args{StrQuery: "a", mapData: map[string]string{"a": "b"}}, want: true},
		{name: "MapDoesNotContain", args: args{StrQuery: "a", mapData: map[string]string{"b": "c"}}, want: false},
		{name: "EmptyMap", args: args{StrQuery: "a", mapData: map[string]string{}}, want: false},
		{name: "EmptyQuery", args: args{StrQuery: "", mapData: map[string]string{"a": "b"}}, want: false},
		{name: "EmptyMapAndQuery", args: args{StrQuery: "", mapData: map[string]string{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapContains(tt.args.StrQuery, tt.args.mapData); got != tt.want {
				t.Errorf("MapContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	type args struct {
		strQuery string
		slData   []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SliceContains", args: args{strQuery: "a", slData: []string{"a", "b"}}, want: true},
		{name: "SliceDoesNotContain", args: args{strQuery: "a", slData: []string{"b", "c"}}, want: false},
		{name: "EmptySlice", args: args{strQuery: "a", slData: []string{}}, want: false},
		{name: "EmptyQuery", args: args{strQuery: "", slData: []string{"a", "b"}}, want: false},
		{name: "EmptySliceAndQuery", args: args{strQuery: "", slData: []string{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContains(tt.args.strQuery, tt.args.slData); got != tt.want {
				t.Errorf("SliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name        string
		args        args
		wantRetDate time.Time
		wantErr     bool
	}{
		{name: "StringToDate", args: args{dateStr: "2020-01-01"}, wantRetDate: time.Date(2020, 1, 1, 15, 04, 05, 0, time.UTC), wantErr: false},
		{name: "StringToDateInvalid", args: args{dateStr: "notadate"}, wantRetDate: time.Time{}, wantErr: true},
		{name: "StringToDateEmpty", args: args{dateStr: ""}, wantRetDate: time.Time{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetDate, err := StringToDate(tt.args.dateStr)
			if !reflect.DeepEqual(gotRetDate, tt.wantRetDate) && (err != nil) != tt.wantErr {
				t.Errorf("StringToDate() gotRetDate = %v, want %v, wantErr = %v, want %v", gotRetDate, tt.wantRetDate, err, tt.wantErr)
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
	type args struct {
		slData []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "ReverseSlice", args: args{slData: []string{"a", "b"}}, want: []string{"b", "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSlice(tt.args.slData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
