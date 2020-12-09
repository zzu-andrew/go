package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX = 0
	frameOY = 32
	// 每个小图像的宽度
	frameWidth = 32
	// 每个小图像的高度
	frameHeight = 32
	frameNum    = 8
)

var runnerImage *ebiten.Image

type aniGame struct {
	count    int
	minCount int
}

// 要保持返回nil否则就认为出错
func (g *aniGame) Update() error {
	g.minCount++
	if g.minCount%5 == 0 {
		g.count++
	}
	return nil
}

func (g *aniGame) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64((frameHeight)/2))
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	i := (g.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	fmt.Println(sx, sy)
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
func (g *aniGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(&aniGame{}); err != nil {
		log.Fatal(err)
	}

}
