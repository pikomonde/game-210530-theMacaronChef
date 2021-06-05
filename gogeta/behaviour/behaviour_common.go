package behaviour

import (
	"github.com/golang/geo/r2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Common struct {
	Position r2.Point
}

func (bhvr *Common) Init() {
}

func (bhvr *Common) Update() {
}

func (bhvr *Common) Draw(screen gm.Screen) {
}
