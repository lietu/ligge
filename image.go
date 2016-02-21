package ligge

import (
	"github.com/veandco/go-sdl2/sdl"
)

var imageCache map[string]*Image = map[string]*Image{}

type Image struct {
	Width    int32
	Height   int32
	Texture  *sdl.Texture
	Rect     *sdl.Rect
}

func (i *Image) Render(x int32, y int32) {
	dst := sdl.Rect{x, y, i.Width, i.Height}
	Gui.Renderer.Copy(i.Texture, i.Rect, &dst)
}

func ImageFromPath(path string) *Image {
	texture := GetImageTexture(path)
	return ImageFromTexture(texture)
}

func ImageFromTexture(texture *sdl.Texture) *Image {
	_, _, w, h, err := texture.Query()

	if err != nil {
		panic(err)
	}

	return NewImage(w, h, texture)
}

func NewImage(w int32, h int32, texture *sdl.Texture) *Image {
	i := Image{
		w,
		h,
		texture,
		&sdl.Rect{0, 0, w, h},
	}

	return &i
}

func GetImage(path string) *Image {
	img, ok := imageCache[path]

	if !ok {
		img = ImageFromPath(path)
		imageCache[path] = img
	}

	return img
}