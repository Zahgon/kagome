package server

import (
	"context"
	"embed"
	"html/template"
	"net/http"
	"time"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

//go:embed asset
var assetFS embed.FS

// assets
var (
	graphT = mustParseTemplate("graph", "asset/graph.html")
	demoT  = mustParseTemplate("demo", "asset/demo.html")
)

func mustParseTemplate(name, path string) *template.Template { _ = "STUB: not implemented"; return nil }

const (
	graphvizCmd = "dot"
	cmdTimeout  = 25 * time.Second
)

// TokenizeDemoHandler represents the tokenizer demo server struct.
type TokenizeDemoHandler struct {
	tokenizer *tokenizer.Tokenizer
}

type record struct {
	Surface       string
	POS           string
	BaseForm      string
	Reading       string
	Pronunciation string
}

func newRecord(t tokenizer.Token) record { _ = "STUB: not implemented"; return *new(record) }

func toRecords(tokens []tokenizer.Token) []record { _ = "STUB: not implemented"; return nil }

//nolint:nonamedreturns
func analyzeGraph(ctx context.Context, tnz *tokenizer.Tokenizer, sen string, mode tokenizer.TokenizeMode) (records []record, svg string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// LatticeHandler represents the lattice graph handler that returns SVG as JSON.
type LatticeHandler struct {
	tokenizer *tokenizer.Tokenizer
}

type latticeResponse struct {
	SVG   string `json:"svg,omitempty"`
	Error string `json:"error,omitempty"`
}

// ServeHTTP serves the lattice SVG as JSON.
func (h *LatticeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// ServeHTTP serves a tokenize demo server.
func (h *TokenizeDemoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Extended uses search mode
