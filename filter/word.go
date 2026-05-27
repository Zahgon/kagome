package filter

import (
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// WordFilter represents a word filter.
type WordFilter struct {
	words map[string]struct{}
}

// NewWordFilter returns a word filter.
func NewWordFilter(words []string) *WordFilter { _ = "STUB: not implemented"; return nil }

// Match returns true if a filter matches a given word.
func (f WordFilter) Match(w string) bool { _ = "STUB: not implemented"; return false }

// Drop drops a token if a filter matches token's surface.
func (f WordFilter) Drop(tokens *[]tokenizer.Token) { _ = "STUB: not implemented"; return }

// Keep keeps a token if a filter matches token's surface.
func (f WordFilter) Keep(tokens *[]tokenizer.Token) { _ = "STUB: not implemented"; return }

func (f WordFilter) apply(tokens *[]tokenizer.Token, drop bool) { _ = "STUB: not implemented"; return }
