package inc

// Exits — the connections between screens.
//
// Where a screen leads is recorded in that screen's own file, and it comes from
// the wiki's location graph (projects/realworld/story#helyszín-gráf-vázlat).
// That graph has two kinds of edge, and so does this file:
//
//   - Physical adjacency. The alley is behind Noodle's house; the club is
//     through the back of the trinket shop; the noodle place is next door to
//     it. These are the arrows the wiki draws between locations, and they are
//     declared with `to`.
//   - The selector. The wiki centres its map on a "Helyszínválasztó" — not a
//     real location but the menu screen every main location connects to both
//     ways. A screen marks itself with onSelector; selector.go builds the other
//     half of each of those edges.
//
// Some of the wiki's arrows are one-way with a condition on them ("beengedés",
// "NG+ kód", "Sumphor liftkódja"). `needs` carries that: the exit is there, and
// says so, but will not open until the flag is set.
//
// Until the doors are measured on the paintings, an exit is a strip along one
// edge of the picture — the convention every 1990s point & click used, and a
// one-field change to re-aim at a real door later.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type side int

const (
	sideLeft  side = iota // off the left edge
	sideRight             // off the right edge
	sideBack              // into the depth of the picture
	sideNear              // out towards the camera
)

// area is the strip of picture the exit occupies, inside the scene region
// (below the top bar at y=20, above the HUD divider at y=262).
func (s side) area() inkwell.Shape {
	switch s {
	case sideLeft:
		return inkwell.Rect(0, 40, 56, 212)
	case sideRight:
		return inkwell.Rect(584, 40, 56, 212)
	case sideBack:
		return inkwell.Rect(232, 28, 176, 84)
	default: // sideNear
		return inkwell.Rect(232, 214, 176, 48)
	}
}

// exit is one way out of a screen.
type exit struct {
	to    string // the screen it leads to; empty means "back the way you came"
	label string // what the status line calls it
	side  side

	// area overrides the edge strip. The selector's map pins use it; nothing
	// else should need to until the art arrives.
	area inkwell.Shape

	// needs is a flag that has to be set before the exit opens, and blocked is
	// what Paul says while it is not.
	needs   string
	blocked string
}

// toSelector is the way back out to the rest of the city, on every screen the
// wiki's Helyszínválasztó lists.
func toSelector() exit {
	return exit{to: ScreenSelector, label: "the rest of the city", side: sideNear}
}

// exitName is how a connection stays machine-readable: the graph can be read
// straight back out of the registered scenes, so a test can walk it.
func exitName(to string) string {
	if to == "" {
		return "exit:back"
	}
	return "exit:" + to
}

func (e exit) hotspot(w *World) inkwell.Hotspot {
	var travel inkwell.Action = inkwell.GoTo(e.to)
	if e.to == "" {
		travel = Back(w)
	}
	if e.needs != "" {
		travel = inkwell.If(inkwell.Flag(e.needs), travel, inkwell.Say(Paul, e.blocked))
	}
	area := e.area
	if area == nil {
		area = e.side.area()
	}
	return inkwell.Hotspot{
		Name:   exitName(e.to),
		Label:  e.label,
		Area:   area,
		Cursor: inkwell.CursorExit,
		OnLook: inkwell.Say(Paul, "That way: "+e.label+"."),
		OnUse:  travel,
		OnTake: inkwell.Say(Paul, "It's a way out, not a thing."),
	}
}
