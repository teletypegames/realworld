# Real World

A point & click adventure built on the inkwell engine (Ebitengine).

## Conventions

### No comments

Go source files carry no comments. Not doc comments, not inline comments, not
section headers. If a piece of code needs explaining, rename it or restructure
it until it does not, and put the reasoning in `README.md` instead.

The single exception is compiler directives (`//go:embed`, `//go:build`), which
are instructions to the toolchain rather than prose.

### Every name is spelled out

No abbreviations, anywhere an identifier is ours: not in a local, not in a
parameter, not in a receiver, not in a constant, not in a field. It is `game`,
not `g`; `theme`, not `th`; `destination`, not `dst`; `context`, not `ctx`;
`mousePoint`, not `mp`; `ScreenWidth`, not `ScreenW`; `SceneNoodleHouse`, not
`SceneNoodle`. A one-letter name is a comment the compiler cannot check, and this
file already forbids the checkable kind.

Receivers are spelled out too, which means Go's one-letter receiver convention
does not apply here: `func (frame *hudFrame)`, `func (slots *slotPanel)`,
`func (fail useWithFail)`. The receiver takes the noun the type is *an instance
of*, so it never has to repeat the type name and never shadows it.

`func (system *system)` is the one place the two words are the same, and it is
not an oversight: there is exactly one system, so the noun an instance of it
takes *is* "system". Nothing inside the methods needs to name the type, so the
shadowing costs nothing.

Three things keep their short form, and all three for the same reason — they are
somebody else's word, not ours:

- **Acronyms that are the domain's own**: `HUD`, `BBS`, `BVK`, `SRP`, `RGB`. An
  acronym is the full name of the thing; `RGB` is what `image/color` calls it.
- **The engine's vocabulary**: `inc.VarMode` and `inc.VarNote` stay `Var`,
  because the engine's own API is `State.Var` and `State.SetVar`, and half a
  translation is worse than none.
- **`err`**, which is the language's own spelling in every Go program.

### English first, then every other language beside it

Identifiers, commit messages, `README.md` and this file are English and only
English. A string the **player** reads is different: it is an `inkwell.Text`,
carrying English first and every other language the game speaks after it, in
the same file, at the same place:

```go
OnLook: inkwell.Say(inc.Paul, inkwell.Text{
	{Lang: "en", Text: "Keypad. Four digits, worn keys."},
	{Lang: "hu", Text: "Billentyűzet. Négy számjegy, kopott gombok."},
}),
```

One translation per line, stacked, and English first because it is the
fallback — a line nobody has translated yet still reaches the player. Never
leave a new player-visible string with only one language, and never lift one
out into a catalogue: the wording belongs beside the thing it names, which is
what lets a translator see what a line is answering. `README.md` has the rest,
under **Languages**.

The design wiki is in Hungarian; when a concept crosses over, translate it and
keep the mapping one-to-one.

### No tests

Do not write tests and do not add test files. Correctness is checked by running
the game.

### Struct literals are always multi-line

Every field of a struct instance goes on its own line, with a trailing comma —
never on the same line as the brace, never two fields on one line.

```go
return inkwell.Asset{
	Name: BgStreet,
	Path: "assets/bg/street.png",
	Kind: inkwell.AssetImage,
}
```

Not `inkwell.Asset{Name: BgStreet, Path: "...", Kind: inkwell.AssetImage}`.

This holds for nested literals too, including small ones such as
`inkwell.Point`. Slice and map literals are not covered by the rule.

An `inkwell.Translation` is the one exception the other way: the two fields
stay on **one** line, so the translations of a string stack into a table and
the eye can run down a single column to compare them.

```go
Title: inkwell.Text{
	{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
	{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
},
```

Inside a `[]inkwell.Text` or a `map[string]inkwell.Text` the element type is
implied and left off, which is what `gofmt -s` produces.

## Layout

One category, one package, one directory, and every category sits in one of two
groups: `inc/system/` is the machinery, `inc/content/` is what the game is made
of. A file keeps the category in its own name anyway — `[category].[name].go`
inside `inc/[group]/[category]/` — because the prefix is what makes a file
findable in a list of editor tabs, a grep, or a diff, where the directory has
already fallen off, and now two directories have.

- The category is always **singular**: `item`, `scene`, `background`,
  `character`, `tape`, `dialog`, `script`, `armilla`, `widget`, `theme`.
- `[category].manager.go` is the file that ties a package together — its entity
  type, its `Manager`, and whatever else the category shares.
- Every other file in a package holds exactly one entity: one scene, one
  item, one screen. The file registers it itself, as a declaration.
- Two things have no category. `inc/constants.go` is the vocabulary every
  package imports, and `main.go` builds the game — flags, one `inkwell.Game`
  literal, and
  `inkwell.Run` — beside the two one-line files a `-tags debug` build picks
  between.
- `inc/system/` is the one directory that is a group *and* a package. The group
  is its subdirectories; the package is the single file `system.go`, which holds
  `System`, the singleton every other package reaches through.
- `inc/util` is in neither group, beside `inc/constants.go` rather than under
  one of them, for the same reason: it is a helper over the standard library
  with no game in it, so it imports nothing of ours and anything may import it.
  `util.color.go` holds `RGB` and `RGBA`, the hex-to-`color.Color` pair the
  themes are written in.

```
main.go                                  flags, one Game literal, inkwell.Run
main.debug.go                            debugBuild = true, behind -tags debug
main.release.go                          debugBuild = false, without it
inc/constants.go                         entity names and world-state keys
inc/util/util.color.go                   RGB and RGBA, hex to color.Color

inc/system/system.go                     System: the binding, the mode, the note
inc/system/widget/widget.manager.go      the Widget alias, HUD layout, the conditions
inc/system/widget/widget.*.go            one widget per file, registering itself
inc/system/armilla/armilla.manager.go    Paul's Armilla: slot two and the strip
inc/system/theme/theme.manager.go        the Theme alias and its Manager
inc/system/theme/theme.realworld.go      realworld-93
inc/system/theme/theme.nokia_punk.go     nokia-punk

inc/content/stage/stage.manager.go       the Stage alias, its Manager, the shared rows
inc/content/stage/stage.game.go          the world, and the key that leaves it
inc/content/stage/stage.*.go             one screen per file: menu, load, save, …
inc/content/scene/scene.manager.go       the Scene alias, its Manager, the floor
inc/content/scene/scene.selector.go      the map screen: pins derived from the graph
inc/content/scene/scene.*.go             one file per scene
inc/content/background/background.manager.go  one asset per registered scene
inc/content/character/character.*.go     the cast: people who stand in a scene
inc/content/item/item.*.go               inventory
inc/content/dialog/dialog.*.go           dialog trees
inc/content/script/script.*.go           named action sequences
inc/content/tape/tape.manager.go         the Tape entity, the Speakers derived from it
inc/content/tape/tape.slots.go           the two Armilla slots, as hotspots
inc/content/tape/tape.*.go               one card per file: label, colour, the dialog it plays
```

`inc` is constants and nothing else, so every package can import it and it can
import none of them. The assembly is the other end of that: it is the one thing
that has to name every package at once, so it can be named by none of them, and
it sits in `main.go`. It cannot live in `system` beside `System` — half the game
imports `system`, and a package cannot be imported by what it imports.

The two groups say who writes a file, not which way an import points — but they
agree, and there is no longer an edge that crosses the other way. A character
names a `theme` colour, a scene names a `system` action, a cassette names the
`widget` manager it registers into, and none of `theme`, `system` or `widget`
knows the caller back.

`widget → tape` used to be the exception, and it was defended here: the tape
slots asked which cassette a tape was in. It is gone, and both halves of it are
worth knowing why. The tape channel does not ask `tape` anything — it asks the
engine which **voice** a speaker has, and a tape is a `VoiceLog` speaker, so
the filter is `Voices: []inkwell.Voice{inkwell.VoiceLog}` and the HUD learns
nothing about cassettes. The tape slots moved the other way: they are a
*cassette* thing, so `tape.slots.go` is a content file that registers a widget,
the way every other content file registers into a manager.

A cassette is something somebody writes — a voice, an item, a dialog, a slot
you push it into — so `tape` is content, and the HUD no longer reaches across
for it. The layering is still the import graph, and the compiler keeps it on its
own whichever directory a package sits in; the directories are there so that
opening `content/` shows the game and opening `system/` shows the machine.

The module is plain `realworld`, with no host in front of it, because nothing
imports this repository — it is an executable, not a library. Own packages are
therefore `realworld/inc/system/widget`, and since that sorts after
`github.com/…`, they go in their own import group, below the dependencies.

## Managers

Every category that owns a collection of entities has a manager, and they are
all the engine's own `inkwell.Manager[T]` — the same registry type the `*Game`
hangs its content off. The game defines no registry of its own.

Each one is declared in its own package's manager file, beside the alias it
holds, and it carries its category in its own name — the one identifier that
keeps the prefix, because a manager is read far from its package and a bare
`Manager` says nothing about what is in it:

```go
// inc/content/character/character.manager.go
package character

type Character = inkwell.Character

var CharacterManager = inkwell.NewManager[Character]()
```

Read from `main.go` the nine of them still line up as a list —
`background.BackgroundManager`, `character.CharacterManager`,
`dialog.DialogManager`, `item.ItemManager`, `scene.SceneManager`,
`script.ScriptManager`, `tape.TapeManager`, `theme.ThemeManager`,
`widget.WidgetManager` — only now the list is a consequence of the imports
rather than a block someone has to keep in step.

Aliases of engine structs already carry `GetName()` and satisfy
`inkwell.Named`, which is the whole of what a manager asks of them.

The methods the game uses are `Register`, `Set`, `Get`, `All` and `Each`.
`Register` panics on a duplicate name — a second registration is a
construction-time bug, not an update — so rewriting an entity that is already
in the registry goes through `Set`, which keeps its position in the order.
`All` and `Each` both hand back the entities in registration order.

Nearly every entity type is an alias, `Scene` included. It was once a struct of
our own, because inkwell's `Scene` could not carry exits; that gap was closed in
the engine, so there is nothing left for a second type to hold.

`Tape` is the exception, and it is a real one: a tape is a card you push into
a device, and it talks back like a chatbot — a thing inkwell has no notion of.
It is **not a character**: nobody on a cassette stands in a scene, walks, or
has a sprite. A tape file is the card in one literal — its `Label`, its
`Color`, the `Item` that carries it, the name of the `Dialog` it plays and
`OnInsert`, the first thing it says when it clicks into slot two. The
dialog itself is a file in `dialog/` like every other tree, because a
conversation is written and read on its own, and the card only has to say
which one is its. `OnInsert` runs after `script.tape_insert.go`, which is the
part of an insertion every card shares — Paul's hand and Dex's remark — so the
generic script never has to know which card went in.

`tape.manager.go` derives the voice. Every tape is an `inkwell.Speaker` with
`VoiceLog`, registered into `tape.Speakers`, which is what `main.go` hands the
engine as its `SpeakerManager`. Dex is one of these cards — the one in slot
one — which is why he is not in `character/` beside Paul and the people Paul
meets. One voice there is a character with no body on purpose: the security
ACP is a room-scale system, so it stands in its scene as a `Character` with
`Voice: inkwell.VoiceLog`, and speaks through the log. The two lookups are
`tape.Of`, the tape an item is, and `tape.Dialog`, the dialog to run for a
loaded item, which is the card's own or Dex's when the card names none.

What a dialog answers with, and when, is the engine's. A choice that is only
offered under a condition says so in `Show`, and the condition is written as a
literal, `inkwell.Given`, with every requirement a field: `Flag`, `NotFlag`,
`Item`, `NotItem`, `Var` with `Is`. An answer that branches is an `inkwell.If`
over the same `Given`; an answer that has to be computed is
`inkwell.Reply(speaker, func(*inkwell.State) inkwell.Text)`, `Say` with the
line decided when it is spoken.

### Entities register themselves

An entity file is a literal and nothing else. No constructor function, no list
somewhere else to keep in step, and no procedure either — the file *is* the
declaration that hands itself to its manager:

```go
package tape

var _ = TapeManager.Register(Tape{
	Name: inc.TapeSupport,
	Label: inkwell.Text{
		{Lang: "en", Text: "TECH SUPPORT"},
		{Lang: "hu", Text: "ÜGYFÉLSZOLGÁLAT"},
	},
	Color: theme.AmberPale,
})
```

`Register` returns what it registered, which is the whole of what makes this
possible: a statement can only stand in a function body, an expression can
stand in a variable initialiser. The result goes to `_` because an entity
needs no name — the file name says which one it is.

This is not only two lines and an indent level shorter than an `init()`. A
package-level variable is initialised in **dependency order**, so the manager
is built before any file that writes into it, whichever one the compiler reads
first — with `init()` that rested on the separate rule that every package
variable is ready before any `init()` runs. The registration order is
unchanged: still the order the files are presented in, which is alphabetical,
and still nothing may depend on it.

Three files keep an `init()`, and all three earn it. `background.manager.go`
loops over the scene deck, and `tape.manager.go` loops over the cards to
derive a speaker from each; a loop is not an expression. `tape.slots.go` builds a
value that has to be wired to itself — `panel.Hotspots = panel.hotspots` —
which a composite literal cannot do. Where a file is assembly before it is an
entity, `init()` is the honest word for it; contorting those two into
declarations would cost more than it saves.

Backgrounds are the one category with no per-entity file, and it is not an
exception to the rule so much as the rule reaching its limit: every one of the
twenty-four was the same three fields, and two of the three were the scene's own
name spelled twice more. `Scene.Background` now defaults to the scene's `Name`
(`inkwell.Scene.BackgroundAsset`), so a scene file says nothing about its
picture, and `background.manager.go` registers one asset per registered scene:

```go
scene.SceneManager.Each(func(entity scene.Scene) {
	BackgroundManager.Register(Background{
		Name: entity.Name,
		Path: "assets/bg/" + entity.Name + ".png",
		Kind: inkwell.AssetImage,
	})
})
```

That is what makes `background → scene` the right direction: the asset is
*derived from* the deck, so the deriving package imports the source. Adding a
scene is now a PNG whose name matches.

Adding an entity is adding a file. Deleting one is deleting a file.

### A file's helpers hang off the file's own type

An entity file that is too long to read as one literal breaks pieces out —
`scene.alley.go` builds its four hotspots that way. Those helpers are **not**
package functions. Twenty-four scene files share one `scene` namespace, so a
package-level `bin()` is the *package's* bin, and the next scene that wants one
cannot have it. The file declares an empty type instead, named after itself, and
hangs its helpers off that:

```go
type alley struct{}

var _ = SceneManager.Register(Scene{
	…
	Hotspots: alley{}.hotspots(),
})

func (place alley) hotspots() []inkwell.Hotspot {
	return []inkwell.Hotspot{place.hidingPlace(), place.backDoor(), place.bin(), place.fireEscape()}
}

func (alley) bin() inkwell.Hotspot { … }
```

`alley{}` in the literal rather than a local is the same rule read from the
other end: the type is the namespace, so naming an instance of it adds a word
that says nothing. Where the helpers are used more than once, a method that
gathers them — `alley{}.hotspots()` — keeps the receiver out of the literal
entirely.

The receiver is the namespace, so the method name drops what the type already
says — `selector.leadsTo`, not `leadsToSelector`. Where the helpers read from
data rather than build it, the type is not empty and the data are its fields:
`script.usewith_fail.go` carries the two `inkwell.Ladder` values as fields of
`useWithFail` and answers with `fail.paul.At(n)` and `fail.dex.At(n)`.

The rule stops where the package is already the namespace: `armilla.Slot2`,
`armilla.Insert` and `tape.Dialog` stay functions, because a receiver there
would be an empty word. `system` is the other way round — the receiver is
the whole point, so everything it offers is a method on `System`.

Initialisation order is file-name order, so **registration order is
alphabetical**. Nothing may depend on it — including the order the debug menu
lists the scenes in, which is simply the order the files sit in.

A package initialises only if something imports it. `main.go` names every one
of them for a manager, which is import enough — `tape` for `tape.Speakers`.
That list is **per package, not per entity** — a new scene file still touches
nothing but itself.

### Widgets name their layer instead of their order

A widget is a file like any other entity — `widget.<name>.go`, one widget,
registering itself — and that includes the engine's own widgets,
which the game registers as literals the same way it registers a background:

```go
var _ = WidgetManager.Register(&inkwell.Cursor{
	Name: "cursor",
})
```

What a widget cannot take from alphabetical registration is its place in the
stack: the frame has to be drawn under the slots that sit on it, the cursor over
everything, and the use-with guard has to see a click after the HUD has had it.
So a widget declares a layer rather than inheriting one from the order it was
registered in — `inkwell.LayerScene`, `LayerPanel`, `LayerHUD`, `LayerSpeech`,
`LayerDialog`, `LayerMenu`, `LayerCurtain`, `LayerCursor`. The engine draws from
the bottom layer up and ticks from the top layer down, and registration order
only decides ties inside one layer:

```go
func (h *hudFrame) Layer() inkwell.Layer { return inkwell.LayerPanel }
```

A widget that says nothing sits on `LayerHUD`; ours all say it, because the
layer is the one thing about a widget the file cannot show.

The other thing a widget declares is **when it is there at all**. The HUD is
hidden for a cutscene, and that is a condition over game state, not a wrapper
and not an `if` at the top of every `Draw`:

```go
var _ = WidgetManager.Register(&inkwell.StatusLine{
	Name: "status",
	When: WhenPlaying,
	Y:    statusY,
})
```

`WhenPlaying`, `whenCutscene` and `whenDebug` are in `widget.manager.go`, and they
read `inc.VarMode` and `inc.FlagDebug` out of the engine's own `State`.
`WhenPlaying` is the one of the three that is exported, because `tape.slots.go`
is a content file that has to name it; the other two are read inside the package
only. A widget of
ours says the same thing with a method — `VisibleWhen() inkwell.Condition` —
because it has no literal to put a field in. A widget that is switched off
neither ticks nor draws nor blocks a click, so nothing else has to ask.

### Handing a category to the engine

Nothing is copied, and nothing is assigned either: the managers are **fields of
the `inkwell.Game` literal** `main.go` hands to `inkwell.New`. The types are
identical, so the engine and the game share one registry per category rather
than keeping two in step:

```go
game := inkwell.New(inkwell.Game{
	…
	AssetManager:     background.BackgroundManager,
	CharacterManager: character.CharacterManager,
	DialogManager:    dialog.DialogManager,
	ItemManager:      item.ItemManager,
	SceneManager:     scene.SceneManager,
	ScriptManager:    script.ScriptManager,
	SpeakerManager:   tape.Speakers,
	ThemeManager:     theme.ThemeManager,
	WidgetManager:    widget.WidgetManager,
})
```

`New` fills in an empty manager for every category a caller leaves out, so the
nine above are the nine we own and the tenth, `VerbManager`, is the engine's
four SCUMM verbs untouched. `SpeakerManager` is the one that is not a package's
own registry but derived from one: the cards in `tape`, each a voice with no
body.

Themes used to be the odd one out — `NewGame` seeded its `ThemeManager` with
four presets, so handing ours over would have thrown them away, and the game
copied its two in with `Each` instead. `inkwell.New` now *adds* the presets to
whichever `ThemeManager` it ends up with, leaving alone any name the game has
taken, so `theme.ThemeManager` goes in the literal like the other seven and the
copy is gone.

One consequence of not copying: a manager swapped in this way must be in place
before `inkwell.Run`, which is where the engine wires up the parts that hold a
registry directly. Nothing else has to run before the hand-off any more — the
selector's derived pins used to, and now `inkwell.Run` asks every scene for its
`Derived` exits itself, before validation, so the map screen needs no line in
the wiring.

The defaults a scene leaves out are no longer written into it at all.
`game.Player` names the character a scene with no actors of its own gets, placed
at his `Start`; `game.Walkboxes` is the floor a scene walks on when none. Both
are fields on the `*Game`, so a scene file that says nothing about either is
saying "the usual", and the selector opts out by declaring both empty.

### What runs at boot

The whole assembly is **one `inkwell.Game` literal**. There is nothing to read
in order and nothing to hold in your head: the game is what the literal says it
is, the same way a scene is what its literal says it is.

```go
game := inkwell.New(inkwell.Game{
	Title:       "Real World",
	Width:       widget.ScreenWidth,
	Height:      widget.ScreenHeight,
	WindowScale: 2,

	StartScene:   *start,
	StartScript:  inc.ScriptOpening,
	FinaleScript: finaleScript,
	StartTheme:   inc.RealWorld,
	Flags:        flags,

	Player:      inc.Paul,
	WalkOnClick: true,
	Walkboxes:   scene.Floor,
	ExitLook:    script.ExitLook,
	ExitTake:    script.ExitTake,
	UseWithFail: script.UseWithFail,
	…
})
```

The opening a player sees is content like every other script, in
`script/script.opening.go`, named by `StartScript` rather than built as a slice
in the wiring. `FinaleScript` is the engine's slot for a closing script, queued
straight after the opening; `-finale` is what fills it, which is why it is a
local computed from the flag rather than a method called afterwards. `Flags` is
the same shape for `-debug`: a boot flag list, not a reach into `State` after
construction.

`ExitLook`, `ExitTake` and `UseWithFail` are the three responses the engine asks
the game for, and all three are `script` functions — a line the player reads
belongs with the rest of the writing, not in the wiring.

The command line is a type rather than four loose pointers: `options` carries
what the flags chose, and the two answers the literal needs that a flag does not
give directly are methods on it — `chosen.finaleScript()`, which is `""` for an
ordinary run, and `chosen.flags()`, which is where `-debug` and `debugBuild`
meet. `chosen.game()` holds the literal. That is the same rule the entity files
follow: a file's helpers hang off the file's own type, so `main` grows methods
rather than package functions with names nobody else can use.

What is left in `main` is exactly what cannot be in the literal: the two
`inkwell` package variables (`Lang`, `Fallback`), `system.System.Attach(game)`,
which needs the pointer the literal is about to become, and `inkwell.Run(game)`.
Six lines — `main` used to be twelve assignments and four method calls.

## Stages

A **stage** is one full-screen thing the game shows, and exactly one shows at
a time: the main menu, the load and save screens, settings, the credits — and
the game itself, which is the point. The world is not the thing every other
screen floats over; it is one of the screens, `inc.StageGame`, the only one
with `World: true`.

Off the world stage inkwell does not paint the scene, does not walk the
characters, does not fire a trigger and does not offer a click to a hotspot.
That is why the main menu can be the first thing a player sees without the
alley running behind it — before stages, the only way to hold the world still
was a script that never finished, which is the trick `ShowEnd` still plays.

A stage file is a screen, both halves of it: the `Stage` literal and the
widget that draws it, in one file. `stage.main_menu.go` registers a
`Stage` and the `inkwell.Menu` that is that stage — `Sticky`, because a menu
with nothing behind it must not be closable — and `stage.credits.go`
registers a `Stage` and an `inkwell.TextCard`. That is the same shape
`tape.slots.go` already has: a content file registering a widget.

The HUD needed one line, not fourteen. `widget.WhenPlaying` is now
`And(OnStage(inc.StageGame), Not(inCutscene))`, so every HUD widget
disappears off the game stage without knowing stages exist. Note which way
round that is: **not being in a cutscene is the default**, so a fresh `State`
after a New Game is a playing one, and nothing has to remember to set the
mode back.

`Restart`, `LoadSlot` and `GoToStage` all land at the top of the next frame,
which is what lets a menu row say the whole of what it means at once:

```go
Do: inkwell.Seq(inkwell.Restart(), inkwell.GoToStage(inc.StageGame)),
```

`-scene` and `-finale` skip the front door: both mean somebody wants to look
at the game, not at its main menu, so `options.startStage` answers
`inc.StageGame` for either.

## The system

There is exactly one game, so there is exactly one system — and since there is
exactly one, it is a value, not a type anyone constructs:

```go
type system struct {
	game *inkwell.Game
}

var System = &system{}
```

The type is unexported, so `System` is the only instance that can exist, and
every name the package offers is a method on it: `System.Do(…)`, `System.Game()`,
`System.Paused(…)`, `System.SetMode(…)`. There are no package functions beside
it, and no `Action` or `Runner` types of its own either — nor any closure now:
`SetMode` is `inkwell.SetVar` over `inc.VarMode`, and `Paused` is
`inkwell.Noting`, the engine's own "hold this var for exactly as long as the
action runs". `System.Func`, the escape hatch to arbitrary Go, is gone with its
last caller, and so is `System.UseTheme`, which is now `inkwell.UseTheme` — a
theme switch mid-script is not this game's idea. One short file
holds the whole package. The mutable state is the struct's
fields rather than package variables, so `Attach` is a hand-off to one value and
not a sweep over several — and nothing is threaded anywhere: a widget holds no
back-reference and a script takes no parameter, because `system.System` is
reachable from anywhere that imports `system`. `main` calls
`System.Attach(game)`, which binds the engine and resets the run state, so
building the game twice is clean.

`system` holds as little as it can get away with, and after the Armilla moved out
it holds nothing but the binding: the engine pointer, the mode, and the top
bar's note. The mode and the note are `State` vars (`inc.VarMode`,
`inc.VarNote`), because state the engine can see is state a `Condition` can read
— that is what makes a widget's `When` possible and what puts "— paused" in the
top bar without anyone pushing it there. `System.Do` hands its action to
`game.Do`, the engine's queue, so the game runs one action at a time without a
pump of its own.

What `system` cannot hold is the assembly. Every content package imports it, so
it can import none of them, and the wiring lives in `main.go` instead.

## The Armilla

The Armilla is the wrist-worn runtime a Personal Tape snaps into — the wiki's
term, and the game's device. Paul wears one, and it is the only piece of the
fiction the code models as a package of its own: slot two and a phosphor text
strip.

```go
armilla.Slot2()          armilla.Insert(item)          armilla.StripIfEmpty(text)
```

Both are `State` vars — `inc.VarSlot2` and `inc.VarArmillaStrip`, both spelled
`armilla.…`, so the save file says the same word the package does. There is no
package state left at all, which is why there is no `Reset` either: it had
nothing to clear once the pending cassette and the waiting tape went.

Those two went for different reasons. `SetPending`/`TakePending` existed only
because a script could not be told which cassette the player had just pushed in
— and now it need not be told, because the slot is a `Hotspot` and the cassette
is the key of its `OnUseWith`, so `armilla.Insert(item)` is built per cassette
with the name already in it. `Offer`/`Waiting`/`TakeOffer` — the tape that wanted
to speak during a cutscene — went because **nothing ever called `Offer`**: the
mechanic was authored, documented as a `SPACE` control, and dead. When content
wants it back, the declarative shape is an engine action that waits for a key,
not a package holding a flag the letterbox polls.

`armilla.Insert` is a `Seq` of three engine actions and no Go: set slot two,
take the cassette out of the inventory, write the strip. `StripIfEmpty` is
`inkwell.SetVarIfEmpty` and nothing else. The strip is an `inkwell.Text` now
rather than a bare English string, because `TopBar` resolves a `Text`-valued
var into the player's language.

`armilla` builds on `system` — slot 2 reads through `system.System.Game()` —
and it knows nothing of `tape`: which dialog a cassette plays is `tape.Dialog`,
asked by the slot, not by the device. It is `inc/system/` because
nobody authors a file into it, diegetic object or not. Its own names stay package
functions, because `armilla` holds no `Armilla` value the way `system` holds
`System` — there is one wrist, and the package is it.

This is what lets an entity file be a literal: the tape-insert script closes
over `system.System`, not over a parameter it would have had to be handed.

## Naming

The package is the namespace, so an identifier never repeats what the package
already says. It is `scene.Floor`, not `scene.SceneFloor`; `tape.Dialog`, not
`tape.TapeDialog`; `System.Do`, not `System.DoAction`. Read the call site, not
the declaration, when choosing a name:

```go
scene.Floor   tape.Of(item)   system.System.Do(action)   widget.ScreenWidth
```

The manager is the one exception, and it keeps its prefix on purpose:
`widget.WidgetManager`, `scene.SceneManager`, `theme.ThemeManager`. A manager is
the one name that is read where the package qualifier has already fallen off —
assigned onto a `*Game` field beside the engine's own `CharacterManager`, or
called from inside its own package's entity files, where `Manager.Register`
would not say what is being registered.

Entities themselves need no name at all — they are anonymous literals inside
their file, handed straight to a manager, and the file name says which one it is.

Exported names are the authoring vocabulary — what a content file spells out:
the constants in `inc/constants.go`, each package's
`<Category>Manager` and entity type, the colour tokens, the tape lookups
(`tape.Of`, `tape.Dialog`) and `tape.Speakers`, the system
(`system.System` and its methods `Do`, `Game`, `Attach`, …), the Armilla
(`armilla.Slot2`, `armilla.Insert`, `armilla.StripIfEmpty`), the action
constructors (`System.Paused`, `System.SetMode`) and the one HUD condition a
content file has to name, `widget.WhenPlaying`. A tape's line is `inkwell.Say`
like anyone else's — the tape voice is on the `Speaker` the manager derives,
not on a second spelling of Say.

Everything else stays unexported, and now the compiler holds the line: the HUD
widgets (`hudFrame`, `slots`, …), the conditions `whenCutscene` and `whenDebug`,
and the file-local helpers.

The word is **scene**, the engine's own. The wiki and the concept-art deck count
*screens*, and this code used to as well, but everything a screen had that a
scene did not now lives in inkwell. "Screen" survives only where it means the
display: `ScreenWidth`, `ScreenHeight`.

## Layering

The layering is the import graph, so the compiler keeps it:

```
inc ← theme ← system ← armilla ← widget ← tape ← main
 ↑      ↑        ↑                  ↑      ↑
util    └── content ────────────────┴──────┘
```

`inc` imports nothing and everything imports it, and `util` sits beside it on
the same footing — standard library only, no game in it. `system` knows nothing
about the HUD or the content; both build on it, never the other way round — and
an arrow pointing back is now an import cycle, not a review comment.

Three edges inside the arrow are worth knowing, and they all point from a
thing to what it is made of or derived from. `tape → widget`: the tape slots
are a cassette thing that registers a widget, so the cassette sits above the
HUD rather than below it — this is the arrow that used to run the other way.
`stage → widget`: a screen registers the menu or card that draws it, for the
same reason. And `background → scene`: an asset is derived from the deck, so
the deriving package imports the source.

`main` is at the very top, and it is the only place that names every package at
once. `theme`, `system`, `armilla` and `widget` are `inc/system/`; the other six
are `inc/content/`, and no arrow leaves `system` for `content` any more.

## Adding a scene

1. One constant in `inc/constants.go`: `Scene<Name>`
2. `inc/content/scene/scene.<name>.go` — a `var _ = SceneManager.Register(Scene{…})`
   with its `Title` and every `Label` written in both languages
3. A 640×380 PNG at `assets/bg/<name>.png`

Nothing else moves. There is no list to update, and nothing names the picture:
a scene that says nothing about its `Background` takes an asset named after
itself, and `background.manager.go` registers one per registered scene. The PNG
file name is the pairing.

## The engine next door

inkwell is checked out beside this repository, and so is the wiki that
documents it. Paths are relative to the game's root:

```
../../engines/inkwell                                           the engine source
../../services/wiki-pages/pages/engines/inkwell/default.en.md   the manual, English
../../services/wiki-pages/pages/engines/inkwell/default.hu.md   the manual, Hungarian
```

The engine's module path is `git.teletypegames.org/engines/inkwell` and `go.mod`
pins a pseudo-version of it, so a local edit is invisible here until it is
pushed:

```
git -C ../../engines/inkwell commit … && git -C ../../engines/inkwell push
go get git.teletypegames.org/engines/inkwell@master
```

A `replace` directive is fine while trying something out, but it never survives
into a commit — drop it and bump the pin instead.

Three things move together when the engine changes: the code, `README.md` in
the inkwell checkout, and **both** wiki pages. The two pages are a translation
pair, section for section — an edit to one is an edit to the other. The
no-comment rule stops at the module boundary: inkwell's own source is commented,
and edits there follow its style, not this one's.

## Commands

```
make build     native binary into bin/
make debug     native binary with the debug menu on, -tags debug
make wasm      js/wasm build into dist/
make watch     rebuild on change
go run . -scene <name>     start on a given scene
go run . -finale           start with the finale
go run . -debug            with the debug menu on d
```

The debug menu is the engine's `Menu` widget with `inkwell.DebugMenuRoot`
handed to it, registered in `widget.debug_menu.go`. Either `-debug` or a
`-tags debug` build sets `inc.FlagDebug`, which is what its `whenDebug`
condition reads; without the flag the widget never ticks. `README.md` has the
rest, under **The debug menu**.

Art is embedded into the binary (`//go:embed assets` in `main.go`), because
js/wasm has no OS filesystem and `make binaries` ships the executable alone.

The screen is 640×380 because the engine stretches a background over the whole
window: any other size squashes the paintings.
