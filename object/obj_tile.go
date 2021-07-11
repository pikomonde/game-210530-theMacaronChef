package object

import (
	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
)

type Tile struct{}

func (obj *Tile) Init() {
	// Set Bevhaviour
	// gm.SetAndInitBehaviours(obj, &behaviour.Common{
	// 	Position: r2.Point{100, 100},
	// }, &behaviour.Sprite{
	// 	Size:   r2.Point{12, 12},
	// 	Anchor: r2.Point{6, 6},
	// })

	// Init Sprite Behaviour
	// bSprite := gm.MustGetBehaviour(obj, &behaviour.Sprite{}).(*behaviour.Sprite)
	// bSprite.FillColor(color.NRGBA{0x00, 0x00, 0x00, 0xff})

}

func (obj *Tile) Update() {

}

func (obj *Tile) Draw(screen gm.Screen) {
	// bSprite := gm.MustGetBehaviour(obj, &behaviour.Sprite{}).(*behaviour.Sprite)
	// bSprite.Draw(screen)
}
