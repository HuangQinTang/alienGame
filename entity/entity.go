package entity

import "github.com/hajimehoshi/ebiten/v2"

// Entity 游戏实体
type Entity interface {
	// Draw 绘制自身
	Draw(screen *ebiten.Image)
	// GetInfo 获取示例长宽，坐标信息
	GetInfo() (Width, Height int, X, Y float64)
}
