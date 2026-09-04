package scene

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"
)

var exitToSelector = Exit{
	To: Selector,
	Label: Text{
		{Lang: "en", Text: "the rest of the city"},
		{Lang: "hu", Text: "a város többi része"},
	},
	Side: inkwell.ExitNear,
}

var Selector = SceneManager.Register(&Scene{
	Name: "selector",
	Title: Text{
		{Lang: "en", Text: "SAN FRANCISCO"},
		{Lang: "hu", Text: "SAN FRANCISCO"},
	},
	Actors:    []SceneActor{},
	Walkboxes: []Polygon{},
	Derived:   selectorDerived,
})

func selectorDerived(game *inkwell.Game, self *Scene) []Exit {
	leadsHere := func(scene *Scene) bool {
		for _, exit := range scene.Exits {
			if exit.To == self {
				return true
			}
		}
		return false
	}
	pin := func(index int) inkwell.Shape {
		column, row := index%pinColumns, index/pinColumns
		return inkwell.Rect(pinX+float64(column)*pinGapX, pinY+float64(row)*pinGapY, pinWidth, pinHeight)
	}
	shortTitle := func(title Text) Text {
		return title.Map(func(title string) string {
			name, _, _ := strings.Cut(title, " — ")
			return name
		})
	}
	var exits []Exit
	for _, entity := range game.SceneManager.All() {
		if !leadsHere(entity) {
			continue
		}
		exits = append(exits, Exit{
			To:    entity,
			Label: shortTitle(entity.Title),
			Area:  pin(len(exits)),
		})
	}
	return exits
}

const (
	pinColumns = 4
	pinWidth   = 148.0
	pinHeight  = 32.0
	pinX       = 16.0
	pinY       = 28.0
	pinGapX    = 156.0
	pinGapY    = 38.0
)
