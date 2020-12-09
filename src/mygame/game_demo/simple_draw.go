package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

type fGame struct {
}

func (g *fGame) Update() error {
	return nil
}

func (g *fGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0xff, A: 0xff})
}

func (g *fGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(&fGame{}); err != nil {
		log.Fatal(err)
	}
}
