package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = DialogManager.Register(Dialog{
	Name:  inc.DialogOfficerOnGuard,
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: inc.OfficerOnGuard,
			Text: inkwell.Text{
				{Lang: "en", Text: "Sir. This is a sealed address."},
				{Lang: "hu", Text: "Uram. Ez lezárt helyszín."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What happened here?"},
					{Lang: "hu", Text: "Mi történt itt?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "Police matter."},
						{Lang: "hu", Text: "Rendőrségi ügy."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Where's the man who lives here?"},
					{Lang: "hu", Text: "Hol van, aki itt lakik?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "One moment."},
						{Lang: "hu", Text: "Egy pillanat."},
					}),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "He's getting a notebook out."},
						{Lang: "hu", Text: "Füzetet vesz elő."},
					}),
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "Would you like to give your details and state your connection to this address, sir?"},
						{Lang: "hu", Text: "Kívánja megadni az adatait, és közölni, mi köze ehhez a címhez, uram?"},
					}),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "…No. No, I would not."},
						{Lang: "hu", Text: "…Nem. Nem kívánom."},
					}),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "Wise. Nothing in that notebook ever got anybody out of anything."},
						{Lang: "hu", Text: "Bölcs. Abból a füzetből még senki nem került ki semmiből."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Long shift?"},
					{Lang: "hu", Text: "Hosszú a műszak?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "Since six. Nobody relieves you on a seal. Nobody brings you anything either."},
						{Lang: "hu", Text: "Hat óta. Pecsételésnél nem vált le senki. Hozni se hoz senki semmit."},
					}),
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "And don't eat in front of me, please. I'm standing here all day and I'm starving."},
						{Lang: "hu", Text: "És kérem, ne egyen itt előttem. Egész nap szobrozok, és éhen halok."},
					}),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "Paul. You heard the man."},
						{Lang: "hu", Text: "Paul. Hallottad az embert."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'll go."},
					{Lang: "hu", Text: "Megyek."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "Do that, sir."},
						{Lang: "hu", Text: "Tegye azt, uram."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
