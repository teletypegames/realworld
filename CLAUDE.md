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

- The category is always **singular**: `item`, `screen`, `background`,
  `character`, `dialog`, `script`, `world`, `ui`, `theme`, `names`, `boot`,
  `content`.
- `[category].manager.go` is the file that ties a category together — its
  manager, its entity type, whatever the category shares.
- Every other file in a category holds exactly one entity: one screen, one
  background, one item. The file registers it itself, in an `init()`.

```
main.go                    flags + inkwell.Run
inc/manager.manager.go     Manager[T]: the generic registry every category uses
inc/manager.interface.go   ManagerInterface and the compile-time assertions
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
inc/background.*.go        one image asset per screen
inc/character.*.go         the cast, tapes included
inc/item.*.go              inventory
inc/dialog.*.go            dialogue trees
inc/script.*.go            named action sequences
inc/screen.manager.go      the Screen type, the deck, the scene builder
inc/screen.exit.go         the connections between screens
inc/screen.*.go            one file per screen
```

`manager` is the one category with no entities of its own: it holds the registry
every other category is built from.

## Managers

Every category that owns a collection of entities has a manager, and they are
all the same generic type — `Manager[T]` in `manager.manager.go`. It keeps one
slice in registration order and one `map[string]int` beside it, so `GetByName`
is a map lookup, not a scan.

Since the entity types are aliases of engine structs, no method can be attached
to them; the manager is told how to read a name instead, which is all it needs:

```go
type Character = inkwell.Character

var CharacterManager = NewManager(func(entity Character) string { return entity.Name })
```

That is the whole of a category's manager file — an alias and one line. There
are seven managers: `BackgroundManager`, `CharacterManager`, `DialogManager`,
`ItemManager`, `ScriptManager`, `ScreenManager`, `ThemeManager`.

`manager.interface.go` holds the contract they all keep, and asserts each one
against it. A new manager goes on that list.

```go
type ManagerInterface[T any] interface {
	Register(entity T)
	GetByName(name string) (T, bool)
	GetAll() []T
}
```

`Register` replaces by name and keeps the entity's position, so registering
twice is an update, never a duplicate. `GetAll` returns the slice itself, in
registration order.

`Screen` is the one entity type that is not an alias: a screen carries exits and
an `OnSelector` flag that inkwell's `Scene` knows nothing about.

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
Nothing may depend on it — including the order the arrow keys walk the screens,
which is simply the order the files sit in.

### Handing a category to the engine

Once the game exists, `registerAll` walks a manager and gives every entity to
the engine's own manager. `registerContent` is the whole of it:

```go
func registerContent() {
	registerAll(BackgroundManager, World.G.AssetManager.Register)
	registerAll(CharacterManager, World.G.CharacterManager.Register)
	registerAll(ItemManager, World.G.ItemManager.Register)
	registerAll(DialogManager, World.G.DialogueManager.Register)
	registerAll(ScriptManager, World.G.ScriptManager.Register)
	registerScreen()
}
```

`registerScreen` is the one that needs its own function, because a screen has to
be turned into an `inkwell.Scene` first, and because the selector's pins are
derived from whichever screens marked themselves `OnSelector`.

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
screenFloor  screenBuild  registerScreen  fillSelectorPins
```

Entities themselves need no name at all — they are anonymous literals inside
their file's `init()`, and the file name says which one it is.

Exported names are the authoring vocabulary — what a content file spells out:
`New` and `Opts`, `World`, the managers and the entity types they hold, the
constants in `names.manager.go`, the colour tokens, and the action constructors
(`TapeSay`, `Paused`, `SetMode`, `EnterScene`, `Back`, `Fn`).

Everything else is machinery and stays unexported: the HUD widgets
(`tapeSlots`, `letterbox`, `hudFrame`, …), the `register…` functions, the
runners, `screenBuild`, `exit`.

The engine's word for a screen is `Scene`. Ours is **screen**, because the
wiki, the concept-art deck and the beat tables all count screens. "Scene"
survives only where the code talks to the engine (`EnterScene`,
`SceneManager`).

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

## Adding a screen

1. Constants in `inc/names.manager.go`: `Screen<Name>`, `Bg<Name>`
2. `inc/screen.<name>.go` — an `init()` registering a `Screen`
3. `inc/background.<name>.go` — an `init()` registering a `Background`
4. A 640×380 PNG in `assets/bg/`

Nothing else moves. There is no list to update.

## Commands

```
make build     native binary into bin/
make wasm      js/wasm build into dist/
make watch     rebuild on change
go run . -screen <name>    start on a given screen
go run . -finale           start with the finale
```

Art is embedded into the binary (`//go:embed assets` in `main.go`), because
js/wasm has no OS filesystem and `make binaries` ships the executable alone.

The screen is 640×380 because the engine stretches a background over the whole
window: any other size squashes the paintings.
