package background

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc/content/scene"
)

type Background = inkwell.Asset

var BackgroundManager = inkwell.NewManager[*Background]()

func init() {
	scene.SceneManager.Each(func(entity *scene.Scene) {
		entity.Background = BackgroundManager.Register(&Background{
			Name: entity.Name,
			Path: "assets/bg/" + entity.Name + ".png",
			Kind: inkwell.AssetImage,
		})
	})
}
