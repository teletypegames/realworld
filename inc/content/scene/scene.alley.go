package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
	"realworld/inc/system/armilla"
)

type alley struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneAlley,
	Title: inkwell.Text{
		{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
		{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
	},
	Exits: []inkwell.Exit{
		{
			To: inc.SceneNoodleHouse,
			Label: inkwell.Text{
				{Lang: "en", Text: "back out to the street"},
				{Lang: "hu", Text: "vissza az utcára"},
			},
			Side: inkwell.ExitLeft,
		},
		{
			To: inc.SceneNoodleSecretRoom,
			Label: inkwell.Text{
				{Lang: "en", Text: "the back door"},
				{Lang: "hu", Text: "a hátsó ajtó"},
			},
			Side:  inkwell.ExitBack,
			Needs: inc.FlagSecretRoomOpen,
			Blocked: inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Locked. Four digits between me and whatever's behind it."},
				{Lang: "hu", Text: "Zárva. Négy számjegy köztem és aközött, ami mögötte van."},
			}),
		},
	},
	Triggers: []inkwell.Trigger{
		{
			Name: "arrival",
			When: inkwell.Given{
				Flag:    inc.FlagPoliceTip,
				NotFlag: inc.FlagHasTape,
			},
			Do: inkwell.Seq(
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Back alley. Just like the clerk said."},
					{Lang: "hu", Text: "Hátsó sikátor. Pont, ahogy az ügyintéző mondta."},
				}),
				inkwell.Say(inc.Dex, inkwell.Text{
					{Lang: "en", Text: "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"},
					{Lang: "hu", Text: "Egy állami hivatal súgta meg, hogy vigyél el valamit egy helyszínről. Ez neked nem furcsa?"},
				}),
			),
			Once: true,
		},
	},
	Actors: []inkwell.SceneActor{
		{
			CharacterName: inc.Paul,
			At: inkwell.Point{
				X: 120,
				Y: 232,
			},
		},
	},
	OnEnter: armilla.StripIfEmpty(inkwell.Text{
		{Lang: "en", Text: "SRP: 1 tape"},
		{Lang: "hu", Text: "SRP: 1 kazetta"},
	}),
	Hotspots: alley{}.hotspots(),
})

func (place alley) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.hidingPlace(),
		place.backDoor(),
		place.bin(),
		place.fireEscape(),
	}
}

func (alley) hidingPlace() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "hiding_place",
		Label: inkwell.Text{
			{Lang: "en", Text: "gap at the foot of the wall"},
			{Lang: "hu", Text: "rés a fal tövében"},
		},
		Area: inkwell.Rect(416, 150, 68, 52),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagHasTape),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Empty. Whatever was in there is on me now."},
				{Lang: "hu", Text: "Üres. Ami benne volt, most nálam van."},
			}),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Loose brick. There's a gap behind it."},
				{Lang: "hu", Text: "Laza tégla. Rés van mögötte."},
			}),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagPoliceTip),
			inkwell.If(inkwell.Flag(inc.FlagHasTape),
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "There's nothing else in there."},
					{Lang: "hu", Text: "Nincs benne más."},
				}),
				inkwell.Seq(
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "“Behind the brick.” All right then."},
						{Lang: "hu", Text: "„A tégla mögött.” Hát jó."},
					}),
					inkwell.Give(inc.ItemMysteryTape),
					inkwell.SetFlag(inc.FlagHasTape),
					inkwell.Say(inc.Paul, inkwell.Text{
						{Lang: "en", Text: "A Personal Tape. No label, no seal."},
						{Lang: "hu", Text: "Egy Személyes Kazetta. Se címke, se pecsét."},
					}),
					inkwell.Say(inc.Dex, inkwell.Text{
						{Lang: "en", Text: "Password-locked. And somebody very much did not want it found."},
						{Lang: "hu", Text: "Jelszóval zárva. És valaki nagyon nem akarta, hogy megtalálják."},
					}),
				),
			),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "One loose brick. I'm not taking the alley apart over it."},
				{Lang: "hu", Text: "Egy laza tégla. Nem szedem szét miatta a sikátort."},
			}),
		),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not carrying a wall."},
			{Lang: "hu", Text: "Nem cipelek el egy falat."},
		}),
	}
}

func (alley) backDoor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "back_door",
		Label: inkwell.Text{
			{Lang: "en", Text: "locked back door"},
			{Lang: "hu", Text: "zárt hátsó ajtó"},
		},
		Area: inkwell.Rect(72, 106, 80, 124),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Keypad. Four digits, worn keys."},
				{Lang: "hu", Text: "Számbillentyűzet. Négy számjegy, kopott gombok."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "The wear is heaviest on the two and the seven. That isn't a code. It's fewer options."},
				{Lang: "hu", Text: "A kettes és a hetes kopott a legjobban. Az nem a kód. Az kevesebb lehetőség."},
			}),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagSecretRoomOpen),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "It's open. I don't need to type it twice."},
				{Lang: "hu", Text: "Nyitva van. Nem kell kétszer beütnöm."},
			}),
			system.System.Paused(inkwell.RunDialog(inc.DialogKeypad)),
		),
		OnTalk: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "To the door? I'm not there yet."},
			{Lang: "hu", Text: "Az ajtóhoz? Még nem tartok ott."},
		}),
		OnUseWith: map[string]inkwell.Action{
			inc.ItemBlackArmilla: inkwell.Seq(
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "A black-market bracelet doesn't open a keypad."},
					{Lang: "hu", Text: "Egy feketepiaci karperec nem nyit ki egy számzárat."},
				}),
				inkwell.Say(inc.Dex, inkwell.Text{
					{Lang: "en", Text: "But you do get an advert out of it."},
					{Lang: "hu", Text: "Viszont kapsz hozzá egy reklámot."},
				}),
			),
		},
	}
}

func (alley) bin() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "bin",
		Label: inkwell.Text{
			{Lang: "en", Text: "toppled bin"},
			{Lang: "hu", Text: "feldöntött kuka"},
		},
		Area: inkwell.Rect(240, 170, 88, 60),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Somebody's been through it. Thoroughly."},
			{Lang: "hu", Text: "Valaki átkutatta. Alaposan."},
		}),
		OnUse: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Already turned out. I'm not going to be the second one."},
				{Lang: "hu", Text: "Már kiforgatták. Nem leszek a második."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "The police don't go through bins. So who did?"},
				{Lang: "hu", Text: "A rendőrség nem túr kukákban. Akkor ki?"},
			}),
		),
	}
}

func (alley) fireEscape() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "fire_escape",
		Label: inkwell.Text{
			{Lang: "en", Text: "fire escape"},
			{Lang: "hu", Text: "tűzlétra"},
		},
		Area: inkwell.Rect(528, 44, 88, 160),
		OnLook: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Runs all the way to the roof. Bottom rung is two metres over my head."},
			{Lang: "hu", Text: "Felvezet egészen a tetőig. A legalsó fok két méterrel a fejem fölött van."},
		}),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Can't reach it. I'd need something to stand on."},
			{Lang: "hu", Text: "Nem érem el. Kéne valami, amire felállhatok."},
		}),
	}
}
