package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var NoodleOffice = SceneManager.Register(&Scene{
	Name: "noodle_office",
	Title: Text{
		{Lang: "en", Text: "NOODLE'S OFFICE"},
		{Lang: "hu", Text: "NOODLE IRODÁJA"},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 120,
				Y: 232,
			},
		},
		{
			Character: character.Noodle,
			At: Point{
				X: 420,
				Y: 240,
			},
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "noodle",
			Label: Text{
				{Lang: "en", Text: "Noodle"},
				{Lang: "hu", Text: "Noodle"},
			},
			Area: inkwell.Rect(392, 160, 56, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Vest open, tie loose, chips in one hand. He looks like the twenty years happened to somebody else."},
				{Lang: "hu", Text: "Nyitott mellény, lógó nyakkendő, chips az egyik kezében. Úgy néz ki, mintha a húsz év mással történt volna."},
			}),
			OnUse: character.Noodle.Say(Text{
				{Lang: "en", Text: "Hands off the merchandise, Paul."},
				{Lang: "hu", Text: "Kezeket el az árutól, Paul."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.NoodleOffice)),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "He weighs more than your whole inventory."},
				{Lang: "hu", Text: "Többet nyom, mint az egész leltárad."},
			}),
			OnUseWithAny: character.Noodle.Say(Text{
				{Lang: "en", Text: "Is that for me? Put it down, I'm eating."},
				{Lang: "hu", Text: "Ez nekem lesz? Tedd le, eszem."},
			}),
		},
		{
			Name: "chair",
			Label: Text{
				{Lang: "en", Text: "Noodle's chair"},
				{Lang: "hu", Text: "Noodle széke"},
			},
			Area: inkwell.Rect(440, 180, 60, 70),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Leather, cracked, moulded to one man. It's not getting up either."},
				{Lang: "hu", Text: "Bőr, repedt, egy emberhez formálódott. Ez sem fog felállni."},
			}),
			OnUse: character.Noodle.Say(Text{
				{Lang: "en", Text: "Occupied."},
				{Lang: "hu", Text: "Foglalt."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "It comes with him attached."},
				{Lang: "hu", Text: "Ő is jönne vele."},
			}),
		},
		{
			Name: "armilla_rack",
			Label: Text{
				{Lang: "en", Text: "rack of cracked Armillas"},
				{Lang: "hu", Text: "állvány feltört Armillákkal"},
			},
			Area: inkwell.Rect(40, 60, 150, 170),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Twenty bracelets, every seal broken. He got them back from the police."},
					{Lang: "hu", Text: "Húsz karperec, mindegyiken törött pecsét. Visszakapta őket a rendőrségtől."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Which tells you what they were actually after. Not these."},
					{Lang: "hu", Text: "Ami elárulja, mit kerestek valójában. Nem ezeket."},
				}),
			),
			OnUse: character.Noodle.Say(Text{
				{Lang: "en", Text: "Those are inventory. You of all people."},
				{Lang: "hu", Text: "Az leltár. Pont te."},
			}),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I've got one of those already, and it's one too many."},
				{Lang: "hu", Text: "Van már egy ilyenem, és az is sok."},
			}),
		},
		{
			Name: "monocle",
			Label: Text{
				{Lang: "en", Text: "monocle on the desk"},
				{Lang: "hu", Text: "monokli az asztalon"},
			},
			Area: inkwell.Rect(300, 190, 40, 30),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "A monocle. Prescription. He wears it instead of glasses because glasses are, I quote, for accountants."},
				{Lang: "hu", Text: "Monokli. Dioptriás. Szemüveg helyett hordja, mert a szemüveg, idézem, könyvelőknek való."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Everything's blurrier through it. Including him."},
				{Lang: "hu", Text: "Minden homályosabb rajta át. Ő is."},
			}),
			OnTake: character.Noodle.Say(Text{
				{Lang: "en", Text: "Paul. I can see you doing that."},
				{Lang: "hu", Text: "Paul. Látom, mit csinálsz."},
			}),
		},
	},
})

var _ = NoodleOffice.Leads(
	Exit{
		To: NoodleHouse,
		Label: Text{
			{Lang: "en", Text: "back out to the street"},
			{Lang: "hu", Text: "vissza az utcára"},
		},
		Side: inkwell.ExitNear,
	},
)
