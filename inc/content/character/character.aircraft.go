package character

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

var aircraftSprite = &inkwell.Sprite{
	Sheet: &inkwell.Asset{
		Name: "aircraft",
		Path: "assets/sprite/aircraft.png",
	},
	Faces: inkwell.FacingRight,
}

var Aircraft = CharacterManager.Register(&Character{
	Name: "aircraft",
	Label: Text{
		{Lang: "en", Text: "aircraft"},
		{Lang: "hu", Text: "repülőgép"},
	},
	Poses: map[inkwell.Pose]*inkwell.Sprite{
		inkwell.PoseIdle: aircraftSprite,
		inkwell.PoseWalk: aircraftSprite,
	},
	Speed: 180,
	W:     150,
	H:     85,
})
