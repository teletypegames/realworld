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
	ScreenSelector        = "selector"
	ScreenPaulShop        = "paul_shop"
	ScreenNoodleHouse     = "noodle_house"
	ScreenNormanApartment = "norman_apartment"
	ScreenPoliceStation   = "police_station"
	ScreenAlley           = "alley"
	ScreenHackerspace     = "hackerspace"
	ScreenIceCreamShop    = "ice_cream_shop"
	ScreenTrinketShop     = "trinket_shop"
	ScreenSmallRestaurant = "small_restaurant"
	ScreenSecretClub      = "secret_club"
	ScreenBBSTerminal     = "bbs_terminal"
	ScreenStreet          = "street"
	ScreenHospital        = "hospital"
	ScreenServerFarm      = "server_farm"
	ScreenSecretLab       = "secret_lab"
	ScreenRooftopHideout  = "rooftop_hideout"
	ScreenPublicBBS       = "public_bbs"
	ScreenBVKBranch       = "bvk_branch"
	ScreenCuratorShop     = "curator_shop"
	ScreenScrapMarket     = "scrap_market"
	ScreenSamizdatPress   = "samizdat_press"
	ScreenShowroom        = "showroom"
	ScreenColumbarium     = "columbarium"
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
