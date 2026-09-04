package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/dialog"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
)

var noodleHouseGone = character.Paul.Say(Text{
	{Lang: "en", Text: "He's gone. Took the tape with him, left the smell of the donuts."},
	{Lang: "hu", Text: "Elment. A szalagot vitte, a fánkszagot hagyta."},
})

var NoodleHouse = SceneManager.Register(&Scene{
	Name: "noodle_house",
	Title: Text{
		{Lang: "en", Text: "NOODLE'S HOUSE — SEALED"},
		{Lang: "hu", Text: "NOODLE HÁZA — LEPECSÉTELVE"},
	},
	Actors: []SceneActor{
		{
			Character: character.Paul,
			At: Point{
				X: 100,
				Y: 232,
			},
		},
		{
			Character: character.OfficerOnGuard,
			At: Point{
				X: 340,
				Y: 236,
			},
		},
	},
	Hotspots: []Hotspot{
		{
			Name: "officer",
			Label: Text{
				{Lang: "en", Text: "officer on guard"},
				{Lang: "hu", Text: "őrt álló rendőr"},
			},
			Area: inkwell.Rect(316, 156, 48, 84),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				noodleHouseGone,
				character.Paul.Say(Text{
					{Lang: "en", Text: "Arms crossed, feet apart, jaw like a shovel. He's been standing there since before the tape went up."},
					{Lang: "hu", Text: "Karba tett kéz, terpesz, lapát-állkapocs. Azelőtt állt ide, hogy a szalag felkerült."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				noodleHouseGone,
				character.Paul.Say(Text{
					{Lang: "en", Text: "I'm not going to push a policeman."},
					{Lang: "hu", Text: "Nem fogok meglökni egy rendőrt."},
				}),
			),
			OnTalk: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				noodleHouseGone,
				system.System.Paused(inkwell.RunDialog(dialog.OfficerOnGuard)),
			),
			OnTake: speaker.Dex.Say(Text{
				{Lang: "en", Text: "That's called kidnapping, Paul, and he's bigger than you."},
				{Lang: "hu", Text: "Azt emberrablásnak hívják, Paul, és nagyobb nálad."},
			}),
			OnUseWith: map[*inkwell.Item]Action{
				item.BreakfastBox: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
					noodleHouseGone,
					inkwell.If(inkwell.Flag(inc.FlagOfficerFed),
						character.OfficerOnGuard.Say(Text{
							{Lang: "en", Text: "I'm good, thanks. One's plenty on duty."},
							{Lang: "hu", Text: "Köszönöm, jól vagyok. Szolgálatban egy is elég."},
						}),
						inkwell.Seq(
							character.OfficerOnGuard.Say(Text{
								{Lang: "en", Text: "For me? …All right. Thanks. Genuinely. I've been out here since six."},
								{Lang: "hu", Text: "Nekem? …Jó. Köszönöm. Komolyan. Hat óta itt állok."},
							}),
							inkwell.TakeAway(item.BreakfastBox),
							inkwell.SetFlag(inc.FlagOfficerFed),
							character.OfficerOnGuard.Say(Text{
								{Lang: "en", Text: "Look. The owner's not here. They took him to the station on Valencia, the SFPD one. You want him, that's where he is."},
								{Lang: "hu", Text: "Nézze. A tulaj nincs itt. Bevitték a Valencia utcai őrsre, az SFPD-hez. Ha őt keresi, ott van."},
							}),
							character.Paul.Say(Text{
								{Lang: "en", Text: "The station. Thanks."},
								{Lang: "hu", Text: "Az őrs. Köszönöm."},
							}),
							speaker.Dex.Say(Text{
								{Lang: "en", Text: "Two dumplings for a man's whereabouts. You've made worse trades."},
								{Lang: "hu", Text: "Két gombóc egy ember hollétéért. Csináltál már rosszabb üzletet."},
							}),
						),
					),
				),
			},
			OnUseWithAny: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				noodleHouseGone,
				character.OfficerOnGuard.Say(Text{
					{Lang: "en", Text: "Put that away, sir."},
					{Lang: "hu", Text: "Tegye el, uram."},
				}),
			),
		},
		{
			Name: "sealed_door",
			Label: Text{
				{Lang: "en", Text: "Noodle's front door"},
				{Lang: "hu", Text: "Noodle bejárati ajtaja"},
			},
			Area: inkwell.Rect(400, 90, 80, 110),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Seal's cut. Somebody's home."},
					{Lang: "hu", Text: "A pecsét elvágva. Valaki otthon van."},
				}),
				character.OfficerOnGuard.Say(Text{
					{Lang: "en", Text: "Nothing to see here, sir. Move along."},
					{Lang: "hu", Text: "Nincs itt semmi látnivaló, uram. Tessék továbbmenni."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Unlocked. He never locked it before either."},
					{Lang: "hu", Text: "Nyitva. Régen sem zárta soha."},
				}),
				inkwell.Seq(
					character.OfficerOnGuard.Say(Text{
						{Lang: "en", Text: "Sir. Step back from the door."},
						{Lang: "hu", Text: "Uram. Lépjen hátra az ajtótól."},
					}),
					character.Paul.Say(Text{
						{Lang: "en", Text: "Stepping."},
						{Lang: "hu", Text: "Lépek."},
					}),
				),
			),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "The door stays with the house."},
				{Lang: "hu", Text: "Az ajtó a házzal marad."},
			}),
		},
		{
			Name: "police_tape",
			Label: Text{
				{Lang: "en", Text: "police tape"},
				{Lang: "hu", Text: "rendőrségi szalag"},
			},
			Area: inkwell.Rect(230, 150, 300, 24),
			OnLook: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				character.Paul.Say(Text{
					{Lang: "en", Text: "Torn and left hanging. They never take it down properly."},
					{Lang: "hu", Text: "Elszakítva, lógva hagyva. Soha nem szedik le rendesen."},
				}),
				character.OfficerOnGuard.Say(Text{
					{Lang: "en", Text: "Nothing to see here, sir. Move along."},
					{Lang: "hu", Text: "Nincs itt semmi látnivaló, uram. Tessék továbbmenni."},
				}),
			),
			OnUse: inkwell.If(inkwell.Flag(inc.FlagNoodleReleased),
				character.Paul.Say(Text{
					{Lang: "en", Text: "I duck under it. Nobody cares any more."},
					{Lang: "hu", Text: "Átbújok alatta. Már senkit nem érdekel."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Touch that and you'll be inside a different building tonight."},
					{Lang: "hu", Text: "Nyúlj hozzá, és ma este egy másik épületben alszol."},
				}),
			),
			OnTake: character.Paul.Say(Text{
				{Lang: "en", Text: "Yellow isn't my colour."},
				{Lang: "hu", Text: "A sárga nem az én színem."},
			}),
		},
		{
			Name: "window",
			Label: Text{
				{Lang: "en", Text: "front window"},
				{Lang: "hu", Text: "utcai ablak"},
			},
			Area: inkwell.Rect(500, 80, 80, 80),
			OnLook: inkwell.Seq(
				character.Paul.Say(Text{
					{Lang: "en", Text: "Blinds down. A rack of Armillas behind them, if the shadows are honest."},
					{Lang: "hu", Text: "Lehúzott redőny. Mögötte egy állvány Armillákkal, ha az árnyékok nem hazudnak."},
				}),
				speaker.Dex.Say(Text{
					{Lang: "en", Text: "Cracked ones. That's what they took him in for. Not what they were looking for."},
					{Lang: "hu", Text: "Feltörtek. Ezért vitték be. Nem ezt keresték."},
				}),
			),
			OnUse: character.Paul.Say(Text{
				{Lang: "en", Text: "I'm not breaking a window in front of a policeman. Or behind one."},
				{Lang: "hu", Text: "Nem török be ablakot egy rendőr előtt. Se mögött."},
			}),
		},
	},
})

var _ = NoodleHouse.Leads(
	exitToSelector,
	Exit{
		To: Alley,
		Label: Text{
			{Lang: "en", Text: "the alley behind the house"},
			{Lang: "hu", Text: "a sikátor a ház mögött"},
		},
		Side: inkwell.ExitRight,
	},
	Exit{
		To: NoodleOffice,
		Label: Text{
			{Lang: "en", Text: "Noodle's front door"},
			{Lang: "hu", Text: "Noodle bejárati ajtaja"},
		},
		Side:  inkwell.ExitBack,
		Needs: inc.FlagNoodleReleased,
		Blocked: inkwell.Seq(
			character.OfficerOnGuard.Say(Text{
				{Lang: "en", Text: "Sir. The tape is there for a reason."},
				{Lang: "hu", Text: "Uram. A szalag okkal van ott."},
			}),
			character.Paul.Say(Text{
				{Lang: "en", Text: "Sealed. And watched."},
				{Lang: "hu", Text: "Lepecsételve. És figyelik."},
			}),
		),
	},
)
