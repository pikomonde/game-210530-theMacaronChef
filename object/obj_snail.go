package object

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/behaviour"
	bhvrCommon "github.com/pikomonde/game-210530-theMacaronChef/gogeta/behaviour/behaviour_common"
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
	frame := &bhvrCommon.Frame{Image: ebiten.NewImage(int(32), int(32))}
	frame.Image.Fill(color.NRGBA{0x00, 0x80, 0x00, 0xff})
	frame.SetAnchorToggle(bhvrCommon.Sprite_FrameAnchor_ToggleDefault)
	obj.Sprite.InsertFrame("", frame)

	// Assign value to the object
	// obj.Vel = r2.Point{0, 0}
}

func (obj *Snail) Update() {
	// frame := obj.Sprite.GetCurrentFrame()
	// frame.Anchor.X = 0
	// frame.Anchor.Y = 0

	obj.Angle += 0.1
	obj.Position.X += 0.5
}

func (obj *Snail) Draw(screen gm.Screen) {
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %v\n", frame.Anchor), 0, 24)
	// bSprite.Angle = sys.AngleFromVector(obj.Vel.X, obj.Vel.Y)
	// obj.Common.Draw(screen)

	// Draw text
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Angle: %.2f", deg), 12, 24)

}
