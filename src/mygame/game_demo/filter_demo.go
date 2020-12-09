package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"image"
	_ "image/png"
	"log"
)

const (
	filterScreenWidth  = 640
	filterScreenHeight = 480
)

var filterImage *ebiten.Image

type filterGame struct {
}

func (g *filterGame) Update() error {
	return nil
}

func (g *filterGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Nearest filter default VS  Linear filter")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(64, 64)
	//	by default nearest filter is used
	screen.DrawImage(filterImage, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(64, 64+240)
	//	specify linear filter  1
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(filterImage, op)
}

func (g *filterGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return filterScreenWidth, filterScreenHeight
}

func main() {

	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}

	filterImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(filterScreenWidth, filterScreenHeight)
	ebiten.SetWindowTitle("Filter (Ebiten Demo)")
	if err := ebiten.RunGame(&filterGame{}); err != nil {
		log.Fatal(err)
	}

}
