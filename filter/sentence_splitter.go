package filter

// SentenceSplitter is a tiny sentence splitter for japanese texts.
type SentenceSplitter struct {
	Delim               []rune // delimiter set. ex. {'。','．'}
	Follower            []rune // allow following after delimiters. ex. {'」','』'}
	SkipWhiteSpace      bool   // eliminate white space or not
	DoubleLineFeedSplit bool   // splite at '\n\n' or not
	MaxRuneLen          int    // max sentence length
}

// default sentence splitter
var defaultSplitter = &SentenceSplitter{
	Delim:               []rune{'。', '．', '！', '!', '？', '?'},
	Follower:            []rune{'.', '｣', '」', '』', ')', '）', '｝', '}', '〉', '》'},
	SkipWhiteSpace:      true,
	DoubleLineFeedSplit: true,
	MaxRuneLen:          128,
}

// ScanSentences implements SplitFunc interface of bufio.Scanner that returns each sentence of text.
// see. https://pkg.go.dev/bufio#SplitFunc
//
//nolint:nonamedreturns
func ScanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (s SentenceSplitter) isDelim(r rune) bool { _ = "STUB: not implemented"; return false }

func (s SentenceSplitter) isFollower(r rune) bool { _ = "STUB: not implemented"; return false }

// ScanSentences is a split function for a Scanner that returns each sentence of text.
//
//nolint:gocyclo,funlen,nonamedreturns
func (s SentenceSplitter) ScanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// nn indicates \n\n

// split

//nolint:gocritic,nestif

// Request more data

// If we're at EOF, we have a final, non-terminated line. Return it.
