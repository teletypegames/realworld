package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func init() {
	WidgetManager.Register(&tapeSlots{
		Name:   "tapes",
		Bounds: inkwell.Rect(192, slotsY, 194, slotsH),
	})
}

type tapeSlots struct {
	Name   string
	Bounds inkwell.Rectangle
}

func (t *tapeSlots) GetName() string                { return t.Name }
func (t *tapeSlots) VisibleWhen() inkwell.Condition { return whenPlaying }
func (t *tapeSlots) Layer() inkwell.Layer           { return inkwell.LayerHUD }

func (t *tapeSlots) BlocksClickAt(p inkwell.Point) bool { return t.Bounds.Contains(p) }

func (t *tapeSlots) slotRects() (inkwell.Rectangle, inkwell.Rectangle) {
	b := t.Bounds
	w := (b.W - Pad) / 2
	return inkwell.Rect(b.X, b.Y, w, b.H), inkwell.Rect(b.X+w+Pad, b.Y, w, b.H)
}

func (t *tapeSlots) Tick(ctx *inkwell.UICtx) {
	g := ctx.Game
	mp := g.Input.Point()
	r1, r2 := t.slotRects()

	switch {
	case r1.Contains(mp):
		g.SetHoverLabel(g.CharacterLabel(Dex))
	case r2.Contains(mp):
		if s := World.Slot2(); s != "" {
			g.SetHoverLabel(itemLabel(g, s))
		} else {
			g.SetHoverLabel("empty tape slot")
		}
	}

	if !g.Input.LeftClicked() {
		return
	}
	if !r1.Contains(mp) && !r2.Contains(mp) {
		return
	}
	g.Input.ConsumeLeft()

	slot2 := r2.Contains(mp)
	sel := g.Inventory.Selected()

	if slot2 && sel != "" {
		g.Inventory.Select("")
		if IsTapeItem(sel) {
			World.SetPending(sel)
			World.Do(inkwell.RunScript(ScriptTapeInsert))
			return
		}
		World.Do(inkwell.Say(Dex, "That isn't a tape, Paul. That's an object. There is a difference."))
		return
	}
	if !slot2 && sel != "" {
		g.Inventory.Select("")
		World.Do(inkwell.Say(Dex, "The first slot is mine. I'm not moving."))
		return
	}

	switch g.SelectedVerb() {
	case "talk":
		if !slot2 {
			World.Do(Paused(inkwell.RunDialogue(DlgDex)))
			return
		}
		if s := World.Slot2(); s != "" {
			World.Do(Paused(inkwell.RunDialogue(TapeDialogue(s))))
			return
		}
		World.Do(inkwell.Say(Dex, "Empty. Nobody to talk to."))
	default:
		if !slot2 {
			World.Do(inkwell.Say(Dex, "Me. Four kilobytes of a dead man. Be grateful."))
			return
		}
		if s := World.Slot2(); s != "" {
			World.Do(inkwell.Say(Dex, "That is the second slot. Be careful what you let in."))
			return
		}
		World.Do(inkwell.Say(Dex, "The second slot is empty. That's the rarer state."))
	}
}

func (t *tapeSlots) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	g := ctx.Game
	th := g.Theme()
	r1, r2 := t.slotRects()

	drawSlot := func(r inkwell.Rectangle, label, body string, filled bool) {
		bg, border, col := th.InventorySlotBG, th.ChatLogSystem, th.ChatLogSystem
		if filled {
			bg, border, col = th.PanelBG, th.CharacterPanelBorder, th.ChatLogResponse
		}
		vector.DrawFilledRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), bg, false)
		vector.StrokeRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), 1, border, false)

		x, inner := int(r.X)+Pad, int(r.W)-Pad*2
		g.DrawText(dst, label, x, int(r.Y)+Pad/2, th.ChatLogSystem)
		g.DrawText(dst, inkwell.ClipText(body, inner), x, int(r.Y)+Pad/2+LineH, col)
	}

	drawSlot(r1, "SLOT 1", g.CharacterLabel(Dex), true)
	if s := World.Slot2(); s != "" {
		drawSlot(r2, "SLOT 2", itemLabel(g, s), true)
	} else {
		drawSlot(r2, "SLOT 2", "empty", false)
	}
}

func itemLabel(g *inkwell.Game, name string) string {
	if it, ok := g.ItemManager.Get(name); ok && it.Description != "" {
		return it.Description
	}
	return name
}
