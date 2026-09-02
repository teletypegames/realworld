package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type smallRestaurant struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneSmallRestaurant,
	Title: inkwell.Text{
		{Lang: "en", Text: "SMALL RESTAURANT — NEXT DOOR"},
		{Lang: "hu", Text: "KIS ÉTTEREM — A SZOMSZÉDBAN"},
	},
	Exits: []inkwell.Exit{
		{
			To: inc.SceneTrinketShop,
			Label: inkwell.Text{
				{Lang: "en", Text: "back to the trinket shop"},
				{Lang: "hu", Text: "vissza a csecsebecseboltba"},
			},
			Side: inkwell.ExitLeft,
		},
		{
			To: inc.SceneRestaurantToilet,
			Label: inkwell.Text{
				{Lang: "en", Text: "the toilet at the back"},
				{Lang: "hu", Text: "a hátsó mosdó"},
			},
			Side: inkwell.ExitBack,
		},
	},
	Hotspots: smallRestaurant{}.hotspots(),
})

func (place smallRestaurant) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.counter(),
		place.table(),
		place.sign(),
	}
}

func (smallRestaurant) counter() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "counter",
		Label: inkwell.Text{
			{Lang: "en", Text: "counter"},
			{Lang: "hu", Text: "pult"},
		},
		Area: inkwell.Rect(40, 130, 200, 70),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Steam, a bell, and nobody behind it. The kitchen's arguing with itself."},
			{Lang: "hu", Text: "Gőz, egy csengő, és senki mögötte. A konyha önmagával vitatkozik."},
		}),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I ring the bell. The arguing stops, then starts again. Nobody comes."},
			{Lang: "hu", Text: "Megnyomom a csengőt. A vita elhal, aztán újrakezdődik. Senki nem jön."},
		}),
		OnTake: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "You can't take a counter, Paul. We've been over this in three buildings."},
			{Lang: "hu", Text: "Egy pultot nem vihetsz el, Paul. Ezt már három épületben megbeszéltük."},
		}),
	}
}

func (smallRestaurant) table() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "table",
		Label: inkwell.Text{
			{Lang: "en", Text: "corner table"},
			{Lang: "hu", Text: "sarokasztal"},
		},
		Area: inkwell.Rect(300, 160, 120, 50),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Four chairs, one bowl, still warm. Someone left in a hurry."},
				{Lang: "hu", Text: "Négy szék, egy tál, még meleg. Valaki sietve ment el."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "“Strange types around the eatery,” the cop said. Strange types don't finish their soup."},
				{Lang: "hu", Text: "„Fura alakok a kajáldánál”, mondta a zsaru. A fura alakok nem eszik meg a levesüket."},
			}),
		),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not sitting down. I'm not hungry, and I'm not staying."},
			{Lang: "hu", Text: "Nem ülök le. Nem vagyok éhes, és nem maradok."},
		}),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Someone else's soup. No."},
			{Lang: "hu", Text: "Valaki más levese. Nem."},
		}),
	}
}

func (smallRestaurant) sign() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "sign",
		Label: inkwell.Text{
			{Lang: "en", Text: "flickering sign"},
			{Lang: "hu", Text: "vibráló felirat"},
		},
		Area: inkwell.Rect(460, 40, 150, 50),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "“TOILET” in neon. The T's out, and the rest is thinking about it."},
				{Lang: "hu", Text: "„MOSDÓ” neonban. Az M kiégett, a többi is gondolkodik rajta."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "The wiring in this block flickers on a rhythm. That isn't a fault, that's a pulse."},
				{Lang: "hu", Text: "Ebben a tömbben a vezetékek ritmusra vibrálnak. Az nem hiba, az pulzus."},
			}),
		),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It's up on the wall. I'd need a ladder and a reason."},
			{Lang: "hu", Text: "Fent van a falon. Kéne egy létra és egy ok."},
		}),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not stealing a neon sign."},
			{Lang: "hu", Text: "Nem lopok el egy neonfeliratot."},
		}),
	}
}
