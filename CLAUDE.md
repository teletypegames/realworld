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

Do not write tests, and do not add test cases to the ones that already exist.
Correctness is checked by running the game.

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
  `register…` function, its shared types, its list of entities.
- Every other file in a category holds exactly one entity: one screen, one
  background, one item.

```
main.go                    flags + inkwell.Run
inc/boot.manager.go        wiring; New(Opts) builds the game
inc/names.manager.go       entity names and world-state keys
inc/theme.manager.go       realworld-93 + nokia-punk
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
inc/screen.manager.go      the deck, the scene builder
inc/screen.exit.go         the connections between screens
inc/screen.*.go            one file per screen
```

## Naming

One package means one namespace, so an entity's constructor carries its
category as a prefix:

```go
screenAlley()            backgroundAlley()
itemMysteryTape()        characterMysteryTape()
dialogDexTalk()          scriptTapeInsert()
```

Category-level functions follow the same shape: `registerScreen`,
`registerBackground`, `registerItem`, and so on. Only `New` and `Opts` are
exported — they are what `main.go` needs, and nothing else leaves the package.

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

1. `inc/screen.<name>.go` — `screen<Name>() screen`
2. `inc/background.<name>.go` — `background<Name>() inkwell.Asset`
3. One line in `screenDeck` (`inc/screen.manager.go`)
4. One line in `registerBackground` (`inc/background.manager.go`)
5. Constants in `inc/names.manager.go`: `Screen<Name>`, `Bg<Name>`
6. A 640×380 PNG in `assets/bg/`

Nothing else moves.

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
