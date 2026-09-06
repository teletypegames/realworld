package character

import (
	"realworld/inc/system/theme"
)

var Secretary = CharacterManager.Register(&Character{
	Name: "secretary",
	Label: Text{
		{Lang: "en", Text: "secretary"},
		{Lang: "hu", Text: "titkár"},
	},
	Lore: Text{
		{Lang: "en", Text: "The only one left in the club of the underwater sun. Knows what the last lecture said about the BVK and what it wants to do with people's heads. Knows nothing about keys or tapes, and says so."},
		{Lang: "hu", Text: "Az egyetlen, aki maradt a víz alatti nap klubjában. Tudja, mit mondott az utolsó előadás a BVK-ról, és arról, mit akar az emberek fejével. Kulcsokról és kazettákról semmit nem tud, és ezt meg is mondja."},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Bone,
})
