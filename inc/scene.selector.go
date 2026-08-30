package inc

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"
)

var exitToSelector = inkwell.Exit{
	To:    SceneSelector,
	Label: "the rest of the city",
	Side:  inkwell.ExitNear,
}

func init() {
	SceneManager.Register(Scene{
		Name:       SceneSelector,
		Title:      "SAN FRANCISCO",
		Background: BgSelector,
		Actors:     []inkwell.SceneActor{},
		Walkboxes:  []inkwell.Polygon{},
	})
}

func fillSelectorPins() {
	selector, ok := SceneManager.Get(SceneSelector)
	if !ok {
		return
	}
	var exits []inkwell.Exit
	for _, entity := range SceneManager.All() {
		if !leadsToSelector(entity) {
			continue
		}
		exits = append(exits, inkwell.Exit{
			To:    entity.Name,
			Label: pinLabel(entity.Title),
			Area:  pin(len(exits)),
		})
	}
	selector.Exits = exits
	SceneManager.Set(selector)
}

func leadsToSelector(s Scene) bool {
	for _, e := range s.Exits {
		if e.To == SceneSelector {
			return true
		}
	}
	return false
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
