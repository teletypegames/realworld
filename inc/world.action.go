package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

func (w *World) Do(a inkwell.Action) {
	if a != nil {
		w.queue = append(w.queue, a)
	}
}

func (w *World) PumpTick(dt float64) {

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
	ctx := &inkwell.Ctx{
		Game: w.G,
		DT:   dt,
	}
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

type ActionPump struct {
	Name string
	W    *World
}

func (p *ActionPump) GetName() string                            { return p.Name }
func (p *ActionPump) Tick(ctx *inkwell.UICtx)                    { p.W.PumpTick(ctx.DT) }
func (p *ActionPump) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

type fnAction struct{ fn func(*inkwell.Ctx) }

func (a *fnAction) Start() inkwell.Runner {
	return &fnRunner{
		spec: a,
	}
}

type fnRunner struct{ spec *fnAction }

func (r *fnRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	r.spec.fn(ctx)
	return inkwell.StatusDone
}

func Fn(f func(*inkwell.Ctx)) inkwell.Action {
	return &fnAction{
		fn: f,
	}
}

func UseTheme(name string) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) { ctx.Game.UseTheme(name) })
}

func SetMode(w *World, m Mode) inkwell.Action {
	return Fn(func(*inkwell.Ctx) { w.SetMode(m) })
}

func EnterScene(w *World, name string) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		if name != w.scene {
			w.prev = w.scene
		}
		w.scene = name
	})
}

func Back(w *World) inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if w.prev == "" {
			return
		}
		inkwell.GoTo(w.prev).Start().Tick(ctx)
	})
}

func TapeSay(speaker, text string) inkwell.Action {
	return &tapeSayAction{
		speaker: speaker,
		text:    text,
	}
}

type tapeSayAction struct{ speaker, text string }

func (a *tapeSayAction) Start() inkwell.Runner {
	return &tapeSayRunner{
		spec: a,
	}
}

type tapeSayRunner struct {
	spec     *tapeSayAction
	started  bool
	elapsed  float64
	duration float64
}

func (r *tapeSayRunner) Tick(ctx *inkwell.Ctx) inkwell.Status {
	if !r.started {
		r.started = true

		r.duration = 1.2 + float64(len([]rune(r.spec.text)))*0.05
		ctx.Game.LogResponse(r.spec.speaker, r.spec.text)
	}
	r.elapsed += ctx.DT
	if r.elapsed >= r.duration {
		return inkwell.StatusDone
	}
	return inkwell.StatusRunning
}

func TapeOffer(w *World, line inkwell.Action) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		w.tapeLine = line
		w.tapeWaiting = true
	})
}

func (w *World) TakeTapeOffer() {
	if !w.tapeWaiting {
		return
	}
	w.tapeWaiting = false
	line := w.tapeLine
	w.tapeLine = nil
	w.Do(line)
}

func Paused(w *World, inner inkwell.Action) inkwell.Action {
	return &pausedAction{
		w:     w,
		inner: inner,
	}
}

type pausedAction struct {
	w     *World
	inner inkwell.Action
}

func (a *pausedAction) Start() inkwell.Runner {
	return &pausedRunner{
		spec:  a,
		inner: a.inner.Start(),
	}
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
