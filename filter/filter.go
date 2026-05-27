package filter

import (
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// Drop drops a token given the provided match function.
func Drop(tokens *[]tokenizer.Token, match func(t tokenizer.Token) bool) {
	_ = "STUB: not implemented"
	return
}

// Keep keeps a token given the provided match function.
func Keep(tokens *[]tokenizer.Token, match func(t tokenizer.Token) bool) {
	_ = "STUB: not implemented"
	return
}

func applyFilter(match func(t tokenizer.Token) bool, tokens *[]tokenizer.Token, drop bool) {
	_ = "STUB: not implemented"
	return
}
