package btn

import "github.com/hajimehoshi/ebiten/v2"

var isPrevMousePressed = make(map[ebiten.MouseButton]bool)
var isMouseButtonClicked = make(map[ebiten.MouseButton]bool)
var isMouseButtonReleased = make(map[ebiten.MouseButton]bool)

func updateMousPrevByMouseID(mouse ebiten.MouseButton) {
	isMouseButtonClicked[mouse] = !isPrevMousePressed[mouse] && ebiten.IsMouseButtonPressed(mouse)
	isMouseButtonReleased[mouse] = isPrevMousePressed[mouse] && !ebiten.IsMouseButtonPressed(mouse)
	isPrevMousePressed[mouse] = ebiten.IsMouseButtonPressed(mouse)
}

func updateMousePrev() {
	updateMousPrevByMouseID(ebiten.MouseButtonLeft)
	updateMousPrevByMouseID(ebiten.MouseButtonRight)
	updateMousPrevByMouseID(ebiten.MouseButtonMiddle)
}

func Update() {
	updateMousePrev()
}

func IsMouseButtonPressed(mouse ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouse)
}

func IsMouseButtonClicked(mouse ebiten.MouseButton) bool {
	return isMouseButtonClicked[mouse]
}

func IsMouseButtonReleased(mouse ebiten.MouseButton) bool {
	return isMouseButtonReleased[mouse]
}
