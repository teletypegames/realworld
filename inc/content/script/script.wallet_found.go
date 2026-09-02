package script

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

var _ = ScriptManager.Register(Script{
	Name: inc.ScriptWalletFound,
	Actions: inkwell.Seq(
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Inside pocket. His wallet."},
			{Lang: "hu", Text: "Belső zseb. A tárcája."},
		}),
		inkwell.Give(inc.ItemNormanWallet),
		inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "A stamp-it hot dog loyalty card. Eight stamps out of ten. The stand is a few blocks from here."},
			{Lang: "hu", Text: "Egy stamp-it hotdog pontgyűjtő kártya. Nyolc pecsét a tízből. A stand pár sarokra van innen."},
		}),
		inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "A man who was two hot dogs from a free one does not leave town."},
			{Lang: "hu", Text: "Egy ember, akinek két hotdog hiányzott az ingyeneshez, nem költözik el a városból."},
		}),
	),
})
