package filter

import (
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type (
	// POS represents a part-of-speech that is a vector of features.
	POS = []string
)

// POSFilter represents a part-of-speech filter.
type POSFilter struct {
	filter *FeaturesFilter
}

// NewPOSFilter returns a part-of-speech filter.
func NewPOSFilter(p ...POS) *POSFilter { _ = "STUB: not implemented"; return nil }

// Match returns true if a filter matches given POS.
func (f POSFilter) Match(p POS) bool { _ = "STUB: not implemented"; return false }

// Drop drops a token if a filter matches token's POS.
func (f POSFilter) Drop(tokens *[]tokenizer.Token) { _ = "STUB: not implemented"; return }

// Keep keeps a token if a filter matches token's POS.
func (f POSFilter) Keep(tokens *[]tokenizer.Token) { _ = "STUB: not implemented"; return }

func (f POSFilter) apply(tokens *[]tokenizer.Token, drop bool) { _ = "STUB: not implemented"; return }
