package world_test

import (
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"

	"realworld/internal/names"
	"realworld/internal/world"
)

// The mode must not leak into saved state: inkwell saves State.Vars but drops
// the in-flight script, so a save taken mid-cutscene would otherwise reload
// into a letterboxed, HUD-less game with nothing running.
func TestModeIsNotSavedState(t *testing.T) {
	w := world.New(inkwell.NewGame("t", 320, 200))

	w.SetMode(world.ModeCutscene)
	for _, k := range []string{"ui.mode", "mode"} {
		if v := w.G.State.Var(k); v != nil {
			t.Errorf("mode leaked into State.Vars (%q = %v)", k, v)
		}
	}

	// Slot 2, by contrast, is real game state and must be saved.
	w.SetSlot2(names.ItemMysteryTape)
	if w.G.State.Var(names.VarSlot2) != names.ItemMysteryTape {
		t.Error("slot 2 is not in State.Vars")
	}
	if got := w.Slot2(); got != names.ItemMysteryTape {
		t.Errorf("Slot2() = %q", got)
	}
}

// The widget marks a tape as pending; the script consumes it — once.
func TestPendingIsConsumedOnce(t *testing.T) {
	w := world.New(inkwell.NewGame("t", 320, 200))
	w.SetPending(names.ItemMysteryTape)
	if got := w.TakePending(); got != names.ItemMysteryTape {
		t.Fatalf("TakePending() = %q", got)
	}
	if got := w.TakePending(); got != "" {
		t.Errorf("pending was not cleared: %q", got)
	}
}
