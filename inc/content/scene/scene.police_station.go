package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var policeStationTip = inkwell.Seq(
	character.Receptionist.Say(Text{
		{Lang: "en", Text: "Sir. A moment."},
		{Lang: "hu", Text: "Uram. Egy pillanat."},
	}),
	character.Receptionist.Say(Text{
		{Lang: "en", Text: "What your friend meant for you isn't in the house. Behind it. The alley, a loose brick at the foot of the wall."},
		{Lang: "hu", Text: "Ami a barátja maga számára szánt, az nincs a házban. Mögötte. A sikátorban, egy laza tégla a fal tövében."},
	}),
	character.Receptionist.Say(Text{
		{Lang: "en", Text: "Take it before somebody else finds it. I never said this."},
		{Lang: "hu", Text: "Vigye el, mielőtt más találja meg. Én ezt nem mondtam."},
	}),
	inkwell.SetFlag(inc.FlagPoliceTip),
	speaker.Dex.Say(Text{
		{Lang: "en", Text: "A government office just tipped you off to remove something from a scene. Hold on to that thought."},
		{Lang: "hu", Text: "Egy állami hivatal most súgta meg, hogy vigyél el valamit egy helyszínről. Tartsd észben ezt a gondolatot."},
	}),
)

var PoliceStation = SceneManager.Register(&Scene{
	Name: "police_station",
	Title: Text{
		{Lang: "en", Text: "SFPD — RECEPTION"},
		{Lang: "hu", Text: "SFPD — RECEPCIÓ"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "police_station",
				Path: "assets/bg/police_station.png",
			},
			Z: -1,
		},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 120,
				Y: 236,
			},
		},
		{
			Character: character.Receptionist,
			At: Point{
				X: 330,
				Y: 214,
			},
		},
		{
			Character: character.Lawyer,
			At: Point{
				X: 520,
				Y: 240,
			},
		},
		{
			Character: character.WorriedRelative,
			At: Point{
				X: 580,
				Y: 240,
			},
		},
	},
	Triggers: []Trigger{
		{
			Name: "tip",
			Once: true,
			When: Given{
				Flag:    inc.FlagVasquezDeal,
				NotFlag: inc.FlagPoliceTip,
			},
			Do: policeStationTip,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "receptionist",
			Label: Text{
				{Lang: "en", Text: "receptionist"},
				{Lang: "hu", Text: "recepciós"},
			},
			Area: inkwell.Rect(306, 134, 48, 84),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Thick glasses, coffee on the shirt, paper up to the chin. Bored, not hostile."},
				{Lang: "hu", Text: "Vastag szemüveg, kávé az ingen, papír az álláig. Unott, nem ellenséges."},
			}),
			OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Receptionist)),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "She's a person. You talk to those."},
				{Lang: "hu", Text: "Ő egy ember. Azokkal beszélni kell."},
			}),
			OnUseWithAny: character.Receptionist.Say(Text{
				{Lang: "en", Text: "Put that away, sir. This is a police station."},
				{Lang: "hu", Text: "Tegye el, uram. Ez egy rendőrség."},
			}),
		},
		{
			Name: "lawyer",
			Label: Text{
				{Lang: "en", Text: "tired lawyer"},
				{Lang: "hu", Text: "fáradt ügyvéd"},
			},
			Area: inkwell.Rect(496, 160, 48, 84),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Wrinkled suit, briefcase in a death grip, eyes on the floor."},
					{Lang: "hu", Text: "Gyűrött öltöny, az aktatáska halálos szorításban, a szeme a padlón."},
				}),
				character.Lawyer.Say(Text{
					{Lang: "en", Text: "Why did we have to buy that Personal Tape off the market… it was a mistake."},
					{Lang: "hu", Text: "Miért kellett a piacról vennünk azt a Personal Tape-et… hiba volt."},
				}),
				character.Lawyer.Say(Text{
					{Lang: "en", Text: "Why does the BVK even bother with these. Nobody could have known."},
					{Lang: "hu", Text: "A BVK miért foglalkozik ilyenekkel. Senki sem tehetett az egészről."},
				}),
			),
			OnTalk: inkwell.Seq(
				character.Lawyer.Say(Text{
					{Lang: "en", Text: "Not now."},
					{Lang: "hu", Text: "Most nem alkalmas."},
				}),
				character.Lawyer.Say(Text{
					{Lang: "en", Text: "I'm waiting."},
					{Lang: "hu", Text: "Várakozom."},
				}),
			),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "I don't need a lawyer. Yet."},
				{Lang: "hu", Text: "Nem kell ügyvéd. Még."},
			}),
		},
		{
			Name: "worried_relative",
			Label: Text{
				{Lang: "en", Text: "worried woman"},
				{Lang: "hu", Text: "aggódó nő"},
			},
			Area: inkwell.Rect(556, 160, 48, 84),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Middle-aged, wringing her coat, watching a closed door."},
					{Lang: "hu", Text: "Középkorú, a kabátját gyűrögeti, egy zárt ajtót néz."},
				}),
				character.WorriedRelative.Say(Text{
					{Lang: "en", Text: "What does the BVK care about some junk bought at a market…"},
					{Lang: "hu", Text: "Mit érdekli a BVK-t egy ilyen piacon vett hülyeség…"},
				}),
				character.WorriedRelative.Say(Text{
					{Lang: "en", Text: "It was just a small business. Not a crime…"},
					{Lang: "hu", Text: "Csak egy kis üzlet volt. Nem bűnözés…"},
				}),
			),
			OnTalk: inkwell.Seq(
				character.WorriedRelative.Say(Text{
					{Lang: "en", Text: "Please leave me alone."},
					{Lang: "hu", Text: "Kérem, hagyjon békén."},
				}),
				character.WorriedRelative.Say(Text{
					{Lang: "en", Text: "I don't want to talk about it."},
					{Lang: "hu", Text: "Nem akarok beszélni róla."},
				}),
			),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Leave her be."},
				{Lang: "hu", Text: "Hagyjuk békén."},
			}),
		},
		{
			Name: "paper_stacks",
			Label: Text{
				{Lang: "en", Text: "stacks of paperwork"},
				{Lang: "hu", Text: "papírkupacok"},
			},
			Area: inkwell.Rect(240, 168, 60, 50),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Folders, forms, carbon copies. Somewhere in there is Noodle, in triplicate."},
				{Lang: "hu", Text: "Dossziék, nyomtatványok, indigós másolatok. Valahol ott van Noodle is, három példányban."},
			}),
			OnUse: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Rifling through police files at the front desk. Bold."},
					{Lang: "hu", Text: "Rendőrségi aktákban túrni a recepción. Merész."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Bold is one word for it."},
					{Lang: "hu", Text: "A merész egy szó rá."},
				}),
			),
			OnTake: character.Receptionist.Say(Text{
				{Lang: "en", Text: "Hands off the desk, sir."},
				{Lang: "hu", Text: "Kezeket le az asztalról, uram."},
			}),
		},
		{
			Name: "monitor",
			Label: Text{
				{Lang: "en", Text: "green-text monitor"},
				{Lang: "hu", Text: "zöld betűs monitor"},
			},
			Area: inkwell.Rect(360, 150, 50, 44),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "A BVK-stamped terminal. Booking records, by the look of it. Facing away from me."},
				{Lang: "hu", Text: "BVK-pecsétes terminál. Előállítási nyilvántartás, úgy nézem. Elfordítva tőlem."},
			}),
			OnUse: character.Receptionist.Say(Text{
				{Lang: "en", Text: "That's mine, sir."},
				{Lang: "hu", Text: "Az az enyém, uram."},
			}),
		},
	},
})

var _ = PoliceStation.Leads(
	Exit{
		To: PoliceStreet,
		Label: Text{
			{Lang: "en", Text: "back out to the street"},
			{Lang: "hu", Text: "vissza az utcára"},
		},
		Side: inkwell.ExitNear,
	},
	Exit{
		To: InterrogationRoom,
		Label: Text{
			{Lang: "en", Text: "the interview rooms"},
			{Lang: "hu", Text: "a kihallgatószobák"},
		},
		Side:  inkwell.ExitBack,
		Needs: inc.FlagReceptionPassed,
		Blocked: character.Receptionist.Say(Text{
			{Lang: "en", Text: "Reason for your visit first, sir."},
			{Lang: "hu", Text: "Először a látogatás okát, uram."},
		}),
	},
)
