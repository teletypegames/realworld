package inc

import (
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"
)

// The mode must not leak into saved state: inkwell saves State.Vars but drops
// the in-flight script, so a save taken mid-cutscene would otherwise reload
// into a letterboxed, HUD-less game with nothing running.
func TestModeIsNotSavedState(t *testing.T) {
	w := newWorld(inkwell.NewGame("t", 320, 200))

	w.SetMode(ModeCutscene)
	for _, k := range []string{"mode", "mode"} {
		if v := w.G.State.Var(k); v != nil {
			t.Errorf("mode leaked into State.Vars (%q = %v)", k, v)
		}
	}

	// Slot 2, by contrast, is real game state and must be saved.
	w.SetSlot2(ItemMysteryTape)
	if w.G.State.Var(VarSlot2) != ItemMysteryTape {
		t.Error("slot 2 is not in State.Vars")
	}
	if got := w.Slot2(); got != ItemMysteryTape {
		t.Errorf("Slot2() = %q", got)
	}
}

// The widget marks a tape as pending; the script consumes it — once.
func TestPendingIsConsumedOnce(t *testing.T) {
	w := newWorld(inkwell.NewGame("t", 320, 200))
	w.SetPending(ItemMysteryTape)
	if got := w.TakePending(); got != ItemMysteryTape {
		t.Fatalf("TakePending() = %q", got)
	}
	if got := w.TakePending(); got != "" {
		t.Errorf("pending was not cleared: %q", got)
	}
}
