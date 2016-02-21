package ligge

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	LastKeyPress sdl.Keycode
}

func (i *Input) RecordKeyPress(event *sdl.KeyUpEvent) {
	i.LastKeyPress = event.Keysym.Sym
	Gui.Scene.KeyPress(event.Keysym.Sym)
}

func (i *Input) ProcessEvent(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		Gui.Stop()
	case *sdl.KeyDownEvent:
		Gui.Scene.KeyDown(t.Keysym.Sym)
	case *sdl.KeyUpEvent:
		Gui.Scene.KeyUp(t.Keysym.Sym)
		i.RecordKeyPress(t)
	case *sdl.WindowEvent:
		fmt.Printf("Got a WindowEvent.\n")
	default:
		// fmt.Printf("Unsupported event type %T\n", event)
	}
}

func (i *Input) Run() {
	println("Waiting for events")
	select {
	case event := <-Gui.Events:
		fmt.Println("Received event", event)
		i.ProcessEvent(event)
	}

}

func NewInput() *Input {
	i := Input{
		sdl.K_UNKNOWN,
	}

	return &i
}