package nasaphotoapi

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

// FetchData makes a GET response to the url string passed to it
func FetchData(url string) (responseBody io.ReadCloser, err error) {
	response, err := http.Get(url)
	//TODO: Need to handle when past rate limit
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// DownloadPhotos downloads all photos contained in photos
func DownloadPhotos(photos Photos) (numberDownloaded int, err error) {
	total := 0
	for _, photo := range photos {
		// Get the data
		resp, err := http.Get(photo.ImgSRC)
		if err != nil {
			return total, err
		}
		defer resp.Body.Close()
		dir, err := os.Getwd()
		if err != nil {
			return total, err
		}
		path := dir + "\\" + photo.Date
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModeDir)
		}

		// Create the file
		stringID := strconv.Itoa(int(photo.ID))
		out, err := os.Create(path + "\\" + photo.Date + "_" + stringID + ".jpg")
		if err != nil {
			return total, err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return 0, err // all or nothing
		}
		total++
	}
	return total, nil

}
