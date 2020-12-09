package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"log"
)

var mImg *ebiten.Image

func init() {
	var err error
	mImg, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type mGame struct {
}

func (g *mGame) Update() error {
	return nil
}

func (g *mGame) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// 移动像素点 50 50
	op.GeoM.Translate(50, 50)
	// 等比缩放
	op.GeoM.Scale(1.5, 1)
	screen.DrawImage(mImg, op)
}

func (g *mGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Geometry Matrix")
	if err := ebiten.RunGame(&mGame{}); err != nil {
		log.Fatal(err)
	}
}
