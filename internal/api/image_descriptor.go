package api

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ImageDescriptor struct {
	OffsetX        int32
	OffsetY        int32
	ImageName      string
	DisplayOnPress sdl.Keycode
}

var (
	SomeImages = []ImageDescriptor{
		//background
		{
			OffsetX:        90,
			OffsetY:        50,
			ImageName:      "pg2_back.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
	}
)
