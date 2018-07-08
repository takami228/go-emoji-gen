package animation

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/soniakeys/quant/median"
)

//RotateImage ...
func RotateImage(inputFilePath string, outputFilePath string, speed float64) error {
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] No Such File.", err)
		return err
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	outputGif := &gif.GIF{
		Image:     []*image.Paletted{},
		Delay:     []int{},
		LoopCount: 0,
	}

	var base = (math.Pi * 2) * 20 * speed / 360

	limit := int(360 / 20 / speed)
	q := median.Quantizer(256)
	p := q.Quantize(make(color.Palette, 0, 256), src)

	for i := 0; i < limit; i++ {
		dst := image.NewPaletted(src.Bounds(), p)
		draw.Draw(dst, src.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
		err = graphics.Rotate(dst, src, &graphics.RotateOptions{Angle: base * float64(i)})
		if err != nil {
			log.Fatal(err)
		}
		outputGif.Image = append(outputGif.Image, dst)
		outputGif.Delay = append(outputGif.Delay, 10)
	}

	out, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] File Create Failed.", err)
		return err
	}

	defer out.Close()
	err = gif.EncodeAll(out, outputGif)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] File Encode Failed.", err)
		return err
	}

	return nil
}
