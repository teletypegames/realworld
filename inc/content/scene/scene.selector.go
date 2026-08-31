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
	SceneManager.Register(Scene{
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

type selector struct{}

func FillSelectorPins() {
	s, ok := SceneManager.Get(inc.SceneSelector)
	if !ok {
		return
	}
	sel := selector{}
	var exits []inkwell.Exit
	for _, entity := range SceneManager.All() {
		if !sel.leadsTo(entity) {
			continue
		}
		exits = append(exits, inkwell.Exit{
			To:    entity.Name,
			Label: sel.pinLabel(entity.Title),
			Area:  sel.pin(len(exits)),
		})
	}
	s.Exits = exits
	SceneManager.Set(s)
}

func (selector) leadsTo(s Scene) bool {
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

func (selector) pin(i int) inkwell.Shape {
	col, row := i%pinCols, i/pinCols
	return inkwell.Rect(pinX+float64(col)*pinGapX, pinY+float64(row)*pinGapY, pinW, pinH)
}

func (selector) pinLabel(title inkwell.Text) inkwell.Text {
	return title.Map(func(s string) string {
		name, _, _ := strings.Cut(s, " — ")
		return name
	})
}
