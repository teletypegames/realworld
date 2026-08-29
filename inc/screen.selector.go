package inc

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"
)

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenSelector,
		Title:      "SAN FRANCISCO",
		Background: BgSelector,
		Actors:     []inkwell.SceneActor{},
		Walkboxes:  []inkwell.Polygon{},
	})
}

func fillSelectorPins() {
	selector, ok := ScreenManager.GetByName(ScreenSelector)
	if !ok {
		return
	}
	var exits []exit
	for _, entity := range ScreenManager.GetAll() {
		if !entity.OnSelector {
			continue
		}
		exits = append(exits, exit{
			to:    entity.Name,
			label: pinLabel(entity.Title),
			area:  pin(len(exits)),
		})
	}
	selector.Exits = exits
	ScreenManager.Register(selector)
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
