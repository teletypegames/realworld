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
	FlagClubEntry = "club_entry"       // wiki: klub_bejutas — the shopkeeper let Paul through

	VarSlot2        = "armilla.slot2"
	VarArmillaStrip = "armilla.strip"
	VarDecay        = "decay_level" // wiki: romlas_szint
)

// Screens.
//
// One constant per screen in the wiki's location catalogue
// (entities#helyszinek), whose order the concept-art deck follows number for
// number (games/realworld_screen_assets); the trailing comment is the wiki's
// Hungarian name. The last seven are the catalogue's "javaslat" locations —
// drawn, so they are wired up, but no beat lands on them yet.
//
// Two screens have no deck number. The alley shares the catalogue's "Noodle
// háza & a sikátor" row with the house; the selector is the wiki's
// Helyszínválasztó (story#helyszínválasztó), the menu screen the location
// graph is centred on.
const (
	ScreenSelector        = "selector"         // wiki: Helyszínválasztó
	ScreenPaulShop        = "paul_shop"        // wiki: Paul boltja / garázsa
	ScreenNoodleHouse     = "noodle_house"     // wiki: Noodle háza
	ScreenNormanApartment = "norman_apartment" // wiki: Norman lakása / műhelye
	ScreenPoliceStation   = "police_station"   // wiki: rendőrség / fogda
	ScreenAlley           = "alley"            // wiki: sikátor (Noodle háza mögött)
	ScreenHackerspace     = "hackerspace"      // wiki: hackerspace
	ScreenIceCreamShop    = "ice_cream_shop"   // wiki: fagyizó
	ScreenTrinketShop     = "trinket_shop"     // wiki: chinatowni csecsebecse-bolt
	ScreenSmallRestaurant = "small_restaurant" // wiki: kajálda
	ScreenSecretClub      = "secret_club"      // wiki: titkos klub („vízalatti nap")
	ScreenBBSTerminal     = "bbs_terminal"     // wiki: a terminál / BBS-hozzáférési pont
	ScreenStreet          = "street"           // wiki: köztér / utca
	ScreenHospital        = "hospital"         // wiki: kórház / pszichiátria
	ScreenServerFarm      = "server_farm"      // wiki: szerverfarm / adatközpont
	ScreenSecretLab       = "secret_lab"       // wiki: titkos kutatólabor
	ScreenRooftopHideout  = "rooftop_hideout"  // wiki: Norman tetőterasz-rejteke
	ScreenPublicBBS       = "public_bbs"       // wiki 💡: nyilvános BBS-terminál
	ScreenBVKBranch       = "bvk_branch"       // wiki 💡: BVK-kirendeltség (San Francisco)
	ScreenCuratorShop     = "curator_shop"     // wiki 💡: kurátor-szerviz („a Mester")
	ScreenScrapMarket     = "scrap_market"     // wiki 💡: selejt-piac
	ScreenSamizdatPress   = "samizdat_press"   // wiki 💡: samizdat-nyomda / P2P-csomópont
	ScreenShowroom        = "showroom"         // wiki 💡: Neumatronic márkakereskedés
	ScreenColumbarium     = "columbarium"      // wiki 💡: kolumbárium / Dex emlékhelye
)

// Backgrounds — one per screen, same order.
const (
	BgSelector        = "bg_selector"
	BgPaulShop        = "bg_paul_shop"
	BgNoodleHouse     = "bg_noodle_house"
	BgNormanApartment = "bg_norman_apartment"
	BgPoliceStation   = "bg_police_station"
	BgAlley           = "bg_alley"
	BgHackerspace     = "bg_hackerspace"
	BgIceCreamShop    = "bg_ice_cream_shop"
	BgTrinketShop     = "bg_trinket_shop"
	BgSmallRestaurant = "bg_small_restaurant"
	BgSecretClub      = "bg_secret_club"
	BgBBSTerminal     = "bg_bbs_terminal"
	BgStreet          = "bg_street"
	BgHospital        = "bg_hospital"
	BgServerFarm      = "bg_server_farm"
	BgSecretLab       = "bg_secret_lab"
	BgRooftopHideout  = "bg_rooftop_hideout"
	BgPublicBBS       = "bg_public_bbs"
	BgBVKBranch       = "bg_bvk_branch"
	BgCuratorShop     = "bg_curator_shop"
	BgScrapMarket     = "bg_scrap_market"
	BgSamizdatPress   = "bg_samizdat_press"
	BgShowroom        = "bg_showroom"
	BgColumbarium     = "bg_columbarium"
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
