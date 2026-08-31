package scene

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var exitToSelector = inkwell.Exit{
	To: inc.SceneSelector,
	Label: inkwell.Text{
		{Lang: "en", Text: "the rest of the city"},
		{Lang: "hu", Text: "a város többi része"},
	},
	Side: inkwell.ExitNear,
}

func init() {
	Manager.Register(Scene{
		Name: inc.SceneSelector,
		Title: inkwell.Text{
			{Lang: "en", Text: "SAN FRANCISCO"},
			{Lang: "hu", Text: "SAN FRANCISCO"},
		},
		Background: inc.BgSelector,
		Actors:     []inkwell.SceneActor{},
		Walkboxes:  []inkwell.Polygon{},
	})
}

func FillSelectorPins() {
	selector, ok := Manager.Get(inc.SceneSelector)
	if !ok {
		return
	}
	var exits []inkwell.Exit
	for _, entity := range Manager.All() {
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
	Manager.Set(selector)
}

func leadsToSelector(s Scene) bool {
	for _, e := range s.Exits {
		if e.To == inc.SceneSelector {
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

func pinLabel(title inkwell.Text) inkwell.Text {
	return title.Map(func(s string) string {
		name, _, _ := strings.Cut(s, " — ")
		return name
	})
}
