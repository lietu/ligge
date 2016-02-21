package ligge

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
)

var scenes map[string]Scene = map[string]Scene{}

type Scene interface {
	Activate()
	Render()
	KeyDown(sdl.Keycode)
	KeyUp(sdl.Keycode)
	KeyPress(sdl.Keycode)
}

type BaseScene struct {
	Initialized bool
	ComponentList []SceneComponent
}

func (b *BaseScene) RenderComponents() {
	for _, component := range b.ComponentList {
		component.Render()
	}
}

func (b *BaseScene) InitComponents() {
	for _, component := range b.ComponentList {
		component.Init()
	}
}

func (b *BaseScene) AddComponent(component SceneComponent) {
	b.ComponentList = append(b.ComponentList, component)
}

func (b *BaseScene) KeyDown(code sdl.Keycode) {
	for _, component := range b.ComponentList {
		component.KeyDown(code)
	}
}

func (b *BaseScene) KeyUp(code sdl.Keycode) {
	for _, component := range b.ComponentList {
		component.KeyUp(code)
	}
}

func (b *BaseScene) KeyPress(code sdl.Keycode) {
	for _, component := range b.ComponentList {
		component.KeyPress(code)
	}
}

func NewBaseScene() *BaseScene {
	b := BaseScene{
		false,
		[]SceneComponent{},
	}

	return &b
}

func GetScene(name string) Scene {
	item, ok := scenes[name]

	if ok == false {
		log.Fatalf("Failed to find scene %s", name)
	}

	return item
}

func RegisterScene(name string, scene Scene) {
	log.Printf("Registered scene %s", name)
	scenes[name] = scene
}