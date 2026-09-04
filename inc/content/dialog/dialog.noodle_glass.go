package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NoodleGlass = DialogManager.Register(&Dialog{
	Name:  "noodle_glass",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: character.Noodle,
			Text: inkwell.Text{
				{Lang: "en", Text: "Paul. You came. You actually came."},
				{Lang: "hu", Text: "Paul. Eljöttél. Tényleg eljöttél."},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Noodle. What did you get yourself into?"},
					{Lang: "hu", Text: "Noodle. Mibe keveredtél?"},
				},
				Show: inkwell.Given{
					NotFlag: inc.FlagVasquezDeal,
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Vasquez, inkwell.Text{
						{Lang: "en", Text: "Not yet."},
						{Lang: "hu", Text: "Még nem."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "She hasn't finished with me."},
						{Lang: "hu", Text: "Még nem végzett velem."},
					}),
					inkwell.EndDialog(),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Noodle. What did you get yourself into?"},
					{Lang: "hu", Text: "Noodle. Mibe keveredtél?"},
				},
				Show: inkwell.Given{
					Flag: inc.FlagVasquezDeal,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.SetFlag(inc.FlagNoodleSeen),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Nothing I couldn't talk my way out of. Usually. They're going to empty my place, Paul. Every shelf."},
						{Lang: "hu", Text: "Semmi olyat, amiből ne beszélném ki magam. Általában. Ki fogják pakolni a helyemet, Paul. Minden polcot."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "What I left you isn't what they think it is. It isn't what I thought it was either."},
						{Lang: "hu", Text: "Amit neked hagytam, az nem az, aminek hiszik. Nem is az, aminek én hittem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Who is Norman?"},
					{Lang: "hu", Text: "Ki az a Norman?"},
				},
				Show: inkwell.Given{
					Flag: inc.FlagVasquezDeal,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.SetFlag(inc.FlagNoodleSeen),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Norman… Norman handed me something to look at. I looked."},
						{Lang: "hu", Text: "Norman… Norman a kezembe adott valamit, hogy nézzem meg. Megnéztem."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "I should not have."},
						{Lang: "hu", Text: "Nem kellett volna."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "He hesitated. Noodle never hesitates. He lies faster than that."},
						{Lang: "hu", Text: "Habozott. Noodle soha nem habozik. Gyorsabban hazudik ennél."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "How do I get you out of here?"},
					{Lang: "hu", Text: "Hogy hozlak ki innen?"},
				},
				Show: inkwell.Given{
					Flag: inc.FlagVasquezDeal,
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.SetFlag(inc.FlagNoodleSeen),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Give her something bigger than me. She wants bigger. You'll find bigger, Paul, you've got the thing."},
						{Lang: "hu", Text: "Adj neki valamit, ami nagyobb nálam. Nagyobbat akar. Találsz nagyobbat, Paul, nálad van az a dolog."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Get me out, and I tell you everything. All of it. There's a room you haven't seen."},
						{Lang: "hu", Text: "Hozz ki, és mindent elmondok. Az egészet. Van egy szoba, amit még nem láttál."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "What room?"},
						{Lang: "hu", Text: "Milyen szoba?"},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Not here. Not through glass."},
						{Lang: "hu", Text: "Nem itt. Nem üvegen át."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Hang in there."},
					{Lang: "hu", Text: "Tarts ki."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "They have chips. I'll live."},
						{Lang: "hu", Text: "Van chipsük. Túlélem."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
