package resize2emoji

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// ResizeImage ...
func ResizeImage(imageFilePath string, outputFilePath string) error {
	//Open File
	file, err := os.Open(imageFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such File.", err)
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	//Set OutPut File
	resizedImageFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such FilePath.", err)
		return err
	}
	defer resizedImageFile.Close()

	//Exec Resize
	resizedImage := resize.Resize(128, 0, img, resize.Lanczos3)
	err = png.Encode(resizedImageFile, resizedImage)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] File Encode Failed.", err)
		return err
	}

	return nil
}
