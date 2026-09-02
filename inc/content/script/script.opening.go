package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptOpening,
	Actions: inkwell.Seq(
		system.System.SetMode(system.ModeCutscene),
		inkwell.Wait(0.6),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Something came under the door. Paper. Nobody sends you paper."},
			{Lang: "hu", Text: "Valami bejött az ajtó alatt. Papír. Neked senki nem küld papírt."},
		}),
		inkwell.Give(inc.ItemNoodleLetter),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "San Francisco postmark. Noodle's handwriting. Twenty years, and it's still terrible."},
			{Lang: "hu", Text: "San Franciscó-i bélyegző. Noodle kézírása. Húsz év, és még mindig rettenetes."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "“Come out, Paul. Not a phone conversation.” That's all."},
			{Lang: "hu", Text: "„Gyere ki, Paul. Ez nem telefontéma.” Ennyi."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "A man who traded cracked Armillas for a living has something he can't say on a phone. Imagine."},
			{Lang: "hu", Text: "Egy ember, aki feltört Armillákból élt, valamit nem mer telefonba mondani. Képzeld."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "He never asked me for anything. Not once, in twenty years."},
			{Lang: "hu", Text: "Soha nem kért tőlem semmit. Húsz év alatt egyszer sem."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "So you're going."},
			{Lang: "hu", Text: "Szóval mész."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm going. Grab what's worth grabbing and lock up."},
			{Lang: "hu", Text: "Megyek. Ami megéri, azt összeszedem, és zárok."},
		}),
		system.System.SetMode(system.ModePlay),
	),
})
