package nasaphotoapi

import (
	"encoding/json"
	"fmt"
	"io"

	"io/ioutil"
)

const (
	baseURL      = "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?"
	keyEarthDate = "earth_date="
	keyAPIKey    = "&api_key="
	valueAPIKey  = "DEMO_KEY"
)

// Photo struct for holding attributes and methods of a photo
type Photo struct {
	ID     float64
	ImgSRC string
	Date   string
}

// Photos is a collection of Photo
type Photos []Photo

// FetchPhotosByDate fetches photos given a day, month,
// and year.
// The following format are accepted:
// Year: YYYY
// Month: MM (X is also accepted for single digits)
// Day: DD (D is also accepted for single digits)
func FetchPhotosByDate(year string, month string, day string) (photos Photos, err error) {
	date := Date{year, month, day}
	url, err := generateURLFromDate(date)
	if err != nil {
		return nil, err
	}
	responseBody, err := FetchData(url)
	if err != nil {
		return nil, err
	}
	photos, err = parseJSONBody(responseBody)
	if err != nil {
		return nil, err
	}
	return photos, nil
}

// parseJSONBody parses a json response and returns a Photos
func parseJSONBody(jsonBody io.ReadCloser) (photos Photos, err error) {
	var jsonData map[string]interface{}

	data, _ := ioutil.ReadAll(jsonBody)
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	collection, err := jsonData["photos"].([]interface{}), nil
	if err != nil {
		return nil, err
	}
	for _, item := range collection {
		photo, err := parseJSONPhoto(item.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

// parseJSONPhoto verifies that expected fields exist, type hints to their appropriate
// values, then creates and returns a Photo struct on success.
func parseJSONPhoto(jsonPhotoData map[string]interface{}) (photo Photo, err error) {
	entry := jsonPhotoData
	// validation
	if _, exist := jsonPhotoData["id"]; !exist {
		return Photo{}, PhotoError{"Index not found", "id"}
	}
	if _, exist := jsonPhotoData["img_src"]; !exist {
		return Photo{}, PhotoError{"Index not found", "img_src"}
	}
	if _, exist := jsonPhotoData["earth_date"]; !exist {
		return Photo{}, PhotoError{"Index not found", "earth_date"}
	}

	// convert and assign
	id := entry["id"].(float64)
	imgSRC := entry["img_src"].(string)
	date := entry["earth_date"].(string)
	return Photo{
		ID:     id,
		ImgSRC: imgSRC,
		Date:   date,
	}, nil
}

// geneateURLFromDate generates an endpoint url based on the
// passed in date struct.
func generateURLFromDate(date Date) (url string, err error) {
	if !date.IsValid() {
		return "", DateError{date, "Date is not valid."}
	}
	return baseURL + keyEarthDate + date.String() + keyAPIKey + valueAPIKey, nil
}

// PhotoError is an error implementation that holds the offending string
// with a message.
type PhotoError struct {
	Message string
	Value   string
}

// Error is an implementation for PhotoError
func (e PhotoError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.Value)
}
