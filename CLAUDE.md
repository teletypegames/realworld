# Underwater Sun Rise

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
- **The engine's vocabulary**: `inc.VarNote` and `inc.VarSlot2` stay `Var`,
  because the engine's own API is `State.Var` and `State.SetVar`, and half a
  translation is worse than none.
- **`err`**, which is the language's own spelling in every Go program.

### English first, then every other language beside it

Identifiers, commit messages, `README.md` and this file are English and only
English. A string the **player** reads is different: it is an `inkwell.Text`,
carrying English first and every other language the game speaks after it, in
the same file, at the same place:

```go
OnLook: character.Paul.Say(Text{
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
return &inkwell.Asset{
	Name: "street",
	Path: "assets/bg/street.png",
}
```

Not `&inkwell.Asset{Name: "street", Path: "..."}`.

This holds for nested literals too, including small ones such as
`inkwell.Point`. Slice and map literals are not covered by the rule.

An `inkwell.Translation` is the one exception the other way: the two fields
stay on **one** line, so the translations of a string stack into a table and
the eye can run down a single column to compare them.

```go
Title: Text{
	{Lang: "en", Text: "ALLEY — BEHIND NOODLE'S HOUSE"},
	{Lang: "hu", Text: "SIKÁTOR — NOODLE HÁZA MÖGÖTT"},
},
```

A recorded line is the third field on that one line, `Voice: &inkwell.Voice{…}`,
because the English and the Hungarian recording are two files and the line
that carries the words carries its own reading of them.

Inside a `[]inkwell.Text` or a `map[string]inkwell.Text` the element type is
implied and left off, which is what `gofmt -s` produces.

## Layout

One category, one package, one directory, and every category sits in one of two
groups: `inc/system/` is the machinery, `inc/content/` is what the game is made
of. A file keeps the category in its own name anyway — `[category].[name].go`
inside `inc/[group]/[category]/` — because the prefix is what makes a file
findable in a list of editor tabs, a grep, or a diff, where the directory has
already fallen off, and now two directories have.

- The category is always **singular**: `item`, `scene`, `character`, `tape`,
  `dialog`, `script`, `armilla`, `widget`, `theme`.
- `[category].manager.go` is the file that ties a package together — its entity
  type, its `Manager`, and whatever else the category shares.
- Every other file in a package holds exactly one entity: one scene, one
  item, one screen. The file registers it itself, as a declaration.
- Two things have no category. `inc/constants.go` holds the world-state keys
  every package imports — flags, vars, notes — and `main.go` builds the game:
  flags, one `inkwell.Game` literal, and `inkwell.Run`, beside the two one-line
  files a `-tags debug` build picks between.
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
main.steam.go                            platformExtensions = the steam.Steam literal, behind -tags steam
main.nosteam.go                          platformExtensions = nil, without it
inc/constants.go                         world-state keys: flags, vars, notes
inc/util/util.color.go                   RGB and RGBA, hex to color.Color

inc/system/system.go                     System: the binding and the note
inc/system/widget/widget.manager.go      the Widget alias, HUD layout, the conditions
inc/system/widget/widget.*.go            one widget per file, registering itself
inc/system/armilla/armilla.manager.go    Paul's Armilla: slot two and the strip
inc/system/theme/theme.manager.go        the Theme alias and its Manager
inc/system/theme/theme.realworld.go      realworld-93
inc/system/theme/theme.nokia_punk.go     nokia-punk

inc/content/stage/stage.manager.go       the Stage alias, its Manager, the shared rows
inc/content/stage/stage.game.go          the world, and the key that leaves it
inc/content/stage/stage.*.go             one screen per file: splash cards, menu, load, save, dossier, …
inc/content/scene/scene.manager.go       the Scene alias, its Manager, the floor
inc/content/scene/scene.selector.go      the map screen: pins derived from the graph
inc/content/scene/scene.*.go             one file per scene
inc/content/achievement/achievement.manager.go   the Achievement alias and its Manager
inc/content/achievement/achievement.*.go one achievement per file: its Steam name, its words, the flag that earns it
inc/content/character/character.*.go     the cast: people who stand in a scene
inc/content/speaker/speaker.*.go         the voices with no body: Dex and the tape cards
inc/content/music/music.manager.go       the Music alias and its Manager: tracks a scene or stage plays
inc/content/sound/sound.manager.go       the Sound alias and its Manager: noises a hotspot plays
inc/content/item/item.*.go               inventory
inc/content/dialog/dialog.*.go           dialog trees
inc/content/script/script.*.go           named action sequences
inc/content/tape/tape.manager.go         the Tape entity and the lookups over it
inc/content/tape/tape.slots.go           the two Armilla slots, as hotspots
inc/content/tape/tape.channel.go         the log the tapes talk into
inc/content/tape/tape.*.go               one card per file: its item, its voice, the dialog it plays
```

`inc` is state keys and nothing else — the strings a `State` is indexed by,
which a save file writes down — so every package can import it and it can
import none of them. Everything that used to be a name constant beside them is
an entity variable in its own package now; see **References are pointers**. The assembly is the other end of that: it is the one thing
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
engine how a speaker's lines are **delivered**, and a tape is a `DeliveryLog`
speaker, so the filter is `Deliveries: []inkwell.Delivery{inkwell.DeliveryLog}`
and the HUD learns nothing about cassettes. The tape slots moved the other way: they are a
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

type (
	Character = inkwell.Character
	Text      = inkwell.Text
	Point     = inkwell.Point
)

var CharacterManager = inkwell.NewManager[*Character]()
```

The manager file also aliases the engine types the package writes as
**literals** — `Text` everywhere, `Hotspot`, `Exit`, `Trigger`, `SceneActor`,
`Point` in `scene`, `DialogNode`, `DialogLine`, `DialogChoice`, `Given` in
`dialog`, `MenuItem` in `stage` — so a content file says `Text{…}` and
`Hotspot{…}` rather than `inkwell.Text{…}` on every line. Only literal types
are aliased, never functions: `inkwell.Seq`, `inkwell.If`, `inkwell.Rect` keep
their qualifier, because a call reads fine with one and a literal drowns in
them. A line is spoken by its speaker, `character.Paul.Say(…)`,
`speaker.Dex.Say(…)`, which is `inkwell.Say` with the first argument moved in
front of the verb, where a script puts it.

Every manager holds **pointers**, `Manager[*Character]` and never
`Manager[Character]`, so the variable an entity file binds and the entry the
engine looks up are one instance. Read from `main.go` the ten of them still
line up as a list — `achievement.AchievementManager`, `character.CharacterManager`, `dialog.DialogManager`, `item.ItemManager`,
`scene.SceneManager`, `script.ScriptManager`, `speaker.SpeakerManager`,
`tape.TapeManager`, `theme.ThemeManager`, `widget.WidgetManager` — only now
the list is a consequence of the imports rather than a block someone has to
keep in step.

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
has a sprite. A tape file is the card in one literal — the `Item` that carries
it, the `Speaker` that is its voice, the `Dialog` it plays and `OnInsert`, the
first thing it says when it clicks into slot two. All three are pointers to the
entity itself, in the package that owns it. The dialog is a file in `dialog/`
like every other tree, because a conversation is written and read on its own,
and the card only has to say which one is its. `OnInsert` runs after
`script.tape_insert.go`, which is the part of an insertion every card shares —
Paul's hand and Dex's remark — so the generic script never has to know which
card went in.

The voice is declared, not derived. `speaker/` holds one `inkwell.Speaker` per
voice with no body — `speaker.Dex`, `speaker.TapeSupport`, … — each with its
`Label`, its `Color` and `DeliveryLog`, so the line lands in the channel and never
in a bubble; `speaker.SpeakerManager` is what `main.go` hands the engine. It
used to be derived from the cards, and that is exactly why a dialog could not
name the voice it was giving a line to: the voice lived one package *above*
it. Dex is such a voice and such a card — `speaker.Dex` says how he sounds,
`tape.Dex` is the cassette in slot one — which is why he is in neither
`character/` nor only in `tape/`. One voice stays a character on purpose: the
security ACP is a room-scale system, so it stands in its scene as a
`Character` with `Delivery: inkwell.DeliveryLog`, and speaks through the log. The
two lookups are `tape.Of`, the card an item is, and `tape.DialogOf`, the
dialog to run for a loaded item, which is the card's own or Dex's when the
card names none.

What a dialog answers with, and when, is the engine's. A choice that is only
offered under a condition says so in `Show`, and the condition is written as a
literal, `inkwell.Given`, with every requirement a field: `Flag`, `NotFlag`,
`Item` and `NotItem` as the `*Item` itself, `Var` with `Is`. An answer that branches is an `inkwell.If`
over the same `Given`; an answer that has to be computed is
`inkwell.Reply(speaker, func(*inkwell.State) inkwell.Text)`, `Say` with the
line decided when it is spoken.

### Entities register themselves

An entity file is a literal and nothing else. No constructor function, no list
somewhere else to keep in step, and no procedure either — the file *is* the
declaration that hands itself to its manager, and binds the result to the one
exported variable every other file refers to it by:

```go
package tape

var Support = TapeManager.Register(&Tape{
	Name:    "tape_support",
	Item:    item.SupportTape,
	Speaker: speaker.TapeSupport,
	Dialog:  dialog.SupportTape,
})
```

`Register` returns what it registered, which is the whole of what makes this
possible: a statement can only stand in a function body, an expression can
stand in a variable initialiser. The variable is the entity's name in code;
the `Name` field is its name in a save file, and nowhere else.

### References are pointers

A reference to an entity is the entity: `Item: item.SupportTape`, not a string
somebody has to look up. The engine's fields and constructors are typed that
way — `Stage.Background *Asset`, `SceneActor.Character *Character`,
`Say(speaker Talker, …)`, `Give(*Item)`, `GoToStage(*Stage)` — so a name
nobody registered is a compile error, the editor follows a field to the struct
it names, and `inc/constants.go` is left with the strings that really are
strings: the keys a `State` is indexed by.

The cost is that a reference is an import, and imports may not form a cycle,
so the content packages stand in a strict order and a reference only ever
points **down**:

```
achievement, character, speaker, music, sound ← item ← script ← dialog ← tape ← scene ← stage
```

A dialog can give a line to `speaker.Dex` and hand out `item.NoodleLetter`,
but it cannot move the player to a scene: `scene` is above it. That is not a
gap in the engine but the price of the compiler holding the graph, and the
answer is the one the layering always gave — the dialog sets a flag, and the
scene's trigger does the moving.

**Exits come after the scene.** Go forbids two package-level variables from
referring to each other in their initialisers, and two scenes that lead to
each other do exactly that. So a scene file declares the scene first and its
exits second, with `Leads`, which appends to `Exits` and hands the scene back:

```go
var Alley = SceneManager.Register(&Scene{ … })

var _ = Alley.Leads(
	inkwell.Exit{To: NoodleHouse, Side: inkwell.ExitLeft, …},
)
```

The second declaration depends on both scenes and neither scene depends on
it, so `Exit.To` holds the `*Scene` itself and the graph may be as tangled as
the city. The selector's pins are `Derived` exits, built at `Run` after every
variable exists, and `Derived` is handed the scene itself so that "who leads
to me" never names its own variable.

**What stays a string.** `Name`, because a save file, a `State` var and the
`-scene` flag have to find the entity again from outside the process — `main`
turns the flag into a scene with `scene.SceneManager.MustGet`. Slot two is a
`State` var holding an item's `Name`, and `armilla.Slot2()` turns it back into
the `*Item` through the engine's `ItemManager`. Hotspot and trigger names, which
are labels inside one literal. Dialog node names, which are inside one tree.
And every flag and var, which is what `inc` is for.

This is not only two lines and an indent level shorter than an `init()`. A
package-level variable is initialised in **dependency order**, so the manager
is built before any file that writes into it, whichever one the compiler reads
first — with `init()` that rested on the separate rule that every package
variable is ready before any `init()` runs. The registration order is
unchanged: still the order the files are presented in, which is alphabetical,
and still nothing may depend on it.

One file keeps an `init()`, and it earns it. `tape.slots.go` builds a value
that has to be wired to itself — `panel.Hotspots = panel.hotspots` — which a
composite literal cannot do. Where a file is assembly before it is an entity,
`init()` is the honest word for it; contorting it into a declaration would cost
more than it saves.

A scene's picture is in the scene file like everything else about it: `Layers`,
full-screen planes at a depth `Z`, with the cast at 0. A negative `Z` is behind
every character and a positive `Z` in front of every character, so the
background is the layer at `Z: -1`, and its `Asset` is a literal nested in
place, named after the scene and paired with a PNG of the same name:

```go
Layers: []SceneLayer{
	{
		Asset: &Asset{
			Name: "alley",
			Path: "assets/bg/alley.png",
		},
		Z: -1,
	},
},
```

What Paul walks behind is one more element of the same slice, at a positive
`Z`, with its own PNG. There is no `background` package and no asset
manager of ours: a picture that belongs to exactly one thing needs neither a
file nor a variable nor a registry, since the engine loads it through the
pointer the layer holds. The splash logo is the same shape, a literal nested
in its stage's `Background`.

Adding an entity is adding a file. Deleting one is deleting a file.

### A file has no type and no methods

An entity file is the variable and the literal, and the variable **is** the
entity: `Alley` is the `*Scene`, `Trinket` the `*Item`, `Keypad` the `*Dialog`.
There is no empty type named after the file, no method hanging off it, and no
helper with a name the next file will not have. The engine's own interfaces
are the only methods in `inc/content/` — the tape slots' `Draw`, a `Tape`'s
`GetName` — and those are the engine's words, not ours.

What a helper used to carry lives in one of three places:

- **A line two hotspots share** is a package-level variable before the
  literal, named with the file's own name in front so twenty-four scene files
  can share one namespace: `policeStreetNoticed`, `noodleHouseGone`,
  `alleyLocked`.
- **A piece that takes a parameter, or has to loop**, is a package-level
  function named the same way, after the field it fills:
  `selectorDerived`, `keypadNodes`, `loadExpand`.
- **The exits** are the second declaration of the file, `var _ = Alley.Leads(…)`,
  under **References are pointers**.

Everything else stays in the literal, however long it gets. A scene file reads
as its shared lines, then the scene, then where it leads; the reader never
meets a name invented for that file alone, and clicking `Alley` anywhere in
the game lands on the scene, not on a type standing in for it.

The one struct with methods in `inc/content/` is the tape slots' widget,
because the engine asks a widget to `Draw` itself; its two rectangles are
fields computed once in `init()`, not helpers recomputed per call. The rule
stops where the package is already the namespace: `armilla.Slot2`,
`armilla.Insert` and `tape.DialogOf` stay functions, because a receiver there
would be an empty word. `system` is the other way round — the receiver is
the whole point, so everything it offers is a method on `System`. `main` is
the same as a content file in spirit: `options` carries the flags, and its
methods are the `Game` fields they answer — `startStage`, `startScene`,
`finaleScript`, `flags` — plus `game`, the literal itself.

Initialisation order is file-name order, so **registration order is
alphabetical**. Nothing may depend on it — including the order the debug menu
lists the scenes in, which is simply the order the files sit in.

A package initialises only if something imports it. `main.go` names every one
of them for a manager, which is import enough; `tape` is the one it imports
blank, because the cards register into `speaker` and `widget` managers and
nothing in the literal is theirs. That list is **per package, not per entity**
— a new scene file still touches nothing but itself.

### Widgets name their layer instead of their order

A widget is a file like any other entity — `widget.<name>.go`, one widget,
registering itself — and that includes the engine's own widgets,
which the game registers as literals the same way it registers a scene:

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
there on the world stage and nowhere else, and that is a condition over game
state, not a wrapper and not an `if` at the top of every `Draw`:

```go
var _ = WidgetManager.Register(&inkwell.StatusLine{
	Name: "status",
	When: WhenPlaying,
	Y:    statusY,
})
```

`WhenPlaying` and `whenDebug` are in `widget.manager.go`: the first is
`inkwell.OnWorld()` and nothing else, the second adds `inc.FlagDebug` read out
of the engine's own `State`. `WhenPlaying` is the one of the two that is
exported, because `tape.slots.go` is a content file that has to name it; the
other is read inside the package only. There is no cutscene mode: a script that
holds the frame holds it with the HUD still on screen, so a cutscene sets
nothing before it starts and clears nothing when it ends. A widget of
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
	AchievementManager: achievement.AchievementManager,
	CharacterManager:   character.CharacterManager,
	DialogManager:      dialog.DialogManager,
	ItemManager:        item.ItemManager,
	SceneManager:       scene.SceneManager,
	ScriptManager:      script.ScriptManager,
	SpeakerManager:     speaker.SpeakerManager,
	StageManager:       stage.StageManager,
	ThemeManager:       theme.ThemeManager,
	WidgetManager:      widget.WidgetManager,
})
```

`New` fills in an empty manager for every category a caller leaves out, so the
ten above are the ten we own and the other two are the engine's:
`VerbManager`, the four SCUMM verbs untouched, and `AssetManager`, which stays
empty because a picture is loaded through the pointer its layer or stage holds
and is registered nowhere.

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
	Title:       "Underwater Sun Rise",
	App:         "realworld",
	Width:       widget.ScreenWidth,
	Height:      widget.ScreenHeight,
	WindowScale: 2,

	StartStage:   chosen.startStage(),
	StartScene:   chosen.scene(),
	StartScript:  script.Opening,
	FinaleScript: chosen.finaleScript(),
	StartTheme:   theme.RealWorld,
	Flags:        chosen.flags(),
	Extensions:   platformExtensions,

	Player:      character.Paul,
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
straight after the opening; `-finale` is what fills it, and `nil` is an
ordinary run. `Flags` is the same shape for `-debug`: a boot flag list, not a
reach into `State` after construction. Every one of these is the entity
itself — `script.Opening`, `theme.RealWorld`, `character.Paul` — and the one
that arrives as a string, the `-scene` flag, becomes one through
`scene.SceneManager.MustGet` before it reaches the literal.

`ExitLook`, `ExitTake` and `UseWithFail` are the three responses the engine asks
the game for, and all three are `script` functions — a line the player reads
belongs with the rest of the writing, not in the wiring.

`Extensions` is the platform: what lives beside the game for the whole run and
is not a widget. It is `platformExtensions`, a variable one of two build-tag
files declares, the way `debugBuild` is — `main.steam.go` behind `-tags steam`
holds the one `steam.Steam` literal, with the AppID, the Steam-name-to-code
language map and the `Store` keys that go to Steam Cloud, and `main.nosteam.go`
holds `nil`. So a build without the tag links none of the engine's `steam/`
package and none of the binding underneath it, and the settings of the
integration are read in one place, in the game and not in the engine. What the
player is recognised for is content: `inc/content/achievement/`, one file per
achievement, each a `When` over an `inc` flag, and the engine reports the
rising edge to whichever extension can take it. `README.md` has the rest,
under **Steam**.

`App` is the game's name to the machine: the engine keeps `settings.json` and
`slots.json` under it in the user's configuration directory, and under the
same prefix in `localStorage` for wasm. Nothing here writes either — `UseLang`
and `UseTheme` save the settings as they run, `SaveSlot` writes a slot — and
`-lang` sets `inkwell.Lang` for the run only, which the engine leaves alone
when it is already set.

The command line is a type rather than four loose pointers: `options` carries
what the flags chose, and the two answers the literal needs that a flag does not
give directly are methods on it — `chosen.finaleScript()`, which is `nil` for
an ordinary run, and `chosen.flags()`, which is where `-debug` and `debugBuild`
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
screen floats over; it is one of the screens, `stage.Game`, the only one
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

The HUD needed one line, not fourteen. `widget.WhenPlaying` is
`inkwell.OnWorld()`, so every HUD widget disappears off the game stage
without knowing stages exist — and without naming one, which it could not:
`stage` registers its menus into `widget`, so `widget` naming `stage.Game`
back would be an import cycle. `inkwell.OnWorld` says the thing the HUD
means, "the Stage showing is the world", whichever Stage that is.

`Restart`, `LoadSlot` and `GoToStage` all land at the top of the next frame,
which is what lets a menu row say the whole of what it means at once:

```go
Do: inkwell.Seq(inkwell.Restart(), inkwell.GoToStage(stage.Game)),
```

The front door is two splash cards before the menu, `stage.splash_studio.go`
— the Teletype Games logo, a 640×380 `Background` nested in the literal as the
asset it is — and `stage.splash_title.go`, and each is a `Stage` that says
how long it shows and where it goes: `Hold: 2`, `Next`, `Skippable` so that
Space, Enter or a click moves on early, and a `Transition` of
`inkwell.Fade(0.5)` it leaves through. A `Transition` is the leaving Stage's,
so the menu and its sub-screens, which name none, keep cutting. The two
numbers are shared in `stage.manager.go`.

`-scene` and `-finale` skip the front door: both mean somebody wants to look
at the game, not at its splash cards or main menu, so `options.startStage`
answers `stage.Game` for either.

**The dossier** is the one screen that reads the world back to the player:
who Paul has met and what he has held, each with a paragraph. It is
`stage.dossier.go`, a `Stage` and the `inkwell.Menu` that is it, with two
`Expand` pages — People, Things — and a `Body` on every row, so a name opens
its text rather than running anything. What a row says is the entity's own
`Lore`, a `Text` beside its `Label` or `Description` in the character, item
or speaker file, stacked in both languages like every other line; an entity
with no `Lore` is never listed. Who is *in* it is the engine's `State.Known`:
a character or a voice once it has said a line, an item once it has been in
the inventory, and anything else once a script says `inkwell.Learn(…)`, which
is how Norman gets a page from a name on a list. The rows come in the order
of discovery, not the order of the files. Paul has a page of his own. The
screen opens from the main menu and from `Tab` on the world stage, and
leaves through `inkwell.GoBack()`, so it returns to whichever of the two
opened it.

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
`System.Paused(…)`. There are no package functions beside it, and no `Action`
or `Runner` types of its own either — nor any closure now: `Paused` is
`inkwell.Noting`, the engine's own "hold this var for exactly as long as the
action runs". `System.Func`, the escape hatch to arbitrary Go, is gone with its
last caller, and so is `System.UseTheme`, which is now `inkwell.UseTheme` — a
theme switch mid-script is not this game's idea. `System.SetMode` went with the
play/cutscene mode itself: a cutscene no longer takes the HUD down, so there
was nothing left for a mode to switch. One short file
holds the whole package. The mutable state is the struct's
fields rather than package variables, so `Attach` is a hand-off to one value and
not a sweep over several — and nothing is threaded anywhere: a widget holds no
back-reference and a script takes no parameter, because `system.System` is
reachable from anywhere that imports `system`. `main` calls
`System.Attach(game)`, which binds the engine and resets the run state, so
building the game twice is clean.

`system` holds as little as it can get away with, and after the Armilla moved out
it holds nothing but the binding: the engine pointer and the top bar's note.
The note is a `State` var (`inc.VarNote`), because state the engine can see is
state a `Condition` can read — that is what puts "— paused" in the top bar
without anyone pushing it there. `System.Do` hands its action to
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
armilla.Slot2()          armilla.Insert(cassette)      armilla.StripIfEmpty(text)
```

Both are `State` vars — `inc.VarSlot2` and `inc.VarArmillaStrip`, both spelled
`armilla.…`, so the save file says the same word the package does. The slot
holds the cassette's `Name`, because a `State` var has to survive a save;
`Slot2()` hands back the `*inkwell.Item` itself, found through the engine's
`ItemManager`, and `Insert` takes one. There is no
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
not a package holding a flag a widget polls.

`armilla.Insert` is a `Seq` of three engine actions and no Go: set slot two
to the cassette's name, take the cassette out of the inventory, write the
strip. `StripIfEmpty` is
`inkwell.SetVarIfEmpty` and nothing else. The strip is an `inkwell.Text` now
rather than a bare English string, because `TopBar` resolves a `Text`-valued
var into the player's language.

`armilla` builds on `system` — slot 2 reads through `system.System.Game()` —
and it knows nothing of `tape`: which dialog a cassette plays is
`tape.DialogOf`, asked by the slot, not by the device. It is `inc/system/` because
nobody authors a file into it, diegetic object or not. Its own names stay package
functions, because `armilla` holds no `Armilla` value the way `system` holds
`System` — there is one wrist, and the package is it.

This is what lets an entity file be a literal: the tape-insert script closes
over `system.System`, not over a parameter it would have had to be handed.

## Naming

The package is the namespace, so an identifier never repeats what the package
already says. It is `scene.Alley`, not `scene.SceneAlley`; `item.Trinket`, not
`item.ItemTrinket`; `tape.DialogOf`, not `tape.TapeDialogOf`; `System.Do`, not
`System.DoAction`. Read the call site, not the declaration, when choosing a
name:

```go
scene.Alley   item.Trinket   speaker.Dex   tape.Of(cassette)   system.System.Do(action)
```

Inside its own package an entity drops the qualifier too — `SupportTape` in an
item file, `MainMenu` in a stage file — because that is what the call site
reads there.

The manager is the one exception, and it keeps its prefix on purpose:
`widget.WidgetManager`, `scene.SceneManager`, `theme.ThemeManager`. A manager is
the one name that is read where the package qualifier has already fallen off —
assigned onto a `*Game` field beside the engine's own `CharacterManager`, or
called from inside its own package's entity files, where `Manager.Register`
would not say what is being registered.

An entity's name in code is the exported variable its file binds —
`scene.Alley`, `dialog.Keypad`, `speaker.TapeSupport` — and it is the name the
file carries too: `scene.alley.go` declares `Alley`, and its helper type is
`alley`. The `Name` field is a different thing, the key a save file uses, and
nobody but the engine reads it.

Exported names are the authoring vocabulary — what a content file spells out:
the entity variables, the state keys in `inc/constants.go`, each package's
`<Category>Manager` and entity type, the colour tokens, the tape lookups
(`tape.Of`, `tape.DialogOf`), the system (`system.System` and its methods
`Do`, `Game`, `Attach`, …), the Armilla (`armilla.Slot2`, `armilla.Insert`,
`armilla.StripIfEmpty`), the action constructor `System.Paused` and the one
HUD condition a content file has to name,
`widget.WhenPlaying`. A tape's line is `speaker.TapeSupport.Say(…)` like
anyone else's — the tape voice is the `Speaker` in `speaker/`, and `Say` is
the engine's one `Say` on every talker, not a second spelling of it.

Everything else stays unexported, and now the compiler holds the line: the HUD
widgets (`hudFrame`, `slots`, …), the condition `whenDebug`, and the
file-local helpers.

The word is **scene**, the engine's own. The wiki and the concept-art deck count
*screens*, and this code used to as well, but everything a screen had that a
scene did not now lives in inkwell. "Screen" survives only where it means the
display: `ScreenWidth`, `ScreenHeight`.

## Layering

The layering is the import graph, so the compiler keeps it:

```
inc ← theme ← system ← armilla ← widget ← main
 ↑      ↑        ↑                  ↑
util    └── achievement, character, speaker, music, sound ← item ← script ← dialog ← tape ← scene ← stage
```

`inc` imports nothing and everything imports it, and `util` sits beside it on
the same footing — standard library only, no game in it. `system` knows nothing
about the HUD or the content; both build on it, never the other way round — and
an arrow pointing back is now an import cycle, not a review comment.

The content row is a strict order because a reference is an import, and it
reads as *what is made of what*: an item is looked at by a character, a script
hands out items, a dialog runs scripts, a card carries an item and plays a
dialog, a scene holds all of them, a stage frames the lot. `tape → widget` and `stage → widget` point the same way:
a cassette and a screen each register the widget that draws them. And the HUD
names no stage back — `widget.WhenPlaying` reads `inkwell.OnWorld()` — which is
the one place the strictness had to be paid for.

`main` is at the very top, and it is the only place that names every package at
once. `theme`, `system`, `armilla` and `widget` are `inc/system/`; the other
eleven are `inc/content/`, and no arrow leaves `system` for `content`.
`achievement` sits at the very bottom of the content row with `character`,
because an achievement names nothing but a flag in `inc`.

## Adding a scene

1. `inc/content/scene/scene.<name>.go`: `var <Name> = SceneManager.Register(&Scene{…})`
   with its `Title` and every `Label` written in both languages, its
   background as the `Layers` element at `Z: -1` whose `Asset` is named after
   the scene, and under it `var _ = <Name>.Leads(…)` with its exits
2. A 640×380 PNG at `assets/bg/<name>.png`

Nothing else moves. There is no constant and no list to update; another scene
reaches this one as `To: <Name>`, and the picture is named nowhere but in the
scene's own literal. Anything Paul walks behind is one more layer in the same
slice, at a positive `Z`, with its own PNG.

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
make steam VERSION=<v> STEAM_SDK=<sdk>   win-x64 and linux-x64 zips with Steam, -tags steam
make steam-debug                         native binary with Steam and the debug menu on
go run . -scene <name>     start on a given scene
go run . -finale           start with the finale
go run . -debug            with the debug menu on d
go run . -explore          the entity explorer, full screen, instead of the game
```

The debug menu is the engine's `Menu` widget with `inkwell.DebugMenuRoot`
handed to it, registered in `widget.debug_menu.go`. Either `-debug` or a
`-tags debug` build sets `inc.FlagDebug`, which is what its `whenDebug`
condition reads; without the flag the widget never ticks. The same flag
switches on the engine's `HotspotDebug` overlay, `widget.hotspot_debug.go`,
which plates every clickable on screen in a colour per kind with its name on
it, `F1` to hide. The menu's last page is the engine's explorer, the `Game`
literal field by field, and `-explore` boots straight into it through
`inkwell.Explore(game)` in place of `inkwell.Run`. `README.md` has the rest,
under **The debug menu**.

Art is embedded into the binary (`//go:embed assets` in `main.go`), because
js/wasm has no OS filesystem and `make binaries` ships the executable alone.

The screen is 640×380 because the engine stretches a background over the whole
window: any other size squashes the paintings.
