package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var CandyCrusher = DialogManager.Register(&Dialog{
	Name:  "candy_crusher",
	Start: "root",
	Nodes: []inkwell.DialogNode{
		{
			Name: "root",
			Lines: []inkwell.DialogLine{{
				Speaker: character.CandyCrusher,
				Text: inkwell.Text{
					{Lang: "en", Text: "Hang on. Hang on. Combo."},
					{Lang: "hu", Text: "Várj. Várj. Kombó."},
				},
			}},
			Choices: []inkwell.DialogChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "I need the terminal for two minutes."},
						{Lang: "hu", Text: "Kéne a terminál két percre."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagTerminalAllowed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Everybody needs the terminal. I'm on level ninety-one. Nobody's on level ninety-one."},
							{Lang: "hu", Text: "Mindenkinek kell a terminál. Kilencvenegyedik szinten vagyok. Senki nincs a kilencvenegyedik szinten."},
						}),
						inkwell.Say(speaker.Dex, inkwell.Text{
							{Lang: "en", Text: "Try caring about something he cares about. It's cheaper than ice cream."},
							{Lang: "hu", Text: "Próbálj érdeklődni valami iránt, ami őt érdekli. Olcsóbb, mint a fagyi."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "What are you playing?"},
						{Lang: "hu", Text: "Mit játszol?"},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagTerminalAllowed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Sugar Cascade. You match the candies. Three gets you a pop, five gets you a bomb."},
							{Lang: "hu", Text: "Sugar Cascade. Párosítod a cukorkákat. Háromért pukk, ötért bomba."},
						}),
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Runs on the terminal's spare cycles. BVK lets you, if it's on the approved list. This one's approved."},
							{Lang: "hu", Text: "A terminál szabad ciklusain fut. A BVK engedi, ha rajta van a jóváhagyott listán. Ez rajta van."},
						}),
						inkwell.GotoNode("game"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "What's the BVK got on that screen anyway?"},
						{Lang: "hu", Text: "Mi van egyáltalán a BVK-nak azon a képernyőn?"},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagTerminalAllowed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Everything! Licences, news, the registry. And the announcements. Did you see the Neural Interface one?"},
							{Lang: "hu", Text: "Mindent! Licenszek, hírek, a nyilvántartás. És a bejelentések. Láttad a Neural Interface-eset?"},
						}),
						inkwell.GotoNode("interface"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Don't let me keep you."},
						{Lang: "hu", Text: "Nem tartalak fel."},
					},
					Actions: []inkwell.Action{inkwell.EndDialog()},
				},
			},
		},
		{
			Name: "game",
			Choices: []inkwell.DialogChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Ninety-one levels. That's dedication."},
						{Lang: "hu", Text: "Kilencvenegy szint. Ez elhivatottság."},
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "It is! It is. Nobody gets that. My wife says it's a waste. It's a discipline."},
							{Lang: "hu", Text: "Az! Az. Ezt senki nem érti. A feleségem szerint pazarlás. Ez fegyelem."},
						}),
						inkwell.Say(speaker.Dex, inkwell.Text{
							{Lang: "en", Text: "Keep going. He's softening. It's horrible to watch."},
							{Lang: "hu", Text: "Folytasd. Puhul. Borzalmas nézni."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Matching candy. On a government terminal."},
						{Lang: "hu", Text: "Cukorkapárosítás. Egy állami terminálon."},
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Approved list. I said. Are you going to stand there judging, or what?"},
							{Lang: "hu", Text: "Jóváhagyott lista. Mondtam. Itt fogsz állni és ítélkezni, vagy mi?"},
						}),
						inkwell.GotoNode("root"),
					},
				},
			},
		},
		{
			Name: "interface",
			Lines: []inkwell.DialogLine{
				{
					Speaker: character.CandyCrusher,
					Text: inkwell.Text{
						{Lang: "en", Text: "The Neural Interface. BVK's own. They're putting the human brain on top of the ACPs, see. Above them."},
						{Lang: "hu", Text: "A Neural Interface. A BVK sajátja. Az emberi agyat teszik az ACP-k fölé, érted. Föléjük."},
					},
				},
				{
					Speaker: character.CandyCrusher,
					Text: inkwell.Text{
						{Lang: "en", Text: "So no machine ever gets clever enough to run the show. Humanity stays in charge. It's in the brochure."},
						{Lang: "hu", Text: "Így egy gép sem lesz sosem elég okos ahhoz, hogy átvegye az irányítást. Az emberiség marad a főnök. Benne van a brosúrában."},
					},
				},
				{
					Speaker: speaker.Dex,
					Text: inkwell.Text{
						{Lang: "en", Text: "Every cage in history came with a brochure."},
						{Lang: "hu", Text: "A történelem minden ketrecéhez járt brosúra."},
					},
				},
			},
			Choices: []inkwell.DialogChoice{
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Brain above the machine. Finally somebody gets it right."},
						{Lang: "hu", Text: "Agy a gép fölött. Végre valaki jól csinálja."},
					},
					Show: inkwell.Given{
						NotFlag: inc.FlagTerminalAllowed,
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "Yes! Exactly! You're the first person in here who isn't a cynic. Fine, two minutes. Don't close my game."},
							{Lang: "hu", Text: "Igen! Pontosan! Te vagy itt az első, aki nem cinikus. Jó, két perc. Ne zárd be a játékomat."},
						}),
						inkwell.SetFlag(inc.FlagTerminalAllowed),
						inkwell.Say(speaker.Dex, inkwell.Text{
							{Lang: "en", Text: "You just agreed with a brochure to get at a keyboard. I'll allow it. Once."},
							{Lang: "hu", Text: "Egyetértettél egy brosúrával, hogy hozzáférj egy billentyűzethez. Ezt elnézem. Egyszer."},
						}),
						inkwell.EndDialog(),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Who says the brain ends up on top?"},
						{Lang: "hu", Text: "Ki mondja, hogy az agy kerül felülre?"},
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "The BVK says. It's their job to say. You're one of those, aren't you."},
							{Lang: "hu", Text: "A BVK mondja. Az a dolguk, hogy mondják. Te is olyan vagy, ugye."},
						}),
						inkwell.GotoNode("root"),
					},
				},
				{
					Text: inkwell.Text{
						{Lang: "en", Text: "Sounds expensive."},
						{Lang: "hu", Text: "Drágán hangzik."},
					},
					Actions: []inkwell.Action{
						inkwell.Say(character.CandyCrusher, inkwell.Text{
							{Lang: "en", Text: "It's a prototype. They've got engineers on it. Real ones, not the licence-mill kind."},
							{Lang: "hu", Text: "Prototípus. Mérnökök dolgoznak rajta. Igaziak, nem a licenszgyári fajta."},
						}),
						inkwell.GotoNode("root"),
					},
				},
			},
		},
	},
})
