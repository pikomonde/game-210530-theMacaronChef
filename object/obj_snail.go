package object

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/sys"
)

type Snail struct {
	Sprite      *ebiten.Image
	SpritePoint r2.Point
	Pos         r2.Point
	Vel         r2.Point
}

func (obj *Snail) D() *gm.ObjectData {
	return &gm.ObjectData{}
}

func (obj *Snail) Init() {
	// Define sprite size
	w, h := 12, 12

	// Assign value to the object
	obj.Sprite = ebiten.NewImage(w, h)
	obj.Sprite.Fill(color.NRGBA{0x00, 0x80, 0x00, 0xff})
	obj.SpritePoint = r2.Point{float64(w) / 2, float64(h) / 2}
	obj.Pos = r2.Point{60, 60}
	obj.Vel = r2.Point{0, 0}
}

func (obj *Snail) Update() {

}

func (obj *Snail) Draw(screen gm.Screen) {
	deg := sys.AngleFromVector(obj.Vel.X, obj.Vel.Y)

	// Draw ant on the main canvas
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-obj.SpritePoint.X, -obj.SpritePoint.Y)
	op.GeoM.Rotate(deg)
	op.GeoM.Translate(obj.Pos.X, obj.Pos.Y)
	(*ebiten.Image)(screen).DrawImage(obj.Sprite, op)

	// Draw text
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Angle: %.2f", deg), 12, 24)

}
