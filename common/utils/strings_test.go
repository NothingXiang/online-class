package utils

import (
	"testing"
)

func TestIsEmptyString(t *testing.T) {

	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"blank space",
			"   ",
			true,
		},
		{
			"no-blank space",
			" 11  ",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyString(tt.args); got != tt.want {
				t.Errorf("IsEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhoneNum(t *testing.T) {

	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"english",
			"wddcac",
			false,
		},
		{
			"chinese",
			"联通移动电信手机号",
			false,
		},
		{
			"number",
			"123456789",
			false,
		},
		{
			"true",
			"13113678929",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhoneNum(tt.args); got != tt.want {
				t.Errorf("IsPhoneNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
