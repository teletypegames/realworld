package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var NoodleOffice = DialogManager.Register(&Dialog{
	Name:  "noodle_office",
	Start: "root",
	Nodes: []inkwell.DialogNode{{
		Name: "root",
		Lines: []inkwell.DialogLine{{
			Speaker: character.Noodle,
			Text: inkwell.Text{
				{Lang: "en", Text: "Paul. Sit, don't sit, whatever you like. Chips?"},
				{Lang: "hu", Text: "Paul. Ülj le, ne ülj le, ahogy tetszik. Chipset?"},
			},
		}},
		Choices: []inkwell.DialogChoice{
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "What happened to you?"},
					{Lang: "hu", Text: "Mi történt veled?"},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Twenty cracked Armillas and a warrant. Routine. Except the detective, Vasquez, didn't care about the bracelets. She wanted a tape."},
						{Lang: "hu", Text: "Húsz feltört Armilla és egy parancs. Rutin. Csak a nyomozót, Vasquezt nem érdekelték a karperecek. Egy kazettát akart."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "I told her I didn't have it. Which, by then, was true."},
						{Lang: "hu", Text: "Mondtam neki, hogy nincs nálam. Ami akkor már igaz is volt."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "Tell me about Norman."},
					{Lang: "hu", Text: "Mesélj Normanról."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Friend of mine. Good engineer, terrible at knowing it. The whole interface thing was a napkin of mine he took seriously."},
						{Lang: "hu", Text: "Egy barátom. Jó mérnök, rettenetesen rossz abban, hogy ezt tudja is. Az egész interfész-ügy egy szalvétám volt, amit komolyan vett."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "I told him to go further. Then he came to have beers with me, and he put his own tape on the table and asked me to look at it."},
						{Lang: "hu", Text: "Mondtam neki, hogy menjen tovább. Aztán eljött sörözni, letette a saját kazettáját az asztalra, és megkért, hogy nézzem meg."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "His own Personal Tape."},
						{Lang: "hu", Text: "A saját Személyes Kazettáját."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "His own. You don't do that, Paul. I knew that much."},
						{Lang: "hu", Text: "A sajátját. Ilyet nem csinál az ember, Paul. Ennyit tudtam."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "The back door in the alley. What's the code?"},
					{Lang: "hu", Text: "A hátsó ajtó a sikátorban. Mi a kód?"},
				},
				Once: true,
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Two, seven, seven, two. I filed the keys so you'd get it. Did you get it?"},
						{Lang: "hu", Text: "Kettő, hét, hét, kettő. Lereszeltem a gombokat, hogy kitaláld. Kitaláltad?"},
					}),
					inkwell.SetFlag(inc.FlagNoodleCode),
					inkwell.If(inkwell.Flag(inc.FlagRoomRaided),
						inkwell.Say(character.Paul, inkwell.Text{
							{Lang: "en", Text: "I got the room. After the police did."},
							{Lang: "hu", Text: "A szobát megtaláltam. A rendőrség után."},
						}),
						inkwell.Seq(
							inkwell.Say(character.Paul, inkwell.Text{
								{Lang: "en", Text: "Not yet."},
								{Lang: "hu", Text: "Még nem."},
							}),
							inkwell.Say(speaker.Dex, inkwell.Text{
								{Lang: "en", Text: "Then hurry. They'll be back for that room, and they don't need a code."},
								{Lang: "hu", Text: "Akkor siess. Visszamennek azért a szobáért, és nekik nem kell kód."},
							}),
						),
					),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "The tape talks now."},
					{Lang: "hu", Text: "A kazetta most már beszél."},
				},
				Show: inkwell.Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "It talks? It never said a word to me. Not one. I had it a week."},
						{Lang: "hu", Text: "Beszél? Nekem egy szót se szólt. Egyet se. Egy hétig volt nálam."},
					}),
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "So either it likes you, Paul, or it wanted something from you. Think about which."},
						{Lang: "hu", Text: "Szóval vagy kedvel téged, Paul, vagy akart tőled valamit. Gondold végig, melyik."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: inkwell.Text{
					{Lang: "en", Text: "I'll come back."},
					{Lang: "hu", Text: "Visszajövök."},
				},
				Actions: []inkwell.Action{
					inkwell.Say(character.Noodle, inkwell.Text{
						{Lang: "en", Text: "Do. Bring beer. Nobody's brought me beer in twenty years."},
						{Lang: "hu", Text: "Gyere. Hozz sört. Húsz éve nem hozott nekem sört senki."},
					}),
					inkwell.EndDialog(),
				},
			},
		},
	}},
})
