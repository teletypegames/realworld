package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var ArmillaFound = ScriptManager.Register(&Script{
	Name: "armilla_found",
	Actions: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "His Armilla. On the desk, charger still in."},
			{Lang: "hu", Text: "Az Armillája. Az asztalon, a töltő még benne."},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "Nobody leaves the house without their Armilla. So he didn't leave. He was taken."},
			{Lang: "hu", Text: "Senki nem megy el otthonról az Armillája nélkül. Tehát nem ment el. Elvitték."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Slot one is empty. His Personal Tape is gone."},
			{Lang: "hu", Text: "Az egyes foglalat üres. A Személyes Kazettája nincs meg."},
		}),
		inkwell.If(inkwell.Flag(inc.FlagTapeSpeaks),
			inkwell.Seq(
				inkwell.Say(speaker.TapeMystery, inkwell.Text{
					{Lang: "en", Text: "The BVK will have confiscated it. They take the tape first. It is what they do."},
					{Lang: "hu", Text: "A BVK lefoglalta. Először a kazettát viszik el. Ez a módszerük."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Convenient."},
					{Lang: "hu", Text: "Kényelmes."},
				}),
			),
			inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Whoever took him wanted the tape more than the wrist."},
				{Lang: "hu", Text: "Aki elvitte, a kazettát akarta, nem a csuklót."},
			}),
		),
		inkwell.SetFlag(inc.FlagArmillaSeen),
	),
})
