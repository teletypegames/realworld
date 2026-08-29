package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// paul is the player character: a junk dealer and street survivor, technically
// competent, living on the edge of the city.
func characterPaul() inkwell.Character {
	return inkwell.Character{
		Name:  Paul,
		Speed: 96,
		W:     28,
		H:     68,
		Start: inkwell.Point{X: 120, Y: 232},
		// Bone white: the world and Paul. If something is amber, a tape said
		// it — see the colour rule in README.
		SpeechColor: Ink,
	}
}
