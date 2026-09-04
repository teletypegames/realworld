package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var PrototypeLogs = ScriptManager.Register(&Script{
	Name: "prototype_logs",
	Actions: inkwell.Seq(
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "The log is mostly Norman talking to himself. “Tape behaving strangely again. Answers before I finish asking.”"},
			{Lang: "hu", Text: "A napló nagyrészt Norman magában beszél. „A kazetta megint furán viselkedik. Válaszol, mielőtt befejezném a kérdést.”"},
		}),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "His own tape, answering before the question. He wrote that down and kept using it anyway."},
			{Lang: "hu", Text: "A saját kazettája, ami a kérdés előtt válaszol. Leírta, és mégis tovább használta."},
		}),
		inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "And the last deployment note. “Prototype two installed on-site. Interior of the Golden Gate.”"},
			{Lang: "hu", Text: "És az utolsó telepítési jegyzet. „Kettes prototípus helyszínre telepítve. A Golden Gate belseje.”"},
		}),
		inkwell.SetFlag(inc.FlagLabLocated),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "A bridge everybody looks at. Of course that's where they hid it."},
			{Lang: "hu", Text: "Egy híd, amit mindenki néz. Persze, hogy oda rejtették."},
		}),
		inkwell.If(inkwell.Flag(inc.FlagTapeSpeaks),
			inkwell.Say(speaker.TapeMystery, inkwell.Text{
				{Lang: "en", Text: "Then that is where we go. Quickly, Paul. Every day in there is a day he does not come back from."},
				{Lang: "hu", Text: "Akkor oda megyünk. Gyorsan, Paul. Minden nap odabent egy nap, ahonnan nem tér vissza."},
			}),
		),
	),
})
