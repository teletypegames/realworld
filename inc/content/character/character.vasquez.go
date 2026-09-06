package character

import (
	"realworld/inc/system/theme"
)

var Vasquez = CharacterManager.Register(&Character{
	Name: "vasquez",
	Label: Text{
		{Lang: "en", Text: "Det. Vasquez"},
		{Lang: "hu", Text: "Vasquez nyomozó"},
	},
	Lore: Text{
		{Lang: "en", Text: "Detective, SFPD. Red hair cut short on one side, a scar across one eye, a blazer instead of a uniform. Doesn't threaten, offers: Noodle's freedom for a phone call when Paul finds something worth one."},
		{Lang: "hu", Text: "Nyomozó, SFPD. Vörös haj, az egyik oldalon rövidre vágva, egy vágás a szemén át, blézer egyenruha helyett. Nem fenyeget, ajánl: Noodle szabadságát egy telefonhívásért, ha Paul talál valamit, ami megér egyet."},
	},
	Speed:       80,
	W:           26,
	H:           68,
	SpeechColor: theme.Amber,
})
