package entity

import (
	"alienGame/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Alien struct {
	image       *ebiten.Image
	Width       int
	Height      int
	X           float64
	Y           float64
	SpeedFactor float64
}

func NewAlien(cfg *config.Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./images/alien.bmp")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		image:       img,
		Width:       width,
		Height:      height,
		X:           0,
		Y:           0,
		SpeedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.X, alien.Y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) GetInfo() (Width, Height int, X, Y float64) {
	return alien.Width, alien.Height, alien.X, alien.Y
}

// OutOfScreen 判断外星人是否处于屏幕之外
func (alien *Alien) OutOfScreen(cfg *config.Config) bool {
	return alien.Y > float64(cfg.ScreenHeight)
}
