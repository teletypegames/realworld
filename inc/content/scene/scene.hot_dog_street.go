package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

type hotDogStreet struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneHotDogStreet,
	Title: inkwell.Text{
		{Lang: "en", Text: "STREET FRONT — THE HOT DOG STAND"},
		{Lang: "hu", Text: "UTCAFRONT — A HOTDOGOS"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
	Actors: []inkwell.SceneActor{
		{
			CharacterName: inc.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 236,
			},
		},
		{
			CharacterName: inc.HotDogVendor,
			At: inkwell.Point{
				X: 420,
				Y: 214,
			},
		},
		{
			CharacterName: inc.Sumphor,
			At: inkwell.Point{
				X: 560,
				Y: 246,
			},
		},
	},
	Hotspots: hotDogStreet{}.hotspots(),
})

func (place hotDogStreet) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.vendor(),
		place.sumphor(),
		place.stand(),
		place.sign(),
	}
}

func (hotDogStreet) vendor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "vendor",
		Label: inkwell.Text{
			{Lang: "en", Text: "hot dog vendor"},
			{Lang: "hu", Text: "hotdogárus"},
		},
		Area: inkwell.Rect(396, 134, 48, 84),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Apron, tongs, the patience of a man who has heard every order there is."},
			{Lang: "hu", Text: "Kötény, fogó, és egy ember türelme, aki már minden létező rendelést hallott."},
		}),
		OnTalk: system.System.Paused(inkwell.RunDialog(inc.DialogHotDogVendor)),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "You talk to a vendor. You don't use him."},
			{Lang: "hu", Text: "Egy árussal beszélsz. Nem használod."},
		}),
		OnTake: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "He's not stock, Paul."},
			{Lang: "hu", Text: "Ő nem árukészlet, Paul."},
		}),
	}
}

func (hotDogStreet) sumphor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "sumphor",
		Label: inkwell.Text{
			{Lang: "en", Text: "man on the kerb"},
			{Lang: "hu", Text: "férfi a járdaszélen"},
		},
		Area: inkwell.Rect(536, 166, 48, 84),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Good shoes, ruined coat. He's arguing with a point about a metre to his left."},
			{Lang: "hu", Text: "Jó cipő, tönkrement kabát. Egy ponttal vitatkozik, ami úgy egy méterre van tőle balra."},
		}),
		OnTalk: system.System.Paused(inkwell.RunDialog(inc.DialogSumphor)),
		OnUse: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Talk to him. Don't poke him."},
			{Lang: "hu", Text: "Beszélj vele. Ne bökdösd."},
		}),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I have enough strays."},
			{Lang: "hu", Text: "Van elég kóborom."},
		}),
	}
}

func (hotDogStreet) stand() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "stand",
		Label: inkwell.Text{
			{Lang: "en", Text: "hot dog stand"},
			{Lang: "hu", Text: "hotdogos bódé"},
		},
		Area: inkwell.Rect(360, 150, 130, 70),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Steam, mustard, a radio tuned between two stations. Same as the one on my street. Same as every street."},
			{Lang: "hu", Text: "Gőz, mustár, egy rádió két állomás közé hangolva. Ugyanaz, mint az én utcámon. Mint minden utcán."},
		}),
		OnUse: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Order from the man, not the counter."},
			{Lang: "hu", Text: "Az embertől rendelj, ne a pulttól."},
		}),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'd need a truck and a worse conscience."},
			{Lang: "hu", Text: "Kéne hozzá egy teherautó és egy rosszabb lelkiismeret."},
		}),
	}
}

func (hotDogStreet) sign() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "sign",
		Label: inkwell.Text{
			{Lang: "en", Text: "stamp-it sign"},
			{Lang: "hu", Text: "stamp-it tábla"},
		},
		Area: inkwell.Rect(370, 60, 110, 60),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "“STAMP-IT. Ten stamps, one free.” The same card that was in Norman's wallet."},
			{Lang: "hu", Text: "„STAMP-IT. Tíz pecsét, egy ingyen.” Ugyanaz a kártya, ami Norman tárcájában volt."},
		}),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It's bolted on. The mustard isn't, but I'm not taking that either."},
			{Lang: "hu", Text: "Odacsavarozták. A mustár nincs, de azt sem viszem el."},
		}),
	}
}
