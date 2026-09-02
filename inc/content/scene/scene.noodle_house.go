package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/system"
)

type noodleHouse struct{}

var _ = SceneManager.Register(Scene{
	Name: inc.SceneNoodleHouse,
	Title: inkwell.Text{
		{Lang: "en", Text: "NOODLE'S HOUSE — SEALED"},
		{Lang: "hu", Text: "NOODLE HÁZA — LEPECSÉTELVE"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
		{
			To: inc.SceneAlley,
			Label: inkwell.Text{
				{Lang: "en", Text: "the alley behind the house"},
				{Lang: "hu", Text: "a sikátor a ház mögött"},
			},
			Side: inkwell.ExitRight,
		},
		{
			To: inc.SceneNoodleOffice,
			Label: inkwell.Text{
				{Lang: "en", Text: "Noodle's front door"},
				{Lang: "hu", Text: "Noodle bejárati ajtaja"},
			},
			Side:  inkwell.ExitBack,
			Needs: inc.FlagNoodleReleased,
			Blocked: inkwell.Seq(
				inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
					{Lang: "en", Text: "Sir. The tape is there for a reason."},
					{Lang: "hu", Text: "Uram. A szalag okkal van ott."},
				}),
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Sealed. And watched."},
					{Lang: "hu", Text: "Lepecsételve. És figyelik."},
				}),
			),
		},
	},
	Actors: []inkwell.SceneActor{
		{
			CharacterName: inc.Paul,
			At: inkwell.Point{
				X: 100,
				Y: 232,
			},
		},
		{
			CharacterName: inc.OfficerOnGuard,
			At: inkwell.Point{
				X: 340,
				Y: 236,
			},
		},
	},
	Hotspots: noodleHouse{}.hotspots(),
})

func (house noodleHouse) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		house.officer(),
		house.sealedDoor(),
		house.policeTape(),
		house.window(),
	}
}

func (noodleHouse) officer() inkwell.Hotspot {
	gone := inkwell.Say(inc.Paul, inkwell.Text{
		{Lang: "en", Text: "He's gone. Took the tape with him, left the smell of the donuts."},
		{Lang: "hu", Text: "Elment. A szalagot vitte, a fánkszagot hagyta."},
	})
	return inkwell.Hotspot{
		Name: "officer",
		Label: inkwell.Text{
			{Lang: "en", Text: "officer on guard"},
			{Lang: "hu", Text: "őrt álló rendőr"},
		},
		Area: inkwell.Rect(316, 156, 48, 84),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			gone,
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Arms crossed, feet apart, jaw like a shovel. He's been standing there since before the tape went up."},
				{Lang: "hu", Text: "Karba tett kéz, terpesz, lapát-állkapocs. Azelőtt állt ide, hogy a szalag felkerült."},
			}),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			gone,
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I'm not going to push a policeman."},
				{Lang: "hu", Text: "Nem fogok meglökni egy rendőrt."},
			}),
		),
		OnTalk: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			gone,
			system.System.Paused(inkwell.RunDialog(inc.DialogOfficerOnGuard)),
		),
		OnTake: inkwell.Say(inc.Dex, inkwell.Text{
			{Lang: "en", Text: "That's called kidnapping, Paul, and he's bigger than you."},
			{Lang: "hu", Text: "Azt emberrablásnak hívják, Paul, és nagyobb nálad."},
		}),
		OnUseWith: map[string]inkwell.Action{
			inc.ItemBreakfastBox: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				gone,
				inkwell.If(inkwell.Flag(inc.FlagOfficerFed),
					inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
						{Lang: "en", Text: "I'm good, thanks. One's plenty on duty."},
						{Lang: "hu", Text: "Köszönöm, jól vagyok. Szolgálatban egy is elég."},
					}),
					inkwell.Seq(
						inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
							{Lang: "en", Text: "For me? …All right. Thanks. Genuinely. I've been out here since six."},
							{Lang: "hu", Text: "Nekem? …Jó. Köszönöm. Komolyan. Hat óta itt állok."},
						}),
						inkwell.TakeAway(inc.ItemBreakfastBox),
						inkwell.SetFlag(inc.FlagOfficerFed),
						inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
							{Lang: "en", Text: "Look. The owner's not here. They took him to the station on Valencia, the SFPD one. You want him, that's where he is."},
							{Lang: "hu", Text: "Nézze. A tulaj nincs itt. Bevitték a Valencia utcai őrsre, az SFPD-hez. Ha őt keresi, ott van."},
						}),
						inkwell.Say(inc.Paul, inkwell.Text{
							{Lang: "en", Text: "The station. Thanks."},
							{Lang: "hu", Text: "Az őrs. Köszönöm."},
						}),
						inkwell.Say(inc.Dex, inkwell.Text{
							{Lang: "en", Text: "Two dumplings for a man's whereabouts. You've made worse trades."},
							{Lang: "hu", Text: "Két gombóc egy ember hollétéért. Csináltál már rosszabb üzletet."},
						}),
					),
				),
			),
		},
		OnUseWithAny: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			gone,
			inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
				{Lang: "en", Text: "Put that away, sir."},
				{Lang: "hu", Text: "Tegye el, uram."},
			}),
		),
	}
}

func (noodleHouse) sealedDoor() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "sealed_door",
		Label: inkwell.Text{
			{Lang: "en", Text: "Noodle's front door"},
			{Lang: "hu", Text: "Noodle bejárati ajtaja"},
		},
		Area: inkwell.Rect(400, 90, 80, 110),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Seal's cut. Somebody's home."},
				{Lang: "hu", Text: "A pecsét elvágva. Valaki otthon van."},
			}),
			inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
				{Lang: "en", Text: "Nothing to see here, sir. Move along."},
				{Lang: "hu", Text: "Nincs itt semmi látnivaló, uram. Tessék továbbmenni."},
			}),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Unlocked. He never locked it before either."},
				{Lang: "hu", Text: "Nyitva. Régen sem zárta soha."},
			}),
			inkwell.Seq(
				inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
					{Lang: "en", Text: "Sir. Step back from the door."},
					{Lang: "hu", Text: "Uram. Lépjen hátra az ajtótól."},
				}),
				inkwell.Say(inc.Paul, inkwell.Text{
					{Lang: "en", Text: "Stepping."},
					{Lang: "hu", Text: "Lépek."},
				}),
			),
		),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "The door stays with the house."},
			{Lang: "hu", Text: "Az ajtó a házzal marad."},
		}),
	}
}

func (noodleHouse) policeTape() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "police_tape",
		Label: inkwell.Text{
			{Lang: "en", Text: "police tape"},
			{Lang: "hu", Text: "rendőrségi szalag"},
		},
		Area: inkwell.Rect(230, 150, 300, 24),
		OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Torn and left hanging. They never take it down properly."},
				{Lang: "hu", Text: "Elszakítva, lógva hagyva. Soha nem szedik le rendesen."},
			}),
			inkwell.Say(inc.OfficerOnGuard, inkwell.Text{
				{Lang: "en", Text: "Nothing to see here, sir. Move along."},
				{Lang: "hu", Text: "Nincs itt semmi látnivaló, uram. Tessék továbbmenni."},
			}),
		),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "I duck under it. Nobody cares any more."},
				{Lang: "hu", Text: "Átbújok alatta. Már senkit nem érdekel."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Touch that and you'll be inside a different building tonight."},
				{Lang: "hu", Text: "Nyúlj hozzá, és ma este egy másik épületben alszol."},
			}),
		),
		OnTake: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "Yellow isn't my colour."},
			{Lang: "hu", Text: "A sárga nem az én színem."},
		}),
	}
}

func (noodleHouse) window() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "window",
		Label: inkwell.Text{
			{Lang: "en", Text: "front window"},
			{Lang: "hu", Text: "utcai ablak"},
		},
		Area: inkwell.Rect(500, 80, 80, 80),
		OnLook: inkwell.Seq(
			inkwell.Say(inc.Paul, inkwell.Text{
				{Lang: "en", Text: "Blinds down. A rack of Armillas behind them, if the shadows are honest."},
				{Lang: "hu", Text: "Lehúzott redőny. Mögötte egy állvány Armillákkal, ha az árnyékok nem hazudnak."},
			}),
			inkwell.Say(inc.Dex, inkwell.Text{
				{Lang: "en", Text: "Cracked ones. That's what they took him in for. Not what they were looking for."},
				{Lang: "hu", Text: "Feltörtek. Ezért vitték be. Nem ezt keresték."},
			}),
		),
		OnUse: inkwell.Say(inc.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not breaking a window in front of a policeman. Or behind one."},
			{Lang: "hu", Text: "Nem török be ablakot egy rendőr előtt. Se mögött."},
		}),
	}
}
