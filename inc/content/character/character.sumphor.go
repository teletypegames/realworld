package character

import (
	"realworld/inc/system/theme"
)

var Sumphor = CharacterManager.Register(&Character{
	Name: "sumphor",
	Label: Text{
		{Lang: "en", Text: "Sumphor"},
		{Lang: "hu", Text: "Sumphor"},
	},
	Lore: Text{
		{Lang: "en", Text: "Sits on the pavement by the hot dog stand, talking to himself about the Project. Worked with Norman on the prototype; the chair in the prototype room is what did this to him. The tape says he owned it once. He says the tape lies."},
		{Lang: "hu", Text: "A járdán ül a hotdogos mellett, magában beszél a Projektről. Normannal dolgozott a prototípuson; a prototípus-szoba széke tette ezt vele. A kazetta szerint egyszer az övé volt. Szerinte a kazetta hazudik."},
	},
	Speed:       80,
	W:           26,
	H:           66,
	SpeechColor: theme.Amber,
})
