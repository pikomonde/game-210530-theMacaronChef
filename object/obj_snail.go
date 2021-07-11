package object

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/behaviour"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Snail struct {
	behaviour.Common
	// behaviour.Platformer
}

func (obj *Snail) Init() {
	// Set Behaviour
	obj.Common.Position = r2.Point{100, 100}

	// Init Sprite Behaviour
	// obj.Sprite.SetSizeAndFillColor(12, 12, color.NRGBA{0x00, 0x80, 0x00, 0xff})
	image := ebiten.NewImage(int(32), int(32))
	image.Fill(color.NRGBA{0x00, 0x80, 0x00, 0xff})
	// image.
	// obj.Sprite.SetSprite(image)
	// obj.Sprite.SetAnchorToggle(bhvrCommon.Sprite_FrameAnchor_ToggleMiddleCenter)
	obj.Sprite.InsertFrameByImage("", image)

	// Assign value to the object
	// obj.Vel = r2.Point{0, 0}
}

func (obj *Snail) Update() {
	obj.Angle += 0.1
	obj.Position.X += 0.5
}

func (obj *Snail) Draw(screen gm.Screen) {
	// bSprite.Angle = sys.AngleFromVector(obj.Vel.X, obj.Vel.Y)
	// obj.Common.Draw(screen)

	// Draw text
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Angle: %.2f", deg), 12, 24)

}
