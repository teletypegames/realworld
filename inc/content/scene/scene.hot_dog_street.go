package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var HotDogStreet = SceneManager.Register(&Scene{
	Name: "hot_dog_street",
	Title: Text{
		{Lang: "en", Text: "STREET FRONT — THE HOT DOG STAND"},
		{Lang: "hu", Text: "UTCAFRONT — A HOTDOGOS"},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 120,
				Y: 236,
			},
		},
		{
			Character: character.HotDogVendor,
			At: Point{
				X: 420,
				Y: 214,
			},
		},
		{
			Character: character.Sumphor,
			At: Point{
				X: 560,
				Y: 246,
			},
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "vendor",
			Label: Text{
				{Lang: "en", Text: "hot dog vendor"},
				{Lang: "hu", Text: "hotdogárus"},
			},
			Area: inkwell.Rect(396, 134, 48, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Apron, tongs, the patience of a man who has heard every order there is."},
				{Lang: "hu", Text: "Kötény, fogó, és egy ember türelme, aki már minden létező rendelést hallott."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.HotDogVendor)),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "You talk to a vendor. You don't use him."},
				{Lang: "hu", Text: "Egy árussal beszélsz. Nem használod."},
			}),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "He's not stock, Paul."},
				{Lang: "hu", Text: "Ő nem árukészlet, Paul."},
			}),
		},
		{
			Name: "sumphor",
			Label: Text{
				{Lang: "en", Text: "man on the kerb"},
				{Lang: "hu", Text: "férfi a járdaszélen"},
			},
			Area: inkwell.Rect(536, 166, 48, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Good shoes, ruined coat. He's arguing with a point about a metre to his left."},
				{Lang: "hu", Text: "Jó cipő, tönkrement kabát. Egy ponttal vitatkozik, ami úgy egy méterre van tőle balra."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Sumphor)),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Talk to him. Don't poke him."},
				{Lang: "hu", Text: "Beszélj vele. Ne bökdösd."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I have enough strays."},
				{Lang: "hu", Text: "Van elég kóborom."},
			}),
		},
		{
			Name: "stand",
			Label: Text{
				{Lang: "en", Text: "hot dog stand"},
				{Lang: "hu", Text: "hotdogos bódé"},
			},
			Area: inkwell.Rect(360, 150, 130, 70),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Steam, mustard, a radio tuned between two stations. Same as the one on my street. Same as every street."},
				{Lang: "hu", Text: "Gőz, mustár, egy rádió két állomás közé hangolva. Ugyanaz, mint az én utcámon. Mint minden utcán."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Order from the man, not the counter."},
				{Lang: "hu", Text: "Az embertől rendelj, ne a pulttól."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I'd need a truck and a worse conscience."},
				{Lang: "hu", Text: "Kéne hozzá egy teherautó és egy rosszabb lelkiismeret."},
			}),
		},
		{
			Name: "sign",
			Label: Text{
				{Lang: "en", Text: "stamp-it sign"},
				{Lang: "hu", Text: "stamp-it tábla"},
			},
			Area: inkwell.Rect(370, 60, 110, 60),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "“STAMP-IT. Ten stamps, one free.” The same card that was in Norman's wallet."},
				{Lang: "hu", Text: "„STAMP-IT. Tíz pecsét, egy ingyen.” Ugyanaz a kártya, ami Norman tárcájában volt."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "It's bolted on. The mustard isn't, but I'm not taking that either."},
				{Lang: "hu", Text: "Odacsavarozták. A mustár nincs, de azt sem viszem el."},
			}),
		},
	},
})

var _ = HotDogStreet.Leads(
	exitToSelector,
)
