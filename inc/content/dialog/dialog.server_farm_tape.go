package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/speaker"
)

var ServerFarmTape = DialogManager.Register(&Dialog{
	Name:  "server_farm_tape",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: speaker.TapeServerFarm,
			Text: inkwell.Text{
				{Lang: "en", Text: "CREDENTIAL. DOOR GROUP: DATACENTER-WEST. AWAITING READER."},
				{Lang: "hu", Text: "HITELESÍTŐ. AJTÓCSOPORT: ADATKÖZPONT-NYUGAT. OLVASÓRA VÁR."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Whose badge were you?"},
					{Lang: "hu", Text: "Kinek a kitűzője voltál?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(speaker.TapeServerFarm, inkwell.Text{
						{Lang: "en", Text: "FIELD NOT PRESENT."},
						{Lang: "hu", Text: "A MEZŐ NEM LÉTEZIK."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Dex, get it to say something. Anything."},
					{Lang: "hu", Text: "Dex, vedd rá, hogy mondjon valamit. Bármit."},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Hello in there. Favourite colour? Regrets? A door you wish you'd opened?"},
						{Lang: "hu", Text: "Halló, odabent. Kedvenc szín? Megbánások? Egy ajtó, amit bárcsak kinyitottál volna?"},
					}),
					inkwell.Say(speaker.TapeServerFarm, inkwell.Text{
						{Lang: "en", Text: "DOOR GROUP: DATACENTER-WEST."},
						{Lang: "hu", Text: "AJTÓCSOPORT: ADATKÖZPONT-NYUGAT."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Nothing in there. I've met stones with more inner life. I'm relieved, honestly."},
						{Lang: "hu", Text: "Semmi nincs benne. Találkoztam már kövekkel, akiknek több belső életük volt. Megkönnyebbültem, komolyan."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Fine."},
					{Lang: "hu", Text: "Jó."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(speaker.TapeServerFarm, inkwell.Text{
						{Lang: "en", Text: "AWAITING READER."},
						{Lang: "hu", Text: "OLVASÓRA VÁR."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
