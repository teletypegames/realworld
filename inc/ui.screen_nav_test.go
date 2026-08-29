package inc

import (
	"testing"

	inkwell "git.teletypegames.org/engines/inkwell"
)

func TestScreenNavWrapsBothWays(t *testing.T) {
	g := inkwell.NewGame("t", ScreenW, ScreenH)
	for _, n := range []string{"a", "b", "c"} {
		g.SceneManager.Register(inkwell.Scene{Name: n, Background: "bg"})
	}
	w := newWorld(g)
	nav := &screenNav{Name: "screen_nav", W: w}
	at := func(name string) { EnterScene(w, name).Start().Tick(&inkwell.Ctx{Game: g}) }

	for _, tc := range []struct {
		from string
		step int
		want string
	}{
		{"a", 1, "b"}, {"b", 1, "c"}, {"c", 1, "a"},
		{"c", -1, "b"}, {"b", -1, "a"}, {"a", -1, "c"},
	} {
		at(tc.from)
		if got := nav.neighbour(g, tc.step); got != tc.want {
			t.Errorf("from %q step %+d = %q, want %q", tc.from, tc.step, got, tc.want)
		}
	}
}
