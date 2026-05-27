package sentence

import (
	"context"
	"flag"
	"io"
	"os"
)

// subcommand property
const (
	CommandName  = "sentence"
	Description  = `tiny sentence splitter`
	usageMessage = "%s [-file filename]"
)

// Stderr writes to stderr
var (
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
)

// options
type option struct {
	file    string
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

func command(_ context.Context, w io.Writer, opt *option) error {
	_ = "STUB: not implemented"
	return nil
}

// Run receives the slice of args and executes the tokenize tool
func Run(ctx context.Context, args []string) error { _ = "STUB: not implemented"; return nil }

// Usage provides information on the use of the tokenize tool
func Usage() { _ = "STUB: not implemented"; return }

// PrintDefaults prints out the default flags
func PrintDefaults(eh flag.ErrorHandling) { _ = "STUB: not implemented"; return }
