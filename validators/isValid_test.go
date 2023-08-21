package validators

import "testing"

func TestValidateInput(t *testing.T) {
	type args struct {
		width   float64
		height  float64
		length  float64
		windows int
		doors   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"All args correct", args{10.0, 10.0, 10.0, 2, 1}, true},
		{"Width must be great than 0", args{0, 10.0, 10.0, 2, 1}, false},
		{"Height must be great than 0", args{10.0, 0, 10.0, 2, 1}, false},
		{"Length must be great than 0", args{10.0, 10.0, 0, 2, 1}, false},
		{"Doors and windows can be 0", args{10.0, 10.0, 10.0, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateInput(tt.args.width, tt.args.height, tt.args.length, tt.args.windows, tt.args.doors); got != tt.want {
				t.Errorf("ValidateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
