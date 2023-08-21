package helpers

import "testing"

func TestCalculateVolume(t *testing.T) {
	type args struct {
		height float64
		width  float64
		length float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Calculate volume", args{4, 10, 20}, 800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateVolume(tt.args.height, tt.args.width, tt.args.length); got != tt.want {
				t.Errorf("CalculateVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
