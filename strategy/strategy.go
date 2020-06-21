package strategy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type PrintStrategy interface {
	Print() error
}

type ConsoleSquare struct{}

type ImageSquare struct {
	DestinationFilePath string
}

func (c *ConsoleSquare) Print() error {
	fmt.Println("Square")
	return nil
}

func (t *ImageSquare) Print() error {
	width := 800
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{
			X: width,
			Y: height,
		},
	})
	bgColor := image.Uniform{
		color.RGBA{
			R: 70,
			G: 70,
			B: 70,
			A: 0,
		},
	}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)
	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("Error opening file")
	}
	defer w.Close()
	if err := jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}
	return nil
}

func Use(s string) {
	var activeStrategy PrintStrategy
	switch s {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"./image.jpg"}
	default:
		activeStrategy = &ConsoleSquare{}
	}
	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
