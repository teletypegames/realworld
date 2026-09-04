package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var SecretClub = SceneManager.Register(&Scene{
	Name: "secret_club",
	Title: inkwell.Text{
		{Lang: "en", Text: "SECRET CLUB — THE UNDERWATER SUN"},
		{Lang: "hu", Text: "TITKOS KLUB — A VÍZ ALATTI NAP"},
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 320,
				Y: 250,
			},
		},
		{
			Character: character.Secretary,
			At: inkwell.Point{
				X: 140,
				Y: 214,
			},
		},
	},
	Hotspots: []inkwell.Hotspot{
		{
			Name: "secretary",
			Label: inkwell.Text{
				{Lang: "en", Text: "secretary"},
				{Lang: "hu", Text: "titkár"},
			},
			Area: inkwell.Rect(116, 134, 48, 84),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "A clipboard, a cardigan, and nobody to take minutes for. He's the whole club tonight."},
				{Lang: "hu", Text: "Egy csíptetős tábla, egy kardigán, és senki, akiről jegyzőkönyvet vezessen. Ma este ő az egész klub."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Secretary)),
			OnUse: inkwell.Say(character.Secretary, inkwell.Text{
				{Lang: "en", Text: "I'm not a fixture, thank you. Ask."},
				{Lang: "hu", Text: "Nem vagyok berendezési tárgy, köszönöm. Kérdezzen."},
			}),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Leave the man his cardigan."},
				{Lang: "hu", Text: "Hagyd meg neki a kardigánját."},
			}),
		},
		{
			Name: "chairs",
			Label: inkwell.Text{
				{Lang: "en", Text: "empty chairs"},
				{Lang: "hu", Text: "üres székek"},
			},
			Area: inkwell.Rect(230, 150, 200, 60),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Forty folding chairs in rows. Forty people sat here two weeks ago and listened to something."},
					{Lang: "hu", Text: "Negyven összecsukható szék sorokban. Két héttel ezelőtt negyven ember ült itt és hallgatott valamit."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "And none of them came back. A talk that empties a room for good is either very boring or very true."},
					{Lang: "hu", Text: "És egyikük sem jött vissza. Egy előadás, ami végleg kiüríti a termet, vagy nagyon unalmas, vagy nagyon igaz."},
				}),
			),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'll stand. Sitting in an empty audience is a kind of surrender."},
				{Lang: "hu", Text: "Állok. Üres nézőtéren ülni egyfajta megadás."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I've got enough chairs. I've got a shop full of chairs."},
				{Lang: "hu", Text: "Van elég székem. Egy egész boltom van tele székkel."},
			}),
		},
		{
			Name: "poster",
			Label: inkwell.Text{
				{Lang: "en", Text: "poster"},
				{Lang: "hu", Text: "plakát"},
			},
			Area: inkwell.Rect(260, 40, 120, 90),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "“THE UNDERWATER SUN — an evening on what they are building inside your head.” Hand-lettered. The sun again."},
					{Lang: "hu", Text: "„A VÍZALATTI NAP — egy este arról, mit építenek a fejedben.” Kézzel írva. Megint a nap."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "A troupe. Not a cult. The terminal called them a cult because the terminal is the BVK's."},
					{Lang: "hu", Text: "Egy társulat. Nem kultusz. A terminál kultusznak hívta, mert a terminál a BVK-é."},
				}),
			),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Taped to the wall. I'll leave it for the next forty."},
				{Lang: "hu", Text: "A falra ragasztva. Meghagyom a következő negyvennek."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It would tear. And it's the only sign anyone was ever here."},
				{Lang: "hu", Text: "Elszakadna. És ez az egyetlen jel, hogy valaki valaha itt volt."},
			}),
		},
		{
			Name: "echo_terminal",
			Label: inkwell.Text{
				{Lang: "en", Text: "ECHO terminal"},
				{Lang: "hu", Text: "ECHO-terminál"},
			},
			Area: inkwell.Rect(500, 80, 110, 120),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "A LECTOR screen, the same as every terminal in the city. Except the cursor blinks wrong."},
					{Lang: "hu", Text: "Egy LECTOR-képernyő, mint a város minden terminálján. Csak a kurzor rosszul villog."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Something else is running underneath the interface. That's ECHO. The BVK doesn't read what goes through it."},
					{Lang: "hu", Text: "Valami más fut a felület alatt. Az az ECHO. Amit ezen át olvasnak, azt a BVK nem olvassa."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNormanNamed),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "I've read enough of it for one night. The name's still on the screen. Norman Reds."},
					{Lang: "hu", Text: "Egy estére eleget olvastam belőle. A név még mindig a képernyőn van. Norman Reds."},
				}),
				inkwell.Seq(
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Boards I've never seen. The world before the BVK. The Resistance, with dates and names."},
						{Lang: "hu", Text: "Fórumok, amiket még sosem láttam. A világ a BVK előtt. Az Ellenállás, dátumokkal, nevekkel."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "And a programme. The BVK calls it “human simulation”. Therapy, training, relaxation, faster learning."},
						{Lang: "hu", Text: "És egy program. A BVK „humán szimulációnak” hívja. Terápia, tréning, relaxáció, gyorsabb tanulás."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "“Ultimately, the preservation of important minds on tape.” They say that like it's a kindness."},
						{Lang: "hu", Text: "„Végső soron fontos elmék megőrzése kazettán.” Úgy mondják, mintha jótétemény volna."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "The developers are listed. Read the names, Paul."},
						{Lang: "hu", Text: "Fel vannak sorolva a fejlesztők. Olvasd a neveket, Paul."},
					}),
					inkwell.Say(character.Paul, inkwell.Text{
						{Lang: "en", Text: "Six names. Five with a status. One with “whereabouts unknown”. Norman Reds."},
						{Lang: "hu", Text: "Hat név. Öt státusszal. Egy „ismeretlen tartózkodási hellyel”. Norman Reds."},
					}),
					inkwell.Say(speaker.Dex, inkwell.Text{
						{Lang: "en", Text: "Norman. Your Norman. The one who never called."},
						{Lang: "hu", Text: "Norman. A te Normanod. Aki soha nem hívott."},
					}),
					inkwell.SetFlag(inc.FlagNormanNamed),
				),
			),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "The only terminal in the city the BVK isn't reading, and you want to carry it out through a shop."},
				{Lang: "hu", Text: "A város egyetlen terminálja, amit a BVK nem olvas, és te ki akarod vinni egy bolton át."},
			}),
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "tape_speaks",
			Once: true,
			When: inkwell.Given{
				Flag:    inc.FlagNormanNamed,
				NotFlag: inc.FlagTapeSpeaks,
				Var:     inc.VarSlot2,
				Is:      item.MysteryTape,
			},
			Do: inkwell.RunScript(script.TapeSpeaks),
		},
		{
			Name: "tape_not_loaded",
			Once: true,
			When: inkwell.And(
				inkwell.Flag(inc.FlagNormanNamed),
				inkwell.Not(inkwell.Flag(inc.FlagTapeSpeaks)),
				inkwell.Not(inkwell.VarEq(inc.VarSlot2, item.MysteryTape)),
			),
			Do: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "The riddle said the key is forged here. If that tape's going to hear anything, it hears it from the slot."},
				{Lang: "hu", Text: "A rejtvény szerint a kulcsot itt kovácsolták. Ha az a kazetta hall valamit, a foglalatból hallja."},
			}),
		},
	},
})

var _ = SecretClub.Leads(
	inkwell.Exit{
		To: BBSTerminal,
		Label: inkwell.Text{
			{Lang: "en", Text: "the terminal in the corner"},
			{Lang: "hu", Text: "a terminál a sarokban"},
		},
		Side: inkwell.ExitBack,
	},
	inkwell.Exit{
		To: TrinketShop,
		Label: inkwell.Text{
			{Lang: "en", Text: "back out through the shop"},
			{Lang: "hu", Text: "vissza a bolton át"},
		},
		Side: inkwell.ExitNear,
	},
)
