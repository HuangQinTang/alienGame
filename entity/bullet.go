package entity

import (
	"alienGame/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Bullet 子弹
type Bullet struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewBullet(cfg *config.Config, ship *Ship) *Bullet {
	// 首先根据配置的宽高创建一个rect对象
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	// 创建*ebiten.Image对象
	img := ebiten.NewImageWithOptions(rect, nil)
	// 填充背景色
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image:       img,
		width:       cfg.BulletWidth,
		height:      cfg.BulletHeight,
		x:           ship.X + float64(ship.Width-cfg.BulletWidth)/2,
		y:           float64(cfg.ScreenHeight - ship.Height - cfg.BulletHeight),
		speedFactor: cfg.BulletSpeedFactor,
	}
}

// Draw 绘制子弹
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}