package inc

// secretClub is the empty room behind the shop where the "underwater sun"
// company gave their talk, with a secretary who knows about the BVK's plans for
// people's minds and nothing about the key or the tape
// (wiki: story, the club beat).
//
// The terminal in here is the one that matters: not under BVK control, ECHO
// running under the standard LECTOR surface. Reading it is where the mystery
// tape finally speaks.
func screenSecretClub() screen {
	return screen{
		name:       ScreenSecretClub,
		title:      "SECRET CLUB — THE UNDERWATER SUN",
		background: BgSecretClub,
		exits: []exit{
			{to: ScreenBBSTerminal, label: "the terminal in the corner", side: sideBack},
			{to: ScreenTrinketShop, label: "back out through the shop", side: sideNear},
		},
	}
}
