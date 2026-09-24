# The UI system

Everything the player looks at that is not the painting: the frame the game is
drawn in, the two panes under it, the controls that fold them away, and the
eight mock screens that exist so the whole thing can be judged by running it.

The code is `inc/system/widget/` for the machinery and
`inc/content/stage/stage.mock_*.go` for the screens. `README.md` covers why the
frame is 640×380 at all; this file covers what is inside it.

---

## The one rule

The layout comes from the **Adventure Game UI** design, and the design has a
single rule that everything else is a consequence of:

> The frame splits **once**, and always at the same line. The top **3/5** is the
> world. The bottom **2/5** is the UI. Inside the UI, the left pane is what the
> character can do and the right pane is a third voice — an advisor that
> comments and never acts. The amber border belongs to the advisor alone, and
> the only text field in the game talks to it.

Before this, the HUD was measured from the bottom up out of whatever its
contents happened to need, so the art got the leftovers. Now the art is
authored against a fixed region and the UI fits itself into what is left.

```
     0                                       416│                      639
   ┌────────────────────────────────────────────────────────────────────┐
 0 │ ALLEY — paused                              SRP: 1 tape            │ TopBar (20)
20 ├────────────────────────────────────────────────────────────────────┤
   │            scene — hotspots, characters, SpeechBubble              │ Scene (208)
228├───────────────────────────────────────────┬────────────────────────┤
   │ Use trinket on: gap at the foot of the wall│ DEX                   │
   │ ┌───┬───┬───┬───┐   ┌───────┐┌────────┐   │ "The trinket is softer │ UI (152)
   │ └───┴───┴───┴───┘   │ 1 DEX ││ 2  —   │   │  than the mortar."     │
379└───────────────────────────────────────────┴────────────────────────┘
     InventoryBar          TapeSlots             TapeChannel
```

## The numbers

All of them are in `inc/system/widget/widget.manager.go`, derived from
`ScreenWidth`, `ScreenHeight` and `LineHeight` (glyph cell + leading) rather
than typed in, so the layout cannot drift out of step with the font or the art:

| constant | value | what it is |
|---|---|---|
| `ArtHeight`, `UITop` | 228 | the split line — `ScreenHeight * 3 / 5` |
| `UIHeight` | 152 | the UI half, open |
| `UICollapsedHeight`, `UIFoldedTop` | 38, 342 | the UI half folded — `ScreenHeight / 10` |
| `ConsoleWidth` | 224 | the right-hand pane, 308 of the design's 880 |
| `DividerX` | 416 | the vertical divider, console open |
| `RailWidth`, `RailX` | 32, 608 | the same divider, console railed |
| `RowHeight`, `RowGap` | 22, 3 | one bordered choice row |
| `PaneInset`, `Padding` | 10, 6 | the left pane's inset, everything else's |
| `HUDTop`, `HUDHeight` | `UITop`, `UIHeight` | what the old HUD widgets read |

`HUDTop` and `HUDHeight` are aliases on purpose: the HUD widgets that predate
the design — `hudFrame`, `StatusLine`, `InventoryBar`, the tape slots, the tape
channel — keep reading the names they always read, and followed the divider
from 394 to 416 and the split from 262 to 228 without an edit each.

The vertical divider is the design's own proportion: 65% world, 35% console.
The dialog box deliberately spans only the left region, so a tape can comment
*alongside* a conversation, and the tape channel sits in the right one, 36
columns wide.

## The two axes

The UI collapses in two independent directions, and either can be in either
state:

**The rail** is the design's own. **TAB** folds the right-hand console into a
32-pixel strip — vertical label, status dot, token pips, the `TAB` legend — and
the left pane grows to `RailX` to take the room. It is what the design's A8
screen shows. The flag is `inc.FlagConsoleCollapsed`.

**The fold** is ours, and the design does not have it. The `↓` chip at the
right of the heading row — or **Space** — drops the *whole* UI half to
`UICollapsedHeight`, 38 pixels, which is exactly one heading row. The art then
runs to y=342 instead of y=228. That is the entire point of it: more picture to
look at and to click. The flag is `inc.FlagUICollapsed`.

```
   ┌───────────────────────────────────────────────────────────┐
 0 │                                                           │
   │              [ CLOSE-UP — THE UNMARKED TAPE ]             │ art (342)
   │                                                           │
   │           “No label, no seal. Somebody wound              │
   │             this one by hand.”                            │
342├────────────────────────────────────┬──────────────────────┤
   │ CUTSCENE   any key to move on   ↑  │ ● DEX · ARMILLA  TAB │ UI (38)
379└────────────────────────────────────┴──────────────────────┘
```

Both flags live in `inc/constants.go` beside every other world-state key, which
is what lets a `Condition` read them and a save file remember them. Neither is
package state in a widget.

Two functions in `widget.manager.go` are the whole of the geometry that moves:

```go
func uiTop(game *inkwell.Game) float64     // UIFoldedTop when folded, else UITop
func paneWidth(game *inkwell.Game) float64 // RailX when railed, else DividerX
```

Every widget below reads those rather than the constants, which is why the two
axes compose without a single `if` about the other one. The fold's toggle is
the one thing that has to be named from outside the package, so it is exported:

```go
func FoldToggle() inkwell.Action
```

`stage.manager.go` binds it to Space, and `FrameSplit` runs it on a click.

## The widgets

One file each in `inc/system/widget/`, all exported and declarative so a stage
file stays a literal — the stage package registers them, and `widget` may not
name a stage back.

| file | type | layer | what it draws | folded |
|---|---|---|---|---|
| `widget.frame_split.go` | `FrameSplit` | `LayerPanel` | the plate, the split line, the amber divider, the fold chip | moves to `UIFoldedTop`, chip flips to `↑` |
| `widget.choice_pane.go` | `ChoicePane` | `LayerHUD` | header, hint, selectable rows — or a `Body`, the cutscene line | heading only; a `Body` moves over the art as a subtitle plate |
| `widget.console.go` | `Console` | `LayerHUD` | the right pane: dot, label, tokens, `TAB` chip, wrapped lines with their cost, the prompt bar — or the rail | its header row only; railed *and* folded leaves the dot |
| `widget.scanlines.go` | `Scanlines` | `LayerCurtain` | the 2px overlay | unchanged, the whole frame |
| `widget.inventory_pane.go` | `InventoryPane` | `LayerHUD` | item column, art box, description, verb chips | heading only |
| `widget.combine_pane.go` | `CombinePane` | `LayerHUD` | two item cards, an operator, rows, the tried-tonight list | heading only |
| `widget.art_mock.go` | `ArtMock` | `LayerScene` | the art half's stand-ins: dashed boxes, a speech bubble, a caption, a hotspot plate, a subtitle | follows the split down |

The first four are meant to outlive the mocks. The last three are the mocks'
own — `ArtMock` in particular stands in for art that does not exist yet, and is
the first thing to delete when it does.

Each pane's `areas()` returns nil when folded, so hover and ↑/↓ stop moving a
cursor the player cannot see. `FrameSplit.BlocksClickAt` cannot reach the game
— the engine's signature is `BlocksClickAt(Point) bool` — so it caches the top
it last drew at, the way `inkwell.Menu` caches its frame.

### The shared drawing vocabulary

`widget.manager.go` holds what more than one widget needs, and nothing that
only one does: `fill`, `stroke`, `strokeDashed`, `label`, `centered`, `plate`
and `platedCenter` for the boxed text over art, `row` with its four states
(normal, selected, dim, struck), `rowCursor` for keyboard-and-hover selection,
and `paneHeader`, which draws the heading line every pane starts with and hands
back the y the pane's own content begins at.

`plate` earned its place by having two callers: `ArtMock`'s caption, tag and
subtitle, and `ChoicePane`'s folded `Body`. It started as a method on `ArtMock`
and moved when the second caller appeared — a move, not a new abstraction.

### Colours

No new theme tokens. Everything reads `game.Theme()`:

| part | token |
|---|---|
| UI plate | `PanelBG` |
| art backdrop | `SceneBackdrop` |
| split line, divider, borders | `CharacterPanelBorder` |
| unselected row border | `ChatLogSystem` |
| selected row | `StatusText` border over `VerbButtonBG` |
| dim text | `TopBarText` |
| console border, label, dot | `TopBarAccent` |
| console body | `ChatLogResponse` |
| console cost suffix, chips | `DialogBorder` |
| console plate | `ChatLogBG` |

## The eight mock screens

All eight of the design's screens are stages of their own, one file each in
`inc/content/stage/`, so the layout is judged by running it:

```bash
go run . -stage mock_menu        # A1 — main menu
go run . -stage mock_dialogue    # A2 — NPC dialogue
go run . -stage mock_cutscene    # A3 — cutscene close-up, opens folded
go run . -stage mock_inventory   # A4 — inventory
go run . -stage mock_use         # A5 — using an item on a hotspot
go run . -stage mock_combine     # A6 — combining two items
go run . -stage mock_fail        # A7 — a combine that fails
go run . -stage mock_rail        # A8 — A2 again, console railed
```

`-stage` is `-scene`'s opposite number: `-scene` reaches a room, `-stage`
reaches a screen, and either skips the splash cards and the main menu. It takes
any stage's `Name`, so `-stage credits` and `-stage dossier` work too.

| control | what it does |
|---|---|
| ↑ / ↓, mouse | move the highlighted row |
| the `↓` / `↑` chip, or **Space** | fold the UI half down to its heading, and back |
| **TAB** | fold the console into its rail, and back |
| **Escape** | back to the main menu |

Nothing is wired to a game action: the mock moves, it does not play.

The copy is the game's own — Paul, Noodle, the alley, the unmarked Personal
Tape, the trinket, the black-market Armilla — rather than the design's
placeholder fiction, so the layout is judged against real string lengths. Every
line carries its `en` and `hu` pair like any other player-visible string.

### How a mock stage is put together

`mockChrome` in `stage.manager.go` is what every one of them goes through:

```go
var MockCutscene = mockChrome(StageManager.Register(&Stage{
	Name:    "mock_cutscene",
	Title:   Text{ … },
	OnEnter: inkwell.SetFlag(inc.FlagUICollapsed),
}))
```

It registers the four things every mock screen has — `FrameSplit`, `Scanlines`,
Escape and the Space fold — and composes `OnEnter` so that both flags are
cleared on entry *before* whatever the stage's own literal asked for. That is
how `mock_cutscene` opens folded and `mock_rail` opens railed with one line
each, and how leaving either one does not leak its state into the next screen.

The stage file then registers its own content — an `ArtMock`, a pane, a
`Console` — each gated on `inkwell.OnStage(…)`, the same shape
`stage.credits.go` and `stage.dossier.go` already had.

`mock_cutscene` opens folded because a cutscene is the case the fold was made
for: there is nothing to pick, only something to watch, so the words belong
over the picture and the plate belongs out of the way.

## Departures from the design

Three, all deliberate:

**The top bar stays.** The design draws art from y=0. The game's `TopBar`
carries the Armilla strip and the note — real content, not chrome — so it keeps
its 20 pixels and overlays the top of the art region on the world stage. The
mock stages are not the world stage, so `widget.WhenPlaying` hides it there and
they show the design's full-height art half.

**The right pane is DEX · ARMILLA.** The design's third voice is called ACP.
In this game the third voice is Dex, on the wrist, and `character.SecurityACP`
is a different thing entirely — a room-scale security system that stands in its
own scene and speaks through the log.

**The token counter is invented.** Nothing in this game spends tokens. It is
drawn because it is a layout element the console header has to leave room for,
and it is the first thing to cut if the mechanic never arrives.

## What it does not do yet

The fold stops at the mock stages. On the world stage the HUD is engine widgets
whose geometry is fixed when they are registered — `StatusLine.Y`,
`InventoryBar.Origin`, the tape slots' rect, the tape channel's `Bounds` — so a
shrinking plate would leave all four painted over the art.

The follow-up is worth recording, because the engine already allows it:
`Game.SceneRect` is a plain exported field that `Game.SceneArea()` reads live
every frame, and `SceneArea` is what exit edge strips and walk-on-click are
measured against. Once the HUD widgets take their geometry from `uiTop` too,
folding on the world stage can grow the *walkable, clickable* scene rather than
merely uncovering more of a background that was always drawn full-frame.

The panes themselves are also still mock-only: `ChoicePane` shows rows, but
nothing hands it a dialog's choices, and `Console` shows lines, but nothing
hands it the tape channel's log. Wiring those is what turns the eight screens
from a layout study into the game's UI.
