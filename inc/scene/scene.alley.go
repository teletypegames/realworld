package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"git.teletypegames.org/games/realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name:       inc.SceneAlley,
		Title:      "ALLEY — BEHIND NOODLE'S HOUSE",
		Background: inc.BgAlley,
		Exits: []inkwell.Exit{
			{
				To:    inc.SceneNoodleHouse,
				Label: "back out to the street",
				Side:  inkwell.ExitLeft,
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
		OnEnter:  setVarIfEmpty(inc.VarArmillaStrip, "SRP: 1 tape"),
		Hotspots: []inkwell.Hotspot{hidingPlace(), backDoor(), bin(), fireEscape()},
	})
}

func hidingPlace() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:  "hiding_place",
		Label: "gap at the foot of the wall",
		Area:  inkwell.Rect(416, 150, 68, 52),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagHasTape),
			inkwell.Say(inc.Paul, "Empty. Whatever was in there is on me now."),
			inkwell.Say(inc.Paul, "Loose brick. There's a gap behind it."),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagPoliceTip),
			inkwell.If(inkwell.Flag(inc.FlagHasTape),
				inkwell.Say(inc.Paul, "There's nothing else in there."),
				inkwell.Seq(
					inkwell.Say(inc.Paul, "“Behind the brick.” All right then."),
					inkwell.Give(inc.ItemMysteryTape),
					inkwell.SetFlag(inc.FlagHasTape),
					inkwell.Say(inc.Paul, "A Personal Tape. No label, no seal."),
					inkwell.Say(inc.Dex, "Password-locked. And somebody very much did not want it found."),
				),
			),
			inkwell.Say(inc.Paul, "One loose brick. I'm not taking the alley apart over it."),
		),
		OnTake: inkwell.Say(inc.Paul, "I'm not carrying a wall."),
	}
}

func backDoor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "back_door",
		Label:  "locked back door",
		Area:   inkwell.Rect(72, 106, 80, 124),
		OnLook: inkwell.Say(inc.Paul, "Keypad. Four digits, worn keys."),
		OnUse: inkwell.Seq(
			inkwell.Say(inc.Paul, "Locked. And I'm not guessing four digits."),
			inkwell.Say(inc.Dex, "The wear is heaviest on the two and the seven. That isn't a code. It's fewer options."),
		),
		OnTalk: inkwell.Say(inc.Paul, "To the door? I'm not there yet."),
		OnUseWith: map[string]inkwell.Action{
			inc.ItemBlackArmilla: inkwell.Seq(
				inkwell.Say(inc.Paul, "A black-market bracelet doesn't open a keypad."),
				inkwell.Say(inc.Dex, "But you do get an advert out of it."),
			),
		},
	}
}

func bin() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "bin",
		Label:  "toppled bin",
		Area:   inkwell.Rect(240, 170, 88, 60),
		OnLook: inkwell.Say(inc.Paul, "Somebody's been through it. Thoroughly."),
		OnUse: inkwell.Seq(
			inkwell.Say(inc.Paul, "Already turned out. I'm not going to be the second one."),
			inkwell.Say(inc.Dex, "The police don't go through bins. So who did?"),
		),
	}
}

func fireEscape() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "fire_escape",
		Label:  "fire escape",
		Area:   inkwell.Rect(528, 44, 88, 160),
		OnLook: inkwell.Say(inc.Paul, "Runs all the way to the roof. Bottom rung is two metres over my head."),
		OnUse:  inkwell.Say(inc.Paul, "Can't reach it. I'd need something to stand on."),
	}
}
