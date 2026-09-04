package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var Receptionist = DialogManager.Register(&Dialog{
	Name:  "receptionist",
	Start: "root",
	Nodes: []DialogNode{
		{
			Name: "root",
			Lines: []DialogLine{{
				Speaker: character.Receptionist,
				Text: Text{
					{Lang: "en", Text: "Yes."},
					{Lang: "hu", Text: "Igen."},
				},
			}},
			Choices: []DialogChoice{
				{
					Text: Text{
						{Lang: "en", Text: "I'm a friend of the man you brought in."},
						{Lang: "hu", Text: "A barátja vagyok annak, akit behoztak."},
					},
					Show: Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []Action{
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "Everybody is somebody's friend."},
							{Lang: "hu", Text: "Mindenki valakinek a barátja."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: Text{
						{Lang: "en", Text: "I'm here about the antiques dealer."},
						{Lang: "hu", Text: "A régiségkereskedő miatt jöttem."},
					},
					Show: Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []Action{
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "Detective Vasquez is handling that. She is not seeing anybody."},
							{Lang: "hu", Text: "Azt Vasquez nyomozó viszi. Nem fogad senkit."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: Text{
						{Lang: "en", Text: "I'm his lawyer."},
						{Lang: "hu", Text: "Én vagyok az ügyvédje."},
					},
					Show: Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []Action{
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "Are you."},
							{Lang: "hu", Text: "Valóban."},
						}),
						character.Paul.Say(Text{
							{Lang: "en", Text: "She's looking at the actual lawyer on the bench. Right."},
							{Lang: "hu", Text: "Az igazi ügyvédet nézi a padon. Értem."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: Text{
						{Lang: "en", Text: "Vasquez is expecting me."},
						{Lang: "hu", Text: "Vasquez vár engem."},
					},
					Show: Given{
						Flag: inc.FlagReceptionPassed,
					},
					Actions: []Action{
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "Through the door, sir."},
							{Lang: "hu", Text: "Az ajtón át, uram."},
						}),
						inkwell.EndDialog(),
					},
				},
				{
					Text: Text{
						{Lang: "en", Text: "Never mind."},
						{Lang: "hu", Text: "Hagyjuk."},
					},
					Actions: []Action{inkwell.EndDialog()},
				},
			},
		},
		{
			Name: "deflected",
			Lines: []DialogLine{{
				Speaker: character.Receptionist,
				Text: Text{
					{Lang: "en", Text: "Anything else."},
					{Lang: "hu", Text: "Még valami."},
				},
			}},
			Choices: []DialogChoice{
				{
					Text: Text{
						{Lang: "en", Text: "Let me try that again."},
						{Lang: "hu", Text: "Hadd próbáljam újra."},
					},
					Actions: []Action{inkwell.GotoNode("root")},
				},
				{
					Text: Text{
						{Lang: "en", Text: "Tell Vasquez that Paul is here. She'll want to know who else was on Noodle's list."},
						{Lang: "hu", Text: "Mondja meg Vasqueznek, hogy Paul van itt. Tudni akarja majd, ki más volt Noodle listáján."},
					},
					Show: Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []Action{
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "…"},
							{Lang: "hu", Text: "…"},
						}),
						character.Paul.Say(Text{
							{Lang: "en", Text: "She picks up the phone. Says two words. Listens for a lot longer."},
							{Lang: "hu", Text: "Felveszi a telefont. Két szót mond. Sokkal tovább hallgat."},
						}),
						character.Receptionist.Say(Text{
							{Lang: "en", Text: "She says go through."},
							{Lang: "hu", Text: "Azt mondja, menjen be."},
						}),
						inkwell.SetFlag(inc.FlagReceptionPassed),
						speaker.Dex.Say(Text{
							{Lang: "en", Text: "There was no list, Paul. There is one now."},
							{Lang: "hu", Text: "Nem volt lista, Paul. Most már van."},
						}),
						inkwell.EndDialog(),
					},
				},
				{
					Text: Text{
						{Lang: "en", Text: "No. Thank you."},
						{Lang: "hu", Text: "Nem. Köszönöm."},
					},
					Actions: []Action{inkwell.EndDialog()},
				},
			},
		},
	},
})
