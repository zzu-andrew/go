package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

const (
	floodScreenWidth  = 320
	floodScreenHeight = 240
)

var (
	floodEbitenImage *ebiten.Image
	colors           = []color.RGBA{
		{0xff, 0xff, 0xff, 0xff},
		{0xff, 0xff, 0x0, 0xff},
		{0xff, 0x0, 0xff, 0xff},
		{0xff, 0x0, 0x0, 0xff},
		{0x0, 0xff, 0xff, 0xff},
		{0x0, 0xff, 0x0, 0xff},
		{0x0, 0x0, 0xff, 0xff},
		{0x0, 0x0, 0x0, 0xff},
	}
)

type floodGame struct {
}

func (g *floodGame) Update() error {
	return nil
}

func (g *floodGame) Draw(screen *ebiten.Image) {
	const (
		ox = 10
		oy = 10
		dx = 60
		dy = 50
	)
	// screen fill填充的画布
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	op := &ebiten.DrawImageOptions{}
	// 向右下方移动 10个像素
	op.GeoM.Translate(ox, oy)
	screen.DrawImage(floodEbitenImage, op)

	for i, c := range colors {
		op := &ebiten.DrawImageOptions{}
		x := i % 4
		y := i/4 + 1
		op.GeoM.Translate(ox+float64(dx*x), oy+float64(dy*y))

		//reset RGB not alpha 0 forcibly
		// 清空原有图片上的 颜色填充
		op.ColorM.Scale(0, 0, 0, 1)
		//set color
		r := float64(c.R) / 0xff
		g := float64(c.G) / 0xff
		b := float64(c.B) / 0xff
		// op 中的参数是针对单个图像对象的，一个screen上可以进行很多图片的绘制
		op.ColorM.Translate(r, g, b, 0)
		screen.DrawImage(floodEbitenImage, op)

	}

}

func (g *floodGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return floodScreenWidth, floodScreenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	floodEbitenImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(floodScreenWidth*2, floodScreenHeight*2)
	ebiten.SetWindowTitle("Flood fill with solid colors (Ebiten Demo)")
	if err := ebiten.RunGame(&floodGame{}); err != nil {
		log.Fatal(err)
	}
}
