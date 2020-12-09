package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"log"
)

//img
var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
}

type rGame struct {
}

func (g *rGame) Update() error {
	return nil
}

func (g *rGame) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *rGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&rGame{}); err != nil {
		log.Fatal(err)
	}
}
