package server

import (
	"net/http"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

// TokenizeHandler represents the tokenizer API server struct
type TokenizeHandler struct {
	tokenizer *tokenizer.Tokenizer
}

// TokenizerRequestBody is the type of the "tokenize" endpoint HTTP request body.
type TokenizerRequestBody struct {
	Input string `json:"sentence"`
	Mode  string `json:"mode,omitempty"`
}

// TokenizerResponseBody is the response type of the "tokenize" endpoint.
type TokenizerResponseBody struct {
	Status bool                  `json:"status"`
	Tokens []tokenizer.TokenData `json:"tokens"`
}

func (h *TokenizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec
