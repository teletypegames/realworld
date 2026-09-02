package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogReceptionist,
	Start: "root",
	Nodes: []inkwell.DialogNode{
		{
			Name: "root",
			Lines: []inkwell.DialogLine{{
				Speaker: inc.Receptionist,
				Text: inkwell.Text{
					{Lang: "en", Text: "Yes."},
					{Lang: "hu", Text: "Igen."},
				},
			}},
			Choices: []inkwell.DialogChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "I'm a friend of the man you brought in."},
						{Lang: "hu", Text: "A barátja vagyok annak, akit behoztak."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "Everybody is somebody's friend."},
							{Lang: "hu", Text: "Mindenki valakinek a barátja."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "I'm here about the antiques dealer."},
						{Lang: "hu", Text: "A régiségkereskedő miatt jöttem."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "Detective Vasquez is handling that. She is not seeing anybody."},
							{Lang: "hu", Text: "Azt Vasquez nyomozó viszi. Nem fogad senkit."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "I'm his lawyer."},
						{Lang: "hu", Text: "Én vagyok az ügyvédje."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "Are you."},
							{Lang: "hu", Text: "Valóban."},
						}),
						inkwell.Say(inc.Paul, inkwell.Text{
							{Lang: "en", Text: "She's looking at the actual lawyer on the bench. Right."},
							{Lang: "hu", Text: "Az igazi ügyvédet nézi a padon. Értem."},
						}),
						inkwell.GotoNode("deflected"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Vasquez is expecting me."},
						{Lang: "hu", Text: "Vasquez vár engem."},
					},
					Show: inkwell.Given{
						Flag: inc.FlagReceptionPassed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "Through the door, sir."},
							{Lang: "hu", Text: "Az ajtón át, uram."},
						}),
						inkwell.EndDialog(),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Never mind."},
						{Lang: "hu", Text: "Hagyjuk."},
					},
					Actions: []inkwell.Action{inkwell.EndDialog()},
				},
			},
		},
		{
			Name: "deflected",
			Lines: []inkwell.DialogLine{{
				Speaker: inc.Receptionist,
				Text: inkwell.Text{
					{Lang: "en", Text: "Anything else."},
					{Lang: "hu", Text: "Még valami."},
				},
			}},
			Choices: []inkwell.DialogChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Let me try that again."},
						{Lang: "hu", Text: "Hadd próbáljam újra."},
					},
					Actions: []inkwell.Action{inkwell.GotoNode("root")},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Tell Vasquez that Paul is here. She'll want to know who else was on Noodle's list."},
						{Lang: "hu", Text: "Mondja meg Vasqueznek, hogy Paul van itt. Tudni akarja majd, ki más volt Noodle listáján."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagReceptionPassed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "…"},
							{Lang: "hu", Text: "…"},
						}),
						inkwell.Say(inc.Paul, inkwell.Text{
							{Lang: "en", Text: "She picks up the phone. Says two words. Listens for a lot longer."},
							{Lang: "hu", Text: "Felveszi a telefont. Két szót mond. Sokkal tovább hallgat."},
						}),
						inkwell.Say(inc.Receptionist, inkwell.Text{
							{Lang: "en", Text: "She says go through."},
							{Lang: "hu", Text: "Azt mondja, menjen be."},
						}),
						inkwell.SetFlag(inc.FlagReceptionPassed),
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "There was no list, Paul. There is one now."},
							{Lang: "hu", Text: "Nem volt lista, Paul. Most már van."},
						}),
						inkwell.EndDialog(),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "No. Thank you."},
						{Lang: "hu", Text: "Nem. Köszönöm."},
					},
					Actions: []inkwell.Action{inkwell.EndDialog()},
				},
			},
		},
	},
})
