package ligge

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

var guiComponents []guiComponent = []guiComponent{}
var Gui *GUI = nil

type guiComponent struct {
	Init func()
}

type GUI struct {
	Running      bool
	Events       chan sdl.Event
	Scene        Scene
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	Input        *Input
	Reload       bool
	Width        int32
	Height       int32
	frameStart   time.Time
	targetFPS    int64
	timePerFrame time.Duration
}

func (g *GUI) SetFrameStart() {
	g.frameStart = time.Now()
}

func (g *GUI) WaitAndSync() {
	maxSleep := 3 * time.Millisecond

	for {
		g.ProcessEvents()

		sleep := g.timePerFrame - time.Since(g.frameStart)

		if sleep > maxSleep {
			time.Sleep(maxSleep)
		} else {
			time.Sleep(sleep)
			break
		}
	}

	g.SetFrameStart()
}

func (g *GUI) Stop() {
	g.Running = false
}

func (g *GUI) GetSize() (int32, int32) {
	return g.Width, g.Height
}

func (g *GUI) ProcessEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		g.Input.ProcessEvent(event)
	}
}

func (g *GUI) SetScene(name string) {
	scene := GetScene(name)
	g.Scene = scene
	g.Scene.Activate()
}

func (g *GUI) RenderScene() {
	g.Scene.Render()
}

func (g *GUI) Run(done chan bool) {
	g.Running = true
	g.SetFrameStart()

	for {
		g.RenderScene()
		g.Renderer.Present()
		g.WaitAndSync()

		if g.Running == false {
			println("Stopping GUI.")
			break
		}

		if g.Reload == true {
			println("Reloading GUI.")
			break
		}
	}
}

func (g *GUI) Destroy() {
	g.Window.Destroy()
}

func (g *GUI) DrawText(fontName string, text string, x int32, y int32, color sdl.Color) {
	DrawText(fontName, text, x, y, color, g.Renderer)
}

func NewGUI(input *Input) *GUI {
	// https://wiki.libsdl.org/SDL_WindowFlags
	var flags uint32 = sdl.WINDOW_SHOWN

	var fps int64 = 60

	var width int32 = 1280
	var height int32 = 720

	g := GUI{
		false,
		make(chan sdl.Event),
		nil,
		nil,
		nil,
		input,
		false,
		width,
		height,
		time.Now(),
		fps,
		time.Second / time.Duration(fps),
	}

	window, err := sdl.CreateWindow("Ligge", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int(width), int(height), flags)
	if err != nil {
		panic(err)
	}

	g.Window = window

	g.Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED);

	if err != nil {
		panic(err);
	}

	for _, component := range guiComponents {
		component.Init()
	}

	return &g
}

func RunGUI(onGUIReady func(), done chan bool, input *Input) bool {

	gui := NewGUI(input)
	defer gui.Destroy()

	Gui = gui

	onGUIReady()

	gui.SetScene("mainmenu")
	gui.Run(done)

	return gui.Reload
}

func RegisterGuiComponent(init func()) {
	guiComponents = append(guiComponents, guiComponent{init})
}
