package inc

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"
)

func screenSelector(rest []screen) screen {
	var exits []exit
	for _, s := range rest {
		if !s.onSelector {
			continue
		}
		exits = append(exits, exit{
			to:    s.name,
			label: pinLabel(s.title),
			area:  pin(len(exits)),
		})
	}
	return screen{
		name:       ScreenSelector,
		title:      "SAN FRANCISCO",
		background: BgSelector,
		exits:      exits,
		actors:     []inkwell.SceneActor{},
		walkboxes:  []inkwell.Polygon{},
	}
}

const (
	pinCols = 3
	pinW    = 192.0
	pinH    = 32.0
	pinX    = 16.0
	pinY    = 28.0
	pinGapX = 204.0
	pinGapY = 38.0
)

func pin(i int) inkwell.Shape {
	col, row := i%pinCols, i/pinCols
	return inkwell.Rect(pinX+float64(col)*pinGapX, pinY+float64(row)*pinGapY, pinW, pinH)
}

func pinLabel(title string) string {
	name, _, _ := strings.Cut(title, " — ")
	return name
}
