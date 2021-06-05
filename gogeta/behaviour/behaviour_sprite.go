package behaviour

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Sprite struct {
	Image  *ebiten.Image
	Size   r2.Point
	Anchor r2.Point
	Angle  float64
}

func (bhvr *Sprite) Init() {
	bhvr.Image = ebiten.NewImage(int(bhvr.Size.X), int(bhvr.Size.Y))
}

func (bhvr *Sprite) Update() {
}

func (bhvr *Sprite) Draw(screen gm.Screen) {
	bhvrCommon := gm.MustGetBehaviourRel(bhvr, &Common{}).(*Common)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-bhvr.Anchor.X, -bhvr.Anchor.Y)
	op.GeoM.Rotate(bhvr.Angle)
	op.GeoM.Translate(bhvrCommon.Position.X, bhvrCommon.Position.Y)
	(*ebiten.Image)(screen).DrawImage(bhvr.Image, op)
}

// Behaviour specific method
func (bhvr *Sprite) FillColor(colorNRGBA color.NRGBA) {
	bhvr.Image.Fill(colorNRGBA)
}
