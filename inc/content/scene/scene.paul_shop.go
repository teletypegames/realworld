package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type paulShop struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.ScenePaulShop,
	Title: inkwell.Text{
		{Lang: "en", Text: "PAUL'S SHOP — JUNK AND GARAGE"},
		{Lang: "hu", Text: "PAUL BOLTJA — ÓCSKAVAS ÉS MŰHELY"},
	},
	Exits: []inkwell.Exit{
		{
			To: inc.SceneBus,
			Label: inkwell.Text{
				{Lang: "en", Text: "the bus to San Francisco"},
				{Lang: "hu", Text: "a busz San Franciscóba"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Hotspots: paulShop{}.hotspots(),
})

func (shop paulShop) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		shop.doorMat(),
		shop.breakfastBox(),
		shop.armillaShelf(),
		shop.junkShelf(),
	}
}

func (paulShop) doorMat() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "door_mat",
		Label: inkwell.Text{
			{Lang: "en", Text: "door mat"},
			{Lang: "hu", Text: "lábtörlő"},
		},
		Area: inkwell.Rect(40, 200, 90, 40),
		OnLook: inkwell.If(inkwell.HasItem(inc.ItemNoodleLetter),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "The letter came under this door. Nothing else has, in years."},
				{Lang: "hu", Text: "Ez alatt az ajtó alatt jött a levél. Más évek óta nem."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Worn through in the middle. Like me."},
				{Lang: "hu", Text: "Középen elkopott. Mint én."},
			}),
		),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I wipe my feet on the way in, not on the way out."},
			{Lang: "hu", Text: "Befelé törlöm a lábam, nem kifelé."},
		}),
		OnTake: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "You are not packing the door mat, Paul."},
			{Lang: "hu", Text: "A lábtörlőt nem csomagolod be, Paul."},
		}),
	}
}

func (paulShop) breakfastBox() inkwell.Hotspot {
	take := inkwell.Seq(
		inkwell.Give(inc.ItemBreakfastBox),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Breakfast. Still warm-ish."},
			{Lang: "hu", Text: "Reggeli. Még többé-kevésbé meleg."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Take it. A man travels badly on an empty stomach, and worse on a full conscience."},
			{Lang: "hu", Text: "Vidd. Üres gyomorral rosszul utazik az ember, tele lelkiismerettel még rosszabbul."},
		}),
	)
	return inkwell.Hotspot{
		Name: "breakfast_box",
		Label: inkwell.Text{
			{Lang: "en", Text: "box of breakfast"},
			{Lang: "hu", Text: "doboz reggeli"},
		},
		Area: inkwell.Rect(250, 150, 60, 40),
		OnLook: inkwell.If(inkwell.HasItem(inc.ItemBreakfastBox),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "A grease ring where the box was."},
				{Lang: "hu", Text: "Zsírkarika, ahol a doboz volt."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Egg roll, two dumplings, a napkin I didn't ask for."},
				{Lang: "hu", Text: "Tojásos tekercs, két gombóc, egy szalvéta, amit nem kértem."},
			}),
		),
		OnUse: inkwell.If(inkwell.HasItem(inc.ItemBreakfastBox),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "It's in my bag already."},
				{Lang: "hu", Text: "Már a táskámban van."},
			}),
			take,
		),
		OnTake: inkwell.If(inkwell.HasItem(inc.ItemBreakfastBox),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "One breakfast is what I ordered, and one is what I have."},
				{Lang: "hu", Text: "Egy reggelit rendeltem, és egy van nálam."},
			}),
			take,
		),
	}
}

func (paulShop) armillaShelf() inkwell.Hotspot {
	take := inkwell.Seq(
		inkwell.Give(inc.ItemBlackArmilla),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Might sell it on the coast. Might not."},
			{Lang: "hu", Text: "A parton talán eladom. Talán nem."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "You bought it to fix it. That was three years ago."},
			{Lang: "hu", Text: "Azért vetted, hogy megjavítsd. Három éve volt."},
		}),
	)
	return inkwell.Hotspot{
		Name: "armilla_shelf",
		Label: inkwell.Text{
			{Lang: "en", Text: "black-market Armilla on the shelf"},
			{Lang: "hu", Text: "feketepiaci Armilla a polcon"},
		},
		Area: inkwell.Rect(400, 90, 70, 50),
		OnLook: inkwell.If(inkwell.HasItem(inc.ItemBlackArmilla),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Dust in the shape of a bracelet."},
				{Lang: "hu", Text: "Karperec alakú por."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Two slots, cracked firmware, a Chinese ad for something inflatable every time it boots."},
				{Lang: "hu", Text: "Két foglalat, feltört firmware, és minden indításnál egy kínai reklám valami felfújhatóról."},
			}),
		),
		OnUse: inkwell.If(inkwell.HasItem(inc.ItemBlackArmilla),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Already on me."},
				{Lang: "hu", Text: "Már nálam van."},
			}),
			take,
		),
		OnTake: inkwell.If(inkwell.HasItem(inc.ItemBlackArmilla),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I only had the one."},
				{Lang: "hu", Text: "Csak egy volt belőle."},
			}),
			take,
		),
	}
}

func (paulShop) junkShelf() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "junk_shelf",
		Label: inkwell.Text{
			{Lang: "en", Text: "shelves of junk"},
			{Lang: "hu", Text: "polcok tele kacattal"},
		},
		Area: inkwell.Rect(500, 60, 120, 170),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Modems, a rotary phone, half a synth. Every one of them somebody's future, once."},
			{Lang: "hu", Text: "Modemek, egy tárcsás telefon, fél szintetizátor. Mindegyik valakinek a jövője volt egyszer."},
		}),
		OnUse: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I know where everything is. That's the business."},
				{Lang: "hu", Text: "Tudom, mi hol van. Ez az üzlet."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "The business is knowing what it's worth. You only do the first half."},
				{Lang: "hu", Text: "Az üzlet az, hogy tudd, mit ér. Te csak az első felét csinálod."},
			}),
		),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm travelling light for once."},
			{Lang: "hu", Text: "Egyszer az életben könnyű csomaggal utazom."},
		}),
	}
}
