// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

package exercises

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
)

// implemented generator functions.
const (
	XOR      = "xor"
	AND      = "and"
	OR       = "or"
	MULTIPLY = "multiply"
	MEAN     = "mean"
)

// SaveFile is the name of the file that the image will
// be saved.
const SaveFile = "generatedImage.png"

// generatorsMap maps each generator for a CLI command to allow
// the user to choose the function.
var generatorsMap = map[string]func(int, int) byte{
	XOR:      byteXor,
	AND:      byteAnd,
	OR:       byteOr,
	MULTIPLY: multiply,
	MEAN:     mean,
}

// Image struct implements the image.Image interface.
// It started as a solution for the [slice](https://go.dev/tour/moretypes/18)
// exercise, but the solution changed to fit the [Image](https://go.dev/tour/methods/25)
// exercise.
type Image struct {
	generator     func(int, int) byte
	width, height int
}

// NewImg create a pointer to a new Image generated by a
// generatorsMap function.
func NewImg(functionName string) (*Image, error) {
	var generator, ok = generatorsMap[functionName]
	if !ok {
		return nil, errors.New("function not recognised")
	}
	return &Image{generator, 255, 255}, nil
}

// Save an Image on the SaveFile file.
func (img *Image) Save() error {
	var file, err = os.Create(SaveFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

// interface functions.

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	var value = img.generator(x, y)
	return color.RGBA{R: value, G: value, B: value, A: 255}
}

// generators implementation.

func byteXor(a, b int) byte {
	return byte(a ^ b)
}

func byteAnd(a, b int) byte {
	return byte(a & b)
}

func byteOr(a, b int) byte {
	return byte(a | b)
}
func multiply(a, b int) byte {
	return byte(a * b)
}

func mean(a, b int) byte {
	return byte((a + b) / 2)
}