package main

import (
	"alienGame/config"
	"alienGame/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	input   *Input
	cfg     *config.Config
	ship    *entity.Ship                //飞船
	bullets map[*entity.Bullet]struct{} //子弹
}

func NewGame() *Game {
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	//避免实际窗口与逻辑窗口伸缩，这里/2
	cfg.ScreenWidth = cfg.ScreenWidth / config.LayoutMultiple
	cfg.ScreenHeight = cfg.ScreenHeight / config.LayoutMultiple

	return &Game{
		input:   &Input{msg: "Ready to play"},
		ship:    entity.NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:     cfg,
		bullets: make(map[*entity.Bullet]struct{}),
	}
}

// Update 每个tick都会被调用。tick是引擎更新的一个时间单位，默认为1/60s。 tick的倒数
// 我们一般称为帧，即游戏的更新频率。默认ebiten游戏是60帧，即每秒更新60次。 该方法主要用
// 来更新游戏的逻辑状态，例如子弹位置更新。注意到Update方法的返回值为error 类型，当 Update
// 方法返回一个非空的error值时，游戏停止。在上面的例子中，我们一直返回nil，故只有关闭窗口时游戏才停止。
func (g *Game) Update() error {
	// 根据键盘操作更新飞船坐标
	g.input.Update(g)

	// 移动子弹
	for bullet := range g.bullets {
		bullet.Y -= bullet.SpeedFactor
	}
	// 清除超出屏幕的子弹
	for bullet := range g.bullets {
		if bullet.OutOfScreen() {
			delete(g.bullets, bullet)
		}
	}
	return nil
}

// Draw 每帧（frame）调用。帧是渲染使用的一个时间单位，依赖显示器的刷新率。
// 如果显示器的刷新率为60Hz，Draw将会每秒被调用60次。 Draw接受一个类型为
// *ebiten.Image的screen对象。ebiten引擎每帧会渲染这个screen。
// 我们调用ebitenutil.DebugPrint函数在screen上渲染一条调试信息。
// 由于调用Draw方法前，screen会被重置，故DebugPrint每次都需要调用。
func (g *Game) Draw(screen *ebiten.Image) {
	//填充背景色
	screen.Fill(g.cfg.BgColor)

	// 绘制飞船
	g.ship.Draw(screen, g.cfg)

	// 绘制子弹
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, g.input.msg)
}

// Layout 该方法接收游戏窗口的尺寸作为参数，返回游戏的逻辑屏幕大小。
// 我们实际上计算坐标是对应这个逻辑屏幕的，Draw将逻辑屏幕渲染到实际窗口上。
// 这个时候可能会出现伸缩。在config.json中配置的游戏窗口大小为(640, 480)，
// Layout返回的逻辑大小为(320, 240)，所以显示会放大1倍。
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

// addBullet 发射子弹
func (g *Game) addBullet(bullet *entity.Bullet) {
	g.bullets[bullet] = struct{}{}
}
