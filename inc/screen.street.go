package inc

// street is the city itself: pavements, a square, a metro entrance, and the
// environmental decay that tracks The AI's spread — glitching systems, odd
// news, infrastructure that slips (wiki: entities#helyszinek).
//
// The wiki's own street front is the one with the hotdog stand and Sumphor on
// it (story#norman-lakása); that is a beat for this screen, not a screen of its
// own.
func screenStreet() screen {
	return screen{
		name:       ScreenStreet,
		title:      "STREET",
		background: BgStreet,
		onSelector: true,
	}
}
