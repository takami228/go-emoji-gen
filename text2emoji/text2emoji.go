package text2emoji

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var (
	backgroundWidth  = 128
	backgroundHeight = 128
	utf8FontSize     = float64(59.0)
	spacing          = float64(2.0)
	dpi              = float64(72)
	ctx              = new(freetype.Context)
	utf8Font         = new(truetype.Font)
	white            = color.RGBA{255, 255, 255, 255}
	black            = color.RGBA{0, 0, 0, 255}
	background       *image.RGBA
)

//GenerateTextImage ...
func GenerateTextImage(inputText string, fontFilePath string, outputFilePath string) error {

	fontBytes, err := ioutil.ReadFile(fontFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such File.", err)
		return err
	}

	utf8Font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] Font Load Error.", err)
		return err
	}

	fontForeGroundColor, fontBackGroundColor := image.NewUniform(black), image.NewUniform(white)

	background = image.NewRGBA(image.Rect(0, 0, backgroundWidth, backgroundHeight))

	draw.Draw(background, background.Bounds(), fontBackGroundColor, image.ZP, draw.Src)

	ctx = freetype.NewContext()
	ctx.SetDPI(dpi)
	ctx.SetFont(utf8Font)
	ctx.SetFontSize(utf8FontSize)
	ctx.SetClip(background.Bounds())
	ctx.SetDst(background)
	ctx.SetSrc(fontForeGroundColor)

	var UTF8text = splitString(inputText)

	pt := freetype.Pt(6, int(ctx.PointToFixed(utf8FontSize)>>6))

	for _, str := range UTF8text {
		_, err := ctx.DrawString(str, pt)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR] Draw Failed", err)
			return err
		}
		pt.Y += ctx.PointToFixed(utf8FontSize + spacing)
	}

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such File.", err)
		return err
	}
	defer outFile.Close()
	buff := bufio.NewWriter(outFile)

	err = png.Encode(buff, background)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such File.", err)
		return err
	}

	err = buff.Flush()
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] Buffer Flush Failed.", err)
		return err
	}

	return nil
}

func splitString(inpuString string) []string {
	var top = string([]rune(inpuString)[:2])
	var bottom = strings.Replace(inpuString, top, "", 1)
	return []string{top, bottom}
}
