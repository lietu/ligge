package ligge

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SceneComponent interface {
	Init()
	Render()
	KeyDown(code sdl.Keycode)
	KeyUp(code sdl.Keycode)
	KeyPress(code sdl.Keycode)
}

type BaseSceneComponent struct {

}

func (b *BaseSceneComponent) KeyDown(code sdl.Keycode) {

}

func (b *BaseSceneComponent) KeyUp(code sdl.Keycode) {

}

func (b *BaseSceneComponent) KeyPress(code sdl.Keycode) {

}

func NewBaseSceneComponent() BaseSceneComponent {
	b := BaseSceneComponent{}

	return b
}
