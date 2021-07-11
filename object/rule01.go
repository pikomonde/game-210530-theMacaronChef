package object

import (
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Rule01 struct {
}

func (obj *Rule01) Init() {
	gm.SetAndInitObject(&Util{})
	gm.SetAndInitObject(&Snail{
		// Common: behaviour.Common{
		// 	Position: r2.Point{100, 100},
		// },
	})
	// gm.SetAndInitObject(&Tile{})
}

func (obj *Rule01) Update() {
	// if btn.IsMouseButtonClicked(ebiten.MouseButtonLeft) {
	// 	x, y := ebiten.CursorPosition()
	// 	gm.SetAndInitObject(&Food{Pos: r2.Point{float64(x), float64(y)}})
	// }
}

func (obj *Rule01) Draw(screen gm.Screen) {
}
