// cmd/dung/main.go — DUNG Phase 1: the real, smallest proof point.
//
// Per DUNG/NORTHSTAR.md's own phased plan: "a real SDL2 window (Go host)
// that (a) renders a visor-style drop-down terminal pane (PTY + vterm,
// hand-ported or vendored from PITVIPER's own real, already-working
// internal/vterm/internal/pty -- the terminal domain doesn't wait on
// BURROW's own emitter work, since it's Go-hosted, not PARENA-compiled),
// and (b) can split that one pane into two via one real i3-primitive... No
// editor pane yet." That is exactly this file's scope -- no PARENA/burrow
// integration, no TextMate highlighting, no Spotlight, no chat pane. Those
// are real, named Phase 2/3+ work, not started here.
//
// Visor behavior: the window starts hidden, docked to the top of the
// screen. F12 slides it down into view / back up out of view -- the same
// real UX Guake/Yakuake/Tilda already establish, named directly in
// NORTHSTAR.md. A true global hotkey (working even when the window doesn't
// have focus) needs X11 XGrabKey outside what SDL2's own event loop gives
// for free -- named as a real, honest simplification here, not solved: F12
// only fires while this window already has focus. Real, later work if this
// needs to behave like an actual visor daemon.
//
// i3-primitive split: Ctrl+Shift+Enter splits the focused pane vertically
// (side-by-side, new pane to the right); Ctrl+Shift+O splits it
// horizontally (stacked, new pane below) -- a binary tree of containers,
// each one horizontal-or-vertical, matching NORTHSTAR.md's own "DUNG only
// needs the real CONTAINER-SPLIT primitive, not i3's own window-manager-
// level concerns." Alt+Arrow moves focus between panes.
package main

//go:generate /home/fatbaby/BURROW/burrow build ../../parena/entry.prn -o ../../internal/burrowgen/entry_gen.go
// k8s_scaling_gen.go is DUNG's own real dogfooding copy of PARENA/stdlib/k8s/scaling.prn --
// proof that burrow's new defstruct/get-field support (added the same day this needed it) works
// end to end for a real, scalar-only Kubernetes decision-logic mod. Not wired into any live DUNG
// feature yet (DUNG has no k8s-facing UI today) -- available for a real future use (e.g. a
// "manage your cluster" pane) rather than forced into one now. Regenerate:
//go:generate /home/fatbaby/BURROW/burrow build /home/fatbaby/PARENA/stdlib/k8s/scaling.prn -o ../../internal/burrowgen/k8s_scaling_gen.go

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"dung/internal/burrowgen"
	"dung/internal/font"
	"dung/internal/pty"
	"dung/internal/vterm"
)

func init() {
	// SDL2 requires all calls happen on the same OS thread that called Init.
	runtime.LockOSThread()
}

// visorHeight is the drop-down window's height in pixels once shown. Width
// tracks the full screen width (a real Guake/Yakuake/Tilda convention).
const visorHeight = 480

// visorAnimStep is how far the window moves per animation frame, in pixels.
const visorAnimStep = 40

// pane is one leaf of the split tree: a live terminal (PTY + vterm) and the
// screen-space rect it currently renders into.
type pane struct {
	term *vterm.Screen
	pty  *pty.PTY
	rect sdl.Rect
}

// splitAxis names which of the two real i3-primitive split directions a
// container node holds its children along.
type splitAxis int

const (
	splitNone       splitAxis = iota
	splitVertical             // side-by-side: children split the rect left/right
	splitHorizontal           // stacked: children split the rect top/bottom
)

// node is one element of the binary split tree -- either a leaf (pane != nil)
// or a container with exactly two children along one axis, matching
// NORTHSTAR.md's own "a tree of containers, each one real, binary
// horizontal-or-vertical split."
type node struct {
	pane     *pane
	axis     splitAxis
	children [2]*node
}

func newLeaf(shell string, cols, rows int, rect sdl.Rect) (*node, error) {
	p, err := pty.Open(shell, cols, rows)
	if err != nil {
		return nil, err
	}
	t := vterm.New(cols, rows)
	pn := &pane{term: t, pty: p, rect: rect}
	go pumpPTY(p, t)
	return &node{pane: pn}, nil
}

// pumpPTY reads PTY output and feeds it to the vterm screen until the PTY
// closes. One goroutine per pane -- vterm.Screen is documented safe for
// concurrent use (its own doc comment: "Safe for concurrent use").
func pumpPTY(p *pty.PTY, t *vterm.Screen) {
	buf := make([]byte, 4096)
	for {
		n, err := p.Master.Read(buf)
		if n > 0 {
			t.Write(buf[:n])
		}
		if err != nil {
			return
		}
	}
}

// leaves walks the tree in split order, collecting every leaf pane -- used
// both for rendering and for Alt+Arrow focus-cycling.
func leaves(n *node) []*node {
	if n == nil {
		return nil
	}
	if n.pane != nil {
		return []*node{n}
	}
	var out []*node
	out = append(out, leaves(n.children[0])...)
	out = append(out, leaves(n.children[1])...)
	return out
}

// layout recomputes every leaf pane's rect from the tree structure and the
// space available to it, then resizes each pane's vterm/PTY to match --
// the real mechanism behind "one shared layout system serves BOTH the
// terminal panes and the editor panes" (NORTHSTAR.md), scoped here to
// terminal panes only (no editor pane yet, per Phase 1).
func layout(n *node, rect sdl.Rect) {
	if n == nil {
		return
	}
	if n.pane != nil {
		n.pane.rect = rect
		cols := int(rect.W) / font.GlyphW
		rows := int(rect.H) / font.GlyphH
		if cols < 1 {
			cols = 1
		}
		if rows < 1 {
			rows = 1
		}
		n.pane.term.Resize(cols, rows)
		_ = n.pane.pty.Resize(cols, rows)
		return
	}
	const borderPx = 2
	switch n.axis {
	case splitVertical:
		// Real decision logic lives in PARENA now, not here -- burrowgen.SplitSize is
		// parena/entry.prn's own split-size, compiled to Go via `burrow build -o *.go`
		// (Phase 6, real, shipped: DUNG is the real host that finally asked for it). This
		// hand-written Go call is the "host owns the plumbing" half of NORTHSTAR.md's own
		// "PARENA owns the decision logic, host owns the plumbing" split -- the mod itself
		// decides the number, this file just wires the SDL2 rect around it.
		leftW := burrowgen.SplitSize(rect.W, borderPx)
		layout(n.children[0], sdl.Rect{X: rect.X, Y: rect.Y, W: leftW, H: rect.H})
		layout(n.children[1], sdl.Rect{X: rect.X + leftW + borderPx, Y: rect.Y, W: rect.W - leftW - borderPx, H: rect.H})
	case splitHorizontal:
		topH := burrowgen.SplitSize(rect.H, borderPx)
		layout(n.children[0], sdl.Rect{X: rect.X, Y: rect.Y, W: rect.W, H: topH})
		layout(n.children[1], sdl.Rect{X: rect.X, Y: rect.Y + topH + borderPx, W: rect.W, H: rect.H - topH - borderPx})
	}
}

// splitFocused replaces the given leaf's slot in the tree with a new
// container holding the old leaf plus a freshly spawned shell, along axis.
// Returns the new node so the caller can retarget focus.
func splitFocused(root **node, target *node, axis splitAxis, shell string) (*node, error) {
	newLeafNode, err := newLeaf(shell, 80, 24, target.pane.rect)
	if err != nil {
		return nil, err
	}
	container := &node{axis: axis, children: [2]*node{{pane: target.pane}, newLeafNode}}
	replaceNode(root, target, container)
	return newLeafNode, nil
}

// replaceNode finds old within the tree rooted at *root (by pointer
// identity) and swaps it for replacement. old is always a leaf here (a
// pane being split), so this only ever needs to check container children.
func replaceNode(root **node, old, replacement *node) {
	if *root == old {
		*root = replacement
		return
	}
	var walk func(*node)
	walk = func(n *node) {
		if n == nil || n.pane != nil {
			return
		}
		for i, c := range n.children {
			if c == old {
				n.children[i] = replacement
				return
			}
			walk(c)
		}
	}
	walk(*root)
}

func main() {
	shell := flag.String("shell", "", "shell to launch in each pane (default: $SHELL)")
	sessionName := flag.String("title", "DUNG", "window title")
	showOnStart := flag.Bool("show", false, "show the visor immediately instead of starting docked above the screen (useful for headless/scripted verification, e.g. under Xvfb with no way to send the F12 toggle)")
	flag.Parse()

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
		fmt.Fprintf(os.Stderr, "sdl.Init: %v\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	dm, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetCurrentDisplayMode: %v\n", err)
		os.Exit(1)
	}
	winW := dm.W

	// Visor starts docked above the top of the screen -- hidden -- per
	// NORTHSTAR.md's own "comes down as a visor" requirement. --show skips
	// that for headless/scripted verification.
	startY := int32(-visorHeight)
	winFlags := sdl.WINDOW_BORDERLESS | sdl.WINDOW_HIDDEN
	if *showOnStart {
		startY = 0
		winFlags = sdl.WINDOW_BORDERLESS | sdl.WINDOW_SHOWN
	}
	win, err := sdl.CreateWindow(
		*sessionName,
		0, startY,
		winW, visorHeight,
		uint32(winFlags),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateWindow: %v\n", err)
		os.Exit(1)
	}
	defer win.Destroy()

	ren, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateRenderer: %v\n", err)
		os.Exit(1)
	}
	defer ren.Destroy()

	cols := int(winW) / font.GlyphW
	rows := visorHeight / font.GlyphH
	root, err := newLeaf(*shell, cols, rows, sdl.Rect{X: 0, Y: 0, W: winW, H: visorHeight})
	if err != nil {
		fmt.Fprintf(os.Stderr, "spawn shell: %v\n", err)
		os.Exit(1)
	}
	focused := root

	visible := *showOnStart
	animating := false
	var animTargetY int32

	running := true
	for running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch e := ev.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_RESIZED {
					w, h := win.GetSize()
					layout(root, sdl.Rect{X: 0, Y: 0, W: w, H: h})
				}
			case *sdl.KeyboardEvent:
				if e.Type != sdl.KEYDOWN {
					continue
				}
				mod := e.Keysym.Mod
				switch {
				case e.Keysym.Sym == sdl.K_F12:
					visible = !visible
					animating = true
					if visible {
						win.Show()
						animTargetY = 0
					} else {
						animTargetY = -visorHeight
					}
				case (mod&sdl.KMOD_CTRL) != 0 && (mod&sdl.KMOD_SHIFT) != 0 && e.Keysym.Sym == sdl.K_RETURN:
					if nn, err := splitFocused(&root, focused, splitVertical, *shell); err == nil {
						focused = nn
						w, h := win.GetSize()
						layout(root, sdl.Rect{X: 0, Y: 0, W: w, H: h})
					}
				case (mod&sdl.KMOD_CTRL) != 0 && (mod&sdl.KMOD_SHIFT) != 0 && e.Keysym.Sym == sdl.K_o:
					if nn, err := splitFocused(&root, focused, splitHorizontal, *shell); err == nil {
						focused = nn
						w, h := win.GetSize()
						layout(root, sdl.Rect{X: 0, Y: 0, W: w, H: h})
					}
				case (mod&sdl.KMOD_ALT) != 0 && (e.Keysym.Sym == sdl.K_LEFT || e.Keysym.Sym == sdl.K_RIGHT ||
					e.Keysym.Sym == sdl.K_UP || e.Keysym.Sym == sdl.K_DOWN):
					focused = nextFocus(root, focused, e.Keysym.Sym)
				default:
					if b := keyToPTYBytes(e.Keysym); b != nil {
						_, _ = focused.pane.pty.Master.Write(b)
					}
				}
			case *sdl.TextInputEvent:
				text := e.GetText()
				if text != "" {
					_, _ = focused.pane.pty.Master.Write([]byte(text))
				}
			}
		}

		if animating {
			_, curY := win.GetPosition()
			if curY < animTargetY {
				curY += visorAnimStep
				if curY > animTargetY {
					curY = animTargetY
				}
			} else if curY > animTargetY {
				curY -= visorAnimStep
				if curY < animTargetY {
					curY = animTargetY
				}
			}
			win.SetPosition(0, curY)
			if curY == animTargetY {
				animating = false
				if !visible {
					win.Hide()
				}
			}
		}

		renderFrame(ren, root)
		time.Sleep(16 * time.Millisecond)
	}

	for _, l := range leaves(root) {
		l.pane.pty.Close()
	}
}

// nextFocus picks the next leaf pane in the given directional order.
// v0 scope cut, named honestly in NORTHSTAR.md's own risk list: this is
// visit-order cycling, not true geometric nearest-neighbor -- Left/Up cycle
// backward and Right/Down cycle forward through leaves() order. Real,
// later upgrade if the split tree gets deep enough for that to matter.
func nextFocus(root, cur *node, sym sdl.Keycode) *node {
	all := leaves(root)
	if len(all) < 2 {
		return cur
	}
	idx := 0
	for i, n := range all {
		if n == cur {
			idx = i
			break
		}
	}
	// Real decision logic lives in PARENA now, not here -- burrowgen.NextFocusIndex is
	// parena/entry.prn's own next-focus-index, compiled to Go via `burrow build -o *.go`.
	// This hand-written Go code is just the SDL2 key-symbol dispatch; the actual
	// wraparound arithmetic is the mod's, same "PARENA decides, host wires" split
	// layout() above already follows.
	switch sym {
	case sdl.K_RIGHT, sdl.K_DOWN:
		idx = int(burrowgen.NextFocusIndex(int32(idx), int32(len(all)), 1))
	case sdl.K_LEFT, sdl.K_UP:
		idx = int(burrowgen.NextFocusIndex(int32(idx), int32(len(all)), -1))
	}
	return all[idx]
}

// renderFrame draws every leaf pane's vterm grid into its own rect, plus a
// 1px border between panes -- the visible expression of the split tree.
func renderFrame(ren *sdl.Renderer, root *node) {
	_ = ren.SetDrawColor(0, 0, 0, 0xff)
	_ = ren.Clear()

	for _, l := range leaves(root) {
		renderPane(ren, l.pane)
	}

	_ = ren.SetDrawColor(0x44, 0x44, 0x44, 0xff)
	for _, l := range leaves(root) {
		_ = ren.DrawRect(&l.pane.rect)
	}

	ren.Present()
}

func renderPane(ren *sdl.Renderer, p *pane) {
	cells, cols, rows, curRow, curCol := p.term.Snapshot()
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cell := cells[row*cols+col]
			ch := cell.Ch
			if ch == 0 {
				ch = ' '
			}
			fg, bg := sdlColor(cell.FG, cell.Bold, false), sdlColor(cell.FG, cell.Bold, true)
			if row == curRow && col == curCol {
				fg, bg = bg, fg
			}
			px := p.rect.X + int32(col*font.GlyphW)
			py := p.rect.Y + int32(row*font.GlyphH)
			_ = ren.SetDrawColor(bg.R, bg.G, bg.B, 0xff)
			_ = ren.FillRect(&sdl.Rect{X: px, Y: py, W: int32(font.GlyphW), H: int32(font.GlyphH)})
			if ch == ' ' {
				continue
			}
			bits := font.GlyphBits(ch)
			_ = ren.SetDrawColor(fg.R, fg.G, fg.B, 0xff)
			for y := 0; y < font.GlyphH; y++ {
				for x := 0; x < font.GlyphW; x++ {
					if bits[y*font.GlyphW+x] != 0 {
						_ = ren.DrawPoint(px+int32(x), py+int32(y))
					}
				}
			}
		}
	}
}

// sdlColor maps a vterm.Color to an SDL color. isBackground selects the
// dim/black default instead of the bright/white default for ColorDefault.
func sdlColor(c vterm.Color, bold, isBackground bool) sdl.Color {
	if c == vterm.ColorDefault {
		if isBackground {
			return sdl.Color{R: 0, G: 0, B: 0, A: 0xff}
		}
		return sdl.Color{R: 0xdd, G: 0xdd, B: 0xdd, A: 0xff}
	}
	base := [16]sdl.Color{
		{R: 0x00, G: 0x00, B: 0x00}, {R: 0xcc, G: 0x33, B: 0x33}, {R: 0x33, G: 0xaa, B: 0x33}, {R: 0xcc, G: 0xaa, B: 0x33},
		{R: 0x33, G: 0x66, B: 0xcc}, {R: 0xaa, G: 0x33, B: 0xaa}, {R: 0x33, G: 0xaa, B: 0xaa}, {R: 0xcc, G: 0xcc, B: 0xcc},
		{R: 0x66, G: 0x66, B: 0x66}, {R: 0xff, G: 0x66, B: 0x66}, {R: 0x66, G: 0xff, B: 0x66}, {R: 0xff, G: 0xff, B: 0x66},
		{R: 0x66, G: 0x99, B: 0xff}, {R: 0xff, G: 0x66, B: 0xff}, {R: 0x66, G: 0xff, B: 0xff}, {R: 0xff, G: 0xff, B: 0xff},
	}
	idx := int(c)
	if idx < 0 || idx > 15 {
		idx = 7
	}
	col := base[idx]
	col.A = 0xff
	return col
}

// keyToPTYBytes maps a small, real subset of non-printable keys to the
// control bytes/escape sequences a shell expects. Printable text arrives
// separately via sdl.TextInputEvent (correct for layout-aware input);
// this covers the keys TextInputEvent never fires for.
func keyToPTYBytes(k sdl.Keysym) []byte {
	if (k.Mod & sdl.KMOD_CTRL) != 0 {
		switch k.Sym {
		case sdl.K_c:
			return []byte{0x03}
		case sdl.K_d:
			return []byte{0x04}
		case sdl.K_l:
			return []byte{0x0c}
		}
	}
	switch k.Sym {
	case sdl.K_RETURN, sdl.K_KP_ENTER:
		return []byte{'\r'}
	case sdl.K_BACKSPACE:
		return []byte{0x7f}
	case sdl.K_TAB:
		return []byte{'\t'}
	case sdl.K_ESCAPE:
		return []byte{0x1b}
	case sdl.K_UP:
		return []byte{0x1b, '[', 'A'}
	case sdl.K_DOWN:
		return []byte{0x1b, '[', 'B'}
	case sdl.K_RIGHT:
		return []byte{0x1b, '[', 'C'}
	case sdl.K_LEFT:
		return []byte{0x1b, '[', 'D'}
	}
	return nil
}
