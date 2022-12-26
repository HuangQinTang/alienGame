package config

type Mode uint8

const (
	LayoutMultiple = 2

	// ModeTitle 启动游戏时
	ModeTitle Mode = iota
	// ModeGame 游戏进行时
	ModeGame
	// ModeOver 游戏结束
	ModeOver
)
