package ui

// TapeSlots — the two slots of Paul's Armilla.
//
// In canon Paul's Armilla has TWO slots where the standard has one: Dex sits
// permanently in the first, and tapes picked up during the game go into the
// second. The wireframe deck does not know about this at all — this is the one
// HUD element that comes purely from the wiki.
//
// Interaction follows the selected verb, like everything else:
//
//	talk  → converse with whoever occupies the slot
//	look  → describe the slot
//	use   + a selected tape → load it into slot 2
//
// Nothing here touches the world: this is the Vox side of the two channels.

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"realworld/internal/names"
	"realworld/internal/world"
)

type TapeSlots struct {
	Name   string
	W      *world.World
	Bounds inkwell.Rectangle
}

func (t *TapeSlots) GetName() string { return t.Name }

func (t *TapeSlots) BlocksClickAt(p inkwell.Point) bool {
	return t.W.HUDVisible() && t.Bounds.Contains(p)
}

// slotRects splits the widget bounds into the two slot rectangles.
func (t *TapeSlots) slotRects() (inkwell.Rectangle, inkwell.Rectangle) {
	b := t.Bounds
	w := (b.W - Pad) / 2
	return inkwell.Rect(b.X, b.Y, w, b.H), inkwell.Rect(b.X+w+Pad, b.Y, w, b.H)
}

func (t *TapeSlots) Tick(ctx *inkwell.UICtx) {
	if !t.W.HUDVisible() {
		return
	}
	g := ctx.Game
	mp := g.Input.Point()
	r1, r2 := t.slotRects()

	// Hover label, so the StatusLine spells out what is about to happen.
	switch {
	case r1.Contains(mp):
		g.SetHoverLabel(names.TapeDisplayName(names.Dex))
	case r2.Contains(mp):
		if s := t.W.Slot2(); s != "" {
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

	// Loading a tape into slot 2.
	if slot2 && sel != "" {
		g.Inventory.Select("")
		if names.IsTapeItem(sel) {
			t.W.SetPending(sel)
			t.W.Do(inkwell.RunScript(names.ScriptTapeInsert))
			return
		}
		t.W.Do(world.TapeSay(names.Dex, "That isn't a tape, Paul. That's an object. There is a difference."))
		return
	}
	if !slot2 && sel != "" {
		g.Inventory.Select("")
		t.W.Do(world.TapeSay(names.Dex, "The first slot is mine. I'm not moving."))
		return
	}

	switch g.SelectedVerb() {
	case "talk":
		if !slot2 {
			t.W.Do(world.Paused(t.W, inkwell.RunDialogue(names.DlgDex)))
			return
		}
		if s := t.W.Slot2(); s != "" {
			t.W.Do(world.Paused(t.W, inkwell.RunDialogue(names.TapeDialogue(s))))
			return
		}
		t.W.Do(world.TapeSay(names.Dex, "Empty. Nobody to talk to."))
	default: // look, take, use all fall back to a description
		if !slot2 {
			t.W.Do(world.TapeSay(names.Dex, "Me. Four kilobytes of a dead man. Be grateful."))
			return
		}
		if s := t.W.Slot2(); s != "" {
			t.W.Do(world.TapeSay(names.Dex, "That is the second slot. Be careful what you let in."))
			return
		}
		t.W.Do(world.TapeSay(names.Dex, "The second slot is empty. That's the rarer state."))
	}
}

func (t *TapeSlots) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {
	if !t.W.HUDVisible() {
		return
	}
	g := ctx.Game
	th := g.Theme()
	r1, r2 := t.slotRects()

	// Two text rows per slot: the slot number on top, its occupant below.
	// Both rows are LineH tall, so nothing can overlap the border.
	drawSlot := func(r inkwell.Rectangle, label, body string, filled bool) {
		bg, border, col := th.InventorySlotBG, th.ChatLogSystem, th.ChatLogSystem
		if filled {
			bg, border, col = th.PanelBG, th.CharacterPanelBorder, th.ChatLogResponse
		}
		vector.DrawFilledRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), bg, false)
		vector.StrokeRect(dst, float32(r.X), float32(r.Y), float32(r.W), float32(r.H), 1, border, false)

		x, inner := int(r.X)+Pad, int(r.W)-Pad*2
		drawTextC(dst, label, x, int(r.Y)+Pad/2, th.ChatLogSystem)
		drawTextC(dst, clip(body, inner), x, int(r.Y)+Pad/2+LineH, col)
	}

	drawSlot(r1, "SLOT 1", names.TapeDisplayName(names.Dex), true)
	if s := t.W.Slot2(); s != "" {
		drawSlot(r2, "SLOT 2", itemLabel(g, s), true)
	} else {
		drawSlot(r2, "SLOT 2", "empty", false)
	}
}

// itemLabel is a short display label for an inventory item.
func itemLabel(g *inkwell.Game, name string) string {
	if it, ok := g.ItemManager.Get(name); ok && it.Description != "" {
		return it.Description
	}
	return name
}
