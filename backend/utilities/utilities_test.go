package utilities

import (
	igDatabase "backend/lifter_data"
	"io"
	"io/fs"
	"log"
	"reflect"
	"testing"
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

func TestLoadCsvFile(t *testing.T) {
	fileHandle, err := igDatabase.InstagramDatabase.Open("ighandles.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(fileHandle fs.File) {
		err := fileHandle.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fileHandle)

	type args struct {
		file io.Reader
	}
	tests := []struct {
		name         string
		args         args
		wantedTypeOf [][]string
	}{
		{name: "LoadCsvFile", args: args{file: fileHandle}, wantedTypeOf: [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadCsvFile(tt.args.file); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.wantedTypeOf)) {
				t.Errorf("LoadCsvFile() = %v, want %v", got, tt.wantedTypeOf)
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
