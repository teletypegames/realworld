package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

func screenAlley() screen {
	return screen{
		name:       ScreenAlley,
		title:      "ALLEY — BEHIND NOODLE'S HOUSE",
		background: BgAlley,
		exits: []exit{
			{
				to:    ScreenNoodleHouse,
				label: "back out to the street",
				side:  sideLeft,
			},
		},
		actors: []inkwell.SceneActor{
			{
				CharacterName: Paul,
				At: inkwell.Point{
					X: 120,
					Y: 232,
				},
			},
		},
		onEnter:  setVarIfEmpty(VarArmillaStrip, "SRP: 1 tape"),
		hotspots: []inkwell.Hotspot{hidingPlace(), backDoor(), bin(), fireEscape()},
	}
}

func hidingPlace() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:  "hiding_place",
		Label: "gap at the foot of the wall",
		Area:  inkwell.Rect(416, 150, 68, 52),
		OnLook: inkwell.If(inkwell.Flag(FlagHasTape),
			inkwell.Say(Paul, "Empty. Whatever was in there is on me now."),
			inkwell.Say(Paul, "Loose brick. There's a gap behind it."),
		),
		OnUse: inkwell.If(inkwell.Flag(FlagPoliceTip),
			inkwell.If(inkwell.Flag(FlagHasTape),
				inkwell.Say(Paul, "There's nothing else in there."),
				inkwell.Seq(
					inkwell.Say(Paul, "“Behind the brick.” All right then."),
					inkwell.Give(ItemMysteryTape),
					inkwell.SetFlag(FlagHasTape),
					inkwell.Say(Paul, "A Personal Tape. No label, no seal."),
					TapeSay(Dex, "Password-locked. And somebody very much did not want it found."),
				),
			),
			inkwell.Say(Paul, "One loose brick. I'm not taking the alley apart over it."),
		),
		OnTake: inkwell.Say(Paul, "I'm not carrying a wall."),
	}
}

func backDoor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "back_door",
		Label:  "locked back door",
		Area:   inkwell.Rect(72, 106, 80, 124),
		OnLook: inkwell.Say(Paul, "Keypad. Four digits, worn keys."),
		OnUse: inkwell.Seq(
			inkwell.Say(Paul, "Locked. And I'm not guessing four digits."),
			TapeSay(Dex, "The wear is heaviest on the two and the seven. That isn't a code. It's fewer options."),
		),
		OnTalk: inkwell.Say(Paul, "To the door? I'm not there yet."),
		OnUseWith: map[string]inkwell.Action{
			ItemBlackArmilla: inkwell.Seq(
				inkwell.Say(Paul, "A black-market bracelet doesn't open a keypad."),
				TapeSay(Dex, "But you do get an advert out of it."),
			),
		},
	}
}

func bin() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "bin",
		Label:  "toppled bin",
		Area:   inkwell.Rect(240, 170, 88, 60),
		OnLook: inkwell.Say(Paul, "Somebody's been through it. Thoroughly."),
		OnUse: inkwell.Seq(
			inkwell.Say(Paul, "Already turned out. I'm not going to be the second one."),
			TapeSay(Dex, "The police don't go through bins. So who did?"),
		),
	}
}

func fireEscape() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "fire_escape",
		Label:  "fire escape",
		Area:   inkwell.Rect(528, 44, 88, 160),
		OnLook: inkwell.Say(Paul, "Runs all the way to the roof. Bottom rung is two metres over my head."),
		OnUse:  inkwell.Say(Paul, "Can't reach it. I'd need something to stand on."),
	}
}
