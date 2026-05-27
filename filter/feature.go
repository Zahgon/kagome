package filter

import (
	"io"
)

type (
	// Feature represents a feature.
	Feature = string
	// Features represents a vector of features.
	Features = []string
)

// Any represents an arbitrary feature.
const Any Feature = "\x00"

type filterEdge struct {
	val Feature
	to  *filterNode
}

type filterNode struct {
	fanout []*filterEdge
}

func (n filterNode) isLeaf() bool { _ = "STUB: not implemented"; return false }

func (n filterNode) has(s Feature) int { _ = "STUB: not implemented"; return 0 }

// FeaturesFilter represents a filter that filters a vector of features.
type FeaturesFilter struct {
	root filterNode
}

// NewFeaturesFilter returns a features filter.
func NewFeaturesFilter(fs ...Features) *FeaturesFilter { _ = "STUB: not implemented"; return nil }

func (f *FeaturesFilter) add(fs Features) { _ = "STUB: not implemented"; return }

// a stronger condition exists.

// clear week conditions.

// Match returns true if a filter matches given features.
func (f *FeaturesFilter) Match(fs Features) bool { _ = "STUB: not implemented"; return false }

// String implements string interface.
func (f *FeaturesFilter) String() string { _ = "STUB: not implemented"; return "" }

func filterString(w io.Writer, indent int, n *filterNode) { _ = "STUB: not implemented"; return }
