package lattice

import (
	"io"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer/lattice/mem"
)

const (
	maximumCost              = 1<<31 - 1
	maximumUnknownWordLength = 1024
	searchModeKanjiLength    = 2
	searchModeKanjiPenalty   = 3000
	searchModeOtherLength    = 7
	searchModeOtherPenalty   = 1700
)

// TokenizeMode represents how to tokenize sentence.
type TokenizeMode int

const (
	// Normal Mode
	Normal TokenizeMode = iota + 1
	// Search Mode
	Search
	// Extended Mode
	Extended
)

var latticePool = mem.NewPool[Lattice](func() *Lattice {
	return new(Lattice)
})

// Lattice represents a grid of morph nodes.
type Lattice struct {
	Input  string
	Output []*Node
	list   [][]*Node
	dic    *dict.Dict
	udic   *dict.UserDict
}

// New returns a new lattice.
func New(d *dict.Dict, u *dict.UserDict) *Lattice { _ = "STUB: not implemented"; return nil }

// Free releases a memory of a lattice.
func (la *Lattice) Free() { _ = "STUB: not implemented"; return }

func (la *Lattice) addNode(pos, id, position, start int, class NodeClass, surface string) {
	_ = "STUB: not implemented"
	return
}

// use default cost

// use default cost

// Build builds a lattice from the inputs.
//
//nolint:gocyclo,funlen
func (la *Lattice) Build(inp string) { _ = "STUB: not implemented"; return }

// (1) USER DIC

// (2) KNOWN DIC

// (3) UNKNOWN DIC

//nolint:nestif

//nolint:wastedassign

// add the string with one character truncated at the end.

// String returns a debug string of a lattice.
func (la *Lattice) String() string { _ = "STUB: not implemented"; return "" }

func kanjiOnly(s string) bool { _ = "STUB: not implemented"; return false }

func additionalCost(n *Node) int { _ = "STUB: not implemented"; return 0 }

// Forward runs forward algorithm of the Viterbi.
func (la *Lattice) Forward(m TokenizeMode) { _ = "STUB: not implemented"; return }

//nolint:gosec // G115: integer overflow conversion int64 -> int32
//nolint:gosec // G115: integer overflow conversion int64 -> int32

// Backward runs backward algorithm of the Viterbi.
func (la *Lattice) Backward(m TokenizeMode) { _ = "STUB: not implemented"; return }

func posFeature(d *dict.Dict, u *dict.UserDict, t *Node) string {
	_ = "STUB: not implemented"
	return ""
}

// undefined

// Dot outputs a lattice in the graphviz dot format.
//
//nolint:gocyclo,funlen
func (la *Lattice) Dot(w io.Writer) { _ = "STUB: not implemented"; return }
