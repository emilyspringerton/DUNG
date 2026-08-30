# DUNG

The BURROW editor — a unified terminal emulator + editor, rewriting `PITVIPER` (SDL2 terminal
emulator) and PARENA's own real `stdlib/editor/*.prn` into one application: a real, ground-up
PARENA-native editor, a Go-hosted terminal (visor-style drop-down, i3-primitive split panes), one
shared window.

New repo (2026-08-30, upstream created ahead of the ask). Bazel-built. The editor domain is
compiled via the real **`burrow`** CLI (`github.com/emilyspringerton/BURROW`'s own Go+PARENA
reimplementation of `parena-c`) — DUNG is BURROW's own real, live, flagship dogfooding consumer.

See `NORTHSTAR.md` for the full real scoping pass: the real UX foundation adopted directly from
`EmilyOS/docs/legacy-archive/gui-v0.1-design-capture.md` (a "tmux × i3 hybrid" layout, a
no-single-click double-click-speed interaction contract, posture-aware non-modal safety feedback),
the real architecture split (ground-up PARENA for the editor domain, Go for SDL2/PTY plumbing).
**Corrected, same day**: this doc originally named `burrow build` as not implemented yet as a real
blocker — `BURROW`'s own Phase 3-4 (region analyzer + emitter parity) shipped the same day, real
and gcc-verified end to end, clearing DUNG's own Phase 0.

## Status

Phase 0 (toolchain) real, cleared. Phase 1 real, shipped (2026-08-30): `cmd/dung` is a real SDL2
window with a visor-style drop-down terminal (PTY + vterm, vendored from `PITVIPER`) and a real
i3-primitive binary split (Ctrl+Shift+Enter/Ctrl+Shift+O, Alt+Arrow focus). `go build`/`go vet`/
`go test` clean, `bazel build //...` clean (real pkg-config/cgo-sandbox fix for the SDL2 dep); ran
for real under Xvfb with a real bash prompt rendering through the full vterm→font→SDL2 pipeline,
screenshot verified, both via plain `go build` and via the real Bazel artifact. First real PARENA
mod also wired in: `internal/burrowgen/entry_gen.go` (checked-in output of `burrow build
parena/entry.prn -o ...`, `BURROW`'s own new Phase 6 Go emission target) supplies `layout()`'s
split-size math and `nextFocus()`'s wraparound math directly — a real Go import, no cgo/FFI
boundary. No editor pane yet — that's Phase 2. See `NORTHSTAR.md` for the full verified proof and
the honest open items (no true global hotkey yet).

## Related

- `BURROW` — the real compiler toolchain this project depends on.
- `PITVIPER`, `PARENA/stdlib/editor/*.prn`, `EmilyOS` — the real things this project rewrites and
  draws its UX foundation from (`PITVIPER` itself is untouched, stays as-is).
- `EMILY` — RSI loop / backlog coordination for cross-repo work.
