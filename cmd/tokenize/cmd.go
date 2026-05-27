package tokenize

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// subcommand property
const (
	CommandName  = "tokenize"
	Description  = `command line tokenize`
	usageMessage = "%s [-file input_file] [-dict dic_file] [-userdict user_dic_file]" +
		" [-sysdict (ipa|uni)] [-simple false] [-mode (normal|search|extended)] [-split] [-json]"
)

var (
	// Stdout is the standard writer.
	Stdout io.Writer = os.Stdout
	// Stderr is the standard error writer.
	Stderr io.Writer = os.Stderr
)

// options
type option struct {
	file    string
	dict    string
	udict   string
	sysdict string
	simple  bool
	mode    string
	split   bool
	json    bool
	flagSet *flag.FlagSet
}

// ContinueOnError ErrorHandling // Return a descriptive error.
// ExitOnError                   // Call os.Exit(2).
// PanicOnError                  // Call panic with a descriptive error.flag.ContinueOnError
func newOption(w io.Writer, eh flag.ErrorHandling) *option { _ = "STUB: not implemented"; return nil }

// option settings

func (o *option) parse(args []string) error { _ = "STUB: not implemented"; return nil }

// validations

// OptionCheck receives a slice of args and returns an error if it was not successfully parsed
func OptionCheck(args []string) error { _ = "STUB: not implemented"; return nil }

func selectDict(path, sysdict string, shrink bool) (*dict.Dict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectMode(mode string) tokenizer.TokenizeMode {
	_ = "STUB: not implemented"
	return *new(tokenizer.TokenizeMode)
}

func command(_ context.Context, opt *option) error { _ = "STUB: not implemented"; return nil }

func printTokens(tokens []tokenizer.Token) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:gosec

//nolint:gosec
//nolint:gosec

//nolint:gosec

//nolint:gosec

func printTokensJSON(tokens []tokenizer.Token) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

//nolint:gosec

//nolint:gosec

//nolint:gosec

// Run receives the slice of args and executes the tokenize tool
func Run(ctx context.Context, args []string) error { _ = "STUB: not implemented"; return nil }

// Usage provides information on the use of the tokenize tool
func Usage() { _ = "STUB: not implemented"; return }

// PrintDefaults prints out the default flags
func PrintDefaults(eh flag.ErrorHandling) { _ = "STUB: not implemented"; return }
