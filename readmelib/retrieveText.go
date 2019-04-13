package readmelib

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

// DetectDocumentText gets the full document text from the Vision API for an image at the given file path.
func DetectDocumentText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotation, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		return err
	}

	if annotation == nil {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Document Text:")
		fmt.Fprintf(w, "%q\n", annotation.Text)

		fmt.Fprintln(w, "Pages:")
		for _, page := range annotation.Pages {
			fmt.Fprintf(w, "\tConfidence: %f, Width: %d, Height: %d\n", page.Confidence, page.Width, page.Height)
			fmt.Fprintln(w, "\tBlocks:")
			for _, block := range page.Blocks {
				fmt.Fprintf(w, "\t\tConfidence: %f, Block type: %v\n", block.Confidence, block.BlockType)
				fmt.Fprintln(w, "\t\tParagraphs:")
				for _, paragraph := range block.Paragraphs {
					fmt.Fprintf(w, "\t\t\tConfidence: %f", paragraph.Confidence)
					fmt.Fprintln(w, "\t\t\tWords:")
					for _, word := range paragraph.Words {
						symbols := make([]string, len(word.Symbols))
						for i, s := range word.Symbols {
							symbols[i] = s.Text
						}
						wordText := strings.Join(symbols, "")
						fmt.Fprintf(w, "\t\t\t\tConfidence: %f, Symbols: %s\n", word.Confidence, wordText)
					}
				}
			}
		}
	}

	return nil
}

// DetectDocumentText gets the full document text from the Vision API for an image at the given file path.
func DetectDocumentTextURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotation, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		return err
	}

	if annotation == nil {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Document Text:")
		fmt.Fprintf(w, "%q\n", annotation.Text)

		fmt.Fprintln(w, "Pages:")
		for _, page := range annotation.Pages {
			fmt.Fprintf(w, "\tConfidence: %f, Width: %d, Height: %d\n", page.Confidence, page.Width, page.Height)
			fmt.Fprintln(w, "\tBlocks:")
			for _, block := range page.Blocks {
				fmt.Fprintf(w, "\t\tConfidence: %f, Block type: %v\n", block.Confidence, block.BlockType)
				fmt.Fprintln(w, "\t\tParagraphs:")
				for _, paragraph := range block.Paragraphs {
					fmt.Fprintf(w, "\t\t\tConfidence: %f", paragraph.Confidence)
					fmt.Fprintln(w, "\t\t\tWords:")
					for _, word := range paragraph.Words {
						symbols := make([]string, len(word.Symbols))
						for i, s := range word.Symbols {
							symbols[i] = s.Text
						}
						wordText := strings.Join(symbols, "")
						fmt.Fprintf(w, "\t\t\t\tConfidence: %f, Symbols: %s\n", word.Confidence, wordText)
					}
				}
			}
		}
	}

	return nil
}
