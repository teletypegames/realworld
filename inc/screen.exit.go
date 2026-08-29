package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type side int

const (
	sideLeft side = iota
	sideRight
	sideBack
	sideNear
)

func (s side) area() inkwell.Shape {
	switch s {
	case sideLeft:
		return inkwell.Rect(0, 40, 56, 212)
	case sideRight:
		return inkwell.Rect(584, 40, 56, 212)
	case sideBack:
		return inkwell.Rect(232, 28, 176, 84)
	default:
		return inkwell.Rect(232, 214, 176, 48)
	}
}

type exit struct {
	to    string
	label string
	side  side

	area inkwell.Shape

	needs   string
	blocked string
}

func toSelector() exit {
	return exit{
		to:    ScreenSelector,
		label: "the rest of the city",
		side:  sideNear,
	}
}

func exitName(to string) string {
	if to == "" {
		return "exit:back"
	}
	return "exit:" + to
}

func (e exit) hotspot() inkwell.Hotspot {
	var travel inkwell.Action = inkwell.GoTo(e.to)
	if e.to == "" {
		travel = Back()
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
