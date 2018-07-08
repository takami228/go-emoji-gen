package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takami228/go-emoji-gen/animation"
	"github.com/takami228/go-emoji-gen/resize2emoji"
	"github.com/takami228/go-emoji-gen/text2emoji"
)

func main() {
	var convertType string
	var imageFilePath string
	var outputFilePath string
	var inputString string
	var fontFilePath string

	flag.StringVar(&convertType, "type", "", "set commad type: resize, text or animation")
	flag.StringVar(&imageFilePath, "image", "-", "resize image filepath.")
	flag.StringVar(&outputFilePath, "out", "output.png", "output image filepath.")
	flag.StringVar(&inputString, "string", "", "input emoji string.")
	flag.StringVar(&fontFilePath, "font", "sample.ttf", "ttf font file path.")

	flag.Parse()

	switch convertType {
	case "resize":
		if err := resize2emoji.ResizeImage(imageFilePath, outputFilePath); err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] File Resize Failed.", err)
		}
	case "text":
		if err := text2emoji.GenerateTextImage(inputString, fontFilePath, outputFilePath); err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Image File Generate Failed.", err)
		}
	case "animation":
		if err := animation.RotateImage(imageFilePath, outputFilePath, 1.1); err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Image File Generate Failed.", err)
		}
	}
}
