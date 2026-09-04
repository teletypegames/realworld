package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var Shopkeeper = DialogManager.Register(&Dialog{
	Name:  "shopkeeper",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Lines: []DialogLine{{
			Speaker: character.Shopkeeper,
			Text: Text{
				{Lang: "en", Text: "Welcome, welcome! Everything authentic. Everything good price. What you looking for?"},
				{Lang: "hu", Text: "Üdv, üdv! Minden eredeti. Minden jó ár. Mit keres?"},
			},
		}},
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "What's that Personal Tape on the shelf?"},
					{Lang: "hu", Text: "Mi az a Személyes Kazetta a polcon?"},
				},
				Show: Given{
					NotItem: item.SupportTape,
				},
				Once: true,
				Actions: []Action{
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Ah! Very useful. Tech support tape. Answers everything. Any question, any device. Twenty dollar."},
						{Lang: "hu", Text: "Á! Nagyon hasznos. Ügyfélszolgálati kazetta. Mindenre válaszol. Bármi kérdés, bármi eszköz. Húsz dollár."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "“Answers everything” is what you say about a thing that answers nothing. Don't."},
						{Lang: "hu", Text: "„Mindenre válaszol” — ezt arról mondják, ami semmire nem. Ne."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Twenty. Fine."},
						{Lang: "hu", Text: "Húsz. Jó."},
					}),
					inkwell.Give(item.SupportTape),
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Very good choice. No refund."},
						{Lang: "hu", Text: "Nagyon jó választás. Visszaváltás nincs."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "And the… thing next to it?"},
					{Lang: "hu", Text: "És az a… valami mellette?"},
				},
				Show: Given{
					NotItem: item.Trinket,
				},
				Once: true,
				Actions: []Action{
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Very rare. Very old. From the mountains. Five dollar, because I like you."},
						{Lang: "hu", Text: "Nagyon ritka. Nagyon régi. A hegyekből. Öt dollár, mert kedvelem."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "What does it do?"},
						{Lang: "hu", Text: "Mire jó?"},
					}),
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Sir, for five dollar, it does what you need."},
						{Lang: "hu", Text: "Uram, öt dollárért arra jó, amire kell."},
					}),
					inkwell.Give(item.Trinket),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "You bought it. You actually bought it. I'm going to remember this longer than you will."},
						{Lang: "hu", Text: "Megvetted. Tényleg megvetted. Erre tovább fogok emlékezni, mint te."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "I'm looking for a club. The underwater sun."},
					{Lang: "hu", Text: "Egy klubot keresek. A vízalatti napot."},
				},
				Show: Given{
					NotFlag: inc.FlagClubEntry,
				},
				Actions: []Action{
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Club? No club. Just a shop. You want a sun, I have sun. Charms, two dollar."},
						{Lang: "hu", Text: "Klub? Nincs klub. Csak bolt. Nap kell? Van napom. Amulett, két dollár."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "He answered a little too fast, and he sold you something on the way out of it."},
						{Lang: "hu", Text: "Kicsit túl gyorsan válaszolt, és közben még el is adott valamit."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "A sun under water. And the word is “Lector”."},
					{Lang: "hu", Text: "Nap a víz alatt. És a szó: „Lector”."},
				},
				Show: Given{
					Flag:    inc.FlagMirrorPassword,
					NotFlag: inc.FlagClubEntry,
				},
				Actions: []Action{
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "…Ah. You are here for the club. Why you don't say so. This way please, this way."},
						{Lang: "hu", Text: "…Á. A klubba jött. Miért nem ezzel kezdte. Erre tessék, erre."},
					}),
					inkwell.RunScript(script.ClubAdmission),
					inkwell.EndDialog(),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Just looking."},
					{Lang: "hu", Text: "Csak nézelődök."},
				},
				Actions: []Action{
					character.Shopkeeper.Say(Text{
						{Lang: "en", Text: "Looking is free. Touching is buying."},
						{Lang: "hu", Text: "Nézni ingyen van. Hozzányúlni vásárlás."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
