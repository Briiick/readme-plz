package main

import (
	"encoding/json"
	"fmt"
	"github.com/papapalapa/readme-plz/readmelib"
	"io"
	"log"
	"net/http"
	"os"
)

// Declare the success response type
type successResponse struct {
	Text    string
	Message string
	Audio   string
}

// Declare the failure response type
type failureResponse struct {
	Message string
}

func main() {
	// Routes consist of a path and a handler function.
	http.HandleFunc("/", ReadImageHandler)
	http.HandleFunc("/upload", UploadHandler)

	// Get environment port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// ReadImageHandler receives binary image data as input and generates an mp3 file
func ReadImageHandler(w http.ResponseWriter, r *http.Request) {
	text := readmelib.DetectDocumentText("result")

	if text != "" {
		audio := readmelib.SynthesizeAudio(text)

		response := &successResponse{
			Text:    text,
			Audio:   audio,
			Message: "SYNTHESIS_SUCCESS"}

		json, _ := json.Marshal(response)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	} else {
		response := &failureResponse{
			Message: "SYNTHESIS_FAILURE"}

		json, _ := json.Marshal(response)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

// UploadHandler receives an image file and returns the binary data of the image
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("./result")
	if err != nil {
		panic(err)
	}

	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}
