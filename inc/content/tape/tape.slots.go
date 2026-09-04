package tape

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/inc/content/script"
	"realworld/inc/content/speaker"
	"realworld/inc/system"
	"realworld/inc/system/armilla"
	"realworld/inc/system/widget"
)

type slotPanel struct {
	inkwell.HotspotPanel
	first  inkwell.Rectangle
	second inkwell.Rectangle
}

func init() {
	bounds := inkwell.Rect(192, widget.SlotsY, 194, widget.SlotsHeight)
	width := (bounds.W - widget.Padding) / 2
	panel := &slotPanel{
		HotspotPanel: inkwell.HotspotPanel{
			Name:   "tapes",
			When:   widget.WhenPlaying,
			Bounds: bounds,
		},
		first:  inkwell.Rect(bounds.X, bounds.Y, width, bounds.H),
		second: inkwell.Rect(bounds.X+width+widget.Padding, bounds.Y, width, bounds.H),
	}
	panel.Hotspots = panel.hotspots
	widget.WidgetManager.Register(panel)
}

func (slots *slotPanel) hotspots(game *inkwell.Game) []Hotspot {
	dexRemark := speaker.Dex.Say(Text{
		{Lang: "en", Text: "Me. Four kilobytes of a dead man. Be grateful."},
		{Lang: "hu", Text: "Én. Egy halott ember négy kilobájtja. Légy hálás."},
	})
	notATape := speaker.Dex.Say(Text{
		{Lang: "en", Text: "That isn't a tape, Paul. That's an object. There is a difference."},
		{Lang: "hu", Text: "Az nem kazetta, Paul. Az egy tárgy. Van különbség."},
	})
	insertions := make(map[*inkwell.Item]Action)
	for _, candidate := range TapeManager.All() {
		if candidate.Item == nil {
			continue
		}
		insertions[candidate.Item] = inkwell.Seq(
			armilla.Insert(candidate.Item),
			inkwell.RunScript(script.TapeInsert),
			candidate.OnInsert,
		)
	}
	dex := Hotspot{
		Name:   "slot1",
		Label:  speaker.Dex.DisplayName(),
		Area:   slots.first,
		OnLook: dexRemark,
		OnUse:  dexRemark,
		OnTake: dexRemark,
		OnTalk: system.System.Paused(inkwell.RunDialog(Dex.Dialog)),
		OnUseWithAny: speaker.Dex.Say(Text{
			{Lang: "en", Text: "The first slot is mine. I'm not moving."},
			{Lang: "hu", Text: "Az első foglalat az enyém. Nem megyek sehová."},
		}),
	}
	loaded := armilla.Slot2()
	if loaded == nil {
		empty := speaker.Dex.Say(Text{
			{Lang: "en", Text: "The second slot is empty. That's the rarer state."},
			{Lang: "hu", Text: "A második foglalat üres. Ez a ritkább állapot."},
		})
		return []Hotspot{
			dex,
			{
				Name:   "slot2",
				Label:  slotEmptyLabel,
				Area:   slots.second,
				OnLook: empty,
				OnUse:  empty,
				OnTake: empty,
				OnTalk: speaker.Dex.Say(Text{
					{Lang: "en", Text: "Empty. Nobody to talk to."},
					{Lang: "hu", Text: "Üres. Nincs kivel beszélni."},
				}),
				OnUseWith:    insertions,
				OnUseWithAny: notATape,
			},
		}
	}
	careful := speaker.Dex.Say(Text{
		{Lang: "en", Text: "That is the second slot. Be careful what you let in."},
		{Lang: "hu", Text: "Az a második foglalat. Vigyázz, mit engedsz be."},
	})
	return []Hotspot{
		dex,
		{
			Name:         "slot2",
			Label:        loaded.Description,
			Area:         slots.second,
			OnLook:       careful,
			OnUse:        careful,
			OnTake:       careful,
			OnTalk:       system.System.Paused(inkwell.RunDialog(DialogOf(loaded))),
			OnUseWith:    insertions,
			OnUseWithAny: notATape,
		},
	}
}

func (slots *slotPanel) Draw(destination *ebiten.Image, context *inkwell.UICtx) {
	game := context.Game
	theme := game.Theme()
	paint := func(area inkwell.Rectangle, label, body Text, filled bool) {
		background, border, textColor := theme.InventorySlotBG, theme.ChatLogSystem, theme.ChatLogSystem
		if filled {
			background, border, textColor = theme.PanelBG, theme.CharacterPanelBorder, theme.ChatLogResponse
		}
		vector.DrawFilledRect(destination, float32(area.X), float32(area.Y),
			float32(area.W), float32(area.H), background, false)
		vector.StrokeRect(destination, float32(area.X), float32(area.Y),
			float32(area.W), float32(area.H), 1, border, false)

		textX, innerWidth := int(area.X)+widget.Padding, int(area.W)-widget.Padding*2
		game.DrawText(destination, label.String(), textX, int(area.Y)+widget.Padding/2, theme.ChatLogSystem)
		game.DrawText(destination, inkwell.ClipText(body.String(), innerWidth), textX,
			int(area.Y)+widget.Padding/2+widget.LineHeight, textColor)
	}
	paint(slots.first, slot1Label, speaker.Dex.DisplayName(), true)
	if loaded := armilla.Slot2(); loaded != nil {
		paint(slots.second, slot2Label, loaded.Description, true)
		return
	}
	paint(slots.second, slot2Label, slotEmpty, false)
}

var (
	slot1Label = Text{
		{Lang: "en", Text: "SLOT 1"},
		{Lang: "hu", Text: "1. FOGLALAT"},
	}
	slot2Label = Text{
		{Lang: "en", Text: "SLOT 2"},
		{Lang: "hu", Text: "2. FOGLALAT"},
	}
	slotEmpty = Text{
		{Lang: "en", Text: "empty"},
		{Lang: "hu", Text: "üres"},
	}
	slotEmptyLabel = Text{
		{Lang: "en", Text: "empty tape slot"},
		{Lang: "hu", Text: "üres kazettafoglalat"},
	}
)
