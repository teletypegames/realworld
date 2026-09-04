package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var NormanFound = ScriptManager.Register(&Script{
	Name: "norman_found",
	Actions: inkwell.Seq(
		system.System.SetMode(system.ModeCutscene),
		inkwell.Wait(0.5),
		character.Paul.Say(Text{
			{Lang: "en", Text: "There's somebody in there. In a chair. Wired into that machine like a lamp into a wall."},
			{Lang: "hu", Text: "Van ott bent valaki. Egy székben. Bekötve abba a gépbe, mint egy lámpa a falba."},
		}),
		inkwell.Wait(0.4),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "Paul. That's Norman."},
			{Lang: "hu", Text: "Paul. Az Norman."},
		}),
		inkwell.Wait(0.6),
		character.Paul.Say(Text{
			{Lang: "en", Text: "He isn't missing. He was never missing. He's been in here the whole time, asleep, running their program in his head."},
			{Lang: "hu", Text: "Nem tűnt el. Soha nem tűnt el. Itt volt egész idő alatt, alva, a fejében futtatva a programjukat."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "The simulation is him. Everything we read about it is happening behind that glass."},
			{Lang: "hu", Text: "A szimuláció ő maga. Minden, amit róla olvastunk, az üveg mögött történik."},
		}),
		inkwell.If(inkwell.Flag(inc.FlagTapeSpeaks),
			speaker.TapeMystery.Say(Text{
				{Lang: "en", Text: "Yes. Him. Now you see."},
				{Lang: "hu", Text: "Igen. Ő az. Most már látod."},
			}),
		),
		character.Paul.Say(Text{
			{Lang: "en", Text: "Then that's the job. Not finding him. Getting him out."},
			{Lang: "hu", Text: "Akkor ez a feladat. Nem megtalálni. Kihozni."},
		}),
		inkwell.SetFlag(inc.FlagNormanFound),
		inkwell.SetVar(inc.VarDecay, inc.DecayNormanFound),
		system.System.SetMode(system.ModePlay),
	),
})
