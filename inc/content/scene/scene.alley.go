package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
	"realworld/inc/system/armilla"
)

var mysteryTapeTaken = character.Paul.Say(Text{
	{Lang: "en", Text: "That’s enough trash for one day."},
	{Lang: "hu", Text: "Ennyi szemét elég volt mára."},
})

var mysteryTapePickUp = inkwell.Seq(
	inkwell.Give(item.MysteryTape),
	character.Paul.Say(Text{
		{Lang: "en", Text: "A tape! This must be Noodle’s."},
		{Lang: "hu", Text: "Egy kazetta! Ez biztos Noodle-é."},
	}),
)

var Alley = SceneManager.Register(&Scene{
	Name: "alley",
	Title: Text{
		{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
		{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "alley",
				Path: "assets/bg/v2_alley.jpeg",
			},
			Z: -1,
		},
	},
	Triggers: []Trigger{
		{
			Name: "arrival",
			When: Given{
				Flag:    inc.FlagPoliceTip,
				NotFlag: inc.FlagHasTape,
			},
			Do: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Back alley. Just like the clerk said."},
					{Lang: "hu", Text: "Hátsó sikátor. Pont, ahogy az ügyintéző mondta."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "A government office tipped you off to remove something from a scene. That doesn't strike you as odd?"},
					{Lang: "hu", Text: "Egy állami hivatal súgta meg, hogy vigyél el valamit egy helyszínről. Ez neked nem furcsa?"},
				}),
			),
			Once: true,
		},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 120,
				Y: 232,
			},
		},
	},
	OnEnter: armilla.StripIfEmpty(Text{
		{Lang: "en", Text: "SRP: 1 tape"},
		{Lang: "hu", Text: "SRP: 1 kazetta"},
	}),
	Hotspots: []Hotspot{
		{
			Name: "hiding_place",
			Label: Text{
				{Lang: "en", Text: "gap at the foot of the wall"},
				{Lang: "hu", Text: "rés a fal tövében"},
			},
			Area: inkwell.Rect(416, 150, 68, 52),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagHasTape),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Empty. Whatever was in there is on me now."},
					{Lang: "hu", Text: "Üres. Ami benne volt, most nálam van."},
				}),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Loose brick. There's a gap behind it."},
					{Lang: "hu", Text: "Laza tégla. Rés van mögötte."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagPoliceTip),
				inkwell.If(inkwell.Flag(inc.FlagHasTape),
					character.Paul.Say(Text{
						{Lang: "en", Text: "There's nothing else in there."},
						{Lang: "hu", Text: "Nincs benne más."},
					}),
					inkwell.Seq(
						character.Paul.Say(Text{
							{Lang: "en", Text: "“Behind the brick.” All right then."},
							{Lang: "hu", Text: "„A tégla mögött.” Hát jó."},
						}),
						inkwell.Give(item.MysteryTape),
						inkwell.SetFlag(inc.FlagHasTape),
						character.Paul.Say(Text{
							{Lang: "en", Text: "A Personal Tape. No label, no seal."},
							{Lang: "hu", Text: "Egy Személyes Kazetta. Se címke, se pecsét."},
						}),
						speaker.Dex.Say(Text{
							{Lang: "en", Text: "Password-locked. And somebody very much did not want it found."},
							{Lang: "hu", Text: "Jelszóval zárva. És valaki nagyon nem akarta, hogy megtalálják."},
						}),
					),
				),
				character.Paul.Say(Text{
					{Lang: "en", Text: "One loose brick. I'm not taking the alley apart over it."},
					{Lang: "hu", Text: "Egy laza tégla. Nem szedem szét miatta a sikátort."},
				}),
			),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "I'm not carrying a wall."},
				{Lang: "hu", Text: "Nem cipelek el egy falat."},
			}),
		},
		{
			Name: "back_door",
			Label: Text{
				{Lang: "en", Text: "locked back door"},
				{Lang: "hu", Text: "zárt hátsó ajtó"},
			},
			Area: inkwell.Rect(72, 106, 80, 124),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Keypad. Four digits, worn keys."},
					{Lang: "hu", Text: "Számbillentyűzet. Négy számjegy, kopott gombok."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "The wear is heaviest on the two and the seven. That isn't a code. It's fewer options."},
					{Lang: "hu", Text: "A kettes és a hetes kopott a legjobban. Az nem a kód. Az kevesebb lehetőség."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagSecretRoomOpen),
				character.Paul.Say(Text{
					{Lang: "en", Text: "It's open. I don't need to type it twice."},
					{Lang: "hu", Text: "Nyitva van. Nem kell kétszer beütnöm."},
				}),
				system.System.Paused(inkwell.RunDialog(dialog.Keypad)),
			),
			OnTalk: character.Paul.Say(Text{
				{Lang: "en", Text: "To the door? I'm not there yet."},
				{Lang: "hu", Text: "Az ajtóhoz? Még nem tartok ott."},
			}),
			OnUseWith: map[*inkwell.Item]Action{
				item.BlackArmilla: inkwell.Seq(
					character.Paul.Say(Text{
						{Lang: "en", Text: "A black-market bracelet doesn't open a keypad."},
						{Lang: "hu", Text: "Egy feketepiaci karperec nem nyit ki egy számzárat."},
					}),
					speaker.Dex.Say(Text{
						{Lang: "en", Text: "But you do get an advert out of it."},
						{Lang: "hu", Text: "Viszont kapsz hozzá egy reklámot."},
					}),
				),
			},
		},
		{
			Name: "bin",
			Label: Text{
				{Lang: "en", Text: "toppled bin"},
				{Lang: "hu", Text: "feldöntött kuka"},
			},
			Area: inkwell.Rect(240, 170, 88, 60),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Somebody's been through it. Thoroughly."},
				{Lang: "hu", Text: "Valaki átkutatta. Alaposan."},
			}),
			OnUse: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Already turned out. I'm not going to be the second one."},
					{Lang: "hu", Text: "Már kiforgatták. Nem leszek a második."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "The police don't go through bins. So who did?"},
					{Lang: "hu", Text: "A rendőrség nem túr kukákban. Akkor ki?"},
				}),
			),
			OnTake: inkwell.If(inkwell.HasItem(item.MysteryTape), mysteryTapeTaken, mysteryTapePickUp),
		},
		{
			Name: "fire_escape",
			Label: Text{
				{Lang: "en", Text: "fire escape"},
				{Lang: "hu", Text: "tűzlétra"},
			},
			Area: inkwell.Rect(528, 44, 88, 160),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Runs all the way to the roof. Bottom rung is two metres over my head."},
				{Lang: "hu", Text: "Felvezet egészen a tetőig. A legalsó fok két méterrel a fejem fölött van."},
			}),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "Can't reach it. I'd need something to stand on."},
				{Lang: "hu", Text: "Nem érem el. Kéne valami, amire felállhatok."},
			}),
		},
	},
})

var _ = Alley.Leads(
	Exit{
		To: NoodleHouse,
		Label: Text{
			{Lang: "en", Text: "back out to the street"},
			{Lang: "hu", Text: "vissza az utcára"},
		},
		Side: inkwell.ExitLeft,
	},
)
