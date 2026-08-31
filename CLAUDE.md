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
`mousePoint`, not `mp`; `ScreenWidth`, not `ScreenW`; `BackgroundAlley`, not
`BgAlley`. A one-letter name is a comment the compiler cannot check, and this
file already forbids the checkable kind.

Receivers are spelled out too, which means Go's one-letter receiver convention
does not apply here: `func (frame *hudFrame)`, `func (channel *tapeChannel)`,
`func (slots *tapeSlots)`. The receiver takes the noun the type is *an instance
of*, so it never has to repeat the type name and never shadows it.

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
  `character`, `tape`, `dialog`, `script`, `world`, `armilla`, `widget`,
  `theme`.
- `[category].manager.go` is the file that ties a package together — its entity
  type, its `Manager`, and whatever else the category shares.
- Every other file in a package holds exactly one entity: one scene, one
  background, one item. The file registers it itself, in an `init()`.
- Two files have no category. `inc/constants.go` is the vocabulary every
  package imports, and `inc/system/boot/boot.go` builds the game.
- `inc/util` is in neither group, beside `inc/constants.go` rather than under
  one of them, for the same reason: it is a helper over the standard library
  with no game in it, so it imports nothing of ours and anything may import it.
  `util.color.go` holds `RGB` and `RGBA`, the hex-to-`color.Color` pair the
  themes are written in.

```
main.go                                  flags + inkwell.Run
inc/constants.go                         entity names and world-state keys
inc/util/util.color.go                   RGB and RGBA, hex to color.Color

inc/system/boot/boot.go                  New(Options) builds the game
inc/system/widget/widget.manager.go      the Widget alias, HUD layout, the conditions
inc/system/widget/widget.*.go            one widget per file, registering itself
inc/system/world/world.manager.go        the engine binding, the mode, the note
inc/system/world/world.action.go         custom actions
inc/system/armilla/armilla.manager.go    Paul's Armilla: two slots, the strip
inc/system/theme/theme.manager.go        the Theme alias and its Manager
inc/system/theme/theme.realworld.go      realworld-93
inc/system/theme/theme.nokia_punk.go     nokia-punk

inc/content/scene/scene.manager.go       the Scene alias, its Manager, the floor
inc/content/scene/scene.selector.go      the map screen: pins derived from the graph
inc/content/scene/scene.*.go             one file per scene
inc/content/background/background.*.go   one image asset per scene
inc/content/character/character.*.go     the cast, tapes included
inc/content/item/item.*.go               inventory
inc/content/dialog/dialog.*.go           dialogue trees
inc/content/script/script.*.go           named action sequences
inc/content/tape/tape.manager.go         the Tape entity and the lookups over it
inc/content/tape/tape.*.go               one tape per file
```

`inc` is constants and nothing else, so every package can import it and it can
import none of them. That is also why `New` sits in `inc/system/boot` rather than
in `inc`: a package cannot be imported by what it imports, and the assembly is
the one thing that has to name every package at once.

The two groups say who writes a file, not which way an import points. `content`
is what a writer adds a file to; `system` is the machine that runs it. Mostly
that agrees with the graph — a character names a `theme` colour, a scene names a
`world` action, and neither `theme` nor `world` knows the caller back — but two
edges cross the other way, and both are meant to. `boot` names every content
package at once, which is what assembly *is*. And `widget` reads `tape`, because
the tape slots ask which cassette a tape is in.

So the grouping is not the layering. The layering is the import graph and the
compiler keeps it on its own, whichever directory a package sits in; the
directories are there so that opening `content/` shows the game and opening
`system/` shows the machine. A cassette is something somebody writes — a voice,
an item, a dialogue — so `tape` is content, even though the HUD reaches across
for it.

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

Read from `inc/system/boot` the nine of them still line up as a list —
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

`Tape` is the exception, and it is a real one: a tape is a cassette that speaks,
a thing inkwell has no notion of. It names the character whose voice it is, the
item that carries it and the dialogue it plays, and `tape.manager.go` holds the
four lookups over the registry — `tape.Is`, `tape.Of`, `tape.IsItem`,
`tape.Dialogue`. What a tape *sounds* like is not in it: the display name and the
log-only voice are `Character.Label` and `Character.Voice`, because those are
facts about a speaker, not about a cassette.

### Entities register themselves

An entity file is a literal and nothing else. No constructor function, no list
somewhere else to keep in step — the file hands itself to its manager in an
`init()`:

```go
package background

func init() {
	BackgroundManager.Register(Background{
		Name: inc.BackgroundServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	})
}
```

Adding an entity is adding a file. Deleting one is deleting a file. Package-level
variables are initialised before any `init()` runs, whatever file each sits in,
so a package's manager exists by the time its first entity file registers into
one.

### A file's helpers hang off the file's own type

An entity file that is too long to read as one literal breaks pieces out —
`scene.alley.go` builds its four hotspots that way. Those helpers are **not**
package functions. Twenty-four scene files share one `scene` namespace, so a
package-level `bin()` is the *package's* bin, and the next scene that wants one
cannot have it. The file declares an empty type instead, named after itself, and
hangs its helpers off that:

```go
type alley struct{}

func init() {
	a := alley{}
	SceneManager.Register(Scene{
		…
		Hotspots: []inkwell.Hotspot{a.hidingPlace(), a.backDoor(), a.bin(), a.fireEscape()},
	})
}

func (alley) bin() inkwell.Hotspot { … }
```

The receiver is the namespace, so the method name drops what the type already
says — `selector.leadsTo`, not `leadsToSelector`. Where the helpers read from
data rather than build it, the type is not empty and the data are its fields:
`script.usewith_fail.go` carries the four line-sets as fields of `useWithFail`
and answers with `fail.paul(n)` and `fail.dex(n)`.

The rule stops at the package's own API. `world.Do`, `armilla.Slot2`,
`tape.Dialogue` and `boot.New` stay functions, because there the package is
already the namespace and the receiver would be an empty word.

`init()` order is file-name order, so **registration order is alphabetical**.
Nothing may depend on it — including the order the arrow keys walk the scenes,
which is simply the order the files sit in.

An `init()` only runs if something imports the package. `inc/system/boot` names
eight of the nine for their manager, which is import enough; `tape` is the one
nobody names, so it is there as a blank import. That list is **per package, not
per entity** — a new scene file still touches nothing but itself.

### Widgets name their layer instead of their order

A widget is a file like any other entity — `widget.<name>.go`, one widget,
registering itself in an `init()` — and that includes the engine's own widgets,
which the game registers as literals the same way it registers a background:

```go
func init() {
	WidgetManager.Register(&inkwell.Cursor{
		Name: "cursor",
	})
}
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
WidgetManager.Register(&inkwell.StatusLine{
	Name: "status",
	When: whenPlaying,
	Y:    statusY,
})
```

`whenPlaying`, `whenCutscene` and `whenUnpaused` are in `widget.manager.go`, and
they read `inc.VarMode` and `inc.VarNote` out of the engine's own `State`. A widget of
ours says the same thing with a method — `VisibleWhen() inkwell.Condition` —
because it has no literal to put a field in. A widget that is switched off
neither ticks nor draws nor blocks a click, so nothing else has to ask.

### Handing a category to the engine

Nothing is copied. `inkwell.NewGame` builds its own empty managers, and `New`
hands ours over in their place — the types are identical, so the engine and the
game end up sharing one registry per category rather than two in step:

```go
g.AssetManager = background.BackgroundManager
g.CharacterManager = character.CharacterManager
g.DialogueManager = dialog.DialogManager
g.ItemManager = item.ItemManager
g.SceneManager = scene.SceneManager
g.ScriptManager = script.ScriptManager
g.WidgetManager = widget.WidgetManager
theme.ThemeManager.Each(g.ThemeManager.Register)
```

Themes are the odd one out and stay a copy: `NewGame` puts four preset themes
into its `ThemeManager`, and replacing it would throw them away. Ours are added
to that set instead.

One consequence of not copying: a manager swapped in this way must be in place
before `inkwell.Run`, which is where the engine wires up the parts that hold a
registry directly. `fillSelectorPins` runs before the hand-off for the same
reason — it writes the derived pins back into `SceneManager` with `Set`.

The defaults a scene leaves out are no longer written into it at all. `g.Player`
names the character a scene with no actors of its own gets, placed at his
`Start`; `g.Walkboxes` is the floor a scene walks on when it declares none. Both
are fields on the `*Game`, so a scene file that says nothing about either is
saying "the usual", and the selector opts out by declaring both empty.

### What runs at boot

`New` names two scripts and nothing else — the opening a player sees is content
like every other script, in `script/script.opening.go`, not a slice built in the
wiring:

```go
g.StartAt(start)
g.OnStart(inc.ScriptOpening)
if o.Finale {
	g.OnFinale(inc.ScriptFinale)
}
```

`OnFinale` is the engine's hook for a closing script, queued straight after the
opening. Here it is what `-finale` uses to drop into the ending, which is why it
is set from `Options` rather than always.

The rest of `New` is fields on the `*Game`, and every one of them replaces
something this package used to carry itself:

```go
g.WindowScale = 2
g.Player = inc.Paul
g.Walkboxes = scene.Floor
g.UseWithFail = script.UseWithFail
```

`UseWithFail` is the response to a use-with pair nobody authored — the engine
asks for it the same way it asks for `ExitLook` and `ExitTake`, and the lines
live with the rest of the writing, in `script/script.usewith_fail.go`.

## The world

There is exactly one game, so there is exactly one world — and since there is
exactly one, the **package is the singleton**. What the singleton holds is a
`state` struct with one unexported instance, so the mutable fields sit together
and `Attach` resets them in a single assignment rather than one line per field —
a fifth field cannot be forgotten. `armilla` does the same. Only the state is a
struct: the exported names stay package functions, because a `*World` threaded
to every widget and script is the thing this design removed. There is no `World` value to pass
around and no back-reference for a widget to hold: `world.Do(…)`,
`world.Game()`, `world.Func(…)` are reachable from anywhere that imports `world`.
`New` calls `world.Attach(g)`, which binds the engine and resets the run state,
so building the game twice is clean.

`world` holds as little as it can get away with, and after the Armilla moved out
it holds nothing but the binding: the engine pointer, the mode, and the top
bar's note. The mode and the note are `State` vars (`inc.VarMode`,
`inc.VarNote`), because state the engine can see is state a `Condition` can read
— that is what makes a widget's `When` possible and what puts "— paused" in the
top bar without anyone pushing it there. `world.Do` hands its action to `g.Do`,
the engine's queue, so the game runs one action at a time without a pump of its
own.

## The Armilla

The Armilla is the wrist-worn runtime a Personal Tape snaps into — the wiki's
term, and the game's device. Paul wears one, and it is the only piece of the
fiction the code models as a package of its own, because it is the only one that
holds state: two slots, a phosphor text strip, and a tape with something to say.

```go
armilla.Slot2()                      armilla.SetSlot2(item)
armilla.SetPending(item)             armilla.TakePending()
armilla.SetStrip("SLOT2: READING")   armilla.StripIfEmpty("SRP: 1 tape")
armilla.Offer(line)                  armilla.Waiting()   armilla.TakeOffer()
```

Slot 2 and the strip are `State` vars — `inc.VarSlot2` and `inc.VarArmillaStrip`,
both spelled `armilla.…`, so the save file says the same word the package does.
The other two are the things the engine has no notion of: the cassette on its way
into the slot, and the tape waiting to speak. Only `armilla.Reset()` clears them,
and `New` calls it straight after `world.Attach`.

`armilla` builds on `world` — slot 2 and the strip go through `world.Game()`, and
`TakeOffer` hands its line to `world.Do` — and it knows nothing of `tape`: which
dialogue a cassette plays is `tape.Dialogue`, asked by the HUD, not by the
device. It is `inc/system/` because nobody authors a file into it, diegetic
object or not.

This is what lets an entity file be a literal: the tape-insert script closes
over the `world` package, not over a parameter it would have had to be handed.

## Naming

The package is the namespace, so an identifier never repeats what the package
already says. It is `scene.Floor`, not `scene.SceneFloor`; `tape.Dialogue`, not
`tape.TapeDialogue`; `world.Do`, not `world.WorldDo`. Read the call site, not
the declaration, when choosing a name:

```go
scene.FillSelectorPins()   tape.Is(speaker)   world.Do(action)   widget.ScreenWidth
```

The manager is the one exception, and it keeps its prefix on purpose:
`widget.WidgetManager`, `scene.SceneManager`, `theme.ThemeManager`. A manager is
the one name that is read where the package qualifier has already fallen off —
assigned onto a `*Game` field beside the engine's own `CharacterManager`, or
called from inside its own package's entity files, where `Manager.Register`
would not say what is being registered.

Entities themselves need no name at all — they are anonymous literals inside
their file's `init()`, and the file name says which one it is.

Exported names are the authoring vocabulary — what a content file spells out:
`boot.New` and `boot.Options`, the constants in `inc/constants.go`, each package's
`<Category>Manager` and entity type, the colour tokens, the tape lookups
(`tape.Is`, `tape.Of`, `tape.IsItem`, `tape.Dialogue`), the world (`world.Do`,
`world.Game`, …), the Armilla (`armilla.Slot2`, `armilla.SetStrip`, …) and the
action constructors (`world.Paused`, `world.SetMode`,
`world.UseTheme`, `armilla.Offer`, `world.Func`). A tape's line is `inkwell.Say`
like anyone else's — the tape voice is on the character, not on a second
spelling of Say.

Everything else stays unexported, and now the compiler holds the line: the HUD
widgets (`tapeSlots`, `letterbox`, `hudFrame`, …), the conditions
(`whenPlaying`, `whenCutscene`, `whenUnpaused`), the runners, `setMode`,
`setNote`, `setVarIfEmpty`.

The word is **scene**, the engine's own. The wiki and the concept-art deck count
*screens*, and this code used to as well, but everything a screen had that a
scene did not now lives in inkwell. "Screen" survives only where it means the
display: `ScreenWidth`, `ScreenHeight`.

## Layering

The layering is the import graph, so the compiler keeps it:

```
inc ← theme ← world ← armilla ← tape ← widget ← boot ← main
 ↑      ↑        ↑                              ↑
util    └── content ───────────────────────────-┘
```

`inc` imports nothing and everything imports it, and `util` sits beside it on
the same footing — standard library only, no game in it. `world` knows nothing about the
HUD or the content; both build on it, never the other way round — and an arrow
pointing back is now an import cycle, not a review comment. The one edge worth
knowing is `widget → tape`: the tape slots ask which cassette a tape is in, so
the tape concept sits below the HUD rather than beside it.

`theme`, `world`, `armilla` and `widget` are `inc/system/`, and so is `boot` —
not because it sits low, it sits at the very top, but because nobody authors
anything in it. The other seven, `tape` among them, are `inc/content/`, which
makes `widget → tape` the one edge that leaves `system` for `content`.

## Adding a scene

1. Constants in `inc/constants.go`: `Scene<Name>`, `Background<Name>`
2. `inc/content/scene/scene.<name>.go` — an `init()` registering a `Scene`, its
   `Title` and every `Label` written in both languages
3. `inc/content/background/background.<name>.go` — an `init()` registering a
   `Background`
4. A 640×380 PNG in `assets/bg/`

Nothing else moves. There is no list to update.

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
make wasm      js/wasm build into dist/
make watch     rebuild on change
go run . -scene <name>     start on a given scene
go run . -finale           start with the finale
```

Art is embedded into the binary (`//go:embed assets` in `main.go`), because
js/wasm has no OS filesystem and `make binaries` ships the executable alone.

The screen is 640×380 because the engine stretches a background over the whole
window: any other size squashes the paintings.
