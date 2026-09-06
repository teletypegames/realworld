package character

import (
	"realworld/inc/system/theme"
)

var DoctorMaren = CharacterManager.Register(&Character{
	Name: "doctor_maren",
	Label: Text{
		{Lang: "en", Text: "Dr. Maren"},
		{Lang: "hu", Text: "Dr. Maren"},
	},
	Lore: Text{
		{Lang: "en", Text: "Psychiatric wing, white coat, no badge. Had Norman for eleven hours before they collected him: no cause, clean bloodwork, an EEG like deep REM that never ended. Was handed the official diagnosis already typed."},
		{Lang: "hu", Text: "Pszichiátriai osztály, fehér köpeny, kitűző nélkül. Tizenegy órán át volt nála Norman, mielőtt elvitték: ok nélkül, tiszta vérképpel, egy EEG-vel, mint a mély REM, ami sosem ért véget. A hivatalos diagnózist gépelve kapta kézhez."},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Bone,
})
