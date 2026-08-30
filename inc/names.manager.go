package inc

const (
	Paul = "paul"
	Dex  = "dex"

	TapeMystery = "tape_mystery"
	TapeSupport = "tape_support"
)

const (
	ItemNoodleLetter = "noodle_letter"
	ItemMysteryTape  = "mystery_tape"
	ItemBlackArmilla = "black_market_armilla"
)

const (
	DlgDex         = "dex_talk"
	DlgMysteryTape = "mystery_tape_silent"
)

const (
	ScriptOpening    = "opening"
	ScriptTapeInsert = "tape_insert"
	ScriptFinale     = "awakening_finale"
)

const (
	FlagPoliceTip = "police_tip"
	FlagHasTape   = "has_mystery_tape"
	FlagTapeSpoke = "tape_spoke"
	FlagClubEntry = "club_entry"

	VarSlot2        = "armilla.slot2"
	VarArmillaStrip = "armilla.strip"
	VarDecay        = "decay_level"
)

const (
	SceneSelector        = "selector"
	ScenePaulShop        = "paul_shop"
	SceneNoodleHouse     = "noodle_house"
	SceneNormanApartment = "norman_apartment"
	ScenePoliceStation   = "police_station"
	SceneAlley           = "alley"
	SceneHackerspace     = "hackerspace"
	SceneIceCreamShop    = "ice_cream_shop"
	SceneTrinketShop     = "trinket_shop"
	SceneSmallRestaurant = "small_restaurant"
	SceneSecretClub      = "secret_club"
	SceneBBSTerminal     = "bbs_terminal"
	SceneStreet          = "street"
	SceneHospital        = "hospital"
	SceneServerFarm      = "server_farm"
	SceneSecretLab       = "secret_lab"
	SceneRooftopHideout  = "rooftop_hideout"
	ScenePublicBBS       = "public_bbs"
	SceneBVKBranch       = "bvk_branch"
	SceneCuratorShop     = "curator_shop"
	SceneScrapMarket     = "scrap_market"
	SceneSamizdatPress   = "samizdat_press"
	SceneShowroom        = "showroom"
	SceneColumbarium     = "columbarium"
)

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

var Tapes = map[string]string{
	Dex:         "DEX",
	TapeMystery: "UNKNOWN TAPE",
	TapeSupport: "TECH SUPPORT",
}

func IsTape(name string) bool { _, ok := Tapes[name]; return ok }

func TapeDisplayName(name string) string {
	if d, ok := Tapes[name]; ok {
		return d
	}
	return name
}

var TapeItems = map[string]string{
	ItemMysteryTape: TapeMystery,
}

func IsTapeItem(item string) bool { _, ok := TapeItems[item]; return ok }

func TapeDialogue(item string) string {
	switch item {
	case ItemMysteryTape:
		return DlgMysteryTape
	}
	return DlgDex
}
