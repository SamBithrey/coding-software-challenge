package helpers

import "testing"

func TestCalculateArea(t *testing.T) {
	type args struct {
		width  float64
		length float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Normal", args{3, 4}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateArea(tt.args.width, tt.args.length); got != tt.want {
				t.Errorf("CalculateArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
