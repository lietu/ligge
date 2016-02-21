package ligge

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var IMAGE_ASSETS = "../assets/images"

var cache map[string]*sdl.Texture = map[string]*sdl.Texture{}

func GetImageTexture(path string) *sdl.Texture {
	texture, ok := cache[path]

	if !ok {
		path = fmt.Sprintf("%s/%s.png", IMAGE_ASSETS, path)
		texture = loadImageTexture(path)
	}

	return texture
}

func loadImageTexture(path string) *sdl.Texture {
	testSurface, err := img.Load(path)

	if err != nil {
		panic(err)
	}

	texture, err := Gui.Renderer.CreateTextureFromSurface(testSurface)

	cache[path] = texture

	return texture
}