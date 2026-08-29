package inc

// hospital is the clinical dead end: Dr. Maren looked for a medical reason for
// Norman's state and did not find one (wiki: entities, timeline 21).
func screenHospital() screen {
	return screen{
		name:       ScreenHospital,
		title:      "HOSPITAL — PSYCHIATRIC WING",
		background: BgHospital,
		onSelector: true,
	}
}
