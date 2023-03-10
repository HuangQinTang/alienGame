package entity

import (
	"alienGame/resources"
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/bmp"
	"log"
)

// Ship 飞船
type Ship struct {
	Image  *ebiten.Image
	Width  int
	Height int
	X      float64 // x坐标
	Y      float64 // y坐标
}

func NewShip(screenWidth, screenHeight int) *Ship {
	//img, _, err := ebitenutil.NewImageFromFile("./images/ship.bmp")
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.ShipBmp))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	//飞船初始坐标，置于屏幕底部居中
	x := (screenWidth - width) / 2
	y := screenHeight - height
	ship := &Ship{
		Image:  img,
		Width:  width,
		Height: height,
		X:      float64(x),
		Y:      float64(y),
	}
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//fmt.Println("cfg.ScreenWidth", cfg.ScreenWidth, ship.Width)
	//fmt.Println("cfg.ScreenHeight", cfg.ScreenHeight, ship.Height)
	//fmt.Println((float64(cfg.ScreenWidth) - ship.X) / 2)

	op.GeoM.Translate(ship.X, ship.Y)

	// 指定飞船坐标相对于原点的偏移
	screen.DrawImage(ship.Image, op)
}

func (ship *Ship) GetInfo() (Width, Height int, X, Y float64) {
	return ship.Width, ship.Height, ship.X, ship.Y
}
