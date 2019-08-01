package nasaphotoapi

import (
	"reflect"
	"testing"
)

func Test_parseJSONPhoto(t *testing.T) {
	type args struct {
		jsonPhotoData map[string]interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantPhoto Photo
		wantErr   bool
	}{
		{"happy path", args{
			map[string]interface{}{"id": 1234.0, "img_src": "http://www.google.com/image.jpg", "earth_date": "2019-01-01"},
		}, Photo{1234, "http://www.google.com/image.jpg", "2019-01-01"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPhoto, err := parseJSONPhoto(tt.args.jsonPhotoData)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseJSONPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPhoto, tt.wantPhoto) {
				t.Errorf("parseJSONPhoto() = %v, want %v", gotPhoto, tt.wantPhoto)
			}
		})
	}
}

func Test_generateURLFromDate(t *testing.T) {
	type args struct {
		date Date
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
		wantErr bool
	}{
		{"Happy path", args{Date{"1999", "1", "1"}}, "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?earth_date=1999-1-1&api_key=DEMO_KEY", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := generateURLFromDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateURLFromDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("generateURLFromDate() = %v, want %v", gotURL, tt.wantURL)
			}
		})
	}
}
