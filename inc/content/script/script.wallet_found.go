package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var WalletFound = ScriptManager.Register(&Script{
	Name: "wallet_found",
	Actions: inkwell.Seq(
		character.Paul.Say(Text{
			{Lang: "en", Text: "Inside pocket. His wallet."},
			{Lang: "hu", Text: "Belső zseb. A tárcája."},
		}),
		inkwell.Give(item.NormanWallet),
		character.Paul.Say(Text{
			{Lang: "en", Text: "A stamp-it hot dog loyalty card. Eight stamps out of ten. The stand is a few blocks from here."},
			{Lang: "hu", Text: "Egy stamp-it hotdog pontgyűjtő kártya. Nyolc pecsét a tízből. A stand pár sarokra van innen."},
		}),
		speaker.Dex.Say(Text{
			{Lang: "en", Text: "A man who was two hot dogs from a free one does not leave town."},
			{Lang: "hu", Text: "Egy ember, akinek két hotdog hiányzott az ingyeneshez, nem költözik el a városból."},
		}),
	),
})
