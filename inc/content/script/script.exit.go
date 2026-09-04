package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
)

var exitLook = Text{
	{Lang: "en", Text: "That way: %s."},
	{Lang: "hu", Text: "Arra van: %s."},
}

var exitTake = Text{
	{Lang: "en", Text: "It's a way out, not a thing."},
	{Lang: "hu", Text: "Az egy kijárat, nem egy tárgy."},
}

func ExitLook(exit inkwell.Exit) Action {
	return character.Paul.Say(inkwell.Format(exitLook, exit.Label))
}

func ExitTake(inkwell.Exit) Action {
	return character.Paul.Say(exitTake)
}
