# DUNG

## What this is

The BURROW editor — a unified terminal emulator + editor rewriting `PITVIPER` and PARENA's own
real `stdlib/editor/*.prn` into one Go+PARENA application. New repo (2026-08-30). **Read
`NORTHSTAR.md` before writing any code** — it has the full real scoping pass, including two real,
founder-driven corrections already baked into the doc's own "Where this came from" section (first
scoped as a BURROW subdirectory, corrected to its own standalone, Bazel-built repo; corrected
again to be a ground-up PARENA-native editor port that uses the real `burrow` CLI as its own
compiler toolchain, not a Go-primary hand-port).

## Status

Phase 0 (toolchain, gated on `BURROW`'s own Phase 3-4) is real and cleared — shipped same-day,
gcc-verified end to end (`EMILY/BACKLOG.md` S206-88, Apple #17060). Phase 1 (a visor-style
drop-down terminal, Go-hosted, i3-primitive split — doesn't wait on `burrow`) real and shipped
same-day: `go build`/`go vet`/`go test` clean, verified running for real under Xvfb (real bash,
real colored PTY output rendered through vterm→font→SDL2). See `NORTHSTAR.md` for the full proof
and honest open items.

## Related Repos

- `BURROW` — this project's own real compiler toolchain dependency (`burrow build` compiles this
  project's own ground-up PARENA editor source).
- `PITVIPER` — the real terminal emulator this project rewrites (unaffected, stays as-is).
- `PARENA` — `stdlib/editor/*.prn` is the real, existing editor logic this project ports ground-up
  into its own real `.prn` source tree.
- `EmilyOS` — `docs/legacy-archive/gui-v0.1-design-capture.md` is the real, authoritative UX spec
  this project adopts directly; `internal/posture/` is the real posture state machine its own
  denial-feedback UX would integrate with.
- `EMILY` — RSI loop / backlog coordination for cross-repo work.

## Founder Real-Time Direction

Whenever the founder gives real-time direction — a new ask, a correction, a "can we also..." —
route it through `emily observe -s info "Founder real-time: <summary>"` first, even if it isn't
this repo's usual domain, then sprint-plan it into `EMILY/BACKLOG.md` (`emily backlog curate`,
scoped into a real SECTION/sub-item, not just a one-line log), and only then implement. See
`EMILY/docs/THE_EMILY_WAY.md` Principle 18 ("Pave the Cow Paths").

## Apple Filing Protocol

After any meaningful change, file an Apple:
```bash
emily apples post -t completion -repo DUNG "<title>" "<body with commit hash>"
```
Then mark the item done in `EMILY/BACKLOG.md` and commit.

## CHANGELOG Protocol

After any meaningful change, update CHANGELOG.md:
```bash
emily changelog add DUNG "<what changed>"
# or manually: append a dated bullet under ## YYYY-MM-DD in DUNG/CHANGELOG.md
```

## Frame-Break Reframing

Founder-sourced prompting technique (REDGARDEN/NORTHSTAR.md §28, full origin in
REDGARDEN/docs2/MULTI_AGENT_RD_RESEARCH_NOTES.md §5): given a request, name the underlying
structural/systemic pattern it's one instance of — one level of abstraction up — as an added
lens during planning/triage/judgment calls. Use it to spot the general case behind a specific
ask. It augments judgment, it does not replace doing the work: direct, concrete execution of
the literal task asked for still happens every time.

## Commit Protocol (standing instruction)

Always commit and push completed work immediately — don't wait to be asked. This is the default for every repo in this monorepo.

Every commit — human-written or produced by automated code paths (git-commit helpers in emily-agent, emily.cli, IDUNA handlers, etc.) — must carry the active `emily session` fingerprint as a `session: <tag>` trailer (blank line, then the trailer). This was silently missing from several independently-implemented automated commit helpers across the monorepo until an audit on 2026-08-10 (founder, real-time: "where in the fuck is my llm session id anywhere"). If you add a new automated git-commit code path anywhere, wire in the session tag the same way — don't assume an existing helper already does it.
