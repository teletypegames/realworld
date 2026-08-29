package boot_test

import (
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/boot"
	"realworld/internal/names"
	"realworld/internal/ui"
)

func build(t *testing.T) *inkwell.Game {
	t.Helper()
	g := boot.New(boot.Opts{})
	if err := g.Validate(); err != nil {
		t.Fatalf("Validate: %v", err)
	}
	return g
}

// The engine's Validate only checks scene→asset/character references. A typo in
// a dialogue, script or item name would surface at runtime as a silent no-op
// (RunDialogue and RunScript fail quietly), so check them here.
func TestReferencesResolve(t *testing.T) {
	g := build(t)

	for _, n := range []string{names.DlgDex, names.DlgMysteryTape} {
		if !g.DialogueManager.Has(n) {
			t.Errorf("missing dialogue: %q", n)
		}
	}
	for _, n := range []string{names.ScriptTapeInsert, names.ScriptFinale} {
		if !g.ScriptManager.Has(n) {
			t.Errorf("missing script: %q", n)
		}
	}
	for _, n := range []string{names.ItemNoodleLetter, names.ItemMysteryTape, names.ItemBlackArmilla} {
		if !g.ItemManager.Has(n) {
			t.Errorf("missing item: %q", n)
		}
	}
	for n := range names.Tapes {
		if !g.CharacterManager.Has(n) {
			t.Errorf("tape character not registered: %q", n)
		}
	}
	// Every loadable tape needs an item, a character and a dialogue.
	for item, char := range names.TapeItems {
		if !g.ItemManager.Has(item) {
			t.Errorf("tape item not registered: %q", item)
		}
		if !g.CharacterManager.Has(char) {
			t.Errorf("tape character not registered: %q", char)
		}
		if !g.DialogueManager.Has(names.TapeDialogue(item)) {
			t.Errorf("tape %q has no dialogue: %q", item, names.TapeDialogue(item))
		}
	}
}

// The two-channel rule: every tape needs its own voice colour, distinct from
// the world's. Without that, "if something is amber, a tape said it" is not
// readable.
func TestTapeVoicesAreDistinct(t *testing.T) {
	g := build(t)
	paul, _ := g.CharacterManager.Get(names.Paul)
	seen := map[[4]uint32]string{}
	for n := range names.Tapes {
		c, _ := g.CharacterManager.Get(n)
		if c.SpeechColor == nil {
			t.Errorf("%s: no SpeechColor", n)
			continue
		}
		if c.SpeechColor == paul.SpeechColor {
			t.Errorf("%s speaks in the same colour as Paul", n)
		}
		r, gr, b, a := c.SpeechColor.RGBA()
		key := [4]uint32{r, gr, b, a}
		if other, dup := seen[key]; dup {
			t.Errorf("%s and %s share a voice colour", n, other)
		}
		seen[key] = n
	}
}

// HUD geometry: the two channels must not overlap, and both must stay inside
// the HUD strip.
func TestHUDGeometry(t *testing.T) {
	g := build(t)

	ch, ok := g.UIManager.Get("channel")
	if !ok {
		t.Fatal("no 'channel' widget")
	}
	tc := ch.(*ui.TapeChannel)
	if tc.Bounds.X < ui.DividerX {
		t.Errorf("tape channel crosses into the world side: X=%v < %v", tc.Bounds.X, ui.DividerX)
	}
	if tc.Bounds.X+tc.Bounds.W > ui.ScreenW {
		t.Errorf("tape channel runs off screen: %v", tc.Bounds.X+tc.Bounds.W)
	}
	if tc.Bounds.Y < ui.HUDTop {
		t.Errorf("tape channel reaches into the scene: Y=%v < %v", tc.Bounds.Y, ui.HUDTop)
	}

	sl, ok := g.UIManager.Get("tapes")
	if !ok {
		t.Fatal("no 'tapes' widget")
	}
	if ts := sl.(*ui.TapeSlots); ts.Bounds.X+ts.Bounds.W > ui.DividerX {
		t.Errorf("tape slots cross the divider: %v > %v", ts.Bounds.X+ts.Bounds.W, ui.DividerX)
	}

	// The dialogue box deliberately covers only the left region, so a tape can
	// comment alongside the conversation.
	db, ok := g.UIManager.Get("dialog")
	if !ok {
		t.Fatal("no 'dialog' widget")
	}
	if box := db.(*inkwell.DialogBox); box.Bounds.X+box.Bounds.W > ui.DividerX {
		t.Errorf("dialogue box would cover the tape channel: %v > %v",
			box.Bounds.X+box.Bounds.W, ui.DividerX)
	}
}

// The guard must be registered FIRST so it ticks LAST among the widgets —
// otherwise it would swallow HUD clicks.
func TestWidgetOrder(t *testing.T) {
	g := build(t)
	n := g.UIManager.Names()
	if len(n) == 0 || n[0] != "usewith_guard" {
		t.Errorf("usewith_guard is not registered first: %v", n)
	}
	if n[len(n)-1] != "cursor" {
		t.Errorf("cursor is not registered last: %v", n)
	}
}
