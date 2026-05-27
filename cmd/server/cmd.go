package server

import (
	"context"
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/ikawaha/kagome-dict/dict"
)

// Stderr is the standard error writer.
var Stderr io.Writer = os.Stderr

// staticFS is an http.FileSystem that serves only non-HTML files and
// denies directory listings. It is used to expose /asset/ without leaking
// server-side template sources (demo.html, graph.html).
type staticFS struct {
	base http.FileSystem
}

func (s staticFS) Open(name string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}

// subcommand property
var (
	CommandName  = "server"
	Description  = `run tokenize server`
	usageMessage = "%s [-http=:6060] [-userdict userdic_file] [-dict (ipa|uni)]"
)

// options
type option struct {
	http    string
	dict    string
	udict   string
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

func command(ctx context.Context, opt *option) error { _ = "STUB: not implemented"; return nil }

// Run receives the slice of args and executes the server
func Run(ctx context.Context, args []string) error { _ = "STUB: not implemented"; return nil }

// Usage provides information on the use of the server
func Usage() { _ = "STUB: not implemented"; return }

// PrintDefaults prints out the default flags
func PrintDefaults(eh flag.ErrorHandling) { _ = "STUB: not implemented"; return }
