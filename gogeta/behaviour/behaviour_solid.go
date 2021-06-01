package behaviour

import "github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"

type Solid struct {
	IsSolid bool
}

func (obj *Solid) D() *gm.BehaviourData {
	return &gm.BehaviourData{}
}

func (obj *Solid) Init() {
}

func (obj *Solid) Update() {
}

func (obj *Solid) Draw(screen gm.Screen) {
}
