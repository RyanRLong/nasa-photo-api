package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	npa "github.com/saltycatfish/nasa-photo-api/pkg"
)

// Result hold the results of the request, including a
// successful boolean (true if successful), and the total
// number of images retrieved.
type Result struct {
	Successful  bool
	TotalPhotos int64
}

func apiCall(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) != 5 {
		http.Error(w, "Malformed url: Format is /photos/yyyy/mm/dd", http.StatusBadRequest)
	}
	photos, err := npa.FetchPhotosByDate(params[2], params[3], params[4])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	total, err := npa.DownloadPhotos(photos)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	results := Result{true, total}
	// marshal response and return
	response, err := json.Marshal(results)
	w.Write(response)
}

func handleRequests() {
	http.HandleFunc("/photos/", apiCall)
	log.Fatal(http.ListenAndServe(":8050", nil))
}

func main() {
	log.Println("Listening...")
	handleRequests()
}
