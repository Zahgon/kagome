package tokenizer

import (
	"io"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer/lattice"
)

// TokenizeMode represents a mode of tokenize.
//
// Kagome has segmentation mode for search such as Kuromoji.
//
//	Normal: Regular segmentation
//	Search: Use a heuristic to do additional segmentation useful for search
//	Extended: Similar to search mode, but also unigram unknown words
type TokenizeMode int

func (m TokenizeMode) String() string { _ = "STUB: not implemented"; return "" }

const (
	// Normal is the normal tokenize mode.
	Normal TokenizeMode = iota + 1
	// Search is the tokenize mode for search.
	Search
	// Extended is the experimental tokenize mode.
	Extended
	// BosEosID means the beginning a sentence (BOS) or the end of a sentence (EOS).
	BosEosID = lattice.BosEosID
)

// Tokenizer represents morphological analyzer.
type Tokenizer struct {
	dict       *dict.Dict     // system dictionary
	userDict   *dict.UserDict // user dictionary
	omitBosEos bool           // omit BOS/EOS
}

// New creates a tokenizer.
func New(d *dict.Dict, opts ...Option) (*Tokenizer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tokenize analyzes a sentence in standard tokenize mode.
func (t Tokenizer) Tokenize(input string) []Token { _ = "STUB: not implemented"; return nil }

// Wakati tokenizes a sentence and returns its divided surface strings.
func (t Tokenizer) Wakati(input string) []string { _ = "STUB: not implemented"; return nil }

// Analyze tokenizes a sentence in the specified mode.
func (t Tokenizer) Analyze(input string, mode TokenizeMode) []Token {
	_ = "STUB: not implemented"
	return nil
}

// Dot returns morphs of a sentence and exports a lattice graph to dot format in standard tokenize mode.
func (t Tokenizer) Dot(w io.Writer, input string) []Token { _ = "STUB: not implemented"; return nil }

// AnalyzeGraph returns morphs of a sentence and exports a lattice graph to dot format.
func (t Tokenizer) AnalyzeGraph(w io.Writer, input string, mode TokenizeMode) []Token {
	_ = "STUB: not implemented"
	return nil
}
