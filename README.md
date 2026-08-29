# Real World

Point & click adventure — game two of the Norman Arc. Engine:
[inkwell](../../engines/inkwell).

The player is **Paul**, a junk dealer investigating what happened to Norman from
the outside. His companion is **Dex**, a Personal Tape riding in slot one of
Paul's **Armilla**. The world, the terminology and the entity catalogue live in
the wiki (`services/wiki-pages/pages/projects/realworld` and
`.../neumatronic-universe`); this file records the game-specific UI decisions.

## Running

inkwell is consumed as an ordinary Go module from the forge. Prerequisites:
Go 1.26+ and SSH access to `git.teletypegames.org`.

```bash
git clone ssh://git@git.teletypegames.org:2222/games/realworld
cd realworld
export GOPRIVATE=git.teletypegames.org   # the Makefile sets this for its own targets
go mod tidy
go run .
```

```bash
go run .                   # the alley (beat 3)
go run . -finale           # the finale — try the Nokia-punk theme switch
go run . -screen selector  # the location selector: the map the city hangs off
go run . -screen hospital  # any screen in the deck, e.g. to look at the art
```

Screen names are the constants in `names.manager.go` (`selector`, `paul_shop`,
`noodle_house`, `alley`, `norman_apartment`, `police_station`, `hackerspace`,
`ice_cream_shop`, `trinket_shop`, `small_restaurant`, `secret_club`,
`bbs_terminal`, `street`, `hospital`, `server_farm`, `secret_lab`,
`rooftop_hideout`, `public_bbs`, `bvk_branch`, `curator_shop`, `scrap_market`,
`samizdat_press`, `showroom`, `columbarium`).

To work against a local inkwell checkout, add a `replace` line temporarily —
and take it out before committing:

```bash
go mod edit -replace git.teletypegames.org/engines/inkwell=../../engines/inkwell
go mod edit -dropreplace git.teletypegames.org/engines/inkwell
```

## Building

```bash
make build                    # native binary into bin/
make wasm                     # dist/game.wasm + wasm_exec.js
make export VERSION=0.1       # zipped HTML/WASM bundle
make binaries VERSION=0.1     # win-x86, win-x64, linux-x64 zips
make clean
```

CI is Woodpecker; `.woodpecker.yaml` only names the platform, and the update
server serves the actual pipeline (`GET /build/config?platform=ebitengine`).
Release metadata lives in `metadata.json`.

| Control | |
|---|---|
| left click | run the selected verb |
| right click | verb coin (Look / Use / Talk / Take) |
| `←` `→` | previous / next screen, in file order, wrapping |
| `F1` | toggle hotspot outlines |
| `SPACE` | during a cutscene: let the tape speak, if it offered |

The screens are connected to each other (see **The map** below); the arrow keys
are a reviewing tool on top of that, walking every screen in turn. The walk is
inert during a cutscene, a menu, or a stopped world, so it can never cut
across an authored beat.

Screenshots: `EBITEN_SCREENSHOT_KEY=q go run .`, then press `q` in the window.

## Layout

Internal resolution is **640×380**: the size of one painted background.
Integer scaling only; the window opens at 2×.

That size is not a taste call. inkwell scales a scene background over the
*whole* screen (`core.engine.go`), not over the region the HUD leaves free — so
at any screen size other than the art's own, every painting is squashed to fit.
The deck is drawn at 640×380, so the screen is 640×380 and the art lands pixel
for pixel. `TestSceneBackgroundsResolve` keeps the two in step: a background of
any other size fails the build.

It still gives the wiki's "320×200 VGA look" brief what it was after — a low,
wide VGA frame at an exact 2× of a 320-wide canvas — and it still keeps the
engine's fixed **6×16** debug font readable, at 4% of the screen height. (At
320×190 it would be 8%: rows overlap and the top bar cannot hold a line.)

The HUD is therefore an opaque strip laid over the foot of the painting, and it
is measured from the bottom and kept as short as its contents allow — a status
row, then the two rows of inventory slots, which are the tallest thing in it.
Every pixel it gives back is a pixel of art.

```
     0                                     394│396                    639
   ┌────────────────────────────────────────────────────────────────────┐
 0 │ ALLEY — paused                              SRP: 1 tape            │ TopBar (20)
20 ├────────────────────────────────────────────────────────────────────┤
   │            scene — hotspots, characters, SpeechBubble              │ Scene (242)
262├─────────────────────────────────────────┬──────────────────────────┤
   │ Use hook on: floor grate                │ DEX                      │
   │ ┌───┬───┬───┬───┐   ┌───────┐┌────────┐ │ "The hook is a hand      │ HUD (118)
   │ └───┴───┴───┴───┘   │ 1 DEX ││ 2  —   │ │  shorter than the gap."  │
379└─────────────────────────────────────────┴──────────────────────────┘
     InventoryBar          TapeSlots           TapeChannel
```

Every vertical measurement derives from `ui.LineH` (glyph cell + leading) and
from `ui.ScreenH`, not from the wireframe deck's ratios, so the layout cannot
drift out of step with the font — or with the art — again;
`TestLayoutFitsTheFont` guards it. The tape channel comes out at 4 rows of 38
columns, which is about one and a half of Dex's remarks on screen at once.

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

The game is one flat package, `inc`. There are no subdirectories: a file's
name says where it belongs, in the form `[category].[name].go`. The category is
always singular, and `[category].manager.go` is the file that ties that category
together — its manager, its shared types, its hand-off to the engine.

```
main.go                    flags + inkwell.Run
inc/manager.manager.go     Manager[T], the generic registry
inc/manager.interface.go   ManagerInterface, the contract every manager keeps
inc/names.manager.go       entity names and world-state keys
inc/theme.*.go             realworld-93 + nokia-punk
inc/world.*.go             unsaved runtime state, custom actions, action pump
inc/ui.*.go                HUD: layout, custom widgets, coloured text
inc/<kind>.<name>.go       one file per registered entity, by kind
inc/boot.manager.go        wiring
```

Nothing is enforced by the compiler any more, so the layering is a rule the
code keeps by hand: the world holds no opinion about the HUD or the content, and
both build on it, never the other way round.

```
names ← theme ← world ← ui
                  ↑      ↑
               content ──┴── boot ← main
```

Content is one registered entity per file, and the category prefix groups them
the way directories used to:

```
background.*.go  one file per screen: background.paul_shop.go, …
character.*.go   character.paul.go  character.dex.go  character.mystery_tape.go
item.*.go        item.noodle_letter.go  item.black_market_armilla.go  …
dialog.*.go      dialog.dex_talk.go  dialog.mystery_tape_silent.go
script.*.go      script.tape_insert.go  script.awakening_finale.go
screen.*.go      one file per screen: screen.alley.go, … + screen.exit.go
```

Adding a screen means adding `screen.<name>.go` and `background.<name>.go`.
Nothing else moves.

Every category owns a manager, and they are all the same generic type,
`Manager[T]` — one slice in registration order, one `map[string]int` beside it,
so a lookup by name is a map hit rather than a scan. The entity types are
aliases of engine structs and cannot carry methods, so the manager is handed a
function that reads the name instead. A category's manager file is an alias and
one line:

```go
type Character = inkwell.Character

var CharacterManager = NewManager(func(entity Character) string { return entity.Name })
```

They all keep the same contract, asserted in `manager.interface.go`:

```go
type ManagerInterface[T any] interface {
	Register(entity T)
	GetByName(name string) (T, bool)
	GetAll() []T
}
```

An entity file is a literal that hands itself over in an `init()`:

```go
func init() {
	BackgroundManager.Register(Background{
		Name: BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	})
}
```

So adding an entity is adding a file, and there is no second list to keep in
step. The price is that registration order is file-name order: the arrow keys
walk the screens alphabetically rather than in the concept-art deck's order.
That was a deliberate trade — the deck order was a list that had to be kept in
step with the files by hand, and the deck is a thing to look at, not a thing to
play through.

The game's own catalogue is readable without going through the engine:
`ScreenManager.GetByName("alley")` answers before a single scene has been handed
over. `registerContent` is where the hand-off happens, one `registerAll` call
per category.

There is one world, and it is a package-level singleton: `World`. Nothing takes
a `*world` parameter and no widget holds a back-reference, which is what lets a
script file be a literal — the tape-insert script closes over `World` rather
than over a parameter someone would have had to thread to it.

**On the word "screen".** inkwell's entity is called a `Scene`, and that is the
type every `screen.*.go` file returns — but the wiki, the concept-art deck and
the beat tables all count *screens*, so the category, the files and the flag say
screen. "Scene" survives only where the code is talking to the engine
(`EnterScene`, `SceneManager`).

Art lives in `assets/bg/` and is embedded into the binary (`main.go`), because
js/wasm has no OS filesystem and `make binaries` packages the executable alone.

## The map

Where a screen leads is recorded **in that screen**, in its own file, and it
comes from the wiki's location graph
(`projects/realworld/story#helyszín-gráf-vázlat`). That graph has two kinds of
edge and so does the code:

- **Physical adjacency** — the alley is behind Noodle's house, the eating place
  is next door to the trinket shop, the club is through its back room, the roof
  is up Norman's stairs. Declared with `to:` in the screen's `exits`.
- **The selector** — the wiki centres its map on a *Helyszínválasztó*, "nem
  valódi helyszín, hanem a menü-képernyő": a screen every main location connects
  to both ways. A screen marks itself `onSelector: true`; `screen/selector.go`
  builds the other half of each of those edges, so the list of locations on the
  map is never written down twice.

The wiki's conditional arrows survive too: `needs:` holds the flag an exit waits
for, which is how *Bolt →|beengedés| Klub* is expressed — the way into the club
is visible and named from the first visit, and stays shut until the shopkeeper
has been shown the underwater sun.

Until the doors are measured on the paintings, an exit is a strip along one edge
of the picture (`sideLeft`, `sideRight`, `sideBack`, `sideNear`) — the
convention every 1990s point & click used, and one field to re-aim later. A
screen reached from several rooms has no fixed way out: the BBS terminal is the
same terminal from the club, the flat or the roof, so its exit is `back`.

The graph stays machine-readable: exit hotspots are named `exit:<target>`, so
the connections can be read straight back out of the registered screens.

## Engine workarounds

Five inkwell limits turned up during implementation that the engine README does
not mention. All five are worked around on the domain side; each is a candidate
for a small engine change.

1. **`drawText` discards colour.** In `asset.text.go` the colour argument is
   `_ = c` and rendering goes through `ebitenutil.DebugPrintAt`, which only
   draws white. Every text colour in `Theme` is therefore inert. *Workaround:*
   `ui.text.go` renders onto a scratch image and blits it tinted with
   `ColorScale`. The custom widgets colour correctly; the built-ins
   (`StatusLine`, `DialogBox`, `TopBar`, `InventoryBar`) are still white.
   *Fix:* move `drawText` to `text/v2` — no call site would change.

2. **`queueAction` is unexported**, so a domain widget cannot start an action.
   *Workaround:* the pump in `world.action.go` drives its own `Runner` through
   the exported `inkwell.Ctx`. *Caveat:* it runs alongside the engine's script
   runner, not instead of it.

3. **The `"Nem ehhez."` flash is hardcoded** in `core.engine.go` for an
   item/hotspot pair with no `OnUseWith`, and cannot be replaced from the
   domain. *Workaround:* `UseWithGuard` — widgets tick before the engine's
   `handleSceneInput` and can consume the click, so unauthored pairs fail in
   character instead, escalating on repeats.

4. **No exported `CurrentScene()`**, but the pump needs `Ctx.Scene`.
   *Workaround:* `EnterScene` is the first step of every scene's
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

- **The HUD still covers the foot of every painting.** The art is no longer
  stretched, but the strip is opaque and 118px of the 380 sit on top of the
  picture. Either the deck gets recut so nothing that matters lives in its
  bottom third, or the HUD gets a translucent panel treatment.
- **The screens are empty.** 22 backgrounds are in and they are wired to each
  other, but only the alley has hotspots, NPCs and dialogue; everywhere else
  there is nothing to do but leave again.
- **The selector has no art.** The wiki's Helyszínválasztó is the centre of the
  map and the deck does not include it, so it renders as a placeholder with one
  labelled pin per location on a plain grid.
- **Two screens are missing from the deck.** The alley (the one authored beat)
  and deck 05, Norman's workplace — which is in the wiki's catalogue but has no
  screen file yet, so the workplace thread of the story has nowhere to happen.
- **Spent dialogue choices are hidden, not struck through.** `DialogueChoice.Once`
  hides them; the deck wanted them struck through but visible, so the list
  becomes a memory of what you already tried. Needs a thin override over
  `DialogBox`.
- **Only beat 3 exists.** The alley is a vertical slice; `bootOpening` sets the
  police-tip flag that beat 2 will eventually set.
- **The render has only been inspected once, by the author of this repo.**
  Geometry is derived from one font cell and one screen size, but this
  environment cannot take a screenshot (both synthetic keystrokes and screen
  capture are blocked by macOS privacy permissions), so every visual judgement
  has to come from you.
