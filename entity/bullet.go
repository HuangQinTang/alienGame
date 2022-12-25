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
	X           float64
	Y           float64
	SpeedFactor float64 //子弹速度
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
		X:           ship.X + float64(ship.Width-cfg.BulletWidth)/2,
		Y:           float64(cfg.ScreenHeight - ship.Height - cfg.BulletHeight),
		SpeedFactor: cfg.BulletSpeedFactor,
	}
}

// Draw 绘制子弹
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.X, bullet.Y)
	screen.DrawImage(bullet.image, op)
}

// 判断子弹是否处于屏幕之外
func (bullet *Bullet) OutOfScreen() bool {
	return bullet.Y < -float64(bullet.height)
}
