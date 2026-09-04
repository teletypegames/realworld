package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/speaker"
)

type policeStreet struct{}

var PoliceStreet = SceneManager.Register(&Scene{
	Name: "police_street",
	Title: inkwell.Text{
		{Lang: "en", Text: "SFPD — STREET FRONT"},
		{Lang: "hu", Text: "SFPD — UTCAFRONT"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
		{
			To: policeStation{},
			Label: inkwell.Text{
				{Lang: "en", Text: "into the station"},
				{Lang: "hu", Text: "be az őrsre"},
			},
			Side: inkwell.ExitBack,
		},
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 236,
			},
		},
		{
			Character: character.FatCop,
			At: inkwell.Point{
				X: 420,
				Y: 244,
			},
		},
		{
			Character: character.ThinFigure,
			At: inkwell.Point{
				X: 500,
				Y: 226,
			},
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "chatter",
			Once: true,
			When: inkwell.Not(inkwell.Flag(inc.FlagPoliceTip)),
			Do:   policeStreet{}.chatter(),
		},
	},
	Hotspots: policeStreet{}.hotspots(),
})

func (street policeStreet) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		street.fatCop(),
		street.thinFigure(),
		street.cruiser(),
	}
}

func (policeStreet) chatter() inkwell.Action {
	return inkwell.Seq(
		inkwell.Say(character.ThinFigure, inkwell.Text{
			{Lang: "en", Text: "That nutcase is back at the hot dog stand. Third day running. Talks to the mustard."},
			{Lang: "hu", Text: "Az a bolond megint ott van a hotdogosnál. Harmadik napja. A mustárral beszélget."},
		}),
		inkwell.Say(character.FatCop, inkwell.Text{
			{Lang: "en", Text: "Not our problem. Hot dogs are a food safety matter."},
			{Lang: "hu", Text: "Nem a mi dolgunk. A hotdog élelmiszer-biztonsági ügy."},
		}),
		inkwell.Say(character.ThinFigure, inkwell.Text{
			{Lang: "en", Text: "And they brought the antiques guy in again. The fat one with the monocle."},
			{Lang: "hu", Text: "És megint behozták a régiségest. A kövéret a monoklival."},
		}),
		inkwell.Say(character.FatCop, inkwell.Text{
			{Lang: "en", Text: "He'll have nothing on him. He never has anything on him. That's his whole talent."},
			{Lang: "hu", Text: "Úgyse lesz nála semmi. Soha nincs nála semmi. Ez az egész tehetsége."},
		}),
		inkwell.Say(character.ThinFigure, inkwell.Text{
			{Lang: "en", Text: "You know that eatery in Chinatown, next to the junk shop? Odd types going in and out at all hours."},
			{Lang: "hu", Text: "Ismered azt a kajáldát a kínai negyedben, a kacatbolt mellett? Fura alakok járnak ki-be mindenféle órán."},
		}),
		inkwell.Say(character.FatCop, inkwell.Text{
			{Lang: "en", Text: "Odd types eat too. Pass me the other donut."},
			{Lang: "hu", Text: "A fura alakok is esznek. Add ide a másik fánkot."},
		}),
	)
}

func (policeStreet) noticed() inkwell.Action {
	return inkwell.Say(character.Paul, inkwell.Text{
		{Lang: "en", Text: "I'd rather they didn't remember me more than necessary."},
		{Lang: "hu", Text: "Inkább nem akarom, hogy megjegyezzenek a kelleténél jobban."},
	})
}

func (street policeStreet) fatCop() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "fat_cop",
		Label: inkwell.Text{
			{Lang: "en", Text: "cop eating a donut"},
			{Lang: "hu", Text: "fánkot evő rendőr"},
		},
		Area:         inkwell.Rect(396, 164, 48, 84),
		OnLook:       street.noticed(),
		OnUse:        street.noticed(),
		OnTalk:       street.noticed(),
		OnTake:       street.noticed(),
		OnUseWithAny: street.noticed(),
	}
}

func (street policeStreet) thinFigure() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "thin_figure",
		Label: inkwell.Text{
			{Lang: "en", Text: "thin figure in the passenger seat"},
			{Lang: "hu", Text: "vékony alak az anyósülésen"},
		},
		Area:         inkwell.Rect(476, 146, 48, 84),
		OnLook:       street.noticed(),
		OnUse:        street.noticed(),
		OnTalk:       street.noticed(),
		OnTake:       street.noticed(),
		OnUseWithAny: street.noticed(),
	}
}

func (policeStreet) cruiser() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "cruiser",
		Label: inkwell.Text{
			{Lang: "en", Text: "parked cruiser"},
			{Lang: "hu", Text: "parkoló rendőrautó"},
		},
		Area: inkwell.Rect(440, 178, 180, 60),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Engine running, doors open, donut crumbs on the dash. Nobody is in a hurry."},
			{Lang: "hu", Text: "Jár a motor, nyitva az ajtók, fánkmorzsa a műszerfalon. Senki nem siet."},
		}),
		OnUse: inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm not getting into a police car voluntarily."},
				{Lang: "hu", Text: "Önként nem szállok be egy rendőrautóba."},
			}),
			inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "That leaves the other way of getting in. Keep walking."},
				{Lang: "hu", Text: "Akkor marad a másik módja a beszállásnak. Menj tovább."},
			}),
		),
		OnTake: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Grand theft, in front of two officers. No."},
			{Lang: "hu", Text: "Autólopás, két rendőr előtt. Nem."},
		}),
	}
}

func (policeStreet) Scene() *inkwell.Scene { return PoliceStreet }
