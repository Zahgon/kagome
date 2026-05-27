package ja

import (
	_ "embed"

	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// Filter represents a japanese token filter.
type Filter struct {
	baserForm *filter.POSFilter
	stopTags  *filter.POSFilter
	stopWords *filter.WordFilter
}

// FilterOption represents an option of the japanese token filter.
type FilterOption func(*Filter)

// BaseFormFilterOption returns a base form filter option.
func BaseFormFilterOption(p []filter.POS) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

// StopTagsFilterOption returns a stop tags filter option.
func StopTagsFilterOption(p []filter.POS) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

// StopWordsFilterOption returns a stop words filter option.
func StopWordsFilterOption(p []string) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

// NewFilter returns a filter with the settings commonly used in lucene.
// To customize, set the options and overwrite the filter.
func NewFilter(opts ...FilterOption) (*Filter, error) { _ = "STUB: not implemented"; return nil, nil }

//go:embed asset/stop_tags.txt
var stopTags []byte

//go:embed asset/stop_words.txt
var stropWords []byte

const (
	posHierarchy      = 4
	defaultPOSFeature = "*"
)

const (
	POS_動詞   = "動詞"   //nolint:asciicheck,gosmopolitan
	POS_形容詞  = "形容詞"  //nolint:asciicheck,gosmopolitan
	POS_形容動詞 = "形容動詞" //nolint:asciicheck,gosmopolitan
)

func newDefaultLuceneFilter() (*Filter, error) { _ = "STUB: not implemented"; return nil, nil }

func newDefaultLuceneStopTagPOSFilter() (*filter.POSFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDefaultLuceneStopWordFilter() (*filter.WordFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Yield returns a filtered word sequence from a token sequence.
func (f Filter) Yield(tokens []tokenizer.Token) []string { _ = "STUB: not implemented"; return nil }

// Drop drops a token given the provided match function (stop-tags and stop-words).
func (f Filter) Drop(tokens *[]tokenizer.Token) { _ = "STUB: not implemented"; return }
