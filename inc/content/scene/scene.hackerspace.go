package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

type hackerspace struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneHackerspace,
	Title: inkwell.Text{
		{Lang: "en", Text: "HACKERSPACE"},
		{Lang: "hu", Text: "HACKERSPACE"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
	},
	Hotspots: hackerspace{}.hotspots(),
})

func (place hackerspace) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.workbench(),
		place.printer(),
		place.notes(),
	}
}

func (hackerspace) workbench() inkwell.Hotspot {
	handOver := inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Half-finished access tape. Datacenter door group, cloned off somebody's badge."},
			{Lang: "hu", Text: "Félkész belépőkazetta. Adatközponti ajtócsoport, valaki kitűzőjéről klónozva."},
		}),
		inkwell.Give(inc.ItemServerFarmTape),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "The kids here would rather you had it than the BVK. They just won't say so to your face."},
			{Lang: "hu", Text: "Az itteni kölykök inkább nálad látják, mint a BVK-nál. Csak nem mondják a szemedbe."},
		}),
	)
	return inkwell.Hotspot{
		Name: "workbench",
		Label: inkwell.Text{
			{Lang: "en", Text: "workbench"},
			{Lang: "hu", Text: "munkapad"},
		},
		Area: inkwell.Rect(360, 150, 200, 90),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Soldering iron still warm. A tape in a cradle, wired to a laptop that's seen better decades."},
			{Lang: "hu", Text: "A forrasztópáka még meleg. Egy kazetta a bölcsőben, egy laptopra kötve, amely jobb évtizedeket is látott."},
		}),
		OnTake: inkwell.If(inkwell.HasItem(inc.ItemServerFarmTape),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I've already got it."},
				{Lang: "hu", Text: "Már megvan."},
			}),
			inkwell.If(inkwell.Flag(inc.FlagTapeSpeaks),
				handOver,
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Nobody here hands a stranger a credential. I'd need a reason they'd believe."},
					{Lang: "hu", Text: "Itt senki nem ad egy idegennek belépőt. Kéne egy ok, amit elhisznek."},
				}),
			),
		),
		OnUse: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "Don't touch the iron. You need the hand for clicking."},
			{Lang: "hu", Text: "Ne nyúlj a pákához. A kezed kell még kattintani."},
		}),
	}
}

func (hackerspace) printer() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "printer",
		Label: inkwell.Text{
			{Lang: "en", Text: "3D printer"},
			{Lang: "hu", Text: "3D nyomtató"},
		},
		Area: inkwell.Rect(60, 140, 100, 100),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Printing an Armilla shell. Two slots. Somebody in here has taste."},
			{Lang: "hu", Text: "Egy Armilla-házat nyomtat. Két foglalattal. Valakinek itt van ízlése."},
		}),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "It's someone's job. I don't interrupt jobs."},
			{Lang: "hu", Text: "Ez valakinek a munkája. Munkát nem szakítok meg."},
		}),
	}
}

func (hackerspace) notes() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "wall_of_notes",
		Label: inkwell.Text{
			{Lang: "en", Text: "wall of notes"},
			{Lang: "hu", Text: "jegyzetekkel teli fal"},
		},
		Area: inkwell.Rect(200, 30, 260, 110),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "String, pins, printouts. Licence numbers, BVK office addresses, and a sun drawn under a wavy line."},
				{Lang: "hu", Text: "Zsinór, gombostűk, nyomatok. Licencszámok, BVK-irodacímek, és egy nap egy hullámos vonal alá rajzolva."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "The underwater sun. It gets around."},
				{Lang: "hu", Text: "A vízalatti nap. Sok helyre eljut."},
			}),
		),
	}
}
