# NORTHSTAR — DUNG (the BURROW editor)

## Where this came from, corrected twice in real time

Founder real-time, 2026-08-30, immediately after BURROW's own Phase 1/2 (lexer + parser parity)
shipped:

> "ok and rewrite pitviper and the parena editor into burrow" → "call it DUNG" → "aka the burrow
> editor" → "and terminal emulator" → "terminal emulator comes down as a visor" → "and there is
> also split pane use i3 primatives" → "like the editor and the terminal in the same window and
> then we like slurp in another file via drag another file onto it or load it from internet it
> opens up a new chat window inside the shared window" → "EMILY os design primatives ux and
> affordances"

**This project's own first scoping pass (written as `BURROW/DUNG.md`, in the BURROW repo) was
corrected directly, twice, in immediate follow-up**:

> "release DUNG as its own repo bazel built parena from the ground up port of parena editor" →
> "that uses burrow" → "use the real burrow cli to build the product"

**Real, corrected scope**: DUNG is its own standalone, real repo (this one — the founder had
already created the real, empty upstream `github.com/emilyspringerton/DUNG.git` before the
correction, same real "upstream created ahead of the ask" pattern `BURROW`'s own repo followed).
Built with **Bazel** (matching PARENA's own real build-system convention, not BURROW's own plain
`go.mod`). The editor half is a **ground-up port of PARENA's own real `stdlib/editor/*.prn`**,
written as real, fresh PARENA source living in this repo — not primarily hand-written Go with
PARENA as a secondary aid, the way BURROW's own `lexer.go`/`parser.go` hand-ports were. **DUNG
uses BURROW**: its own real build process compiles that PARENA source via the real `burrow` CLI
(BURROW's own Go+PARENA reimplementation of `parena-c`), not the original C-based `parena-c`
directly — DUNG is BURROW's own real, live, flagship dogfooding consumer.

**Scoping only, this pass.** No DUNG code exists yet.

## Real, explicit non-goal: PITVIPER itself is untouched

Same real precedent named in this project's own first draft, unchanged by the repo move:
**`PITVIPER`'s own Go implementation stays exactly as-is** — DUNG is a new, separate rewrite, not
a rewrite-in-place of the existing repo. The same real pattern `SAND` already set relative to
`PITVIPER` (`/home/fatbaby/CLAUDE.md`'s own SAND row), applied here a second time.

## What DUNG actually is

A single, real, unified application combining:

1. **A real terminal emulator** — the real successor to `PITVIPER`'s own SDL2 terminal (PTY
   handling, glyph rendering, vterm/scrollback). One real, new UX requirement named directly by
   the founder: **it comes down as a visor** — a real, quake-style drop-down overlay (the same
   real UX `Guake`/`Yakuake`/`Tilda` already establish), not a plain, permanently-windowed
   terminal.
2. **A real editor** — the real, ground-up PARENA-native successor to PARENA's own real
   `stdlib/editor/*.prn` (buffer management, TextMate-based syntax highlighting + theming, a
   Vim-modal keybinding model, a Spotlight/`Ctrl+T` command palette with a real, hot-swappable
   plugin-backend API, a reusable widget system) — real, substantial, already-designed logic
   this project ports forward, not started from a blank page.
3. **A real, shared, i3-primitive split-pane layout system** — the founder's own explicit
   correction over `PITVIPER/docs/NORTHSTAR.md`'s own existing, real, but different Milestone 3
   plan (vim-style `Ctrl+W |`/`Ctrl+W -` splits). DUNG's own real pane model follows i3's own real
   tiling primitives instead: a tree of containers, each one real, binary horizontal-or-vertical
   split, panes and splits nested arbitrarily, real focus movement between panes in the tree. One
   shared layout system serves BOTH the terminal panes and the editor panes.

## Why this is architecturally coherent, not a forced merger

Checked before writing this doc, not assumed: `PITVIPER` and PARENA's own editor **already share
the identical real technical foundation** — `PARENA/stdlib/editor/ui.prn`'s own header comment
states it directly: "gutter/diagnostics/status-bar/popups... render via sdl2's own real
primitives..., same architecture as PITVIPER's own SDL2 terminal-emulator precedent, not a new
toolkit." Both are already SDL2-based; `PARENA/stdlib/editor/events.prn`'s own header comment
cites PITVIPER's own real "ConPTY + mouse-drag-selection + clipboard" work as its own
architecture's precedent directly. DUNG isn't inventing a merger between two unrelated things —
it's the real, natural next step of two things that were already converging.

## Real UX foundation — adopted directly, not invented (`EmilyOS/docs/legacy-archive/gui-v0.1-design-capture.md`)

Founder real-time: "EMILY os design primatives ux and affordances." Found a direct, real match
before writing anything new: `EmilyOS/docs/legacy-archive/gui-v0.1-design-capture.md` is a
complete, already-written GUI v0.1 spec whose own layout model is explicitly named **"tmux × i3
hybrid"** — predating and directly matching the founder's own "use i3 primitives" instruction.
Adopted as DUNG's own real, authoritative UX foundation directly, not re-derived. The real,
load-bearing pieces (full detail in that doc):

- **Visual system**: a real, hard-coded palette — deep navy/near-black background, EGSHELL
  (white-ish) reserved exclusively for directory/container tiles, colored (never white) button
  tiles darker than EGSHELL, all-caps blocky typography, single-shot (never looping) activation
  flashes only.
- **Interaction contract — no single-click actions, ever**: a real, deterministic **fast
  double-click (≤220ms) activates**, a real, deterministic **slow double-click (350-800ms) enters
  label-edit mode**. Real, exact timing thresholds already specified, not left to be invented.
- **Layout — the real "tmux × i3 hybrid" panes**: the screen is a real tiling space; panes hold a
  directory view, a file view, a process/log panel, or a command-line/verb-bar panel; real
  keyboard-first pane operations. No freeform dragging by default.
- **Safety/posture hooks, real and already EmilyOS-integrated**: postures
  (`NORMAL`/`SIEGE`/`MERCY`/`INCIDENT`/`GAME`, `EmilyOS/internal/posture/`) gate which actions are
  AVAILABLE, never the visuals; a denied action gets a real, quiet, one-frame "deny flash" —
  explicitly no modal dialogs, no toast notifications. Real, direct relevance: EmilyOS is already
  named as PITVIPER's own real, related repo ("PITVIPER is the operator interface" per EmilyOS's
  own CLAUDE.md) — DUNG inheriting this same posture-aware interaction contract is a real,
  natural continuation, not a new one being forced.

**Real, new addition, reconciled with the original spec's own real "no freeform dragging by
default" rule**: "like the editor and the terminal in the same window and then we like slurp in
another file via drag another file onto it or load it from internet it opens up a new chat window
inside the shared window" — confirms DUNG is one real, shared window, and names a real, new pane
TYPE the original EmilyOS spec doesn't yet have: a **chat pane** (the real Emily Prime pane
`PITVIPER`'s own `docs/NORTHSTAR.md` already plans as its own Milestone 4), opened by dragging a
file onto the DUNG window OR loading one from a URL — the dropped/loaded file becomes that new
chat pane's own real, scoped context. A real, explicit verb-shaped action, not silently violating
the original spec's own no-casual-drag default.

## Real architecture: Bazel, ground-up PARENA editor, Go for what PARENA genuinely can't do yet

Real, corrected split (see "Where this came from" above):

- **PARENA, ground-up, primary** for the editor domain: a real, fresh port of `PARENA/stdlib/
  editor/*.prn`'s own real logic (buffer manipulation, TextMate tokenization + theming rules, the
  Spotlight fuzzy-match/plugin-dispatch algorithm, the widget system's own state transitions,
  i3-primitive split-pane tree logic) — written as real `.prn` source living in this repo, not a
  hand-port fallback.
- **The real `burrow` CLI is this project's own compiler toolchain** — `DUNG`'s own Bazel build
  invokes the real `burrow` binary (`BURROW`, this same session's own Go+PARENA reimplementation
  of `parena-c`) to compile that PARENA source, not the original C-based `parena-c`. **Real,
  honest, current blocker, named directly, not glossed over**: `burrow build` today reports "not
  yet implemented" (`BURROW/NORTHSTAR.md`'s own Phase 3-4 — region analyzer + emitter parity —
  hasn't started yet). DUNG's own real build cannot succeed end to end until that work lands.
  This makes BURROW's own Phase 3/4 a real, concrete, now-higher-priority prerequisite for DUNG,
  not a parallel, independent track.
- **Go, real and direct, for what SDL2/PTY genuinely need** (matching `PITVIPER`'s own
  already-proven real stack — SDL2 window/event-loop/renderer via cgo, FreeType2 glyph rendering,
  PTY via `openpty(3)`, the vterm/scrollback state machine): real, substantial systems-level
  plumbing no PARENA target has ever attempted, and not the ground-up-PARENA-port's own real
  scope (that's the editor domain specifically). This Go layer is the real host DUNG's own
  PARENA-compiled editor logic runs inside — the same real "PARENA owns the decision logic, host
  code owns the plumbing" split every mod in this monorepo already uses.

## Real scale, honestly

- `PITVIPER`: **4066+ real lines of already-working Go** — `cmd/pitviper/main.go` (689 lines),
  `internal/vterm` (842 lines + 555 lines of tests), `internal/pty`, `internal/font` (218 lines +
  emoji/shiny variants), `internal/gfdapi` (235 lines), `internal/mudconn`, `internal/scrollmod`.
- PARENA's editor: **13 real `.prn` files** (`buffer`/`render`/`spotlight`/`textmate`/
  `textmate_markdown`/`textmate_parena`/`textmate_loader`/`theme`/`ui`/`widget`/`plugin`/`events`/
  `construct_split`) plus a real, already-built `editor-demo` binary in the PARENA repo.
- **A full rewrite of both, in one sitting, is not realistic** — the same honest scale assessment
  `BURROW/NORTHSTAR.md` already makes for the compiler rewrite itself. This doc's own real job is
  scoping the real shape and the real first slice, not claiming a finished design.

## Real, phased plan

**DUNG Phase 0 — unblock the real toolchain**: this project's own real Bazel build is gated on
`BURROW`'s own Phase 3-4 (region analyzer + emitter parity) landing enough real emit capability
for `burrow build` to produce something runnable from real `.prn` source. Not this project's own
work to do — a real, named dependency on BURROW's own separate track.

**DUNG Phase 1 — the real, smallest proof point**: a real SDL2 window (Go host) that (a) renders
a visor-style drop-down terminal pane (PTY + vterm, hand-ported or vendored from `PITVIPER`'s own
real, already-working `internal/vterm`/`internal/pty` — the terminal domain doesn't wait on
BURROW's own emitter work, since it's Go-hosted, not PARENA-compiled), and (b) can split that one
pane into two via one real i3-primitive — matching the same "smallest real proof point"
discipline every other real Phase 0/1 in this monorepo already follows. No editor pane yet.

**DUNG Phase 2 — the real first ground-up PARENA editor slice**, gated on Phase 0: port
`stdlib/editor/buffer.prn`'s own real logic first (the most foundational real editor domain,
everything else depends on it), compiled via `burrow build`, real end-to-end proof that the
BURROW→DUNG toolchain relationship actually works.

**DUNG Phase 3+ (design only, not detailed here)**: TextMate highlighting, Spotlight, the plugin
API, the chat pane, Emily Prime integration (matching `PITVIPER`'s own already-planned
Milestone 4).

## Real acceptance test file, named ahead of the code that will need it

Founder real-time: "make sure we are testing with the backlog when we dev DUNG its big." Once any
real file-viewing/editor rendering exists (Phase 2+), `EMILY/BACKLOG.md` is the real, standing
stress-test file — genuinely huge (25000+ lines as of 2026-08-30, mixed-language, real, organic
growth, not a synthetic benchmark), and the same real file the founder's own "whatever the browser
is doing to make this file viewable is like cliutch it works so good" comment (and the raw
GitHub link that followed it) already names as the real bar to clear. Named here now so it isn't
forgotten once real rendering code exists to test it against.

## Real risks and open questions, named honestly

- **Real, hard sequencing dependency on BURROW's own unstarted Phase 3-4** — DUNG's own real
  editor-domain build is blocked until that lands; not this project's own work to accelerate
  unilaterally.
- **Bazel rules mix**: this repo needs real C/Go/PARENA build integration (rules_cc-equivalent for
  any Go+cgo/SDL2 pieces, plus genrule-style steps invoking the real `burrow` binary against
  `.prn` source) — not yet designed at the `BUILD.bazel` level.
- **cgo linking against real PARENA-emitted output** (once `burrow build` can emit something,
  whatever target it lands on first — C is the most likely first real target, matching `parena-c`'s
  own default) needs the same real `parena_runtime.h`/`.c` bridging every other real PARENA-mod
  host in this monorepo already uses — real, proven pattern, not yet applied to a GUI/SDL2 host.
- **i3's own real tiling model is more general than DUNG strictly needs** (i3 manages whole
  windows/workspaces across an entire desktop session) — DUNG only needs the real CONTAINER-SPLIT
  primitive, not i3's own window-manager-level concerns.
- **No real acceptance bar named yet for DUNG** (unlike `BURROW`, which has the founder's own
  explicit "pass all that parena c tests" bar) — a real, open question for the next scoping pass.
- **The chat-pane's own real backend contract isn't specified yet** — presumably the same real
  `EMILY_AGENT_URL`/Emily Prime agent connection `PITVIPER/docs/NORTHSTAR.md`'s own Milestone 4
  already plans, but not confirmed for DUNG specifically; "load [a file] from internet" as a
  chat-context source has no real fetch/security model specified yet either.

## Related

- `BURROW` — this project's own real compiler toolchain dependency (`burrow build` compiles
  DUNG's own ground-up PARENA editor source); `BURROW/NORTHSTAR.md`'s own Phase 3-4 is a real,
  named prerequisite this project is gated on.
- `PITVIPER/CLAUDE.md`, `PITVIPER/docs/NORTHSTAR.md` — the real, existing terminal emulator DUNG
  rewrites (unaffected, stays as-is); its own Milestone 4 (Emily Prime pane) is the real precedent
  for DUNG's own new chat-pane feature.
- `PARENA/stdlib/editor/*.prn` — the real, existing editor logic this project ports, ground-up,
  into its own real `.prn` source tree; `PARENA/editor-demo` is the real, already-built proof this
  logic already compiles and runs (via `parena-c`, not yet via `burrow`).
- `EmilyOS/docs/legacy-archive/gui-v0.1-design-capture.md` — the real, already-written,
  authoritative UX/interaction spec this project adopts directly.
- `EmilyOS/internal/posture/` — the real posture state machine DUNG's own denial-feedback UX
  would need to read from for a real, live integration.
- `SAND` — the real precedent for "a new, separate fork of an existing mission, not a
  rewrite-in-place," applied here a second time (DUNG relative to `PITVIPER`).
