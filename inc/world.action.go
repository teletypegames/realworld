package inc

import (
	inkwell "git.teletypegames.org/engines/inkwell"
	"github.com/hajimehoshi/ebiten/v2"
)

func (w *world) Do(a inkwell.Action) {
	if a != nil {
		w.queue = append(w.queue, a)
	}
}

func (w *world) PumpTick(dt float64) {

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

func (w *world) sceneTitle() string {
	s, ok := w.G.SceneManager.Get(w.scene)
	if !ok {
		return ""
	}
	if s.Title != "" {
		return s.Title
	}
	return s.Name
}

type actionPump struct {
	Name string
}

func (p *actionPump) GetName() string                            { return p.Name }
func (p *actionPump) Tick(ctx *inkwell.UICtx)                    { World.PumpTick(ctx.DT) }
func (p *actionPump) Draw(dst *ebiten.Image, ctx *inkwell.UICtx) {}

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

func SetMode(m Mode) inkwell.Action {
	return Fn(func(*inkwell.Ctx) { World.SetMode(m) })
}

func EnterScene(name string) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		if name != World.scene {
			World.prev = World.scene
		}
		World.scene = name
	})
}

func Back() inkwell.Action {
	return Fn(func(ctx *inkwell.Ctx) {
		if World.prev == "" {
			return
		}
		inkwell.GoTo(World.prev).Start().Tick(ctx)
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

func TapeOffer(line inkwell.Action) inkwell.Action {
	return Fn(func(*inkwell.Ctx) {
		World.tapeLine = line
		World.tapeWaiting = true
	})
}

func (w *world) TakeTapeOffer() {
	if !w.tapeWaiting {
		return
	}
	w.tapeWaiting = false
	line := w.tapeLine
	w.tapeLine = nil
	w.Do(line)
}

func Paused(inner inkwell.Action) inkwell.Action {
	return &pausedAction{
		inner: inner,
	}
}

type pausedAction struct {
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
		World.paused = true
	}
	s := r.inner.Tick(ctx)
	if s != inkwell.StatusRunning {
		World.paused = false
	}
	return s
}
