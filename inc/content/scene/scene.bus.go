package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var Bus = SceneManager.Register(&Scene{
	Name: "bus",
	Title: Text{
		{Lang: "en", Text: "BUS — AIRPORT TO THE CITY"},
		{Lang: "hu", Text: "BUSZ — A REPTÉRRŐL A VÁROSBA"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "bus",
				Path: "assets/bg/bus.jpeg",
			},
			Z: -1,
		},
	},
	Triggers: []Trigger{
		{
			Name: "tutorial",
			When: Given{
				NotFlag: inc.FlagBusTutorialDone,
			},
			Do:   inkwell.RunScript(script.BusTutorial),
			Once: true,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "window",
			Label: Text{
				{Lang: "en", Text: "bus window"},
				{Lang: "hu", Text: "buszablak"},
			},
			Area: inkwell.Rect(60, 40, 520, 110),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Same hills. New billboards. Half of them are for things you wear on your wrist."},
					{Lang: "hu", Text: "Ugyanazok a dombok. Új plakátok. A felük olyasmit hirdet, amit a csuklódon hordasz."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "And the other half are for the licence to use them. Welcome back."},
					{Lang: "hu", Text: "A másik felük meg az engedélyt, hogy használhasd őket. Isten hozott itthon."},
				}),
			),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "It doesn't open. They never do any more."},
				{Lang: "hu", Text: "Nem nyílik. Már sosem nyílnak."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "The city stays where it is."},
				{Lang: "hu", Text: "A város marad, ahol van."},
			}),
		},
		{
			Name: "seat",
			Label: Text{
				{Lang: "en", Text: "empty seat"},
				{Lang: "hu", Text: "üres ülés"},
			},
			Area: inkwell.Rect(380, 160, 100, 90),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Somebody carved a phone number into it. Area code from before the split."},
				{Lang: "hu", Text: "Valaki egy telefonszámot karcolt bele. A körzetszám még a szétválás előttről."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "I'm sitting. That's the whole plan for the next forty minutes."},
				{Lang: "hu", Text: "Ülök. Ez az egész terv a következő negyven percre."},
			}),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Bolted down. The city learned that one early."},
				{Lang: "hu", Text: "Lecsavarozva. Ezt a város korán megtanulta."},
			}),
		},
		{
			Name: "armilla",
			Label: Text{
				{Lang: "en", Text: "Paul's Armilla"},
				{Lang: "hu", Text: "Paul Armillája"},
			},
			Area: inkwell.Rect(150, 170, 70, 60),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Two slots. Slot one is Dex. Slot two is whatever I find and shouldn't."},
				{Lang: "hu", Text: "Két foglalat. Az egyesben Dex. A kettesben az, amit találok, és nem kéne."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Push a tape into the second slot down there and it talks. Push me out of the first and I don't."},
				{Lang: "hu", Text: "Nyomj egy kazettát a kettes foglalatba odalent, és beszélni fog. Nyomj ki engem az egyesből, és én nem."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Dex)),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "It's on my wrist. That's as taken as it gets."},
				{Lang: "hu", Text: "A csuklómon van. Ennél felvettebb már nem lesz."},
			}),
		},
	},
})

var _ = Bus.Leads(
	Exit{
		To: NoodleHouse,
		Label: Text{
			{Lang: "en", Text: "get off at Noodle's street"},
			{Lang: "hu", Text: "leszállni Noodle utcájánál"},
		},
		Side: inkwell.ExitRight,
	},
)
