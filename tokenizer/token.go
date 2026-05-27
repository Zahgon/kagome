package tokenizer

import (
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer/lattice"
)

// TokenClass represents the token class.
type TokenClass lattice.NodeClass

const (
	// DUMMY represents the dummy token.
	DUMMY = TokenClass(lattice.DUMMY)
	// KNOWN represents the token in the dictionary.
	KNOWN = TokenClass(lattice.KNOWN)
	// UNKNOWN represents the token which is not in the dictionary.
	UNKNOWN = TokenClass(lattice.UNKNOWN)
	// USER represents the token in the user dictionary.
	USER = TokenClass(lattice.USER)
)

// String returns string representation of a token class.
func (c TokenClass) String() string { _ = "STUB: not implemented"; return "" }

// Token represents a morph of a sentence.
type Token struct {
	Index    int
	ID       int
	Class    TokenClass
	Position int // byte position
	Start    int
	End      int
	Surface  string
	dict     *dict.Dict
	udict    *dict.UserDict
}

// Features returns contents of a token.
func (t Token) Features() []string { _ = "STUB: not implemented"; return nil }

// FeatureAt returns the i th feature if exists.
//
//nolint:gocyclo
func (t Token) FeatureAt(i int) (string, bool) { _ = "STUB: not implemented"; return "", false }

// UserExtra represents custom segmentation and custom reading for user entries.
type UserExtra struct {
	Tokens   []string
	Readings []string
}

// UserExtra returns extra data if token comes from a user dict.
func (t Token) UserExtra() *UserExtra { _ = "STUB: not implemented"; return nil }

// POS returns POS elements of features.
func (t Token) POS() []string { _ = "STUB: not implemented"; return nil }

// EqualFeatures returns true, if the features of tokens are equal.
func (t Token) EqualFeatures(tt Token) bool { _ = "STUB: not implemented"; return false }

// EqualPOS returns true, if the POSs of tokens are equal.
func (t Token) EqualPOS(tt Token) bool { _ = "STUB: not implemented"; return false }

// EqualFeatures returns true, if the features are equal.
func EqualFeatures(lhs, rhs []string) bool { _ = "STUB: not implemented"; return false }

// InflectionalType returns the inflectional type feature if exists.
func (t Token) InflectionalType() (string, bool) { _ = "STUB: not implemented"; return "", false }

// InflectionalForm returns the inflectional form feature if exists.
func (t Token) InflectionalForm() (string, bool) { _ = "STUB: not implemented"; return "", false }

// BaseForm returns the base form features if exists.
func (t Token) BaseForm() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Reading returns the reading feature if exists.
func (t Token) Reading() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Pronunciation returns the pronunciation feature if exists.
func (t Token) Pronunciation() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (t Token) pickupFromFeatures(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// String returns a string representation of a token.
func (t Token) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if tokens are equal.
func (t Token) Equal(v Token) bool { _ = "STUB: not implemented"; return false }

// TokenData is a data format with all the contents of the token.
type TokenData struct {
	ID            int      `json:"id"`
	Start         int      `json:"start"`
	End           int      `json:"end"`
	Surface       string   `json:"surface"`
	Class         string   `json:"class"`
	POS           []string `json:"pos"`
	BaseForm      string   `json:"base_form"`
	Reading       string   `json:"reading"`
	Pronunciation string   `json:"pronunciation"`
	Features      []string `json:"features"`
}

// NewTokenData returns a data which has with all the contents of the token.
func NewTokenData(t Token) TokenData { _ = "STUB: not implemented"; return *new(TokenData) }
