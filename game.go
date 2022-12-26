package main

import (
	"alienGame/config"
	"alienGame/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"time"
)

// Game 我们规定如果击杀config.SuccessNum个外星人则游戏胜利，如果有config.FailNum个外星人移出屏幕外或者碰撞到飞船则游戏失败。
type Game struct {
	input          *Input //键盘事件
	cfg            *config.Config
	mode           config.Mode                 //游戏状态
	ship           *entity.Ship                //飞船
	bullets        map[*entity.Bullet]struct{} //子弹
	aliens         map[*entity.Alien]struct{}  //外星人
	lastAliensTime time.Time                   //外星人上次出现的时间
	failCount      int                         //被外星人碰撞和移出屏幕的外星人数量之和
	successCount   int                         //消灭的外星人数
	overMsg        string                      //游戏结束提示
}

func (g *Game) init() {
	// 创建字体
	g.CreateFonts()
	// 初始化游戏状态
	g.mode = config.ModeTitle
	// 游戏提示重空
	g.overMsg = ""
	g.successCount = 0
	g.failCount = 0
	g.aliens = make(map[*entity.Alien]struct{})
	g.bullets = make(map[*entity.Bullet]struct{})
	// 创建一组外星人
	g.CreateAliens()
}

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	//smallArcadeFont font.Face
)

// CreateFonts 创建字体
func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	//smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
	//	Size:    float64(g.cfg.SmallFontSize),
	//	DPI:     dpi,
	//	Hinting: font.HintingFull,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func NewGame() *Game {
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	// 避免实际窗口与逻辑窗口伸缩，这里/2
	//cfg.ScreenWidth = cfg.ScreenWidth / config.LayoutMultiple
	//cfg.ScreenHeight = cfg.ScreenHeight / config.LayoutMultiple

	game := &Game{
		input: &Input{msg: "Ready to play"},
		mode:  config.ModeTitle,
		ship:  entity.NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:   cfg,
	}

	game.init()

	return game
}

// Update 每个tick都会被调用。tick是引擎更新的一个时间单位，默认为1/60s。 tick的倒数
// 我们一般称为帧，即游戏的更新频率。默认ebiten游戏是60帧，即每秒更新60次。 该方法主要用
// 来更新游戏的逻辑状态，例如子弹位置更新。注意到Update方法的返回值为error 类型，当 Update
// 方法返回一个非空的error值时，游戏停止。在上面的例子中，我们一直返回nil，故只有关闭窗口时游戏才停止。
func (g *Game) Update() error {
	switch g.mode {
	// 游戏启动
	case config.ModeTitle:
		if g.input.IsKeyPressed() { //当收到键盘空格，鼠标左击，设置游戏状态为进行
			g.mode = config.ModeGame
		}

	// 游戏进行
	case config.ModeGame:
		// 根据键盘操作更新飞船坐标
		g.input.Update(g)

		// 移动子弹
		for bullet := range g.bullets {
			bullet.Y -= bullet.SpeedFactor
			if bullet.OutOfScreen() { // 清除超出屏幕的子弹
				delete(g.bullets, bullet)
			}
		}

		for alien := range g.aliens {
			// 移动外星人
			alien.Y += alien.SpeedFactor

			// 清除和统计超出屏幕的外星人
			if alien.OutOfScreen(g.cfg) {
				g.failCount++ //未消灭的外星人+1
				delete(g.aliens, alien)
				continue
			}

			// 检测外星人与飞船是否碰撞
			if CheckCollision(alien, g.ship) {
				g.failCount++ //未消灭的外星人+1
				delete(g.aliens, alien)
				continue
			}

			// 检测子弹与外星人是否碰撞
			for bullet := range g.bullets {
				if CheckCollision(alien, bullet) { //删除碰撞的子弹和外星人
					delete(g.aliens, alien)
					delete(g.bullets, bullet)
					g.successCount++ //消灭的外星人+1
				}
			}
		}

		//判断游戏通关情况
		if g.failCount >= g.cfg.FailNum {
			g.overMsg = "Game Over!"
		} else if g.successCount >= g.cfg.SuccessNum {
			g.overMsg = "You Win!"
		}
		if len(g.overMsg) > 0 {
			g.mode = config.ModeOver
			break
		}

		// 创建新的外星人
		if time.Now().Sub(g.lastAliensTime).Milliseconds() > g.cfg.AlienInterval {
			g.CreateAliens()
		}

	// 游戏结束
	case config.ModeOver:
		if g.input.IsKeyPressed() {
			g.init()
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

	var titleTexts []string
	var texts []string

	switch g.mode {
	// 游戏启动
	case config.ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}

	// 游戏进行
	case config.ModeGame:
		// 绘制飞船
		g.ship.Draw(screen)
		// 绘制子弹
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		// 绘制外星人
		for alien := range g.aliens {
			alien.Draw(screen)
		}
		//打印操作，方便调试
		ebitenutil.DebugPrint(screen, "")

	// 游戏结束
	case config.ModeOver:
		texts = []string{"", g.overMsg}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}
}

// Layout 该方法接收游戏窗口的尺寸作为参数，返回游戏的逻辑屏幕大小。
// 我们实际上计算坐标是对应这个逻辑屏幕的，Draw将逻辑屏幕渲染到实际窗口上。
// 这个时候可能会出现伸缩。
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

// CreateAliens 创建一批外星人
func (g *Game) CreateAliens() {
	alien := entity.NewAlien(g.cfg)

	// 左右各留一个外星人宽度的空间
	availableSpaceX := g.cfg.ScreenWidth - 2*alien.Width
	// 两个外星人之间留一个外星人宽度的空间。所以一行可以创建的外星人的数量为
	numAliens := availableSpaceX / (2 * alien.Width)

	for row := 0; row < 2; row++ {
		for i := 0; i <= numAliens; i++ {
			alien = entity.NewAlien(g.cfg)
			alien.X = float64(alien.Width + 2*alien.Width*i)
			alien.Y = float64(alien.Height*row) * 1.5
			g.addAlien(alien)
		}
	}
	g.lastAliensTime = time.Now()
}

// addBullet 添加子弹对象，发射子弹
func (g *Game) addBullet(bullet *entity.Bullet) {
	g.bullets[bullet] = struct{}{}
}

// addAlien 添加外线人对象，创建外星人
func (g *Game) addAlien(alien *entity.Alien) {
	g.aliens[alien] = struct{}{}
}
