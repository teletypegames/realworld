// Package names is the single source of entity names and world-state keys.
//
// The names follow the wiki entity catalogue
// (services/wiki-pages/pages/projects/realworld/entities) so that content and
// design share one vocabulary. Note that the wiki's world-state table spells
// its keys in Hungarian (hatosagi_tipp, titokzatos_tape_megvan, romlas_szint,
// tape_behelyezes, sikator); the constants below are their English
// equivalents — the mapping is one-to-one, in catalogue order.
//
// This package imports nothing from the project, so every layer can depend on
// it without creating a cycle.
package names

// Characters.
const (
	Paul = "paul"
	Dex  = "dex" // Personal Tape, permanently in slot 1 of Paul's Armilla

	// Tapes (Personal Tapes) are characters, not props: the wiki is explicit
	// that they have their own personality, motivation and voice.
	TapeMystery = "tape_mystery" // = The AI, but that only surfaces in the finale
	TapeSupport = "tape_support" // the "Indian tech support" tape
)

// Inventory items.
const (
	ItemNoodleLetter = "noodle_letter"
	ItemMysteryTape  = "mystery_tape"
	ItemBlackArmilla = "black_market_armilla"
)

// Dialogues.
const (
	DlgDex         = "dex_talk"
	DlgMysteryTape = "mystery_tape_silent"
)

// Scripts.
const (
	ScriptTapeInsert = "tape_insert"
	ScriptFinale     = "awakening_finale"
)

// World state — flags and variables.
const (
	FlagPoliceTip = "police_tip"       // wiki: hatosagi_tipp
	FlagHasTape   = "has_mystery_tape" // wiki: titokzatos_tape_megvan
	FlagTapeSpoke = "tape_spoke"       // wiki: tape_megszolalt

	VarSlot2        = "armilla.slot2"
	VarArmillaStrip = "armilla.strip"
	VarDecay        = "decay_level" // wiki: romlas_szint
)

// Scenes.
const (
	SceneAlley = "alley" // wiki: sikator
)

// Assets.
const (
	AssetAlleyBG = "bg_alley"
)

// Tapes maps a character name to the label shown in the tape channel header.
// TapeChannel uses this to decide whether a line belongs in the tape channel
// or stays in the scene.
var Tapes = map[string]string{
	Dex:         "DEX",
	TapeMystery: "UNKNOWN TAPE",
	TapeSupport: "TECH SUPPORT",
}

// IsTape reports whether a character is a tape.
func IsTape(name string) bool { _, ok := Tapes[name]; return ok }

// TapeDisplayName is the label shown in the tape channel header.
func TapeDisplayName(name string) string {
	if d, ok := Tapes[name]; ok {
		return d
	}
	return name
}

// TapeItems maps an inventory item that can go into slot 2 to the character
// living on it.
var TapeItems = map[string]string{
	ItemMysteryTape: TapeMystery,
}

// IsTapeItem reports whether an item can be loaded into slot 2.
func IsTapeItem(item string) bool { _, ok := TapeItems[item]; return ok }

// TapeDialogue is the dialogue belonging to a tape in slot 2.
func TapeDialogue(item string) string {
	switch item {
	case ItemMysteryTape:
		return DlgMysteryTape
	}
	return DlgDex
}
