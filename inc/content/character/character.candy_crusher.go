package character

import (
	"realworld/inc/system/theme"
)

var CandyCrusher = CharacterManager.Register(&Character{
	Name: "candy_crusher",
	Label: Text{
		{Lang: "en", Text: "man with the handheld"},
		{Lang: "hu", Text: "kézikonzolos fickó"},
	},
	Lore: Text{
		{Lang: "en", Text: "The man with the handheld at the terminal in the ice cream shop. A believer: the BVK's neural interface will put the human mind above the ACPs, and he'll tell you so. Lets you search if you buy him a cone."},
		{Lang: "hu", Text: "A kézikonzolos ember a fagyizó termináljánál. Hívő: a BVK neurális interfésze az emberi elmét az ACP-k fölé emeli, és ezt el is mondja. Enged keresni, ha veszel neki egy tölcsért."},
	},
	Speed:       80,
	W:           30,
	H:           60,
	SpeechColor: theme.Bone,
})
