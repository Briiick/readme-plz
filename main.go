package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/papapalapa/readme-plz/readmelib"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", ReadImageHandler)
	r.HandleFunc("/upload", UploadHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

// ReadImageHandler receives binary image data as input and generates an mp3 file
func ReadImageHandler(w http.ResponseWriter, r *http.Request) {
	text := readmelib.DetectDocumentText("result")
	readmelib.SynthesizeAudio(text)
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

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}
