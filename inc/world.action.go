package inc

// The game's own actions, and the domain action pump.
//
// inkwell's queueAction is unexported, so domain widgets (TapeSlots) cannot put
// an action on the engine's queue — the built-in widgets only manage it because
// they live inside the package. inkwell.Ctx, however, is exported with exported
// fields, so the domain can drive a Runner itself. That is what the pump does:
// it runs HUD-originated actions (starting a dialogue, a tape line) alongside
// the engine's own script runner. See README, "Engine workarounds".

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

// ----- pump -------------------------------------------------------------

// Do queues a HUD-originated action.
func (w *World) Do(a inkwell.Action) {
	if a != nil {
		w.queue = append(w.queue, a)
	}
}

// PumpTick advances the pump. The ActionPump widget calls it every frame.
func (w *World) PumpTick(dt float64) {
	// The TopBar's left section: location plus a single word for a stopped
	// world. No pause overlay, no blur — one word says time has stopped.
	if w.top != nil {
		title := w.sceneTitle()
		if w.paused {
			title += " — paused"
		}
		w.top.LeftText = title
	}

	if w.running == nil {
		if len(w.queue) == 0 {
			return
		}
		w.running = w.queue[0].Start()
		w.queue = w.queue[1:]
	}
	ctx := &inkwell.Ctx{Game: w.G, DT: dt}
	if s, ok := w.G.SceneManager.Get(w.scene); ok {
		ctx.Scene = &s
	}
	if w.running.Tick(ctx) != inkwell.StatusRunning {
		w.running = nil
	}
}

func (w *World) sceneTitle() string {
	s, ok := w.G.SceneManager.Get(w.scene)
	if !ok {
		return ""
	}
	if s.Title != "" {
		return s.Title
	}
	return s.Name
}

// ActionPump is an invisible widget that advances the pump. The HUD registers
// it; the world itself knows nothing about the widget tree.
type ActionPump struct {
	Name string
	W    *World
}

func (p *ActionPump) GetName() string                            { return p.Name }
func (p *ActionPump) Tick(ctx *inkwell.UICtx)                    { p.W.PumpTick(ctx.DT) }
func (p *ActionPump) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

// ----- helpers ----------------------------------------------------------

type fnAction struct{ fn func(*inkwell.Ctx) }

func (a *fnAction) Start() inkwell.Runner { return &fnRunner{spec: a} }

type fnRunner struct{ spec *fnAction }

func (r *fnRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	r.spec.fn(ctx)
	return inkwell.StatusDone
}

// Fn turns a function into a one-shot action that completes immediately.
func Fn(f func(*inkwell.Ctx)) inkwell.Action { return &fnAction{fn: f} }

// UseTheme switches the theme at runtime. This is the single line that makes
// the Nokia-punk texture break in during the finale. inkwell saves ActiveTheme,
// so the switch survives a save without needing a flag of its own.
func UseTheme(name string) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) { ctx.Game.UseTheme(name) })
}

// SetMode switches the HUD mode (play / cutscene / menu).
func SetMode(w *World, m Mode) inkwell.Action {
	return Fn(func(*inkwell.Ctx) { w.SetMode(m) })
}

// EnterScene tells the domain which scene we are in. inkwell exposes no
// CurrentScene accessor, and the pump needs Ctx.Scene.
func EnterScene(w *World, name string) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		if name != w.scene {
			w.prev = w.scene
		}
		w.scene = name
	})
}

// Back returns to the screen we came from.
//
// Most exits name their destination, because a door leads where it leads. A
// screen that is reached from several different rooms cannot: the BBS terminal
// is the same terminal whether you walked to it from the club, from Norman's
// flat or from the roof, so its way out is "back", not a place.
func Back(w *World) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if w.prev == "" {
			return
		}
		inkwell.GoTo(w.prev).Start().Tick(ctx)
	})
}

// ----- TapeSay ----------------------------------------------------------

// TapeSay is a line spoken by a tape.
//
// It is deliberately not inkwell.Say: Say puts a SpeechBubble above the
// speaker's head, but tapes have no body in the scene — they speak into your
// ear. This is the two-channel rule enforced in code: a tape can only reach the
// tape channel, never the scene.
func TapeSay(speaker, text string) inkwell.Action {
	return &tapeSayAction{speaker: speaker, text: text}
}

type tapeSayAction struct{ speaker, text string }

func (a *tapeSayAction) Start() inkwell.Runner { return &tapeSayRunner{spec: a} }

type tapeSayRunner struct {
	spec     *tapeSayAction
	started  bool
	elapsed  float64
	duration float64
}

func (r *tapeSayRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	if !r.started {
		r.started = true
		// Same pacing as Say, so the two breathe together inside a Seq.
		r.duration = 1.2 + float64(len([]rune(r.spec.text)))*0.05
		ctx.Game.LogResponse(r.spec.speaker, r.spec.text)
	}
	r.elapsed += ctx.DT
	if r.elapsed >= r.duration {
		return inkwell.StatusDone
	}
	return inkwell.StatusRunning
}

// ----- TapeOffer --------------------------------------------------------

// TapeOffer offers a tape line during a cutscene: the Letterbox shows that a
// tape would speak, but it only speaks if the player asks. A tape never
// interrupts — the floor is the player's.
func TapeOffer(w *World, line inkwell.Action) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		w.tapeLine = line
		w.tapeWaiting = true
	})
}

// TakeTapeOffer queues the offered line and clears the prompt.
func (w *World) TakeTapeOffer() {
	if !w.tapeWaiting {
		return
	}
	w.tapeWaiting = false
	line := w.tapeLine
	w.tapeLine = nil
	w.Do(line)
}

// ----- Paused -----------------------------------------------------------

// Paused marks the world as stopped for the duration of the wrapped action.
// inkwell exposes no dialogueActive accessor, so we mark it on the content
// side: every dialogue start is wrapped in this.
func Paused(w *World, inner inkwell.Action) inkwell.Action {
	return &pausedAction{w: w, inner: inner}
}

type pausedAction struct {
	w     *World
	inner inkwell.Action
}

func (a *pausedAction) Start() inkwell.Runner {
	return &pausedRunner{spec: a, inner: a.inner.Start()}
}

type pausedRunner struct {
	spec    *pausedAction
	inner   inkwell.Runner
	started bool
}

func (r *pausedRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	if !r.started {
		r.started = true
		r.spec.w.paused = true
	}
	s := r.inner.Tick(ctx)
	if s != inkwell.StatusRunning {
		r.spec.w.paused = false
	}
	return s
}
