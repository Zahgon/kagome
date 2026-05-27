package lattice

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// subcommand property
var (
	CommandName            = "lattice"
	Description            = `lattice viewer`
	UsageMessage           = "%s [-udict userdict_file] [-dict (ipa|uni)] [-mode (normal|search|extended)] [-output output_file] [-v] sentence"
	Stdout       io.Writer = os.Stdout
	Stderr       io.Writer = os.Stderr
)

// options
type option struct {
	udict   string
	dict    string
	mode    string
	output  string
	verbose bool
	input   string
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

func selectDict(name string) (*dict.Dict, error) { _ = "STUB: not implemented"; return nil, nil }

func selectMode(mode string) tokenizer.TokenizeMode {
	_ = "STUB: not implemented"
	return *new(tokenizer.TokenizeMode)
}

//nolint:nonamedreturns
func command(_ context.Context, opt *option) (err error) { _ = "STUB: not implemented"; return nil }

// Run receives the slice of args and executes the lattice tool
func Run(ctx context.Context, args []string) error { _ = "STUB: not implemented"; return nil }

// Usage provides information on the use of the lattice tool
func Usage() { _ = "STUB: not implemented"; return }

// PrintDefaults prints out the default flags
func PrintDefaults(eh flag.ErrorHandling) { _ = "STUB: not implemented"; return }
