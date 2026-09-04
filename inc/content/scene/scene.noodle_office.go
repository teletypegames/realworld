package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

type noodleOffice struct{}

var _ place = noodleOffice{}

var NoodleOffice = SceneManager.Register(&Scene{
	Name: "noodle_office",
	Title: inkwell.Text{
		{Lang: "en", Text: "NOODLE'S OFFICE"},
		{Lang: "hu", Text: "NOODLE IRODÁJA"},
	},
	Exits: []inkwell.Exit{
		{
			To: noodleHouse{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back out to the street"},
				{Lang: "hu", Text: "vissza az utcára"},
			},
			Side: inkwell.ExitNear,
		},
	},
	Actors: []inkwell.SceneActor{
		{
			Character: character.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 232,
			},
		},
		{
			Character: character.Noodle,
			At: inkwell.Point{
				X: 420,
				Y: 240,
			},
		},
	},
	Hotspots: noodleOffice{}.hotspots(),
})

func (noodleOffice) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		{
			Name: "noodle",
			Label: inkwell.Text{
				{Lang: "en", Text: "Noodle"},
				{Lang: "hu", Text: "Noodle"},
			},
			Area: inkwell.Rect(392, 160, 56, 84),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Vest open, tie loose, chips in one hand. He looks like the twenty years happened to somebody else."},
				{Lang: "hu", Text: "Nyitott mellény, lógó nyakkendő, chips az egyik kezében. Úgy néz ki, mintha a húsz év mással történt volna."},
			}),
			OnUse: inkwell.Say(character.Noodle, inkwell.Text{
				{Lang: "en", Text: "Hands off the merchandise, Paul."},
				{Lang: "hu", Text: "Kezeket el az árutól, Paul."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.NoodleOffice)),
			OnTake: inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "He weighs more than your whole inventory."},
				{Lang: "hu", Text: "Többet nyom, mint az egész leltárad."},
			}),
			OnUseWithAny: inkwell.Say(character.Noodle, inkwell.Text{
				{Lang: "en", Text: "Is that for me? Put it down, I'm eating."},
				{Lang: "hu", Text: "Ez nekem lesz? Tedd le, eszem."},
			}),
		},
		{
			Name: "chair",
			Label: inkwell.Text{
				{Lang: "en", Text: "Noodle's chair"},
				{Lang: "hu", Text: "Noodle széke"},
			},
			Area: inkwell.Rect(440, 180, 60, 70),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Leather, cracked, moulded to one man. It's not getting up either."},
				{Lang: "hu", Text: "Bőr, repedt, egy emberhez formálódott. Ez sem fog felállni."},
			}),
			OnUse: inkwell.Say(character.Noodle, inkwell.Text{
				{Lang: "en", Text: "Occupied."},
				{Lang: "hu", Text: "Foglalt."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It comes with him attached."},
				{Lang: "hu", Text: "Ő is jönne vele."},
			}),
		},
		{
			Name: "armilla_rack",
			Label: inkwell.Text{
				{Lang: "en", Text: "rack of cracked Armillas"},
				{Lang: "hu", Text: "állvány feltört Armillákkal"},
			},
			Area: inkwell.Rect(40, 60, 150, 170),
			OnLook: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Twenty bracelets, every seal broken. He got them back from the police."},
					{Lang: "hu", Text: "Húsz karperec, mindegyiken törött pecsét. Visszakapta őket a rendőrségtől."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "Which tells you what they were actually after. Not these."},
					{Lang: "hu", Text: "Ami elárulja, mit kerestek valójában. Nem ezeket."},
				}),
			),
			OnUse: inkwell.Say(character.Noodle, inkwell.Text{
				{Lang: "en", Text: "Those are inventory. You of all people."},
				{Lang: "hu", Text: "Az leltár. Pont te."},
			}),
			OnTake: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "I've got one of those already, and it's one too many."},
				{Lang: "hu", Text: "Van már egy ilyenem, és az is sok."},
			}),
		},
		{
			Name: "monocle",
			Label: inkwell.Text{
				{Lang: "en", Text: "monocle on the desk"},
				{Lang: "hu", Text: "monokli az asztalon"},
			},
			Area: inkwell.Rect(300, 190, 40, 30),
			OnLook: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "A monocle. Prescription. He wears it instead of glasses because glasses are, I quote, for accountants."},
				{Lang: "hu", Text: "Monokli. Dioptriás. Szemüveg helyett hordja, mert a szemüveg, idézem, könyvelőknek való."},
			}),
			OnUse: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Everything's blurrier through it. Including him."},
				{Lang: "hu", Text: "Minden homályosabb rajta át. Ő is."},
			}),
			OnTake: inkwell.Say(character.Noodle, inkwell.Text{
				{Lang: "en", Text: "Paul. I can see you doing that."},
				{Lang: "hu", Text: "Paul. Látom, mit csinálsz."},
			}),
		},
	}
}

func (noodleOffice) Scene() *inkwell.Scene { return NoodleOffice }
