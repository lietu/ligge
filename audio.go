package ligge

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

var AUDIO_ASSETS = "../assets/audio"
var AUDIO_CATEGORIES map[string]int = map[string]int{}

type Sound struct {
	Category string
	Chunk *mix.Chunk
}

func (a *Sound) Play() {
	a.Chunk.Volume(AUDIO_CATEGORIES[a.Category])
	a.Chunk.Play(-1, 0)
}

func PlayMusic(path string) {
	mix.PauseMusic()
	music, err := mix.LoadMUS(fmt.Sprintf("%s/%s.wav", AUDIO_ASSETS, path))
	mix.VolumeMusic(AUDIO_CATEGORIES["music"])

	if err != nil {
		panic(err)
	}

	music.Play(9999999999)
}

func NewSound(category string, path string) *Sound {
	chunk, err := mix.LoadWAV(fmt.Sprintf("%s/%s.wav", AUDIO_ASSETS, path))

	if err != nil {
		panic(err)
	}

	a := Sound{
		category,
		chunk,
	}

	return &a
}

func RegisterAudioCategory(name string, volume int) {
	AUDIO_CATEGORIES[name] = volume
}

func FloatToVolume(f float32) int {
	return int(f * 128)
}
