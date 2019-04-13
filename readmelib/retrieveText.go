package readmelib

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

// DetectDocumentText gets the full document text from the Vision API for an image at the given file path.
func DetectDocumentText(file string) string {
	ctx := context.Background()

	client, _ := vision.NewImageAnnotatorClient(ctx)
	f, _ := os.Open(file)
	defer f.Close()

	image, _ := vision.NewImageFromReader(f)
	annotation, _ := client.DetectDocumentText(ctx, image, nil)

	if annotation == nil {
		fmt.Print("No text found.")
	} else {
		return annotation.Text
	}

	return ""
}
