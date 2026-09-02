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

type selector struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneSelector,
	Title: inkwell.Text{
		{Lang: "en", Text: "SAN FRANCISCO"},
		{Lang: "hu", Text: "SAN FRANCISCO"},
	},
	Actors:    []inkwell.SceneActor{},
	Walkboxes: []inkwell.Polygon{},
	Derived:   selector{}.fromGraph,
})

func (selector) fromGraph(game *inkwell.Game) []inkwell.Exit {
	pins := selector{}
	var exits []inkwell.Exit
	for _, entity := range game.SceneManager.All() {
		if !pins.leadsTo(entity) {
			continue
		}
		exits = append(exits, inkwell.Exit{
			To:    entity.Name,
			Label: pins.label(entity.Title),
			Area:  pins.area(len(exits)),
		})
	}
	return exits
}

func (selector) leadsTo(scene Scene) bool {
	for _, exit := range scene.Exits {
		if exit.To == inc.SceneSelector {
			return true
		}
	}
	return false
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

func (selector) area(index int) inkwell.Shape {
	column, row := index%pinColumns, index/pinColumns
	return inkwell.Rect(pinX+float64(column)*pinGapX, pinY+float64(row)*pinGapY, pinWidth, pinHeight)
}

func (selector) label(title inkwell.Text) inkwell.Text {
	return title.Map(func(title string) string {
		name, _, _ := strings.Cut(title, " — ")
		return name
	})
}
