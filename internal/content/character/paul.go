package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
	"realworld/internal/theme"
)

// paul is the player character: a junk dealer and street survivor, technically
// competent, living on the edge of the city.
func paul() inkwell.Character {
	return inkwell.Character{
		Name:  names.Paul,
		Speed: 96,
		W:     28,
		H:     68,
		Start: inkwell.Point{X: 120, Y: 232},
		// Bone white: the world and Paul. If something is amber, a tape said
		// it — see the colour rule in README.
		SpeechColor: theme.Ink,
	}
}
