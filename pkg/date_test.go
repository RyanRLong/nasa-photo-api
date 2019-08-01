package nasaphotoapi

import (
	"testing"
)

func TestDate_String(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want string
	}{
		{"", Date{"2019", "01", "01"}, "2019-01-01"},
		{"", Date{"xyz", "a", "b"}, "xyz-a-b"}, // does not validate, only concats
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.String(); got != tt.want {
				t.Errorf("Date.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want bool
	}{
		{"with valid date returns true", Date{"2019", "01", "01"}, true},
		{"with valid date but single digit month and day returns true", Date{"2019", "1", "1"}, true},
		{"with invalid year returns false", Date{"12190", "1", "1"}, false},
		{"with invalid month returns false", Date{"1999", "365", "1"}, false},
		{"with invalid day returns false", Date{"1999", "1", "365"}, false},
		{"with valid leap year date returns true", Date{"2020", "2", "29"}, true},
		{"with invalid leap year date returns true", Date{"2029", "2", "29"}, false},
		{"with characters as month returns false", Date{"1999", "Jan", "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsValid(); got != tt.want {
				t.Errorf("Date.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_padToTwoDigits(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Given 1 digit string - pads string with one zero", args{"1"}, "01"},
		{"Given 3 digit string - does not pad", args{"123"}, "123"},
		{"Given empty string - pads to 00", args{""}, "00"},
		{"Given 2 digit string - adds no padding", args{"12"}, "12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padToTwoDigits(tt.args.value); got != tt.want {
				t.Errorf("padToTwoDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
