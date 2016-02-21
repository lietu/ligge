package ligge

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var fonts map[string]*ttf.Font = map[string]*ttf.Font{}

var fontSizes map[string]int = map[string]int{}
var fontPaths map[string]string = map[string]string{}

func LoadFont(name string, size int, path string) *ttf.Font {
	font, err := ttf.OpenFont(path, size)

	if err != nil {
		panic(err)
	}

	return font
}

func RegisterFont(name string, size int, path string) {
	_, ok := fontSizes[name]

	if ok {
		log.Fatalf("Trying to re-register font %s", name)
	}

	fontSizes[name] = size
	fontPaths[name] = path
}

func GetFont(name string) *ttf.Font {
	font, ok := fonts[name]

	if ok == false {
		font = LoadFont(name, fontSizes[name], fontPaths[name])
		fonts[name] = font
	}

	return font
}

func DrawText(fontName string, text string, x int32, y int32, color sdl.Color, renderer *sdl.Renderer) {

	font := GetFont(fontName)

	surface, err := font.RenderUTF8_Blended(text, color)
	texture, err := renderer.CreateTextureFromSurface(surface)

	if err != nil {
		panic(err)
	}

	_, _, w, h, err := texture.Query()

	if err != nil {
		panic(err)
	}

	src := sdl.Rect{0, 0, w, h}
	dst := sdl.Rect{x, y, w, h}
	renderer.Copy(texture, &src, &dst)
}

func UnloadFonts() {
}

func init() {

}
