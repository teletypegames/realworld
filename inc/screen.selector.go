package inc

import (
	"strings"

	inkwell "git.teletypegames.org/engines/inkwell"
)

// selector is the wiki's Helyszínválasztó: "nem valódi helyszín, hanem a
// menü-képernyő, ahonnan az összes fő helyszín elérhető"
// (story#helyszínválasztó). The location graph is centred on it — every main
// location connects to it both ways, and physical adjacency takes over from
// there inwards — so without it the city is a handful of disconnected rooms.
//
// It is not in the concept-art deck and has no painting, so it renders as an
// inkwell placeholder with one labelled pin per destination, laid out on a
// grid. Whatever the map ends up looking like, the connections it carries are
// the ones the screens declare, not a list kept here: screenSelector() is built from
// whichever screens marked themselves onSelector.
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
		// Nobody stands on a map.
		actors:    []inkwell.SceneActor{},
		walkboxes: []inkwell.Polygon{},
	}
}

// The pin grid: three columns down the scene region, filled top to bottom.
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

// pinLabel is the screen's title without the em-dash subtitle: on a map a
// location wants its name, not its description.
func pinLabel(title string) string {
	name, _, _ := strings.Cut(title, " — ")
	return name
}
