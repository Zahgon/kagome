package tokenizer

import (
	"github.com/ikawaha/kagome-dict/dict"
)

// Option represents an option for the tokenizer.
type Option func(*Tokenizer) error

// Nop represents a no operation option.
func Nop() Option { _ = "STUB: not implemented"; return *new(Option) }

// UserDict is a tokenizer option to sets a user dictionary.
func UserDict(d *dict.UserDict) Option { _ = "STUB: not implemented"; return *new(Option) }

// OmitBosEos is a tokenizer option to omit BOS/EOS from output tokens.
func OmitBosEos() Option { _ = "STUB: not implemented"; return *new(Option) }
