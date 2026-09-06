package character

import (
	"realworld/inc/system/theme"
)

var IceCreamVendor = CharacterManager.Register(&Character{
	Name: "ice_cream_vendor",
	Label: Text{
		{Lang: "en", Text: "ice cream vendor"},
		{Lang: "hu", Text: "fagyiárus"},
	},
	Lore: Text{
		{Lang: "en", Text: "Sells pistachio where Paul's old bar used to be, and insists there never was a bar. Keeps a BVK-registered terminal in the corner and a regular who never gets off it."},
		{Lang: "hu", Text: "Pisztáciát árul ott, ahol Paul régi kocsmája volt, és állítja, hogy sosem volt ott kocsma. A sarokban egy BVK-bejegyzett terminál, és egy törzsvendég, aki sosem kel fel mellőle."},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.AmberPale,
})
