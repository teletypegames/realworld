package inc

import (
	"image"
	"image/color"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	glyphW = 6
	glyphH = 16
)

var scratch *ebiten.Image

func textW(s string) int { return utf8.RuneCountInString(s) * glyphW }

func drawTextC(dst *ebiten.Image, s string, x, y int, c color.Color) {
	if s == "" {
		return
	}
	w, h := textW(s)+glyphW, glyphH
	if scratch == nil || scratch.Bounds().Dx() < w || scratch.Bounds().Dy() < h {
		nw, nh := w, h
		if nw < 320 {
			nw = 320
		}
		scratch = ebiten.NewImage(nw, nh)
	}
	scratch.Clear()
	ebitenutil.DebugPrintAt(scratch, s, 0, 0)

	if c == nil {
		c = color.White
	}
	r, g, b, a := c.RGBA()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.Scale(
		float32(r)/0xffff, float32(g)/0xffff, float32(b)/0xffff, float32(a)/0xffff,
	)
	dst.DrawImage(scratch.SubImage(image.Rect(0, 0, w, h)).(*ebiten.Image), op)
}

func wrap(s string, maxPx int) []string {
	if maxPx <= glyphW || textW(s) <= maxPx {
		return []string{s}
	}
	var out []string
	line, word := "", ""
	flush := func() {
		if line != "" {
			out = append(out, line)
			line = ""
		}
	}
	emit := func() {
		if word == "" {
			return
		}
		switch {
		case line == "":
			line = word
		case textW(line+" "+word) <= maxPx:
			line += " " + word
		default:
			flush()
			line = word
		}
		word = ""
	}
	for _, r := range s {
		switch r {
		case ' ':
			emit()
		case '\n':
			emit()
			flush()
		default:
			word += string(r)
		}
	}
	emit()
	flush()
	if len(out) == 0 {
		return []string{s}
	}
	return out
}

func clip(s string, maxPx int) string {
	r := []rune(s)
	max := maxPx / glyphW
	switch {
	case max <= 0:
		return ""
	case len(r) <= max:
		return s
	case max == 1:
		return string(r[:1])
	default:
		return string(r[:max-1]) + "…"
	}
}
