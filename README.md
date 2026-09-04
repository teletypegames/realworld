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
go run .                   # the two splash cards, then the main menu
go run . -lang hu          # the same game in Hungarian, for this run only
go run . -finale           # the finale — try the Nokia-punk theme switch
go run . -scene alley      # straight into the alley (beat 3), no main menu
go run . -scene selector   # the location selector: the map the city hangs off
go run . -scene hospital   # any scene in the deck, e.g. to look at the art
go run . -debug            # with the debug menu on d — see The debug menu below
```

Scene names are the `Name` fields of the scene files (`selector`, `paul_shop`,
`noodle_house`, `alley`, `norman_apartment`, `police_station`, `hackerspace`,
`ice_cream_shop`, `trinket_shop`, `small_restaurant`, `secret_club`,
`bbs_terminal`, `street`, `hospital`, `server_farm`, `secret_lab`,
`rooftop_hideout`, `public_bbs`, `bvk_branch`, `curator_shop`, `scrap_market`,
`samizdat_press`, `showroom`, `columbarium`).

Settings and saves outlive the process. The engine keeps two JSON documents
under the game's name, `realworld`, in the user's configuration directory —
`~/Library/Application Support/realworld/` on macOS, `~/.config/realworld/` on
Linux, `%AppData%\realworld\` on Windows — and in `localStorage` for the wasm
build. `settings.json` is the language and the theme, written the moment they
are changed on the Settings screen; `slots.json` is every save slot in one
table, the session itself a base64 string beside the time and the scene a
load screen lists. `-lang` chooses the language for one run without writing it
back; with no flag and nothing saved yet the game speaks English.

To work against a local inkwell checkout, add a `replace` line temporarily —
and take it out before committing:

```bash
go mod edit -replace git.teletypegames.org/engines/inkwell=../../engines/inkwell
go mod edit -dropreplace git.teletypegames.org/engines/inkwell
```

## Building

```bash
make build                    # native binary into bin/
make debug                    # the same, with the debug menu switched on
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
| `F1` | toggle the debug overlay: a coloured plate and its name over every clickable — on from the start in a debug run or build |
| `d` | the debug menu, in a debug run or build — see below |
| `ESC` | in the game: the main menu. On a menu screen: back to the main menu |

Screenshots: `EBITEN_SCREENSHOT_KEY=q go run .`, then press `q` in the window.

## The debug menu

A window down the right half of the screen, opened and closed with `d`: the
scene deck, the item catalogue, and everything clickable in the scene the
player is standing in. It is the engine's `Menu` widget with the engine's
`DebugMenuRoot` handed to it, registered in `widget.debug_menu.go` like any
other widget — the menu system is generic, and the debug menu is one use of it.

It is off unless the build or the run switches it on:

```bash
go run . -debug     # this run
make debug          # a binary with it on from the start (-tags debug)
```

Both do the same one thing: set `inc.FlagDebug`, which is what the widget's
`whenDebug` reads. A release build with no flag never ticks the widget, so `d`
does nothing at all.

The same flag switches on the engine's `HotspotDebug` overlay, registered in
`widget.hotspot_debug.go` with `When: whenDebug` and `Enabled: true`: every
clickable on screen gets a translucent plate with its name written across it,
a different colour per kind — hotspots, exits, the characters in the scene,
the items in the inventory, the two Armilla slots. The HUD's own targets get
the plate alone, since the HUD already prints its labels, and the menu opens
on top of the overlay. `F1` hides and shows it during a debug run; outside
one the widget never ticks, like the menu.

```
DEBUG
├── Scenes         every scene in the deck — Enter teleports
├── Items          give / take away / select, and the use-with pairs
├── Action points  the current scene's hotspots and exit strips, verb by verb
└── Exits          where each one leads, and the flag that unlocks it
```

Arrows walk the rows, Enter opens one, Backspace goes back a level and closes
at the root, Escape closes; the mouse does the same, and the title bar carries
a back arrow and a close box. The game keeps running underneath — clicks
outside the window still reach it — and the menu stays live during a cutscene,
which is what makes it closable once one starts.

## Layout

Internal resolution is **640×380**: the size of one painted background.
Integer scaling only; the window opens at 2×.

That size is not a taste call. inkwell scales a scene background over the
*whole* screen (`core.engine.go`), not over the region the HUD leaves free — so
at any screen size other than the art's own, every painting is squashed to fit.
The deck is drawn at 640×380, so the screen is 640×380 and the art lands pixel
for pixel. Nothing enforces it — correctness here is checked by looking at the
game — so a background of any other size will simply come out squashed.

It still gives the wiki's "320×200 VGA look" brief what it was after — a low,
wide VGA frame at an exact 2× of a 320-wide canvas — and it still keeps the
engine's fixed **6×16** font readable, at 4% of the screen height. (At
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

Every vertical measurement derives from `LineHeight` (glyph cell + leading) and
from `ScreenHeight`, both in `widget/widget.manager.go`, not from the wireframe deck's ratios, so
the layout cannot drift out of step with the font — or with the art — again.
The tape channel comes out at 4 rows of 38
columns, which is about one and a half of Dex's remarks on screen at once.

The vertical divider at x=394 is the one proportion carried over from the
wireframe deck: 61.6% world, 38.4% tape channel. The deck's *vertical*
proportions are not carried over.

The dialog box deliberately spans only the left region, so a tape can comment
*alongside* a conversation.

## Languages

The game is written twice, in place. Every string the player reads is an
`inkwell.Text` — the same string in as many languages as somebody has written
it in — and it stays in the file that owns the thing it names:

```go
OnLook: inkwell.Say(character.Paul, inkwell.Text{
	{Lang: "en", Text: "Keypad. Four digits, worn keys."},
	{Lang: "hu", Text: "Billentyűzet. Négy számjegy, kopott gombok."},
}),
```

One translation per line, stacked, so the languages read as a table. There is
no catalogue and no message id: a line lives beside the hotspot it answers, so
a translator reading `scene.alley.go` can see what Paul is looking at while he
says it. Adding a third language is adding one more line to each array, and it
touches nothing else. `-lang` picks which one is spoken:

```bash
go run . -lang en    # the language it was authored in, and the fallback
go run . -lang hu
```

`main` sets `inkwell.Lang` from the flag and `inkwell.Fallback` to `"en"`,
so a language a line has not been translated into yet still reaches the player
in English rather than vanishing. Resolution happens at the moment of display,
never at construction, so the declaration that registers an entity does not
care when the language was set.

Two things do **not** get translated, and both are load-bearing:

- **Names.** A scene's `Name`, `inc.FlagHasTape`, `inc.NotePaused` are
  keys the engine and the conditions compare against. `inc.NotePaused`
  is the one that looks like prose — it is a `State` var a `Condition` reads,
  so the value stays `"paused"` in every language and `TopBar.Notes` in
  `widget.top_bar.go` maps it to what the top bar shows.
- **Sentence frames built around another string.** `inkwell.Format` resolves
  the frame and its arguments in one language, because English puts the
  preposition between the verb and its object and Hungarian puts it after
  them — "Use hook on grate" against "Használd a kampót ezen: rács". A join
  cannot say both; a frame per language can.

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

One category, one package, one directory, under one of two groups: `inc/system/`
is the machinery, `inc/content/` is what the game is made of. A file still says
which entity it holds in its own name — `[category].[name].go`, the prefix
repeated inside the directory that already carries it, because a file called
`alley.go` tells you nothing in a list of open editor tabs.

```
main.go                                flags, one Game literal, inkwell.Run
main.debug.go  main.release.go         debugBuild, the -tags debug pair
inc/constants.go                       the state keys — flags, vars, notes; imports nothing
inc/util/                              RGB and RGBA over image/color; imports nothing

inc/system/system.go                   System: the binding, the mode, the note
inc/system/widget/                     HUD layout, visibility conditions, 13 widgets
inc/system/armilla/                    Paul's Armilla: slot two and the strip
inc/system/theme/                      the Theme alias, realworld-93, nokia-punk

inc/content/stage/                     one file per full-screen screen
inc/content/scene/  …/background/       one file per location, twice
inc/content/tape/                      the Tape entity, the slots, the channel
inc/content/speaker/                   the voices with no body: Dex and the cards
inc/content/item/  …/character/  …/dialog/  …/script/
```

`inc` itself holds nothing but state keys, which is why every package can import
it and it can import none of them; `inc/util` is in neither group for the same
reason, a helper over `image/color` that names nothing of the game's. The
consequence at the other end is that the assembly cannot sit in `inc/system`
beside `System`: half the game imports `system`, and a package cannot be imported
by what it imports, so the wiring is `main.go`.

The groups say who writes a file rather than which way an import points:
`content` is what a writer adds a file to, `system` is the machine that runs it.
They agree with the graph, and there is no edge crossing the other way any
more: a character names a `theme` colour, a scene names a `system` action, a
cassette names the `widget` manager it registers into, and none of them is
known back.

`widget → tape` used to be the exception, because the tape slots asked which
cassette a tape was in. Both halves of it went. The tape channel does not ask
`tape` anything — it asks the engine which **voice** a speaker has, and a tape
is a `VoiceLog` speaker, so the filter is one field
(`Voices: []inkwell.Voice{inkwell.VoiceLog}`) and the HUD learns nothing about
cassettes. The tape slots went the other way: a slot you push a cassette into
is a cassette thing, so `tape/tape.slots.go` is a content file that registers a
widget, like every other content file registering into a manager.

The layering itself is still the import graph, and the compiler keeps it
whichever directory a package sits in. `inc/system` is the one directory that is
both a group and a package — the group is its subdirectories, the package is
`System`.

The module is plain `realworld` — no host in front of it, because nothing
imports this repository; it is an executable, not a library. Own packages read
`realworld/inc/system/widget`, and since that sorts after `github.com/…` they
take their own import group under the dependencies.

The layering is no longer a rule kept by hand — it is the import graph, and the
compiler rejects the arrow that points the wrong way:

```
inc ← theme ← system ← armilla ← widget ← tape ← main
        ↑        ↑                          ↑
        └── content ──────────────────────--┘
```

Adding a scene means adding `inc/content/scene/scene.<name>.go` and a
`assets/bg/<name>.png` beside it. Nothing else moves — a scene that says nothing
about its `Background` takes an asset named after itself, and
`background/background.manager.go` registers one per registered scene, so the
file name is the whole of the pairing.

Every category owns a manager, and they are all the engine's own
`inkwell.Manager[T]` — the same registry the `*Game` hangs its content off, kept
in registration order and addressed by name. There is no second registry type
here: the entity types are aliases of engine structs, so they already satisfy
`inkwell.Named` and need no help reading their own name. Each one lives in its
own package's manager file, next to the alias it holds, and it is the one
identifier that keeps its category in its own name — a manager is read where the
package qualifier has already fallen off, on a `*Game` field or inside its own
entity files, and a bare `Manager` says nothing about what is in it:

```go
// inc/content/character/character.manager.go
type Character = inkwell.Character

var CharacterManager = inkwell.NewManager[*Character]()
```

An entity file is a literal that hands itself over, and it is a declaration
rather than a procedure. What it declares is the variable every other file
refers to the entity by:

```go
package tape

var Support = TapeManager.Register(&Tape{
	Name:    "tape_support",
	Item:    item.SupportTape,
	Speaker: speaker.TapeSupport,
	Dialog:  dialog.SupportTape,
})
```

**References are pointers.** This used to read `Item: inc.ItemSupportTape`, a
string constant the engine looked up by name at run time — the inheritance of
the Lua-scripted game before this one, where a name was all a reference could
be. It meant that clicking `Background` in an editor landed on a constant, not
on the asset, and that a typo in a name was found by `Validate` at boot rather
than by the compiler. Now every manager holds pointers, every engine field
that names another entity is a pointer to it — `Stage.Background *Asset`,
`Exit.To`, `Say(speaker Talker, …)`, `Give(*Item)` — and the variable a file
binds is the entity's name in code. `Name` is left with the one job only a
string can do: the key a save file, a `State` var and the `-scene` flag use to
find the entity again from outside the process. `inc/constants.go` shrank to
those keys.

The price is that a reference is an import, so the content packages stand in
a strict order and only ever point down: `character` and `speaker`, then
`item`, `script`, `dialog`, `tape`, `scene`, `background`, `stage`. The old
graph had two edges pointing up, and both went through the same thing — the
tape voices, which `tape.manager.go` used to *derive* from the cards while
`dialog` and `script` gave them lines. A derived thing cannot be pointed at
from below it, so the voices became declared entities in a package of their
own, `speaker/`, and both edges turned downward. Scenes are the other place
Go pushes back: two package-level variables may not refer to each other in
their initialisers, and a round trip between two rooms does exactly that.
So an exit names its target through the scene file's own empty type,
`To: street{}`, whose `Scene()` method returns the variable; the engine asks
at run time, when every variable exists, and a type is not a dependency.

A tape is a card, not a character. It is pushed into a device and talks back
like a chatbot, so nothing about it stands in a scene, and the file is the
card: the item it travels in, the voice it speaks with, the dialog it plays,
which is a file in `dialog/` like every other tree, and `OnInsert`, its first
words when it clicks into slot two — run after the generic `tape_insert`
script, which is the part of an insertion every card shares. The voice is an
`inkwell.Speaker` in `speaker/` — `VoiceLog`, so the line lands in the channel
and never in a bubble — and `speaker.SpeakerManager` is what `main.go` hands
the engine. Dex is such a voice and such a card, so he is in neither
`character/`, which holds Paul and everyone he meets, nor only in `tape/` — and
one voice with no body stays a character, the security ACP, a `Character` with
`Voice: VoiceLog` because it is a room, not a cassette.

A choice a card only offers under a condition says so in `Show`, and the
condition is a literal: `inkwell.Given{Item: …, NotFlag: …}`, every set field
a requirement. An answer that branches is an `inkwell.If` over the same
`Given`; one that depends on more than a branch is `inkwell.Reply(speaker,
func(*State) Text)`, `Say` with the line decided when it is spoken.

`Register` returns what it registered, which is what makes that possible — a
statement can only stand in a function body, an expression can stand in a
variable initialiser. It buys more than two lines and an indent level: a
package-level variable is initialised in dependency order, so the manager is
built before any file that writes into it, whichever the compiler reads first.
Two files keep an `init()` and earn it — `background.manager.go` loops over
the scene deck, and `tape.slots.go` builds a value that has to be wired to
itself, neither of which is an expression.

A package initialises only if something imports it, so a package whose
entities nobody names by symbol needs a blank import in `main.go`.
Exactly one does — `tape`, everything else is imported for its manager — and
the list is per package, not per entity, so adding a file is still adding a
file.

Widgets go the same way, the engine's own included: `widget/widget.cursor.go` registers
an `inkwell.Cursor` exactly as `tape/tape.mystery.go` registers a cassette. What
they cannot take from alphabetical order is their place in the stack, so each
widget names a **layer** — `LayerScene`, `LayerPanel`, `LayerHUD`, `LayerSpeech`,
`LayerDialog`, `LayerMenu`, `LayerCurtain`, `LayerCursor` — and inkwell draws
from the bottom layer up, ticks from the top layer down, and lets registration
order decide only ties inside one layer. That is why the HUD frame comes out
under the tape slots and the cursor over everything, whatever the file names
happen to be. The layer was added to inkwell for this; the alternative was a
registration list, which is the thing this package spent its life deleting.

The second thing a widget declares is when it is on screen at all: `When` is an
`inkwell.Condition` over game state, and the HUD's is `WhenPlaying`, which reads
`VarMode` out of the engine's `State`. Hiding the HUD for a cutscene is
therefore one word per widget rather than a wrapper around each of them and an
`if` at the top of every `Draw`.

**Stages.** A stage is one full-screen thing the game shows, and exactly one
shows at a time: the main menu, load, save, settings, credits — and the game
itself, which is the point. The world is not the thing every other screen
floats over; it is one of them, `stage.Game`, the only one with
`World: true`. Off it inkwell paints no scene, walks no character, fires no
trigger and offers no click to a hotspot, which is what lets the main menu be
the first thing a player sees without the alley running behind it.

Each stage file is a whole screen: the `Stage` literal and the widget that
draws it, in one file — a `Sticky` `inkwell.Menu` for the four menu
screens, an `inkwell.TextCard` for the credits and for the two splash cards.
Sticky because a menu with nothing behind it must not be closable.

The splash cards are what the game opens on: `stage.SplashStudio`, the
Teletype Games logo, and `stage.SplashTitle`, "Real World" centred, each
two seconds, each fading to the next. None of that is code in this
repository: a `Stage` says `Hold: 2` and `Next`, and the fade is its
`Transition`, `inkwell.Fade(0.5)`. The logo is the Stage's `Background`, a
picture the engine stretches over the whole window, so
`assets/stage/splash_studio.png` is already 640×380: the logo scaled to
150 px tall on a black canvas, the logo's own background — and the colour
the `Fade` goes through, so the card leaves without a visible edge. The asset
is a literal nested in the stage's `Background`, registered in place —
`background.manager.go` derives an asset per *scene*, and a stage is not one,
and a picture that belongs to exactly one thing needs no name of its own. The two numbers are `splashHold` and
`splashTransition` in `stage.manager.go`, shared so that the two cards move
at one pace. A `Transition` belongs to the Stage that *leaves* through it,
which is what keeps the menu and its sub-screens cutting: neither names one.
`-scene` and `-finale` skip the cards along with the menu, because both mean
somebody wants to look at the game. The HUD needed one line to follow all
of this: `WhenPlaying` is `And(OnWorld(), Not(inCutscene))`, and because
*not being in a cutscene* is the default, a fresh `State` after a New Game is
already a playing one. `OnWorld` rather than `OnStage(stage.Game)` because the
HUD cannot name the stage that registers into it — that would be an import
cycle — and "the Stage showing is the world" is what it meant anyway.

Tapes are the one entity type this game invents. `Tape` holds the voice it
speaks with, the cassette that carries it and the dialog it plays, one file
each — what a tape *sounds* like stays on the `Speaker`, as its `Label`,
`Color` and `Voice`, because that is a fact about a voice.
A tape's line is `inkwell.Say` like anybody else's; the engine sends it to the
log instead of a speech bubble because the character says so.

So adding an entity is adding a file, and there is no second list to keep in
step. The price is that registration order is file-name order: the debug
menu lists the scenes alphabetically rather than in the concept-art deck's
order.
That was a deliberate trade — the deck order was a list that had to be kept in
step with the files by hand, and the deck is a thing to look at, not a thing to
play through.

The game's own catalogue is readable without going through the engine:
`scene.SceneManager.Get("alley")` answers long before there is a `*Game` to ask.
And when the game is built, nothing is copied into it and nothing is assigned
either — the eight managers are fields of the one `inkwell.Game` literal
`main.go` hands to `inkwell.New`, so engine and game share one registry per
category instead of keeping two of them in step. Themes used to be the
exception, because the engine seeded its theme manager with four presets and
handing ours over would have thrown them away; `inkwell.New` adds the presets to
whichever manager it ends up with, so `theme.ThemeManager` goes in the literal
like the other seven.

The defaults a scene may leave out are not written into it at all any more:
`g.Player` names the character a scene with no actors of its own receives, at
his `Start`, and `g.Walkboxes` is the floor a scene without one walks on. Both
are fields on the `*Game`, so a scene file that says nothing about either means
"the usual", and the selector opts out by declaring both empty. `WalkOnClick`
sits beside them: a left click on empty floor walks Paul there through the
walkboxes, and the engine leaves it off unless the literal says so. What is still
derived half of the map is no longer written back from the wiring at all: the
selector declares `Derived: pins.fromGraph`, and inkwell asks every scene for
its derived exits once, before validation. The map screen is a literal like
every other scene file, and `main` says nothing about it.

There is one game, so there is one system, and it is a value rather than a type
anyone constructs: `var System = &system{}`, with `system` unexported so no
second one can exist. Everything the package offers is a method on it —
`System.Do(…)`, `System.Game()`, `System.SetMode(…)`, `System.Paused(…)` — and
there are no package functions beside them, no hand-written `Action` or `Runner`
types, and now no closures either: `SetMode` is `inkwell.SetVar` over
`VarMode`, and `Paused` is `inkwell.Noting`, the engine's own "hold this var for
exactly as long as the action runs". `System.Func`, the escape hatch to
arbitrary Go, went with its last caller; `System.UseTheme` became
`inkwell.UseTheme`, because a theme switch mid-script is not this game's idea.
Nothing takes a system parameter and no
widget holds a back-reference, which is what lets a script file be a literal.
`System` itself is down to the binding:
the engine pointer, the mode and the top bar's note, the last two being `State`
vars the engine can see for itself. What it deliberately does not hold is the
assembly — every content package imports `system`, so the wiring lives in
`main.go`.

The device came out of it. **Armilla** is the wiki's word for the wrist-worn
runtime a Personal Tape snaps into, and Paul's is the one piece of the fiction
with enough state to earn a package — though barely, now: `armilla.Slot2()`,
`armilla.Insert(cassette)`, `armilla.StripIfEmpty(text)`, and nothing else.
Both values are `State` vars already spelled `armilla.slot2` and
`armilla.strip` in the save file, so there is no package state left and no
`Reset` either. The slot var holds the cassette's `Name`, because it has to
survive a save; `Slot2()` hands the `*inkwell.Item` back, found through the
engine's `ItemManager`.

Two things went. `SetPending`/`TakePending` existed only because a script could
not be told which cassette the player had just pushed in; now the slot is a
`Hotspot` and the cassette is the key of its `OnUseWith`, so
`armilla.Insert(item)` is built per cassette with the name already in it, out of
three engine actions and no Go. And `Offer`/`Waiting`/`TakeOffer` — the tape
that wanted to speak during a cutscene, once documented in the table above as a
`SPACE` control — went because nothing ever called `Offer`. The mechanic was
written, never wired, and therefore never seen. When content wants it back, the
declarative shape is an engine action that waits for a key, not a package
holding a flag the letterbox polls each frame.

The strip is an `inkwell.Text` now rather than a bare English string, because
`TopBar` resolves a `Text`-valued var into the player's language — so the one
player-visible string in the game that had no Hungarian has one.

What `armilla` does not know is `tape`: which dialog a cassette plays is
`tape.DialogOf`, asked by the slot, because that is a fact about the cassette
and not about the thing it is plugged into.

**On the word "scene".** The wiki, the concept-art deck and the beat tables all
count *screens*, and this code used to as well: it carried its own `Screen`
struct, because inkwell's `Scene` could not hold the things a screen needs —
its exits above all. Those went into the engine instead (inkwell `Exit`,
`Game.SceneRect`, `CurrentScene`/`PreviousScene`, `Back()`), and with the gap
closed there was nothing left for a second type to carry. The code says scene,
because that is what the thing is. "Screen" survives only where it means the
display: `ScreenWidth`, `ScreenHeight`.

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
an exit can only name a scene that exists, because it holds it.

## What this game pushed into the engine

Everything below started as a workaround in this package and ended up in
inkwell, because in each case the thing being worked around was a fact an
entity should have carried in the first place. The domain side of each is now a
field in a literal:

| was worked around here | is now |
|---|---|
| a scratch-image text blitter, because `drawText` dropped the colour | `Game.DrawText` renders in colour; `inkwell.TextWidth`, `WrapText`, `ClipText`, `GlyphW/H` are the library's |
| a pump widget driving its own `Runner`, because `queueAction` was unexported | `g.Do(action)` — one queue, in order, engine-side |
| `TapeSay`, a second spelling of `Say` that skipped the speech bubble | `Speaker.Voice` — `VoiceLog` sends a speaker's lines to the log |
| a `Tapes` map of display names | `Speaker.Label` |
| a `Character` per cassette, so that it could speak | `inkwell.Speaker` and `Game.SpeakerManager`: a voice with no body |
| `inc.SceneAlley`, `inc.ItemMysteryTape`: a string constant per entity, looked up by name at boot | every manager holds pointers and every reference is typed — `*Asset`, `*Item`, `*Stage`, `Talker`, `SceneRef` — so the compiler holds the graph and the editor follows a field to its struct |
| a dialog node per game state a cassette answers from | `inkwell.Reply(speaker, func(*State) Text)`: the line decided when spoken |
| a `gate` wrapper and `if !World.HUDVisible()` at the top of every `Draw` | `Widget.When`, an `inkwell.Condition` the engine evaluates |
| a `titleBar` widget that pushed the scene title into the top bar | `TopBar` already discovers the title; `NoteVar` adds the "— paused" |
| a `UseWithGuard` widget stealing clicks to answer unauthored use-with pairs | `Game.UseWithFail`, beside `ExitLook` and `ExitTake` |
| a `windowSizer` widget resizing the window on its first tick | `Game.WindowScale` |
| a `sceneNav` widget walking the scene catalogue with the arrow keys | `inkwell.Menu` + `inkwell.DebugMenuRoot`: a menu system, and a debug menu that lists the deck rather than stepping through it |
| nothing yet — but a menu could not have taken the arrow keys without `SceneNav` taking them too, and could not have stayed open across a cutscene | `Input.KeyPressed` / `ClaimKey`, the consume-on-use protocol extended to the keyboard, and `Awake`, the widget that ticks while a script holds the frame |
| `sceneDefaults` rewriting every scene with Paul and a floor | `Game.Player` and `Game.Walkboxes` |
| nothing yet — but a second language would have meant a parallel catalogue keyed by id | `inkwell.Text`: every displayed field carries its own languages, in place |
| nothing yet — the debug font's atlas stops at U+00FF, so `ő` and `ű` drew as blanks (and so did the em dashes already in these scene titles) | inkwell bakes its own atlas through Latin Extended-A |

Also: the inkwell README gives the module path as
`git.teletypegames.org/games/inkwell`; the real path per its `go.mod` is
`git.teletypegames.org/engines/inkwell`.

## Known gaps

- **The HUD still covers the foot of every painting.** The art is no longer
  stretched, but the strip is opaque and 118px of the 380 sit on top of the
  picture. Either the deck gets recut so nothing that matters lives in its
  bottom third, or the HUD gets a translucent panel treatment.
- **The scenes are empty.** 22 backgrounds are in and they are wired to each
  other, but only the alley has hotspots, NPCs and dialog; everywhere else
  there is nothing to do but leave again.
- **The selector has no art.** The wiki's Helyszínválasztó is the centre of the
  map and the deck does not include it, so it renders as a placeholder with one
  labelled pin per location on a plain grid.
- **Two scenes are missing from the deck.** The alley (the one authored beat)
  and deck 05, Norman's workplace — which is in the wiki's catalogue but has no
  scene file yet, so the workplace thread of the story has nowhere to happen.
- **Spent dialog choices are hidden, not struck through.** `DialogChoice.Once`
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
