package entity

import (
	"alienGame/config"
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
	img, _, err := ebitenutil.NewImageFromFile("./images/ship.bmp")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		Image:  img,
		Width:  width,
		Height: height,
		X:      float64(width),
		Y:      float64(height),
	}
	return ship
}

// Draw 绘制自身
func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	op := &ebiten.DrawImageOptions{}
	//fmt.Println("cfg.ScreenWidth", cfg.ScreenWidth, ship.Width)
	//fmt.Println("cfg.ScreenHeight", cfg.ScreenHeight, ship.Height)
	//fmt.Println((float64(cfg.ScreenWidth) - ship.X) / 2)

	//初始时，飞船x等于width；飞船y等于height，以下公式将飞船至于屏幕底部居中
	op.GeoM.Translate(
		(float64(cfg.ScreenWidth)-ship.X)/2,
		float64(cfg.ScreenHeight)-ship.Y)

	// 指定飞船坐标相对于原点的偏移
	screen.DrawImage(ship.Image, op)
}
