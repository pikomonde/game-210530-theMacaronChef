package object

import (
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/behaviour"
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Snail struct{}

func (obj *Snail) Init() {
	// Set Bevhaviour
	gm.SetAndInitBehaviours(obj, &behaviour.Common{
		Position: r2.Point{60, 60},
	}, &behaviour.Sprite{
		Size:   r2.Point{12, 12},
		Anchor: r2.Point{6, 6},
	})

	// Init Sprite Behaviour
	// bSprite := obj.D().GetBehaviour(&behaviour.Sprite{}).(*behaviour.Sprite)
	bSprite := gm.MustGetBehaviour(obj, &behaviour.Sprite{}).(*behaviour.Sprite)
	bSprite.FillColor(color.NRGBA{0x00, 0x80, 0x00, 0xff})

	// Assign value to the object
	// obj.Vel = r2.Point{0, 0}
}

func (obj *Snail) Update() {

}

func (obj *Snail) Draw(screen gm.Screen) {
	// bSprite := obj.D().GetBehaviour(&behaviour.Sprite{}).(*behaviour.Sprite)
	bSprite := gm.MustGetBehaviour(obj, &behaviour.Sprite{}).(*behaviour.Sprite)
	// bSprite.Angle = sys.AngleFromVector(obj.Vel.X, obj.Vel.Y)
	bSprite.Draw(screen)

	// Draw text
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Angle: %.2f", deg), 12, 24)

}
