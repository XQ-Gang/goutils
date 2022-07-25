package goutils

import "testing"

func TestIf(t *testing.T) {
	tests := []struct {
		condition         bool
		trueVal, falseVal int
		want              int
	}{
		{
			condition: true,
			trueVal:   1,
			falseVal:  0,
			want:      1,
		},
		{
			condition: false,
			trueVal:   1,
			falseVal:  0,
			want:      0,
		},
	}
	for _, tt := range tests {
		got := If(tt.condition, tt.trueVal, tt.falseVal)
		if got != tt.want {
			t.Errorf("If(%t, %d, %d) = %v, want %v", tt.condition, tt.trueVal, tt.falseVal, got, tt.want)
		}
	}
}
