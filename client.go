package ligge

import (
	"runtime"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

func InitializeClientComponents() {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		panic(err)
	}

	img.Init(img.INIT_PNG)

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)

	if err != nil {
		panic(err)
	}
}

func RunClient(onGUIReady func()) {
	runtime.LockOSThread()
	done := make(chan bool)

	i := NewInput()

	for {
		reload := RunGUI(onGUIReady, done, i)
		if reload == false {
			break
		}
	}

	UnloadFonts()
}

func UninitializeClientComponents() {
	ttf.Quit()
	img.Quit()
    sdl.Quit()
}
