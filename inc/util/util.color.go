package util

import (
	"image/color"
)

func RGB(hex uint32) color.Color {
	return color.RGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: 0xff,
	}
}

func RGBA(hex uint32, alpha uint8) color.Color {
	return color.RGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: alpha,
	}
}
