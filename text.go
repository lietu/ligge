package ligge

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Text struct {
	Image *Image
}

func (t *Text) Render(x int32, y int32) {
	t.Image.Render(x, y)
}

func CreateText(fontName string, text string, color sdl.Color) *Text {
	font := GetFont(fontName)

	surface, err := font.RenderUTF8_Blended(text, color)
	texture, err := Gui.Renderer.CreateTextureFromSurface(surface)

	if err != nil {
		panic(err)
	}

	image := ImageFromTexture(texture)

	return NewText(image)
}

func NewText(image *Image) *Text {
	t := Text{
		image,
	}

	return &t
}


