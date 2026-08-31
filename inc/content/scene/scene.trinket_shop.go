package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
)

func init() {
	Manager.Register(Scene{
		Name: inc.SceneTrinketShop,
		Title: inkwell.Text{
			{Lang: "en", Text: "CHINATOWN — TRINKET SHOP"},
			{Lang: "hu", Text: "KÍNAI NEGYED — CSECSEBECSEBOLT"},
		},
		Background: inc.BgTrinketShop,
		Exits: []inkwell.Exit{
			exitToSelector,
			{
				To: inc.SceneSmallRestaurant,
				Label: inkwell.Text{
					{Lang: "en", Text: "the eating place next door"},
					{Lang: "hu", Text: "a szomszédos falatozó"},
				},
				Side: inkwell.ExitRight,
			},
			{
				To: inc.SceneSecretClub,
				Label: inkwell.Text{
					{Lang: "en", Text: "the way in at the back"},
					{Lang: "hu", Text: "a hátsó bejárat"},
				},
				Side:  inkwell.ExitBack,
				Needs: inc.FlagClubEntry,
				Blocked: inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Just a shop, as far as the man behind the counter is concerned."},
					{Lang: "hu", Text: "Csak egy bolt, már ami a pult mögötti embert illeti."},
				}),
			},
		},
	})
}
