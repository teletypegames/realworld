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
go run . -scene selector   # the location selector: the map the city hangs off
go run . -scene hospital   # any scene in the deck, e.g. to look at the art
```

Scene names are the constants in `names.manager.go` (`selector`, `paul_shop`,
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
| `←` `→` | previous / next scene, in file order, wrapping |
| `F1` | toggle hotspot outlines |
| `SPACE` | during a cutscene: let the tape speak, if it offered |

The scenes are connected to each other (see **The map** below); the arrow keys
are a reviewing tool on top of that, walking every scene in turn. The walk is
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

Every vertical measurement derives from `LineH` (glyph cell + leading) and from
`ScreenH`, both in `widget/widget.manager.go`, not from the wireframe deck's ratios, so
the layout cannot drift out of step with the font — or with the art — again;
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

One category, one package, one directory. A file still says which entity it
holds in its own name — `[category].[name].go`, the prefix repeated inside the
directory that already carries it, because a file called `alley.go` tells you
nothing in a list of open editor tabs.

```
main.go                          flags + inkwell.Run
inc/constants.go                 every name and state key; imports nothing
inc/boot/                        New(Opts): the hand-off to the engine
inc/theme/                       the Theme alias, RGB, realworld-93, nokia-punk
inc/world/                       the run state and the custom actions
inc/tape/                        the Tape entity and the lookups over it
inc/widget/                      HUD layout, the visibility conditions, 14 widgets
inc/scene/  inc/background/      one file per location, twice
inc/item/  inc/character/  inc/dialog/  inc/script/
```

`inc` itself holds nothing but constants, which is why every package can import
it and it can import none of them. The consequence is that the assembly moved
down rather than up: `New` lives in `inc/boot`, because a package cannot be
imported by what it imports.

The layering is no longer a rule kept by hand — it is the import graph, and the
compiler rejects the arrow that points the wrong way:

```
inc ← theme ← world ← tape ← widget ← boot ← main
        ↑        ↑                      ↑
        └── content ───────────────────-┘
```

Adding a scene means adding `inc/scene/scene.<name>.go` and
`inc/background/background.<name>.go`. Nothing else moves.

Every category owns a manager, and they are all the engine's own
`inkwell.Manager[T]` — the same registry the `*Game` hangs its content off, kept
in registration order and addressed by name. There is no second registry type
here: the entity types are aliases of engine structs, so they already satisfy
`inkwell.Named` and need no help reading their own name. Each one lives in its
own package's manager file, next to the alias it holds, and it needs no prefix
because the package already is one:

```go
// inc/character/character.manager.go
type Character = inkwell.Character

var Manager = inkwell.NewManager[Character]()
```

An entity file is a literal that hands itself over in an `init()`:

```go
package background

func init() {
	Manager.Register(Background{
		Name: inc.BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	})
}
```

The `init()` only runs if something imports the package, so a package whose
entities nobody names by symbol needs a blank import in `inc/boot`. Exactly one
does — `tape`, everything else is imported for its `Manager` — and the list is
per package, not per entity, so adding a file is still adding a file.

Widgets go the same way, the engine's own included: `widget/widget.cursor.go` registers
an `inkwell.Cursor` exactly as `background/background.street.go` registers an image. What
they cannot take from alphabetical order is their place in the stack, so each
widget names a **layer** — `LayerScene`, `LayerPanel`, `LayerHUD`, `LayerSpeech`,
`LayerDialog`, `LayerMenu`, `LayerCurtain`, `LayerCursor` — and inkwell draws
from the bottom layer up, ticks from the top layer down, and lets registration
order decide only ties inside one layer. That is why the HUD frame comes out
under the tape slots and the cursor over everything, whatever the file names
happen to be. The layer was added to inkwell for this; the alternative was a
registration list, which is the thing this package spent its life deleting.

The second thing a widget declares is when it is on screen at all: `When` is an
`inkwell.Condition` over game state, and the HUD's is `whenPlaying`, which reads
`VarMode` out of the engine's `State`. Hiding the HUD for a cutscene is
therefore one word per widget rather than a wrapper around each of them and an
`if` at the top of every `Draw`.

Tapes are the one entity type this game invents. `Tape` names the character
whose voice it is, the cassette that carries it and the dialogue it plays, one
file each — what a tape *sounds* like stays on the character, as
`Character.Label` and `Character.Voice`, because that is a fact about a speaker.
A tape's line is `inkwell.Say` like anybody else's; the engine sends it to the
log instead of a speech bubble because the character says so.

So adding an entity is adding a file, and there is no second list to keep in
step. The price is that registration order is file-name order: the arrow keys
walk the scenes alphabetically rather than in the concept-art deck's order.
That was a deliberate trade — the deck order was a list that had to be kept in
step with the files by hand, and the deck is a thing to look at, not a thing to
play through.

The game's own catalogue is readable without going through the engine:
`scene.Manager.Get("alley")` answers long before there is a `*Game` to ask. And
when the game is built, nothing is copied into it — `New` assigns our managers
onto the `*Game` in place of the empty ones `NewGame` made, so engine and game
share one registry per category instead of keeping two of them in step. Themes
are the exception: `NewGame` seeds its theme manager with four presets, so ours
are added to that set rather than replacing it.

The defaults a scene may leave out are not written into it at all any more:
`g.Player` names the character a scene with no actors of its own receives, at
his `Start`, and `g.Walkboxes` is the floor a scene without one walks on. Both
are fields on the `*Game`, so a scene file that says nothing about either means
"the usual", and the selector opts out by declaring both empty. What is still
written back before the hand-off is the derived half of the map:
`scene.FillSelectorPins` reads the exit graph backwards and `Set`s the selector.

There is one world, and now the package *is* the singleton: `world.Do(…)`,
`world.Slot2()`, `world.SetPending(…)`. Nothing takes a world parameter and no
widget holds a back-reference, which is what lets a script file be a literal —
the tape-insert script closes over the `world` package rather than over a
parameter someone would have had to thread to it. What little it still keeps is
in package-level variables that only `world.Attach` may reset; the mode and the
top bar's note are not among them, because those are `State` vars the engine
itself can see.

**On the word "scene".** The wiki, the concept-art deck and the beat tables all
count *screens*, and this code used to as well: it carried its own `Screen`
struct, because inkwell's `Scene` could not hold the things a screen needs —
its exits above all. Those went into the engine instead (inkwell `Exit`,
`Game.SceneRect`, `CurrentScene`/`PreviousScene`, `Back()`), and with the gap
closed there was nothing left for a second type to carry. The code says scene,
because that is what the thing is. "Screen" survives only where it means the
display: `ScreenW`, `ScreenH`.

Art lives in `assets/bg/` and is embedded into the binary (`main.go`), because
js/wasm has no OS filesystem and `make binaries` packages the executable alone.

## The map

Where a scene leads is recorded **in that scene**, in its own file, and it
comes from the wiki's location graph
(`projects/realworld/story#helyszín-gráf-vázlat`). That graph has two kinds of
edge and so does the code:

- **Physical adjacency** — the alley is behind Noodle's house, the eating place
  is next door to the trinket shop, the club is through its back room, the roof
  is up Norman's stairs. Declared with `To:` in the scene's `Exits`.
- **The selector** — the wiki centres its map on a *Helyszínválasztó*, "nem
  valódi helyszín, hanem a menü-képernyő": a scene every main location connects
  to both ways. A main location lists `exitToSelector` among its exits, and
  `scene/scene.selector.go` derives the other half of each edge by reading the graph
  backwards — every scene with an exit to the selector gets a pin on it. The
  list of locations on the map is never written down twice.

The wiki's conditional arrows survive too: `Needs:` holds the flag an exit waits
for, which is how *Bolt →|beengedés| Klub* is expressed — the way into the club
is visible and named from the first visit, and stays shut until the shopkeeper
has been shown the underwater sun.

Until the doors are measured on the paintings, an exit is a strip along one edge
of the picture (`inkwell.ExitLeft`, `ExitRight`, `ExitBack`, `ExitNear`), sized
from `Game.SceneRect` so it lands inside the painting rather than under the HUD
— the convention every 1990s point & click used, and one field to re-aim later.
A scene reached from several rooms has no fixed way out: the BBS terminal is the
same terminal from the club, the flat or the roof, so its exit leaves `To` empty
and the engine binds it to `Back()`.

The graph stays machine-readable: exit hotspots are named `exit:<target>`, so
the connections can be read straight back out of the registered scenes — and
`Validate` rejects an exit that names a scene nobody registered.

## What this game pushed into the engine

Everything below started as a workaround in this package and ended up in
inkwell, because in each case the thing being worked around was a fact an
entity should have carried in the first place. The domain side of each is now a
field in a literal:

| was worked around here | is now |
|---|---|
| a scratch-image text blitter, because `drawText` dropped the colour | `Game.DrawText` renders in colour; `inkwell.TextWidth`, `WrapText`, `ClipText`, `GlyphW/H` are the library's |
| a pump widget driving its own `Runner`, because `queueAction` was unexported | `g.Do(action)` — one queue, in order, engine-side |
| `TapeSay`, a second spelling of `Say` that skipped the speech bubble | `Character.Voice` — `VoiceLog` sends a character's lines to the log |
| a `Tapes` map of display names | `Character.Label` |
| a `gate` wrapper and `if !World.HUDVisible()` at the top of every `Draw` | `Widget.When`, an `inkwell.Condition` the engine evaluates |
| a `titleBar` widget that pushed the scene title into the top bar | `TopBar` already discovers the title; `NoteVar` adds the "— paused" |
| a `UseWithGuard` widget stealing clicks to answer unauthored use-with pairs | `Game.UseWithFail`, beside `ExitLook` and `ExitTake` |
| a `windowSizer` widget resizing the window on its first tick | `Game.WindowScale` |
| a `sceneNav` widget walking the scene catalogue | `inkwell.SceneNav`, a built-in dev widget |
| `sceneDefaults` rewriting every scene with Paul and a floor | `Game.Player` and `Game.Walkboxes` |

Also: the inkwell README gives the module path as
`git.teletypegames.org/games/inkwell`; the real path per its `go.mod` is
`git.teletypegames.org/engines/inkwell`.

## Known gaps

- **The HUD still covers the foot of every painting.** The art is no longer
  stretched, but the strip is opaque and 118px of the 380 sit on top of the
  picture. Either the deck gets recut so nothing that matters lives in its
  bottom third, or the HUD gets a translucent panel treatment.
- **The scenes are empty.** 22 backgrounds are in and they are wired to each
  other, but only the alley has hotspots, NPCs and dialogue; everywhere else
  there is nothing to do but leave again.
- **The selector has no art.** The wiki's Helyszínválasztó is the centre of the
  map and the deck does not include it, so it renders as a placeholder with one
  labelled pin per location on a plain grid.
- **Two scenes are missing from the deck.** The alley (the one authored beat)
  and deck 05, Norman's workplace — which is in the wiki's catalogue but has no
  scene file yet, so the workplace thread of the story has nowhere to happen.
- **Spent dialogue choices are hidden, not struck through.** `DialogueChoice.Once`
  hides them; the deck wanted them struck through but visible, so the list
  becomes a memory of what you already tried. Needs a thin override over
  `DialogBox`.
- **Only beat 3 exists.** The alley is a vertical slice; the opening script sets
  the police-tip flag that beat 2 will eventually set.
- **The render has only been inspected once, by the author of this repo.**
  Geometry is derived from one font cell and one screen size, but this
  environment cannot take a screenshot (both synthetic keystrokes and screen
  capture are blocked by macOS privacy permissions), so every visual judgement
  has to come from you.
