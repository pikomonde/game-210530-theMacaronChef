package behaviour_common

import (
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
)

type Frame struct {
	Image  *ebiten.Image
	Anchor r2.Point
}

type Animation []*Frame

type Animations map[string]Animation

type Sprite struct {
	Animations       Animations
	CurrentAnimation string
	CurrentFrame     int
}

func (bhvr *Sprite) Init() {
	// bhvr.Image = ebiten.NewImage(int(32), int(32))
	bhvr.Animations = make(Animations)
	bhvr.Animations[DefaultAnimationName] = make(Animation, 0)
	bhvr.CurrentAnimation = DefaultAnimationName
	bhvr.CurrentFrame = 0
}

func (bhvr *Sprite) GetCurrentFrame() *Frame {
	return bhvr.Animations[bhvr.CurrentAnimation][bhvr.CurrentFrame]
}

// === Sprite ===

const (
	DefaultAnimationName = "default"
)

// InsertFrame insert frame(s) at the end of the animation named
// "animationName" by Image. If "animationName" is empty, then it will use
// default value.
func (bhvr *Sprite) InsertFrame(animationName string, newFrames ...*Frame) {
	// validation(s)
	if animationName == "" {
		animationName = DefaultAnimationName
	}

	bhvr.Animations[animationName] = append(bhvr.Animations[animationName], newFrames...)
}

// InsertFrameByImage insert frame(s) at the end of the animation named
// "animationName" by Image. If "animationName" is empty, then it will use
// default value.
func (bhvr *Sprite) InsertFrameByImage(animationName string, images ...*ebiten.Image) {
	newFrames := make([]*Frame, 0)

	for _, image := range images {
		newFrame := &Frame{Image: image}
		newFrame.SetAnchorToggle(Sprite_FrameAnchor_ToggleDefault)
		newFrames = append(newFrames, newFrame)
	}

	bhvr.InsertFrame(animationName, newFrames...)
}

// === Behaviour specific method:  ===
// func (bhvr *Sprite) SetSize(x, y int) {
// 	bhvr.Image = ebiten.NewImage(x, y)
// }

// func (bhvr *Sprite) FillColor(colorNRGBA color.NRGBA) {
// 	bhvr.Image.Fill(colorNRGBA)
// }

// func (bhvr *Sprite) SetSizeAndFillColor(x, y int, colorNRGBA color.NRGBA) {
// 	bhvr.SetSize(x, y)
// 	bhvr.FillColor(colorNRGBA)
// }

// func (bhvr *Sprite) SetSprite(image *ebiten.Image) {
// 	bhvr.Image = image
// }

// === Frame ===

func (bhvr *Frame) SetAnchorToggle(pos FrameAnchorToggle) {
	w, h := bhvr.Image.Size()
	switch pos {
	case Sprite_FrameAnchor_ToggleTopLeft:
		bhvr.Anchor = r2.Point{0, 0}
	case Sprite_FrameAnchor_ToggleTopCenter:
		bhvr.Anchor = r2.Point{float64(w) / 2, 0}
	case Sprite_FrameAnchor_ToggleTopRight:
		bhvr.Anchor = r2.Point{float64(w), 0}
	case Sprite_FrameAnchor_ToggleMiddleLeft:
		bhvr.Anchor = r2.Point{0, float64(h) / 2}
	case Sprite_FrameAnchor_ToggleMiddleCenter:
		bhvr.Anchor = r2.Point{float64(w) / 2, float64(h) / 2}
	case Sprite_FrameAnchor_ToggleMiddleRight:
		bhvr.Anchor = r2.Point{float64(w), float64(h) / 2}
	case Sprite_FrameAnchor_ToggleBottomLeft:
		bhvr.Anchor = r2.Point{0, float64(h)}
	case Sprite_FrameAnchor_ToggleBottomCenter:
		bhvr.Anchor = r2.Point{float64(w) / 2, float64(h)}
	case Sprite_FrameAnchor_ToggleBottomRight:
		bhvr.Anchor = r2.Point{float64(w), float64(h)}
	}
}

type FrameAnchorToggle int

const (
	Sprite_FrameAnchor_ToggleTopLeft FrameAnchorToggle = iota
	Sprite_FrameAnchor_ToggleTopCenter
	Sprite_FrameAnchor_ToggleTopRight
	Sprite_FrameAnchor_ToggleMiddleLeft
	Sprite_FrameAnchor_ToggleMiddleCenter
	Sprite_FrameAnchor_ToggleMiddleRight
	Sprite_FrameAnchor_ToggleBottomLeft
	Sprite_FrameAnchor_ToggleBottomCenter
	Sprite_FrameAnchor_ToggleBottomRight

	Sprite_FrameAnchor_ToggleDefault = Sprite_FrameAnchor_ToggleMiddleCenter
)
