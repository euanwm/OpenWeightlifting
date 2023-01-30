package lifter

import "testing"

func TestCheckUserList(t *testing.T) {
	type args struct {
		lifterName string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{name: "Expected Match", args: args{lifterName: "Euan Meston"}, want: true, want1: "scream_and_jerk"},
		{"No User Exists", args{lifterName: "Jesus Christ"}, false, ""},
		{"No Match on Case", args{lifterName: "euan meston"}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckUserList(tt.args.lifterName)
			if got != tt.want {
				t.Errorf("CheckUserList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckUserList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
