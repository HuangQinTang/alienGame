package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
	"time"
)

type Game struct {
	input *Input
}

func NewGame() *Game {
	return &Game{
		input: &Input{msg: "Hello, World!"},
	}
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rand.Seed(time.Now().UnixNano())
	screen.Fill(color.RGBA{
		R: uint8(rand.Intn(254) + 1),
		G: uint8(rand.Intn(254) + 1),
		B: uint8(rand.Intn(254) + 1),
		A: uint8(rand.Intn(254) + 1),
	})
	ebitenutil.DebugPrint(screen, g.input.msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("外星人入侵")
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Input struct {
	msg string
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("←←←←←←←←←←←←←←←←←←←←←←←")
		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("→→→→→→→→→→→→→→→→→→→→→→→")
		i.msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("-----------------------")
		i.msg = "space pressed"
	}
}
