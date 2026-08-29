package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

// supportTape answers everything with basic information, insufferably. Dex
// hates it. Acquired in Chinatown (beat 5) — not reachable yet.
func characterSupportTape() inkwell.Character {
	return inkwell.Character{
		Name:        TapeSupport,
		SpeechColor: RGB(0xE8C86A),
	}
}
