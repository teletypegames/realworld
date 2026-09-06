package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var normanApartmentSeen = character.Paul.Say(Text{
	{Lang: "en", Text: "His Armilla. Slot one empty. Still here, and he isn't."},
	{Lang: "hu", Text: "Az Armillája. Az egyes foglalat üres. Ez itt maradt, ő nem."},
})

var normanApartmentGone = character.Paul.Say(Text{
	{Lang: "en", Text: "Just the jacket now. The wallet's on me."},
	{Lang: "hu", Text: "Már csak a zakó. A tárca nálam van."},
})

var NormanApartment = SceneManager.Register(&Scene{
	Name: "norman_apartment",
	Title: Text{
		{Lang: "en", Text: "NORMAN'S APARTMENT"},
		{Lang: "hu", Text: "NORMAN LAKÁSA"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "norman_apartment",
				Path: "assets/bg/norman_apartment.png",
			},
			Z: -1,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "armilla",
			Label: Text{
				{Lang: "en", Text: "Norman's Armilla"},
				{Lang: "hu", Text: "Norman Armillája"},
			},
			Area:   inkwell.Rect(300, 150, 60, 40),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagArmillaSeen), normanApartmentSeen, inkwell.RunScript(script.ArmillaFound)),
			OnUse:  inkwell.If(inkwell.Flag(inc.FlagArmillaSeen), normanApartmentSeen, inkwell.RunScript(script.ArmillaFound)),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "One wrist, one Armilla. And it's his."},
				{Lang: "hu", Text: "Egy csukló, egy Armilla. És ez az övé."},
			}),
		},
		{
			Name: "terminal",
			Label: Text{
				{Lang: "en", Text: "Norman's computer"},
				{Lang: "hu", Text: "Norman számítógépe"},
			},
			Area: inkwell.Rect(376, 120, 90, 72),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Left on. Mail client, a diary folder, and a fan that has been dying for months."},
				{Lang: "hu", Text: "Bekapcsolva hagyva. Levelező, egy naplómappa, és egy ventilátor, ami hónapok óta haldoklik."},
			}),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNormanTerminalRead),
				character.Paul.Say(Text{
					{Lang: "en", Text: "I've read it. Norman didn't believe in himself, he believed in Noodle and in his tape. And there's the lift code."},
					{Lang: "hu", Text: "Elolvastam. Norman nem magában hitt, hanem Noodle-ban és a kazettájában. És ott a liftkód."},
				}),
				inkwell.RunScript(script.NormanTerminal),
			),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I sell these. I don't steal them."},
				{Lang: "hu", Text: "Ilyeneket árulok. Nem lopok."},
			}),
		},
		{
			Name: "jacket",
			Label: Text{
				{Lang: "en", Text: "jacket on the chair"},
				{Lang: "hu", Text: "zakó a széken"},
			},
			Area: inkwell.Rect(200, 160, 60, 70),
			OnLook: inkwell.If(inkwell.HasItem(item.NormanWallet),
				normanApartmentGone,
				character.Paul.Say(Text{
					{Lang: "en", Text: "His jacket, over the chair. Whoever took him didn't let him dress."},
					{Lang: "hu", Text: "A zakója, a széken. Aki elvitte, nem hagyta felöltözni."},
				}),
			),
			OnUse:  inkwell.If(inkwell.HasItem(item.NormanWallet), normanApartmentGone, inkwell.RunScript(script.WalletFound)),
			OnTake: inkwell.If(inkwell.HasItem(item.NormanWallet), normanApartmentGone, inkwell.RunScript(script.WalletFound)),
		},
		{
			Name: "bed",
			Label: Text{
				{Lang: "en", Text: "unmade bed"},
				{Lang: "hu", Text: "bevetetlen ágy"},
			},
			Area: inkwell.Rect(40, 150, 130, 80),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Slept in. Not recently."},
				{Lang: "hu", Text: "Aludtak benne. Nem mostanában."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "You are not napping in a missing man's bed, Paul."},
				{Lang: "hu", Text: "Nem fogsz egy eltűnt ember ágyában szunyókálni, Paul."},
			}),
		},
		{
			Name: "poster",
			Label: Text{
				{Lang: "en", Text: "poster on the wall"},
				{Lang: "hu", Text: "poszter a falon"},
			},
			Area: inkwell.Rect(500, 40, 100, 90),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "A BVK recruitment poster. “Your mind, certified.” Somebody has drawn a moustache on the certificate."},
				{Lang: "hu", Text: "BVK toborzóplakát. „Az elméd, hitelesítve.” Valaki bajuszt rajzolt a tanúsítványra."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "The moustache is the only good part, and it's in pen."},
				{Lang: "hu", Text: "A bajusz az egyetlen jó rész, és tollal van."},
			}),
		},
	},
})

var _ = NormanApartment.Leads(
	exitToSelector,
	Exit{
		To: RooftopHideout,
		Label: Text{
			{Lang: "en", Text: "the stairs up to the roof"},
			{Lang: "hu", Text: "a lépcső fel a tetőre"},
		},
		Side: inkwell.ExitBack,
	},
	Exit{
		To: BBSTerminal,
		Label: Text{
			{Lang: "en", Text: "Norman's terminal"},
			{Lang: "hu", Text: "Norman terminálja"},
		},
		Side: inkwell.ExitRight,
	},
)
