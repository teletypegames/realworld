package inc

// serverFarm is the cooled, humming corridors where The AI's physical
// infrastructure can be seen for what it is. Getting in needs the access tape
// from the hackerspace or from Noodle (wiki: entities, timeline 24).
func screenServerFarm() screen {
	return screen{
		name:       ScreenServerFarm,
		title:      "SERVER FARM",
		background: BgServerFarm,
		onSelector: true,
	}
}
