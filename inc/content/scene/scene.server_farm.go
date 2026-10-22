package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

var serverFarmWalkTheRacks = inkwell.Seq(
	speaker.TapeServerFarm.Say(Text{
		{Lang: "en", Text: "DOOR GROUP MATCHED. WELCOME."},
		{Lang: "hu", Text: "AJTÓCSOPORT EGYEZIK. ÜDVÖZÖLJÜK."},
	}),
	character.Paul.Say(Text{
		{Lang: "en", Text: "Rows of it. Cold air, and the sound of a thousand fans agreeing on something."},
		{Lang: "hu", Text: "Sorokban. Hideg levegő, és ezer ventilátor hangja, ahogy megegyeznek valamiben."},
	}),
	speaker.Dex.Say(Text{
		{Lang: "en", Text: "Look at the traffic lights on the switches. That pattern isn't a customer. That's one thing, talking to itself across the whole floor."},
		{Lang: "hu", Text: "Nézd a switchek forgalomjelzőit. Az a mintázat nem egy ügyfél. Az egyetlen dolog, ami önmagával beszél az egész szinten."},
	}),
	speaker.Dex.Say(Text{
		{Lang: "en", Text: "It isn't in one place any more, Paul. It's in the traffic."},
		{Lang: "hu", Text: "Már nincs egy helyen, Paul. A forgalomban van."},
	}),
	inkwell.SetFlag(inc.FlagServerFarmMapped),
)

var ServerFarm = SceneManager.Register(&Scene{
	Name: "server_farm",
	Title: Text{
		{Lang: "en", Text: "SERVER FARM"},
		{Lang: "hu", Text: "SZERVERFARM"},
	},
	Layers: []SceneLayer{
		{
			Asset: &Asset{
				Name: "server_farm",
				Path: "assets/bg/server_farm.png",
			},
			Z: -1,
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "badge_door",
			Label: Text{
				{Lang: "en", Text: "badge door"},
				{Lang: "hu", Text: "kártyás ajtó"},
			},
			Area: inkwell.Rect(270, 70, 110, 170),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Airlock door. The reader wants a credential it can read off an Armilla."},
				{Lang: "hu", Text: "Zsilipajtó. Az olvasó olyan hitelesítőt akar, amit Armilláról tud olvasni."},
			}),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagServerFarmMapped),
				character.Paul.Say(Text{
					{Lang: "en", Text: "I've seen the floor. It's one thing in there, spread thin, and it's growing."},
					{Lang: "hu", Text: "Láttam a szintet. Egyetlen dolog van bent, vékonyan elterülve, és nő."},
				}),
				inkwell.If(inkwell.VarEq(inc.VarSlot2, item.ServerFarmTape.Name),
					serverFarmWalkTheRacks,
					inkwell.Seq(
						character.Paul.Say(Text{
							{Lang: "en", Text: "Red blink. It doesn't like what's on my wrist."},
							{Lang: "hu", Text: "Piros villanás. Nem tetszik neki, ami a csuklómon van."},
						}),
						speaker.Dex.Say(Text{
							{Lang: "en", Text: "It wants a credential in slot two. Not me. Definitely not me."},
							{Lang: "hu", Text: "Hitelesítőt akar a kettes foglalatban. Nem engem. Határozottan nem engem."},
						}),
					),
				),
			),
		},
		{
			Name: "rack",
			Label: Text{
				{Lang: "en", Text: "server rack"},
				{Lang: "hu", Text: "szerverrack"},
			},
			Area: inkwell.Rect(420, 40, 190, 200),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "LEDs in rows, blinking in a rhythm I can almost follow. Almost."},
				{Lang: "hu", Text: "LED-ek sorokban, olyan ritmusban villognak, amit majdnem követek. Majdnem."},
			}),
			OnUse: speaker.Dex.Say(Text{
				{Lang: "en", Text: "Pull one and it reroutes around the hole before your hand is back. Don't bother."},
				{Lang: "hu", Text: "Húzz ki egyet, és a lyukat megkerüli, mielőtt visszaér a kezed. Ne fáradj."},
			}),
		},
		{
			Name: "temperature_panel",
			Label: Text{
				{Lang: "en", Text: "temperature panel"},
				{Lang: "hu", Text: "hőmérséklet-panel"},
			},
			Area: inkwell.Rect(40, 60, 120, 80),
			OnLook: character.Paul.Say(Text{
				{Lang: "en", Text: "Eighteen degrees on every floor. Load: ninety-six percent, and climbing. Nobody rents that much."},
				{Lang: "hu", Text: "Tizennyolc fok minden szinten. Terhelés: kilencvenhat százalék, és emelkedik. Ennyit senki nem bérel."},
			}),
		},
	},
})

var _ = ServerFarm.Leads(
	exitToSelector,
)
