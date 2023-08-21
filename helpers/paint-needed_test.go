package helpers

import "testing"

func TestCalculatePaint(t *testing.T) {
	type args struct {
		width   float64
		height  float64
		length  float64
		windows int
		doors   int
	}
	tests := []struct {
		name             string
		args             args
		wantMaximumPaint float64
		wantWallPaint    float64
	}{
		{"name", args{0, 0, 0, 0, 0}, 0, 0},
		{"Normal input", args{3, 3, 4, 1, 1}, 1.4, 1.34},
		{"No windows or doors", args{3, 3, 4, 0, 0}, 1.4, 1.4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMaximumPaint, gotWallPaint := CalculatePaint(tt.args.width, tt.args.height, tt.args.length, tt.args.windows, tt.args.doors)
			if gotMaximumPaint != tt.wantMaximumPaint {
				t.Errorf("CalculatePaint() gotMaximumPaint = %v, want %v", gotMaximumPaint, tt.wantMaximumPaint)
			}
			if gotWallPaint != tt.wantWallPaint {
				t.Errorf("CalculatePaint() gotWallPaint = %v, want %v", gotWallPaint, tt.wantWallPaint)
			}
		})
	}
}
