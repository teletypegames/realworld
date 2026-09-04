package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var IceCreamVendor = DialogManager.Register(&Dialog{
	Name:  "ice_cream_vendor",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: character.IceCreamVendor,
			Text: inkwell.Text{
				{Lang: "en", Text: "Welcome to Scoops! What can I get you?"},
				{Lang: "hu", Text: "Üdv a Scoopsban! Mit adhatok?"},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "This used to be a bar. Manny's."},
					{Lang: "hu", Text: "Ez régen kocsma volt. Mannyé."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.IceCreamVendor, inkwell.Text{
						{Lang: "en", Text: "No, sir. This has always been Scoops. Family business, forty years."},
						{Lang: "hu", Text: "Nem, uram. Ez mindig is a Scoops volt. Családi vállalkozás, negyven év."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "I drank here. Six years ago. That wall had a dartboard."},
						{Lang: "hu", Text: "Itt ittam. Hat éve. Azon a falon céltábla volt."},
					}),
					inkwell.Say(character.IceCreamVendor, inkwell.Text{
						{Lang: "en", Text: "That wall has always had the flavour board, sir. Would you like to try the pistachio?"},
						{Lang: "hu", Text: "Azon a falon mindig is az íztábla volt, uram. Megkóstolná a pisztáciát?"},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "He isn't lying, Paul. That's the part that should bother you."},
						{Lang: "hu", Text: "Nem hazudik, Paul. Ez az, aminek zavarnia kéne."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "One ice cream, please."},
					{Lang: "hu", Text: "Egy fagyit kérek."},
				},
				Show: inkwell.Given{
					NotItem: item.IceCream,
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.IceCreamVendor, inkwell.Text{
						{Lang: "en", Text: "Pistachio it is! Extra sprinkles, no charge."},
						{Lang: "hu", Text: "Akkor pisztácia! Extra szórat, ingyen."},
					}),
					inkwell.Give(item.IceCream),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Thirty-four years old and buying a cone with sprinkles. I'd look away if I had eyes."},
						{Lang: "hu", Text: "Harmincnégy éves, és szóratos tölcsért vesz. Elfordulnék, ha volna szemem."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Who's the man at the terminal?"},
					{Lang: "hu", Text: "Ki az a fickó a terminálnál?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.IceCreamVendor, inkwell.Text{
						{Lang: "en", Text: "A regular. Comes in at ten, leaves at six. Buys one sundae. Plays his little game."},
						{Lang: "hu", Text: "Törzsvendég. Tízkor jön, hatkor megy. Vesz egy fagyikelyhet. Játszik a kis játékával."},
					}),
					inkwell.Say(character.IceCreamVendor, inkwell.Text{
						{Lang: "en", Text: "He likes his sweets. If you want him to move, that's my only advice."},
						{Lang: "hu", Text: "Szereti az édességet. Ha azt akarja, hogy elmozduljon, ez az egyetlen tanácsom."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Nothing. Thanks."},
					{Lang: "hu", Text: "Semmit. Köszönöm."},
				},
				Actions: []inkwell.Action{inkwell.EndDialog()},
			},
		},
	}},
})
