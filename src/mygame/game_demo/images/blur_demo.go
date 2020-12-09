package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"image"
	_ "image/jpeg"
	"log"
)

const (
	screenWidth  = 826
	screenHeight = 480
)

var gophersImage *ebiten.Image

type blurGame struct {
	count int
}

func (g *blurGame) Update() error {

	return nil
}

func (g *blurGame) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(gophersImage, op)

	for j := -3; j <= 3; j++ {
		for i := -3; i <= 3; i++ {
			op := &ebiten.DrawImageOptions{}
			// 图像偏移的位置
			op.GeoM.Translate(float64(i), 244+float64(j))

			op.ColorM.Scale(1, 1, 1, 1.0/25.0)
			screen.DrawImage(gophersImage, op)
		}
	}

	/*for i := -3; i <= 1; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 240)

		//op.ColorM.Scale(1, 1, 1, 1.0/25.0)
		screen.DrawImage(gophersImage, op)
	}*/

}

func (g *blurGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.FiveYears_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Blur (Ebiten Demo)")
	if err := ebiten.RunGame(&blurGame{}); err != nil {
		log.Fatal(err)
	}

}
