package character

import (
	"realworld/inc/system/theme"
)

var Noodle = CharacterManager.Register(&Character{
	Name: "noodle",
	Label: Text{
		{Lang: "en", Text: "Noodle"},
		{Lang: "hu", Text: "Noodle"},
	},
	Lore: Text{
		{Lang: "en", Text: "Paul's friend from college, antiques out of the corner house and something shadier out of the back room. Brilliant with a machine, careless with everything else. Wrote a letter, then wasn't there to explain it."},
		{Lang: "hu", Text: "Paul egyetemi barátja, régiségek a sarki házban, valami sötétebb a hátsó szobában. Zseniális a gépekkel, gondatlan minden mással. Levelet írt, aztán nem volt ott, hogy megmagyarázza."},
	},
	Speed:       80,
	W:           40,
	H:           62,
	SpeechColor: theme.Bone,
})
