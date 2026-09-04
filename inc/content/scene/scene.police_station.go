package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

type policeStation struct{}

var PoliceStation = SceneManager.Register(&Scene{
	Name: "police_station",
	Title: inkwell.Text{
		{Lang: "en", Text: "SFPD — RECEPTION"},
		{Lang: "hu", Text: "SFPD — RECEPCIÓ"},
	},
	Exits: []inkwell.Exit{
		{
			To: policeStreet{},
			Label: inkwell.Text{
				{Lang: "en", Text: "back out to the street"},
				{Lang: "hu", Text: "vissza az utcára"},
			},
			Side: inkwell.ExitNear,
		},
		{
			To: interrogationRoom{},
			Label: inkwell.Text{
				{Lang: "en", Text: "the interview rooms"},
				{Lang: "hu", Text: "a kihallgatószobák"},
			},
			Side:  inkwell.ExitBack,
			Needs: inc.FlagReceptionPassed,
			Blocked: inkwell.Say(character.Receptionist, inkwell.Text{
				{Lang: "en", Text: "Reason for your visit first, sir."},
				{Lang: "hu", Text: "Először a látogatás okát, uram."},
			}),
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
			Character: character.Receptionist,
			At: inkwell.Point{
				X: 330,
				Y: 214,
			},
		},
		{
			Character: character.Lawyer,
			At: inkwell.Point{
				X: 520,
				Y: 240,
			},
		},
		{
			Character: character.WorriedRelative,
			At: inkwell.Point{
				X: 580,
				Y: 240,
			},
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "tip",
			Once: true,
			When: inkwell.Given{
				Flag:    inc.FlagVasquezDeal,
				NotFlag: inc.FlagPoliceTip,
			},
			Do: policeStation{}.tip(),
		},
	},
	Hotspots: policeStation{}.hotspots(),
})

func (station policeStation) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		station.receptionist(),
		station.lawyer(),
		station.worriedRelative(),
		station.paperStacks(),
		station.monitor(),
	}
}

func (policeStation) tip() inkwell.Action {
	return inkwell.Seq(
		inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "Sir. A moment."},
			{Lang: "hu", Text: "Uram. Egy pillanat."},
		}),
		inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "What your friend meant for you isn't in the house. Behind it. The alley, a loose brick at the foot of the wall."},
			{Lang: "hu", Text: "Ami a barátja maga számára szánt, az nincs a házban. Mögötte. A sikátorban, egy laza tégla a fal tövében."},
		}),
		inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "Take it before somebody else finds it. I never said this."},
			{Lang: "hu", Text: "Vigye el, mielőtt más találja meg. Én ezt nem mondtam."},
		}),
		inkwell.SetFlag(inc.FlagPoliceTip),
		inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "A government office just tipped you off to remove something from a scene. Hold on to that thought."},
			{Lang: "hu", Text: "Egy állami hivatal most súgta meg, hogy vigyél el valamit egy helyszínről. Tartsd észben ezt a gondolatot."},
		}),
	)
}

func (policeStation) receptionist() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "receptionist",
		Label: inkwell.Text{
			{Lang: "en", Text: "receptionist"},
			{Lang: "hu", Text: "recepciós"},
		},
		Area: inkwell.Rect(306, 134, 48, 84),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Thick glasses, coffee on the shirt, paper up to the chin. Bored, not hostile."},
			{Lang: "hu", Text: "Vastag szemüveg, kávé az ingen, papír az álláig. Unott, nem ellenséges."},
		}),
		OnTalk: system.System.Paused(inkwell.RunDialog(dialog.Receptionist)),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "She's a person. You talk to those."},
			{Lang: "hu", Text: "Ő egy ember. Azokkal beszélni kell."},
		}),
		OnUseWithAny: inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "Put that away, sir. This is a police station."},
			{Lang: "hu", Text: "Tegye el, uram. Ez egy rendőrség."},
		}),
	}
}

func (policeStation) lawyer() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "lawyer",
		Label: inkwell.Text{
			{Lang: "en", Text: "tired lawyer"},
			{Lang: "hu", Text: "fáradt ügyvéd"},
		},
		Area: inkwell.Rect(496, 160, 48, 84),
		OnLook: inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Wrinkled suit, briefcase in a death grip, eyes on the floor."},
				{Lang: "hu", Text: "Gyűrött öltöny, az aktatáska halálos szorításban, a szeme a padlón."},
			}),
			inkwell.Say(character.Lawyer, inkwell.Text{
				{Lang: "en", Text: "Why did we have to buy that Personal Tape off the market… it was a mistake."},
				{Lang: "hu", Text: "Miért kellett a piacról vennünk azt a Personal Tape-et… hiba volt."},
			}),
			inkwell.Say(character.Lawyer, inkwell.Text{
				{Lang: "en", Text: "Why does the BVK even bother with these. Nobody could have known."},
				{Lang: "hu", Text: "A BVK miért foglalkozik ilyenekkel. Senki sem tehetett az egészről."},
			}),
		),
		OnTalk: inkwell.Seq(
			inkwell.Say(character.Lawyer, inkwell.Text{
				{Lang: "en", Text: "Not now."},
				{Lang: "hu", Text: "Most nem alkalmas."},
			}),
			inkwell.Say(character.Lawyer, inkwell.Text{
				{Lang: "en", Text: "I'm waiting."},
				{Lang: "hu", Text: "Várakozom."},
			}),
		),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I don't need a lawyer. Yet."},
			{Lang: "hu", Text: "Nem kell ügyvéd. Még."},
		}),
	}
}

func (policeStation) worriedRelative() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "worried_relative",
		Label: inkwell.Text{
			{Lang: "en", Text: "worried woman"},
			{Lang: "hu", Text: "aggódó nő"},
		},
		Area: inkwell.Rect(556, 160, 48, 84),
		OnLook: inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Middle-aged, wringing her coat, watching a closed door."},
				{Lang: "hu", Text: "Középkorú, a kabátját gyűrögeti, egy zárt ajtót néz."},
			}),
			inkwell.Say(character.WorriedRelative, inkwell.Text{
				{Lang: "en", Text: "What does the BVK care about some junk bought at a market…"},
				{Lang: "hu", Text: "Mit érdekli a BVK-t egy ilyen piacon vett hülyeség…"},
			}),
			inkwell.Say(character.WorriedRelative, inkwell.Text{
				{Lang: "en", Text: "It was just a small business. Not a crime…"},
				{Lang: "hu", Text: "Csak egy kis üzlet volt. Nem bűnözés…"},
			}),
		),
		OnTalk: inkwell.Seq(
			inkwell.Say(character.WorriedRelative, inkwell.Text{
				{Lang: "en", Text: "Please leave me alone."},
				{Lang: "hu", Text: "Kérem, hagyjon békén."},
			}),
			inkwell.Say(character.WorriedRelative, inkwell.Text{
				{Lang: "en", Text: "I don't want to talk about it."},
				{Lang: "hu", Text: "Nem akarok beszélni róla."},
			}),
		),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Leave her be."},
			{Lang: "hu", Text: "Hagyjuk békén."},
		}),
	}
}

func (policeStation) paperStacks() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "paper_stacks",
		Label: inkwell.Text{
			{Lang: "en", Text: "stacks of paperwork"},
			{Lang: "hu", Text: "papírkupacok"},
		},
		Area: inkwell.Rect(240, 168, 60, 50),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Folders, forms, carbon copies. Somewhere in there is Noodle, in triplicate."},
			{Lang: "hu", Text: "Dossziék, nyomtatványok, indigós másolatok. Valahol ott van Noodle is, három példányban."},
		}),
		OnUse: inkwell.Seq(
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Rifling through police files at the front desk. Bold."},
				{Lang: "hu", Text: "Rendőrségi aktákban túrni a recepción. Merész."},
			}),
			inkwell.Say(speaker.Dex, inkwell.Text{
				{Lang: "en", Text: "Bold is one word for it."},
				{Lang: "hu", Text: "A merész egy szó rá."},
			}),
		),
		OnTake: inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "Hands off the desk, sir."},
			{Lang: "hu", Text: "Kezeket le az asztalról, uram."},
		}),
	}
}

func (policeStation) monitor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "monitor",
		Label: inkwell.Text{
			{Lang: "en", Text: "green-text monitor"},
			{Lang: "hu", Text: "zöld betűs monitor"},
		},
		Area: inkwell.Rect(360, 150, 50, 44),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "A BVK-stamped terminal. Booking records, by the look of it. Facing away from me."},
			{Lang: "hu", Text: "BVK-pecsétes terminál. Előállítási nyilvántartás, úgy nézem. Elfordítva tőlem."},
		}),
		OnUse: inkwell.Say(character.Receptionist, inkwell.Text{
			{Lang: "en", Text: "That's mine, sir."},
			{Lang: "hu", Text: "Az az enyém, uram."},
		}),
	}
}

func (policeStation) Scene() *inkwell.Scene { return PoliceStation }
