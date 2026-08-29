package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
	"realworld/internal/world"
)

// alley is the alley behind Noodle's house — beat 3, and the vertical slice
// the whole HUD is measured against: hotspots, a gated pickup, the two-slot
// Armilla, tape dialogue and an authored failure.
func alley(w *world.World) inkwell.Scene {
	return inkwell.Scene{
		Name:       names.SceneAlley,
		Title:      "ALLEY — BEHIND NOODLE'S HOUSE",
		Background: names.AssetAlleyBG,
		Actors: []inkwell.SceneActor{
			{CharacterName: names.Paul, At: inkwell.Point{X: 120, Y: 232}},
		},
		Walkboxes: []inkwell.Polygon{
			inkwell.Poly(
				inkwell.Point{X: 16, Y: 196}, inkwell.Point{X: 624, Y: 196},
				inkwell.Point{X: 624, Y: 258}, inkwell.Point{X: 16, Y: 258},
			),
		},
		OnEnter: inkwell.Seq(
			world.EnterScene(w, names.SceneAlley),
			setVarIfEmpty(names.VarArmillaStrip, "SRP: 1 tape"),
		),
		Hotspots: []inkwell.Hotspot{
			hidingPlace(),
			backDoor(),
			bin(),
			fireEscape(),
		},
	}
}

// hidingPlace holds the mystery tape. It only opens once Paul has the tip from
// the police station (beat 2) — without it he does not know what to look for.
func hidingPlace() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:  "hiding_place",
		Label: "gap at the foot of the wall",
		Area:  inkwell.Rect(416, 150, 68, 52),
		OnLook: inkwell.If(inkwell.Flag(names.FlagHasTape),
			inkwell.Say(names.Paul, "Empty. Whatever was in there is on me now."),
			inkwell.Say(names.Paul, "Loose brick. There's a gap behind it."),
		),
		OnUse: inkwell.If(inkwell.Flag(names.FlagPoliceTip),
			inkwell.If(inkwell.Flag(names.FlagHasTape),
				inkwell.Say(names.Paul, "There's nothing else in there."),
				inkwell.Seq(
					inkwell.Say(names.Paul, "“Behind the brick.” All right then."),
					inkwell.Give(names.ItemMysteryTape),
					inkwell.SetFlag(names.FlagHasTape),
					inkwell.Say(names.Paul, "A Personal Tape. No label, no seal."),
					world.TapeSay(names.Dex, "Password-locked. And somebody very much did not want it found."),
				),
			),
			inkwell.Say(names.Paul, "One loose brick. I'm not taking the alley apart over it."),
		),
		OnTake: inkwell.Say(names.Paul, "I'm not carrying a wall."),
	}
}

// backDoor is the New Game+ hook: the code can be entered on the first visit
// too, if the player already knows it.
func backDoor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "back_door",
		Label:  "locked back door",
		Area:   inkwell.Rect(72, 106, 80, 124),
		OnLook: inkwell.Say(names.Paul, "Keypad. Four digits, worn keys."),
		OnUse: inkwell.Seq(
			inkwell.Say(names.Paul, "Locked. And I'm not guessing four digits."),
			world.TapeSay(names.Dex, "The wear is heaviest on the two and the seven. That isn't a code. It's fewer options."),
		),
		OnTalk: inkwell.Say(names.Paul, "To the door? I'm not there yet."),
		OnUseWith: map[string]inkwell.Action{
			// An AUTHORED failure — not the engine's hardcoded flash.
			names.ItemBlackArmilla: inkwell.Seq(
				inkwell.Say(names.Paul, "A black-market bracelet doesn't open a keypad."),
				world.TapeSay(names.Dex, "But you do get an advert out of it."),
			),
		},
	}
}

func bin() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "bin",
		Label:  "toppled bin",
		Area:   inkwell.Rect(240, 170, 88, 60),
		OnLook: inkwell.Say(names.Paul, "Somebody's been through it. Thoroughly."),
		OnUse: inkwell.Seq(
			inkwell.Say(names.Paul, "Already turned out. I'm not going to be the second one."),
			world.TapeSay(names.Dex, "The police don't go through bins. So who did?"),
		),
	}
}

func fireEscape() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name:   "fire_escape",
		Label:  "fire escape",
		Area:   inkwell.Rect(528, 44, 88, 160),
		OnLook: inkwell.Say(names.Paul, "Runs all the way to the roof. Bottom rung is two metres over my head."),
		OnUse:  inkwell.Say(names.Paul, "Can't reach it. I'd need something to stand on."),
	}
}
