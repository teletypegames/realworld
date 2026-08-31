package theme

import (
	"image/color"

	inkwell "git.teletypegames.org/engines/inkwell"
)

type Theme = inkwell.Theme

var ThemeManager = inkwell.NewManager[Theme]()

func RGB(hex uint32) color.Color {
	return color.RGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: 0xff,
	}
}

func RGBA(hex uint32, a uint8) color.Color {
	return color.RGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: a,
	}
}
