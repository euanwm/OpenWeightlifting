package lifter

import "testing"

func TestCheckUserList(t *testing.T) {
	tests := []struct {
		testName   string
		lifterName string
		want       string
	}{
		{testName: "Expected Match", lifterName: "Euan Meston", want: "scream_and_jerk"},
		{testName: "No User Exists", lifterName: "Jesus Christ", want: ""},
		{"No Match on Case", "euan meston", ""},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := CheckUserList(tt.lifterName); got != tt.want {
				t.Errorf("CheckUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}
