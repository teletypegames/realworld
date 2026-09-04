package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
)

var normanApartmentSeen = inkwell.Say(character.Paul, inkwell.Text{
	{Lang: "en", Text: "His Armilla. Slot one empty. Still here, and he isn't."},
	{Lang: "hu", Text: "Az Armillája. Az egyes foglalat üres. Ez itt maradt, ő nem."},
})

var normanApartmentGone = inkwell.Say(character.Paul, inkwell.Text{
	{Lang: "en", Text: "Just the jacket now. The wallet's on me."},
	{Lang: "hu", Text: "Már csak a zakó. A tárca nálam van."},
})

var NormanApartment = SceneManager.Register(&Scene{
	Name: "norman_apartment",
	Title: inkwell.Text{
		{Lang: "en", Text: "NORMAN'S APARTMENT"},
		{Lang: "hu", Text: "NORMAN LAKÁSA"},
	},
	Hotspots: []inkwell.Hotspot{
		{
			Name: "armilla",
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman's Armilla"},
				{Lang: "hu", Text: "Norman Armillája"},
			},
			Area:   inkwell.Rect(300, 150, 60, 40),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagArmillaSeen), normanApartmentSeen, inkwell.RunScript(script.ArmillaFound)),
			OnUse:  inkwell.If(inkwell.Flag(inc.FlagArmillaSeen), normanApartmentSeen, inkwell.RunScript(script.ArmillaFound)),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "One wrist, one Armilla. And it's his."},
				{Lang: "hu", Text: "Egy csukló, egy Armilla. És ez az övé."},
			}),
		},
		{
			Name: "terminal",
			Label: inkwell.Text{
				{Lang: "en", Text: "Norman's computer"},
				{Lang: "hu", Text: "Norman számítógépe"},
			},
			Area: inkwell.Rect(376, 120, 90, 72),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Left on. Mail client, a diary folder, and a fan that has been dying for months."},
				{Lang: "hu", Text: "Bekapcsolva hagyva. Levelező, egy naplómappa, és egy ventilátor, ami hónapok óta haldoklik."},
			}),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNormanTerminalRead),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "I've read it. Norman didn't believe in himself, he believed in Noodle and in his tape. And there's the lift code."},
					{Lang: "hu", Text: "Elolvastam. Norman nem magában hitt, hanem Noodle-ban és a kazettájában. És ott a liftkód."},
				}),
				inkwell.RunScript(script.NormanTerminal),
			),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I sell these. I don't steal them."},
				{Lang: "hu", Text: "Ilyeneket árulok. Nem lopok."},
			}),
		},
		{
			Name: "jacket",
			Label: inkwell.Text{
				{Lang: "en", Text: "jacket on the chair"},
				{Lang: "hu", Text: "zakó a széken"},
			},
			Area: inkwell.Rect(200, 160, 60, 70),
			OnLook: inkwell.If(inkwell.HasItem(item.NormanWallet),
				normanApartmentGone,
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "His jacket, over the chair. Whoever took him didn't let him dress."},
					{Lang: "hu", Text: "A zakója, a széken. Aki elvitte, nem hagyta felöltözni."},
				}),
			),
			OnUse:  inkwell.If(inkwell.HasItem(item.NormanWallet), normanApartmentGone, inkwell.RunScript(script.WalletFound)),
			OnTake: inkwell.If(inkwell.HasItem(item.NormanWallet), normanApartmentGone, inkwell.RunScript(script.WalletFound)),
		},
		{
			Name: "bed",
			Label: inkwell.Text{
				{Lang: "en", Text: "unmade bed"},
				{Lang: "hu", Text: "bevetetlen ágy"},
			},
			Area: inkwell.Rect(40, 150, 130, 80),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Slept in. Not recently."},
				{Lang: "hu", Text: "Aludtak benne. Nem mostanában."},
			}),
			OnUse: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "You are not napping in a missing man's bed, Paul."},
				{Lang: "hu", Text: "Nem fogsz egy eltűnt ember ágyában szunyókálni, Paul."},
			}),
		},
		{
			Name: "poster",
			Label: inkwell.Text{
				{Lang: "en", Text: "poster on the wall"},
				{Lang: "hu", Text: "poszter a falon"},
			},
			Area: inkwell.Rect(500, 40, 100, 90),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "A BVK recruitment poster. “Your mind, certified.” Somebody has drawn a moustache on the certificate."},
				{Lang: "hu", Text: "BVK toborzóplakát. „Az elméd, hitelesítve.” Valaki bajuszt rajzolt a tanúsítványra."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "The moustache is the only good part, and it's in pen."},
				{Lang: "hu", Text: "A bajusz az egyetlen jó rész, és tollal van."},
			}),
		},
	},
})

var _ = NormanApartment.Leads(
	exitToSelector,
	inkwell.Exit{
		To: RooftopHideout,
		Label: inkwell.Text{
			{Lang: "en", Text: "the stairs up to the roof"},
			{Lang: "hu", Text: "a lépcső fel a tetőre"},
		},
		Side: inkwell.ExitBack,
	},
	inkwell.Exit{
		To: BBSTerminal,
		Label: inkwell.Text{
			{Lang: "en", Text: "Norman's terminal"},
			{Lang: "hu", Text: "Norman terminálja"},
		},
		Side: inkwell.ExitRight,
	},
)
