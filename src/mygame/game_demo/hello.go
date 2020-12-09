package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type Game struct {
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello world!")

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 720, 480
}

func main() {

	ebiten.SetWindowSize(720, 480)
	ebiten.SetWindowTitle("Hello world!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
