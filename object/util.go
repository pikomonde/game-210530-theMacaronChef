package object

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Util struct {
}

func (obj *Util) Init() {
}

func (obj *Util) Update() {
}

func (obj *Util) Draw(screen gm.Screen) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f\n", ebiten.CurrentTPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f\n", ebiten.CurrentFPS()), 0, 12)
}
