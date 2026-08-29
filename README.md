# Real World

Point & click adventure — game two of the Norman Arc. Engine:
[inkwell](../../engines/inkwell).

The player is **Paul**, a junk dealer investigating what happened to Norman from
the outside. His companion is **Dex**, a Personal Tape riding in slot one of
Paul's **Armilla**. The world, the terminology and the entity catalogue live in
the wiki (`services/wiki-pages/pages/projects/realworld` and
`.../neumatronic-universe`); this file records the game-specific UI decisions.

## Running

```bash
go run .                 # the alley scene (beat 3)
go run . -finale         # the finale — try the Nokia-punk theme switch
go run . -scene alley    # pick a starting scene
go test ./...            # content and layout checks, no window needed
```

`go.mod` carries a `replace` pointing at the neighbouring inkwell checkout so
the engine and the game move together. Drop that line for a release build.

| Control | |
|---|---|
| left click | run the selected verb |
| right click | verb coin (Look / Use / Talk / Take) |
| `F1` | toggle hotspot outlines |
| `SPACE` | during a cutscene: let the tape speak, if it offered |

Screenshots: `EBITEN_SCREENSHOT_KEY=q go run .`, then press `q` in the window.

## Layout

Internal resolution is **640×400** — the same 16:10 as the wiki's "320×200 VGA
look" brief and exactly 2× it, so art drawn at 320×200 upscales cleanly.
Integer scaling only; the window opens at 2×.

The doubling is forced by type, not taste. inkwell renders text through a fixed
**6×16** debug font that the domain cannot replace. At 320×200 a 16px glyph is
8% of the screen height: rows overlap, and a top bar sized from the wireframe
deck's proportions cannot hold a single line. At 640×400 the same glyph lands
at 4%, which is the proportion the design was drawn for.

```
     0                                     394│396                    639
   ┌────────────────────────────────────────────────────────────────────┐
 0 │ ALLEY — paused                              SRP: 1 tape            │ TopBar (20)
20 ├────────────────────────────────────────────────────────────────────┤
   │            scene — hotspots, characters, SpeechBubble              │ Scene (244)
264├─────────────────────────────────────────┬──────────────────────────┤
   │ Use hook on: floor grate                │ DEX                      │
   │ ┌───┬───┬───┬───┐   ┌───────┐┌────────┐ │ "The hook is a hand      │ HUD (136)
   │ └───┴───┴───┴───┘   │ 1 DEX ││ 2  —   │ │  shorter than the gap."  │
399└─────────────────────────────────────────┴──────────────────────────┘
     InventoryBar          TapeSlots           TapeChannel
```

Every vertical measurement derives from `ui.LineH` (glyph cell + leading), not
from the deck's ratios, so the layout cannot drift out of step with the font
again — `TestLayoutFitsTheFont` guards it. The tape channel comes out at 5 rows
of 38 columns, which is about two of Dex's remarks on screen at once.

The vertical divider at x=394 is the one proportion carried over from the
wireframe deck: 61.6% world, 38.4% tape channel. The deck's *vertical*
proportions are not carried over.

The dialogue box deliberately spans only the left region, so a tape can comment
*alongside* a conversation.

## The rule everything follows

Canon's central idea is the dual architecture: every machine has a digital side
and an ACP side, and the "agent-mediated focus" turns that into a UI paradigm —
one main text focus plus a side status panel, with the hand typing at the
digital side and the voice speaking to the ACP.

> **The hand holds the world. The voice speaks to the tape.**
>
> `WORLD` — mouse, the four verbs, inventory → Paul acts.
> `TAPE` — the `talk` verb and tape commentary → nothing happens in the world.
>
> There is always a visible, continuous separator between them, and the tape
> side never carries a button that reaches into the world.

Two consequences worth stating: there is **no text input anywhere** (the tape
channel is voice, opened with `talk`), and **colour carries the rule** — the
world and Paul speak in bone white, tapes in amber. If something is amber, a
tape said it.

There is also no "token" economy. The wireframe deck metered advice, but canon
has no such concept and the wiki describes Dex as commenting continuously. If
scarcity is ever wanted, the canon-native lever is the operator licence —
without one a Personal Tape runs in restricted mode, and Paul certainly has no
licence.

## Structure

```
main.go                  flags + inkwell.Run
internal/names/          entity names and world-state keys
internal/theme/          realworld-93 + nokia-punk
internal/world/          unsaved runtime state, custom actions, action pump
internal/ui/             HUD: layout, custom widgets, coloured text
internal/content/        one file per registered entity, by kind
internal/boot/           wiring
```

`world` deliberately knows nothing about the HUD or the content — both build on
it, so the reverse would be a cycle:

```
names ← theme ← world ← ui
                  ↑      ↑
               content ──┴── boot ← main
```

Content is one registered entity per file, grouped by kind:

```
internal/content/
  asset/      alley_background.go
  character/  paul.go  dex.go  mystery_tape.go  support_tape.go
  item/       noodle_letter.go  black_market_armilla.go  mystery_tape.go
  dialog/     dex_talk.go  mystery_tape_silent.go
  script/     tape_insert.go  awakening_finale.go
  scene/      alley.go
```

Adding a scene means adding `scene/<name>.go` and one line in `scene/scene.go`.
Nothing else moves.

## Engine workarounds

Five inkwell limits turned up during implementation that the engine README does
not mention. All five are worked around on the domain side; each is a candidate
for a small engine change.

1. **`drawText` discards colour.** In `asset.text.go` the colour argument is
   `_ = c` and rendering goes through `ebitenutil.DebugPrintAt`, which only
   draws white. Every text colour in `Theme` is therefore inert. *Workaround:*
   `ui/text.go` renders onto a scratch image and blits it tinted with
   `ColorScale`. The custom widgets colour correctly; the built-ins
   (`StatusLine`, `DialogBox`, `TopBar`, `InventoryBar`) are still white.
   *Fix:* move `drawText` to `text/v2` — no call site would change.

2. **`queueAction` is unexported**, so a domain widget cannot start an action.
   *Workaround:* the pump in `world/action.go` drives its own `Runner` through
   the exported `inkwell.Ctx`. *Caveat:* it runs alongside the engine's script
   runner, not instead of it.

3. **The `"Nem ehhez."` flash is hardcoded** in `core.engine.go` for an
   item/hotspot pair with no `OnUseWith`, and cannot be replaced from the
   domain. *Workaround:* `UseWithGuard` — widgets tick before the engine's
   `handleSceneInput` and can consume the click, so unauthored pairs fail in
   character instead, escalating on repeats.

4. **No exported `CurrentScene()`**, but the pump needs `Ctx.Scene`.
   *Workaround:* `world.EnterScene` is the first step of every scene's
   `OnEnter`; the domain tracks it.

5. **Widgets have no `Visible` field and the `Manager` cannot unregister**, so
   the built-in HUD cannot be hidden during a cutscene. *Workaround:* the
   `gate` wrapper in `ui/ui.go` forwards `Tick`/`Draw`/`BlocksClickAt` only
   while the HUD is visible.

6. **`Run` hardcodes a 4× window** (`core.dsl.go`), which at 640×400 would be
   2560×1600 — bigger than most laptop screens. *Workaround:* the `windowSizer`
   widget resizes once on the first tick, since Run sets the size before
   entering the loop.

Also: the inkwell README gives the module path as
`git.teletypegames.org/games/inkwell`; the real path per its `go.mod` is
`git.teletypegames.org/engines/inkwell`.

## Known gaps

- **No art.** Everything is an inkwell placeholder.
- **Spent dialogue choices are hidden, not struck through.** `DialogueChoice.Once`
  hides them; the deck wanted them struck through but visible, so the list
  becomes a memory of what you already tried. Needs a thin override over
  `DialogBox`.
- **Only beat 3 exists.** The alley is a vertical slice; `boot.opening` sets the
  police-tip flag that beat 2 will eventually set.
- **The render has only been inspected once, by the author of this repo.**
  Geometry is covered by tests and a layout dump, but this environment cannot
  take a screenshot (both synthetic keystrokes and screen capture are blocked by
  macOS privacy permissions), so every visual judgement has to come from you.
