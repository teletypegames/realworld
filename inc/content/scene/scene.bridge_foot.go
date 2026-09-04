package scene

import (
	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/inc"
	"realworld/inc/content/character"
	"realworld/inc/content/item"
	"realworld/inc/content/speaker"
)

type bridgeFoot struct{}

var BridgeFoot = SceneManager.Register(&Scene{
	Name: "bridge_foot",
	Title: inkwell.Text{
		{Lang: "en", Text: "GOLDEN GATE — THE FOOT OF THE BRIDGE"},
		{Lang: "hu", Text: "GOLDEN GATE — A HÍD LÁBA"},
	},
	Exits: []inkwell.Exit{
		exitToSelector,
		{
			To: securityRoom{},
			Label: inkwell.Text{
				{Lang: "en", Text: "up into the structure"},
				{Lang: "hu", Text: "fel a szerkezetbe"},
			},
			Side:  inkwell.ExitBack,
			Needs: inc.FlagBridgeOpen,
			Blocked: inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Steel door, no handle. It opens for a card, not for me."},
				{Lang: "hu", Text: "Acélajtó, kilincs nélkül. Kártyának nyílik, nem nekem."},
			}),
		},
	},
	Hotspots: bridgeFoot{}.hotspots(),
})

func (place bridgeFoot) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{
		place.door(),
		place.reader(),
		place.bay(),
	}
}

func (bridgeFoot) door() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "service_door",
		Label: inkwell.Text{
			{Lang: "en", Text: "service door"},
			{Lang: "hu", Text: "szolgálati ajtó"},
		},
		Area: inkwell.Rect(268, 90, 104, 140),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "A maintenance door in the base of a national monument. Fresh paint over the seams."},
			{Lang: "hu", Text: "Karbantartó ajtó egy nemzeti műemlék tövében. Friss festék a hézagokon."},
		}),
		OnUse: inkwell.If(inkwell.Flag(inc.FlagBridgeOpen),
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "It's open. Nobody is going to invite me twice."},
				{Lang: "hu", Text: "Nyitva. Kétszer senki nem fog behívni."},
			}),
			inkwell.Say(character.Paul, inkwell.Text{
				{Lang: "en", Text: "Card reader. Nothing to push."},
				{Lang: "hu", Text: "Kártyaolvasó. Nincs mit megnyomni."},
			}),
		),
		OnUseWith: map[*inkwell.Item]inkwell.Action{
			item.SumphorCard: inkwell.Seq(
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Sumphor's card. Let's see if they ever revoked it."},
					{Lang: "hu", Text: "Sumphor kártyája. Lássuk, visszavonták-e valaha."},
				}),
				inkwell.Wait(0.6),
				inkwell.SetFlag(inc.FlagBridgeOpen),
				inkwell.Say(character.Paul, inkwell.Text{
					{Lang: "en", Text: "Green. They never did."},
					{Lang: "hu", Text: "Zöld. Soha nem vonták."},
				}),
				inkwell.Say(speaker.Dex, inkwell.Text{
					{Lang: "en", Text: "A bridge the whole world takes pictures of, and nobody looks at the door in the bottom of it."},
					{Lang: "hu", Text: "Egy híd, amit az egész világ fotóz, és senki nem nézi meg az ajtót az aljában."},
				}),
			),
		},
	}
}

func (bridgeFoot) reader() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "card_reader",
		Label: inkwell.Text{
			{Lang: "en", Text: "card reader"},
			{Lang: "hu", Text: "kártyaolvasó"},
		},
		Area: inkwell.Rect(380, 140, 28, 40),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Not municipal. This is BVK hardware, and a new model."},
			{Lang: "hu", Text: "Nem önkormányzati. Ez BVK-hardver, méghozzá új típus."},
		}),
		OnUse: inkwell.Say(speaker.Dex, inkwell.Text{
			{Lang: "en", Text: "It wants a card, Paul. Waving your wrist at it just tells it you have an Armilla."},
			{Lang: "hu", Text: "Kártyát akar, Paul. Ha a csuklódat lengeted neki, csak annyit tud meg, hogy van Armillád."},
		}),
	}
}

func (bridgeFoot) bay() inkwell.Hotspot {
	return inkwell.Hotspot{
		Name: "bay",
		Label: inkwell.Text{
			{Lang: "en", Text: "the bay"},
			{Lang: "hu", Text: "az öböl"},
		},
		Area: inkwell.Rect(24, 40, 220, 120),
		OnLook: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "Fog on the water. Somewhere under all that orange steel, Norman is sleeping."},
			{Lang: "hu", Text: "Köd a vízen. Valahol az alatt a sok narancssárga acél alatt Norman alszik."},
		}),
		OnUse: inkwell.Say(character.Paul, inkwell.Text{
			{Lang: "en", Text: "I'm not swimming to it."},
			{Lang: "hu", Text: "Nem úszom oda."},
		}),
	}
}

func (bridgeFoot) Scene() *inkwell.Scene { return BridgeFoot }
