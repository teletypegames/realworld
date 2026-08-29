module realworld

go 1.26.5

// Fejlesztés közben a szomszédos checkout-ra mutatunk, hogy a motor és a
// játék együtt mozogjon. Kiadáskor ezt a sort kell törölni.
replace git.teletypegames.org/engines/inkwell => ../../engines/inkwell

require (
	git.teletypegames.org/engines/inkwell v0.0.0-00010101000000-000000000000
	github.com/hajimehoshi/ebiten/v2 v2.9.10
)

require (
	github.com/ebitengine/gomobile v0.0.0-20250923094054-ea854a63cce1 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/oto/v3 v3.4.1 // indirect
	github.com/ebitengine/purego v0.9.0 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/jfreymuth/oggvorbis v1.0.5 // indirect
	github.com/jfreymuth/vorbis v1.0.2 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.44.0 // indirect
)
