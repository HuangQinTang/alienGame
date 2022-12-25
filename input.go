package main

import (
	"alienGame/config"
	"alienGame/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg        string  //用于调试
	LMoveLimit float64 //飞船左移时x轴最大值
	RMoveLimit float64 //飞船右移时允许的x轴最小值
}

func (i *Input) Update(ship *entity.Ship, config *config.Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		//fmt.Println("←←←←←←←←←←←←←←←←←←←←←←←")
		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		//fmt.Println("→→→→→→→→→→→→→→→→→→→→→→→")
		i.msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		//fmt.Println("-----------------------")
		i.msg = "space pressed"
	}

	// 左右移动时更新飞船坐标
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if ship.X <= i.LMoveLimit {
			ship.X += config.ShipSpeedFactor
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if ship.X >= i.RMoveLimit {
			ship.X -= config.ShipSpeedFactor
		}
	}
}
