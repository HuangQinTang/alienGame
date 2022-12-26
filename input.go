package main

import (
	"alienGame/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	msg            string    // 用于调试
	lastBulletTime time.Time //上一次发射子弹的时间
}

func (i *Input) Update(game *Game) {
	// 左右移动更新飞船坐标
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if game.ship.X >= 0-float64(game.ship.Width/2) { //允许最大左移半个身位
			game.ship.X -= game.cfg.ShipSpeedFactor
		}
		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if game.ship.X <= float64(game.cfg.ScreenWidth-game.ship.Width/2) { //允许最大右移半个身为
			game.ship.X += game.cfg.ShipSpeedFactor
		}
		i.msg = "right pressed"
	}

	// 空格发射子弹
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		//限制同时存在的子弹，以及子弹发射的间隔
		if len(game.bullets) < game.cfg.MaxBulletNum &&
			time.Now().Sub(i.lastBulletTime).Milliseconds() > game.cfg.BulletInterval {
			bullet := entity.NewBullet(game.cfg, game.ship)
			game.addBullet(bullet)
			i.lastBulletTime = time.Now()
			i.msg = "space pressed"
		}
	}
}

func (i *Input) IsKeyPressed() bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}
