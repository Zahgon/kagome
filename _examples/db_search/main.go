/*
# Full-text search with Kagome and SQLite3

This example provides a practical example of how to work with Japanese text data and perform efficient full-text search using Kagome and SQLite3.

For details and acknowledgements, see the README.md file in the same directory.
*/
package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	_ = "STUB: not implemented"
	// Contents to be inserted into the database. Each element represents a line
	// of text and will be inserted into a row of the database.
	return nil
}

// Create a database. In-memory database is used for simplicity.

// Create tables.
// The first table "contents_fts" is for storing the original content, and
// the second table "fts" is for storing the tokenized content.

// Insert contents

// Search by word.
// Note the difference between words and character search.

// "人魚" should be found at line 1, 4, 6.

// "人" exists as a character but should not be found since it is not used
// as a word.

// "北方" should be found at line 3.

// "北" should be found at line 2.
// The character "北" itself exists in line 3 as well, but it is not used
// as a word. Therefore, line 3 should not match.

// Print search results

// Output:
// Searching for: 人魚
//   Found content: 人魚は、南の方の海にばかり棲んでいるのではありません。 at line: 1
//   Found content: ある時、岩の上に、女の人魚があがって、 at line: 4
//   Found content: 小川未明 『赤い蝋燭と人魚』 at line: 6
// Searching for: 人
//   No results found
// Searching for: 北方
//   Found content: 北方の海の色は、青うございました。 at line: 3
// Searching for: 北
//   Found content: 北の海にも棲んでいたのであります。 at line: 2

func insertContent(db *sql.DB, content string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func insertSearchToken(db *sql.DB, rowID int64, content string) error {
	_ = "STUB: not implemented"
	// This example uses the IPA dictionary, but it may be more efficient to use
	// the 'Uni' dictionary if memory is available.
	return nil
}

// remove duplicate segment tokens

func retrieveContent(db *sql.DB, rowID int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func searchFTS4(db *sql.DB, searchWord string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Debug
// fmt.Printf("- Table: fts, RowID: %d, Value: %s\n", lineID, words)
