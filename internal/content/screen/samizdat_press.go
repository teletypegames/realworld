package screen

import "git.teletypegames.org/games/realworld/internal/names"

// samizdatPress is a 💡 suggestion: the Resistance's physical node — parody
// BVK databases, underground publications, a P2P node humming in the cellar.
// The samizdat intellectuals are the first to write down that the world is
// slipping (wiki: entities#potenciális-további-helyszínek).
func samizdatPress() screen {
	return screen{
		name:       names.ScreenSamizdatPress,
		title:      "SAMIZDAT PRINTING HOUSE",
		background: names.BgSamizdatPress,
		onSelector: true,
	}
}
