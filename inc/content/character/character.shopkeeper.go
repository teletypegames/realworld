package character

import (
	"realworld/inc/system/theme"
)

var Shopkeeper = CharacterManager.Register(&Character{
	Name: "shopkeeper",
	Label: Text{
		{Lang: "en", Text: "shopkeeper"},
		{Lang: "hu", Text: "boltos"},
	},
	Lore: Text{
		{Lang: "en", Text: "Chinatown trinket shop, at the address where a club used to give lectures. Sells brass things with no purpose and Personal Tapes with too much. Knows the way to the club, for the right customer."},
		{Lang: "hu", Text: "Csecsebecse-bolt a kínai negyedben, azon a címen, ahol egy klub előadásokat tartott. Céltalan rézvacakokat árul és túl sok célú Personal Tape-eket. Tudja az utat a klubba, a megfelelő vevőnek."},
	},
	Speed:       80,
	W:           26,
	H:           62,
	SpeechColor: theme.Amber,
})
