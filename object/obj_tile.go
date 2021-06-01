package object

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/sys"
)

type Tile struct {
	Sprite      *ebiten.Image
	SpritePoint r2.Point
	Pos         r2.Point
	Vel         r2.Point
}

func (obj *Tile) D() *gm.ObjectData {
	return &gm.ObjectData{}
}

func (obj *Tile) Init() {
	// Define sprite size
	w, h := 16, 16

	// Assign value to the object
	obj.Sprite = ebiten.NewImage(w, h)
	obj.Sprite.Fill(color.Black)
	obj.SpritePoint = r2.Point{float64(w) / 2, float64(h) / 2}
	obj.Pos = r2.Point{100, 100}
	obj.Vel = r2.Point{0, 0}

	// Draw object's sprite
	headSprite := ebiten.NewImage(1, 4)
	headSprite.Fill(color.NRGBA{0x80, 0x00, 0x00, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(3, 0)
	obj.Sprite.DrawImage(headSprite, op)

}

func (obj *Tile) Update() {

}

func (obj *Tile) Draw(screen gm.Screen) {
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
