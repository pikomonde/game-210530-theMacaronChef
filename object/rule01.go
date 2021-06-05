package object

import (
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Rule01 struct {
}

func (obj *Rule01) Init() {
	gm.SetObjectAndInit(&Util{})
	gm.SetObjectAndInit(&Snail{})
	gm.SetObjectAndInit(&Tile{})
}

func (obj *Rule01) Update() {
	// if btn.IsMouseButtonClicked(ebiten.MouseButtonLeft) {
	// 	x, y := ebiten.CursorPosition()
	// 	gm.SetObjectAndInit(&Food{Pos: r2.Point{float64(x), float64(y)}})
	// }
}

func (obj *Rule01) Draw(screen gm.Screen) {
}
