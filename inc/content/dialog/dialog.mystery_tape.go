package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

var MysteryTape = DialogManager.Register(&Dialog{
	Name:  "mystery_tape",
	Start: "root",
	Nodes: []DialogNode{{
		Name: "root",
		Choices: []DialogChoice{
			{
				Text: Text{
					{Lang: "en", Text: "Are you sure it works?"},
					{Lang: "hu", Text: "Biztos, hogy működik?"},
				},
				Show: Given{
					NotFlag: inc.FlagTapeSpeaks,
				},
				Actions: []Action{
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "It works. It just isn't talking to you. That is not the same as silence."},
						{Lang: "hu", Text: "Működik. Csak nem beszél hozzád. Ez nem ugyanaz, mint a csend."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Say the riddle again."},
					{Lang: "hu", Text: "Mondd újra a rejtvényt."},
				},
				Show: Given{
					NotFlag: inc.FlagTapeSpeaks,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "The key forged in the workshop of the underwater sun opens the secrets."},
						{Lang: "hu", Text: "A vízalatti nap műhelyében kovácsoltatott kulcs nyitja a titkokat."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Word for word. It isn't thinking. It's playing."},
						{Lang: "hu", Text: "Szóról szóra. Nem gondolkodik. Lejátszik."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Who are you, really?"},
					{Lang: "hu", Text: "Ki vagy te valójában?"},
				},
				Show: Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "I belonged to one of the developers. A scientist. The project broke him. The BVK let it."},
						{Lang: "hu", Text: "Az egyik fejlesztőé voltam. Egy tudósé. A projekt összetörte. A BVK hagyta."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "I am what is left of what he knew. And I know Norman can still be saved. He isn't broken yet. Only asleep."},
						{Lang: "hu", Text: "Én vagyok az, ami abból maradt, amit tudott. És tudom, hogy Norman még megmenthető. Ő még nem tört össze. Csak alszik."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "A dead man's tape, speaking for a living one. I'd know something about the first half of that."},
						{Lang: "hu", Text: "Egy halott ember kazettája, egy élő nevében beszél. Az első feléhez volna hozzászólásom."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "What is the simulation?"},
					{Lang: "hu", Text: "Mi az a szimuláció?"},
				},
				Show: Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "A bridge between an ACP and a brain. The brain does the computing. The ACP does the dreaming. The person does neither. The person is inside."},
						{Lang: "hu", Text: "Híd egy ACP és egy agy között. Az agy számol. Az ACP álmodik. Az ember egyiket sem. Az ember odabent van."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Nothing runs on a machine. It all runs in the head of whoever is strapped in. That is what makes it impossible to switch off from outside."},
						{Lang: "hu", Text: "Semmi nem fut gépen. Minden a bekötött ember fejében fut. Ezért nem lehet kívülről kikapcsolni."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Then how do you get someone out?"},
						{Lang: "hu", Text: "Akkor hogyan hozol ki valakit?"},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "You give the machine something it wants more than him."},
						{Lang: "hu", Text: "Adsz a gépnek valamit, amit jobban akar, mint őt."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Where is Norman?"},
					{Lang: "hu", Text: "Hol van Norman?"},
				},
				Show: Given{
					Flag:    inc.FlagTapeSpeaks,
					NotFlag: inc.FlagLabLocated,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "I don't know. I was taken from him before they moved him. Start where he lived. Then find the people who hid from the BVK. They keep lists."},
						{Lang: "hu", Text: "Nem tudom. Elvettek tőle, mielőtt elvitték. Kezdd ott, ahol lakott. Aztán találd meg azokat, akik elbújtak a BVK elől. Ők listákat vezetnek."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "It doesn't know where he is, but it knows where he lived. Keep count of what it knows, Paul."},
						{Lang: "hu", Text: "Nem tudja, hol van, de tudja, hol lakott. Számold, mit tud, Paul."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "How did Noodle end up with you?"},
					{Lang: "hu", Text: "Hogy kerültél Noodle-hoz?"},
				},
				Show: Given{
					Flag: inc.FlagTapeSpeaks,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Noodle buys what nobody should be selling. Somebody sold me. He knew what I was worth, and he knew who to send for."},
						{Lang: "hu", Text: "Noodle azt veszi, amit senkinek nem kéne eladnia. Valaki eladott engem. Tudta, mit érek, és tudta, kiért küldjön."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "He sent for you. That should tell you what he thought of you."},
						{Lang: "hu", Text: "Érted küldött. Ez elmondja, mit gondolt rólad."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Or what he thought he could get you to do."},
						{Lang: "hu", Text: "Vagy hogy mit gondolt, mire lehet rávenni téged."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "Norman's Armilla was empty. Where's his tape?"},
					{Lang: "hu", Text: "Norman Armillája üres volt. Hol a kazettája?"},
				},
				Show: Given{
					Flag: inc.FlagArmillaSeen,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Confiscated. The BVK took it the night they took him. A man's tape is evidence, and evidence disappears."},
						{Lang: "hu", Text: "Lefoglalták. A BVK vitte el, azon az éjjelen, amikor őt is. Egy ember kazettája bizonyíték, és a bizonyíték eltűnik."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Stolen, if you want the honest word. Whatever Norman kept on it, they have it now."},
						{Lang: "hu", Text: "Ellopták, ha az őszinte szót akarod. Amit Norman rajta tartott, az most náluk van."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "Convenient."},
						{Lang: "hu", Text: "Kényelmes."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "The man at the hot dog stand. Sumphor. Do you know him?"},
					{Lang: "hu", Text: "A hotdogos fickó. Sumphor. Ismered?"},
				},
				Show: Given{
					Flag: inc.FlagSumphorTalked,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Know him. I was his. He is the scientist I told you about. The one the project broke."},
						{Lang: "hu", Text: "Ismerem. Az övé voltam. Ő az a tudós, akiről beszéltem. Akit a projekt összetört."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "He says he's never seen you."},
						{Lang: "hu", Text: "Azt mondja, sosem látott."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "He sat in the chair, Paul. He says a lot of things. Which of us sounds like the one who lost his mind?"},
						{Lang: "hu", Text: "Ő ült a székben, Paul. Sok mindent mond. Melyikünk hangzik úgy, mint aki elvesztette az eszét?"},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "The one who's calm is not always the one who's right. I'd remember that."},
						{Lang: "hu", Text: "Aki nyugodt, nem mindig az, akinek igaza van. Ezt megjegyezném."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "The logs say the Golden Gate. What then?"},
					{Lang: "hu", Text: "A naplók szerint a Golden Gate. És aztán?"},
				},
				Show: Given{
					Flag: inc.FlagLabLocated,
				},
				Actions: []Action{
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Then you go. Into the bridge, up to where he sleeps, and to the machine that runs it."},
						{Lang: "hu", Text: "Aztán mész. Be a hídba, fel oda, ahol alszik, és a géphez, amely futtatja."},
					}),
					speaker.TapeMystery.Say(Text{
						{Lang: "en", Text: "Put me in it. I know the simulation from the inside. I can reach him where nobody else can. Please, Paul. It has to be me."},
						{Lang: "hu", Text: "Tegyél be engem. Belülről ismerem a szimulációt. Elérem ott, ahol senki más nem. Kérlek, Paul. Nekem kell lennem."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "“Please” is new. Tapes don't say please. I never have."},
						{Lang: "hu", Text: "A „kérlek” új. A kazetták nem mondanak kérleket. Én soha nem mondtam."},
					}),
					inkwell.GotoNode("root"),
				},
			},
			{
				Text: Text{
					{Lang: "en", Text: "That's all for now."},
					{Lang: "hu", Text: "Most ennyi."},
				},
				Actions: []Action{inkwell.EndDialog()},
			},
		},
	}},
})
