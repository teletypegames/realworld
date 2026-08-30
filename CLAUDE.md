# Real World

A point & click adventure built on the inkwell engine (Ebitengine).

## Conventions

### No comments

Go source files carry no comments. Not doc comments, not inline comments, not
section headers. If a piece of code needs explaining, rename it or restructure
it until it does not, and put the reasoning in `README.md` instead.

The single exception is compiler directives (`//go:embed`, `//go:build`), which
are instructions to the toolchain rather than prose.

### English only

Everything written in the repository is in English: identifiers, string
literals, commit messages, `README.md`, this file. The design wiki is in
Hungarian; when a concept crosses over, translate it and keep the mapping
one-to-one.

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

## Layout

All game code lives in one flat package, `inc`. There are no subdirectories: a
file's name carries the structure, in the form `[category].[name].go`.

- The category is always **singular**: `item`, `scene`, `background`,
  `character`, `dialog`, `script`, `world`, `ui`, `theme`, `names`, `boot`,
  `content`.
- `[category].manager.go` is the file that ties a category together — its
  manager, its entity type, whatever the category shares.
- Every other file in a category holds exactly one entity: one scene, one
  background, one item. The file registers it itself, in an `init()`.

```
main.go                    flags + inkwell.Run
inc/boot.manager.go        wiring; New(Opts) builds the game
inc/names.manager.go       entity names and world-state keys
inc/theme.manager.go       the theme manager and the colour helpers
inc/theme.realworld.go     realworld-93
inc/theme.nokia_punk.go    nokia-punk
inc/world.manager.go       unsaved runtime state
inc/world.action.go        custom actions and the action pump
inc/ui.manager.go          HUD layout and widget registration
inc/ui.*.go                custom widgets, coloured text
inc/content.manager.go     composition root; calls every register…
inc/background.*.go        one image asset per scene
inc/character.*.go         the cast, tapes included
inc/item.*.go              inventory
inc/dialog.*.go            dialogue trees
inc/script.*.go            named action sequences
inc/scene.manager.go       the Scene alias, the defaults every scene gets
inc/scene.selector.go      the map screen: its pins are derived from the graph
inc/scene.*.go             one file per scene
```

## Managers

Every category that owns a collection of entities has a manager, and they are
all the engine's own `inkwell.Manager[T]` — the same registry type the `*Game`
hangs its content off. The game defines no registry of its own.

The entity types are aliases of engine structs, so they already carry
`GetName()` and satisfy `inkwell.Named`; the manager needs nothing else:

```go
type Character = inkwell.Character

var CharacterManager = inkwell.NewManager[Character]()
```

That is the whole of a category's manager file — an alias and one line. There
are seven managers: `BackgroundManager`, `CharacterManager`, `DialogManager`,
`ItemManager`, `ScriptManager`, `SceneManager`, `ThemeManager`.

The methods the game uses are `Register`, `Set`, `Get`, `All` and `Each`.
`Register` panics on a duplicate name — a second registration is a
construction-time bug, not an update — so rewriting an entity that is already
in the registry goes through `Set`, which keeps its position in the order.
`All` and `Each` both hand back the entities in registration order.

Every entity type is an alias, `Scene` included. It was once a struct of our
own, because inkwell's `Scene` could not carry exits; that gap was closed in the
engine, so there is nothing left for a second type to hold.

### Entities register themselves

An entity file is a literal and nothing else. No constructor function, no list
somewhere else to keep in step — the file hands itself to its manager in an
`init()`:

```go
package inc

func init() {
	BackgroundManager.Register(Background{
		Name: BgServerFarm,
		Path: "assets/bg/server_farm.png",
		Kind: inkwell.AssetImage,
	})
}
```

Adding an entity is adding a file. Deleting one is deleting a file. Package-level
variables are initialised before any `init()` runs, so the managers exist by the
time the first file registers into one.

`init()` order is file-name order, so **registration order is alphabetical**.
Nothing may depend on it — including the order the arrow keys walk the scenes,
which is simply the order the files sit in.

### Handing a category to the engine

Once the game exists, `Each` walks a manager and gives every entity to the
engine's own manager. `registerContent` is the whole of it:

```go
func registerContent() {
	BackgroundManager.Each(World.G.AssetManager.Register)
	CharacterManager.Each(World.G.CharacterManager.Register)
	ItemManager.Each(World.G.ItemManager.Register)
	DialogManager.Each(World.G.DialogueManager.Register)
	ScriptManager.Each(World.G.ScriptManager.Register)
	registerScene()
}
```

`registerScene` is the one that needs its own function: it fills in the defaults
a scene may leave out (Paul's starting position, the floor walkbox), and it
derives the selector's pins by reading the exit graph backwards.

## The world

There is exactly one game, so there is exactly one world: `World`, a
package-level singleton in `world.manager.go`. Nothing takes a `*world`
parameter and no widget holds a back-reference — `World.Do(…)`,
`World.HUDVisible()`, `World.Slot2()` are reachable from anywhere in the
package. `New` calls `World.attach(g)`, which binds the engine and resets the
runtime state, so building the game twice is clean.

This is what lets an entity file be a literal: the tape-insert script closes
over `World`, not over a parameter it would have had to be handed.

## Naming

One package means one namespace, so an entity's constructor carries its
category as a prefix:

```go
sceneFloor  sceneDefaults  registerScene  fillSelectorPins
```

Entities themselves need no name at all — they are anonymous literals inside
their file's `init()`, and the file name says which one it is.

Exported names are the authoring vocabulary — what a content file spells out:
`New` and `Opts`, `World`, the managers and the entity types they hold, the
constants in `names.manager.go`, the colour tokens, and the action constructors
(`TapeSay`, `Paused`, `SetMode`, `EnterScene`, `Back`, `Fn`).

Everything else is machinery and stays unexported: the HUD widgets
(`tapeSlots`, `letterbox`, `hudFrame`, …), the `register…` functions, the
runners, `sceneDefaults`, `fillSelectorPins`.

The word is **scene**, the engine's own. The wiki and the concept-art deck count
*screens*, and this code used to as well, but everything a screen had that a
scene did not now lives in inkwell. "Screen" survives only where it means the
display: `ScreenW`, `ScreenH`.

## Layering

Nothing is enforced by the compiler any more, so the layering is a rule kept by
hand:

```
names ← theme ← world ← ui
                  ↑      ↑
               content ──┴── boot ← main
```

`world` knows nothing about the HUD or the content. Both build on it, never the
other way round.

## Adding a scene

1. Constants in `inc/names.manager.go`: `Scene<Name>`, `Bg<Name>`
2. `inc/scene.<name>.go` — an `init()` registering a `Scene`
3. `inc/background.<name>.go` — an `init()` registering a `Background`
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
