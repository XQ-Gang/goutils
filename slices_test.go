package goutils

import (
	"reflect"
	"testing"
)

func TestGetBatches(t *testing.T) {
	tests := []struct {
		slice     []string
		batchSize int
		want      [][]string
	}{
		{
			slice:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			batchSize: 1,
			want:      [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}, {"7"}, {"8"}, {"9"}},
		},
		{
			slice:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			batchSize: 2,
			want:      [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}, {"7", "8"}, {"9"}},
		},
		{
			slice:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
			batchSize: 3,
			want:      [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
		},
	}
	for _, tt := range tests {
		got := GetBatches(tt.slice, tt.batchSize)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetBatches(%v, %d) = %v, want %v", tt.slice, tt.batchSize, got, tt.want)
		}
	}
}
