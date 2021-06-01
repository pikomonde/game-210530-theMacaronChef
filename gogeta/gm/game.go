package gm

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/constant"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/btn"
)

var gm game

type game struct {
	Objects Objects
}

func (g *game) Update() error {
	btn.Update()
	g.Objects.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// Color the background
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	// Draw object
	g.Objects.Draw(Screen(screen))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.CanvasWidth, constant.CanvasHeight
}

func Init() error {
	gm.Objects = make(map[string][]Object)
	ebiten.SetWindowSize(constant.WindowWidth, constant.WindowHeight)
	return nil
}

func Run() error {
	return ebiten.RunGame(&gm)
}
