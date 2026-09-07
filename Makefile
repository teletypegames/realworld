# -----------------------------------------
# Makefile – Ebitengine project builder
# -----------------------------------------

PROJECT = realworld

# Our own modules live on the forge, not on proxy.golang.org: fetch them
# directly and skip the public checksum database.
export GOPRIVATE = git.teletypegames.org

BIN_DIR     = bin
DIST_DIR    = dist
WASM_NAME   = game.wasm
OUTPUT_BIN  = $(BIN_DIR)/$(PROJECT)
OUTPUT_WASM = $(DIST_DIR)/$(WASM_NAME)
OUTPUT_JS   = $(DIST_DIR)/wasm_exec.js
OUTPUT_ZIP  = $(PROJECT)-$(VERSION).html.zip

INDEX_HTML_URL = https://git.teletypegames.org/tools/ebitengine-tools/raw/branch/master/web/index.html

all: build

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(OUTPUT_BIN) .

# Same binary with the debug menu switched on from the start — no -debug flag
# needed. See README, "The debug menu".
debug:
	@mkdir -p $(BIN_DIR)
	go build -tags debug -o $(OUTPUT_BIN) .

# Headless: composes the game, validates every cross-manager name reference
# and checks the HUD layout. No window opens.
test:
	go test ./...

wasm:
	@mkdir -p $(DIST_DIR)
	GOOS=js GOARCH=wasm go build -o $(OUTPUT_WASM) .
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(OUTPUT_JS)

export: wasm
	@if [ -z "$(VERSION)" ]; then \
		echo "ERROR: VERSION not set!"; exit 1; \
	fi
	@echo "==> Downloading index.html"
	curl -sSL $(INDEX_HTML_URL) -o $(DIST_DIR)/index.html
	@echo "==> Packaging HTML/WASM for $(VERSION)"
	zip -r $(OUTPUT_ZIP) -j $(DIST_DIR)/$(WASM_NAME) $(DIST_DIR)/wasm_exec.js $(DIST_DIR)/index.html
	@echo "==> Cleaning temporary files"
	rm -f $(DIST_DIR)/$(WASM_NAME) $(DIST_DIR)/wasm_exec.js $(DIST_DIR)/index.html

# --- native binaries (lásd: teletypegames/NOTES_BINARY_PLAN.md) -------------
# win-x86 / win-x64: pure Go cross-compile (Windowson nem kell cgo)
# linux-x64: cgo build, linux/amd64 hoston fut (builder image, X11/GL dev libekkel)
BINARY_TARGETS = win-x86 win-x64 linux-x64

binaries:
	@if [ -z "$(VERSION)" ]; then \
		echo "ERROR: VERSION not set!"; exit 1; \
	fi
	@for target in $(BINARY_TARGETS); do \
		$(MAKE) binary-$$target VERSION=$(VERSION) || exit 1; \
	done

binary-win-x86:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=windows B_GOARCH=386 B_CGO=0 B_EXT=.exe B_TARGET=win-x86

binary-win-x64:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=windows B_GOARCH=amd64 B_CGO=0 B_EXT=.exe B_TARGET=win-x64

binary-linux-x64:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=linux B_GOARCH=amd64 B_CGO=1 B_EXT= B_TARGET=linux-x64

# --- Steam builds -----------------------------------------------------------
# The same packages with -tags steam, which links the engine's steam/ package
# (purego over Valve's redistributable), plus that redistributable beside the
# binary, where Steam launches the game from. 64-bit only: Valve ships the
# library for win64, linux64 and osx alone, so win-x86 and wasm never get the
# tag. STEAM_SDK is the unpacked Steamworks SDK, the directory that holds
# redistributable_bin/. The dev-only steam_appid.txt is never packaged — see
# README, "Steam".
STEAM_TARGETS = win-x64 linux-x64

steam:
	@if [ -z "$(VERSION)" ]; then \
		echo "ERROR: VERSION not set!"; exit 1; \
	fi
	@for target in $(STEAM_TARGETS); do \
		$(MAKE) steam-$$target VERSION=$(VERSION) || exit 1; \
	done

steam-win-x64:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=windows B_GOARCH=amd64 B_CGO=0 B_EXT=.exe B_TARGET=steam-win-x64 B_TAGS=steam B_STEAM_LIB=win64/steam_api64.dll

steam-linux-x64:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=linux B_GOARCH=amd64 B_CGO=1 B_EXT= B_TARGET=steam-linux-x64 B_TAGS=steam B_STEAM_LIB=linux64/libsteam_api.so

# macOS builds natively, on a Mac, for the Mac it runs on.
steam-mac:
	@$(MAKE) binary-build VERSION=$(VERSION) B_GOOS=darwin B_GOARCH=$$(go env GOARCH) B_CGO=1 B_EXT= B_TARGET=steam-mac-$$(go env GOARCH) B_TAGS=steam B_STEAM_LIB=osx/libsteam_api.dylib

# Local Steam-enabled binary for development: the library is found through
# STEAMWORKS_LIB_PATH and the app id through STEAM_APPID, so nothing has to
# be copied next to it. Steam must be running.
steam-debug:
	@mkdir -p $(BIN_DIR)
	go build -tags "steam debug" -o $(OUTPUT_BIN) .

# belső helper: egy target buildje + zip egyetlen gyökérmappával
# (a zipet unixos zip készíti, így a végrehajtási bit megmarad)
# B_TAGS: build tagek; B_STEAM_LIB: a Steam redistributable útja a
# $(STEAM_SDK)/redistributable_bin alatt, a bináris mellé másolva
binary-build:
	@set -e; \
	PKG_DIR="$(PROJECT)-$(VERSION)-$(B_TARGET)"; \
	echo "==> Building $$PKG_DIR"; \
	rm -rf "$$PKG_DIR" "$$PKG_DIR.zip"; \
	mkdir -p "$$PKG_DIR"; \
	CGO_ENABLED=$(B_CGO) GOOS=$(B_GOOS) GOARCH=$(B_GOARCH) go build -tags "$(B_TAGS)" -o "$$PKG_DIR/$(PROJECT)$(B_EXT)" .; \
	if [ -n "$(B_STEAM_LIB)" ]; then \
		if [ -z "$(STEAM_SDK)" ]; then echo "ERROR: STEAM_SDK not set!"; exit 1; fi; \
		cp "$(STEAM_SDK)/redistributable_bin/$(B_STEAM_LIB)" "$$PKG_DIR/"; \
	fi; \
	if [ -f LICENSE.md ]; then cp LICENSE.md "$$PKG_DIR/"; fi; \
	if [ -f README.md ]; then cp README.md "$$PKG_DIR/"; fi; \
	zip -r "$$PKG_DIR.zip" "$$PKG_DIR" >/dev/null; \
	rm -rf "$$PKG_DIR"; \
	echo "==> $$PKG_DIR.zip kesz"

watch:
	fswatch -o . | while read; do make build; done

clean:
	rm -rf $(BIN_DIR) $(DIST_DIR)

.PHONY: all build debug test wasm export watch clean binaries binary-win-x86 binary-win-x64 binary-linux-x64 binary-build steam steam-win-x64 steam-linux-x64 steam-mac steam-debug
