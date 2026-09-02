package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptMirrorPassword,
	Actions: inkwell.Seq(
		inkwell.Say(inc.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "I see the device is flickering, sir. Have you tried turning it off and on again?"},
			{Lang: "hu", Text: "Látom, az eszköz vibrál, uram. Próbálta már ki- és bekapcsolni?"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Don't you dare listen to it."},
			{Lang: "hu", Text: "Ne merészelj hallgatni rá."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Off. And… on."},
			{Lang: "hu", Text: "Ki. És… be."},
		}),
		inkwell.Wait(0.8),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It didn't come back. The tube's gone. The glass is black, and there's something glowing in it."},
			{Lang: "hu", Text: "Nem jött vissza. A cső kiment. Az üveg fekete, és világít benne valami."},
		}),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "A sun. Under waves. And a word painted beneath it: Lector."},
			{Lang: "hu", Text: "Egy nap. Hullámok alatt. Alatta egy szó festve: Lector."},
		}),
		inkwell.Say(inc.TapeSupport, inkwell.Text{
			{Lang: "en", Text: "I'm glad I could resolve your issue today. Is there anything else I can help you with?"},
			{Lang: "hu", Text: "Örülök, hogy ma megoldhattam a problémáját. Segíthetek még valamiben?"},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "It was right. Twenty dollars of scripted nonsense was right, and I have to live in the slot next to it."},
			{Lang: "hu", Text: "Igaza volt. Húsz dollár betanult hablatynak igaza volt, és nekem a mellette lévő foglalatban kell élnem."},
		}),
		inkwell.SetFlag(inc.FlagMirrorPassword),
	),
})
