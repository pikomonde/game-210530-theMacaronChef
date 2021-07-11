package behaviour

import (
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	bhvrCommon "github.com/pikomonde/game-210530-theMacaronChef/gogeta/behaviour/behaviour_common"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Common struct {
	Sprite   bhvrCommon.Sprite
	Position r2.Point // Position of the object based on cartesian room
	Angle    float64  // Angle of the object based on sprite anchor
}

func (bhvr *Common) Init() {
	bhvr.Sprite.Init()
}

func (bhvr *Common) Update() {
}

func (bhvr *Common) Draw(screen gm.Screen) {
	// 	bhvrCommon := gm.MustGetBehaviourRel(bhvr, &Common{}).(*Common)
	frame := bhvr.Sprite.GetCurrentFrame()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-frame.Anchor.X, -frame.Anchor.Y)
	op.GeoM.Rotate(bhvr.Angle)
	op.GeoM.Translate(bhvr.Position.X, bhvr.Position.Y)
	// op.GeoM.Translate(60, 60)
	(*ebiten.Image)(screen).DrawImage(frame.Image, op)
}
