package main

import (
	_ "image/png"
	"log"

	// "github.com/hajimehoshi/ebiten/v2/examples/resources/images"

	"github.com/pikomonde/game-210530-theMacaronChef/gogeta/gm"
	"github.com/pikomonde/game-210530-theMacaronChef/object"
)

func main() {
	// Initialize objects
	gm.Init()
	gm.SetAndInitObject(&object.Rule01{})
	// gm.Game.SetObjectAndInit(&object.Colony{})

	// Run game
	if err := gm.Run(); err != nil {
		log.Fatal("error game run: ", err)
	}
}
