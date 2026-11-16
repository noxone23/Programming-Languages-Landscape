// Package readinglist is the Reading List logic, ported from TypeScript
// (see ../typescript/books.ts and ../SPEC.md).
package readinglist

import "strings"

// Book is one book of the reading list.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"` // "want-to-read", "reading" or "done"
	Rating *int   `json:"rating"` // 1 to 5, nil means "not rated"
	Notes  string `json:"notes"`
}

// StatusLabels maps every status to its label.
var StatusLabels = map[string]string{
	"want-to-read": "Want to read",
	"reading":      "Reading",
	"done":         "Done",
}

// FilterBooks returns the books with the given status that match the query.
// status is "want-to-read", "reading", "done" or "all".
func FilterBooks(books []Book, status string, query string) []Book {
	normalizedQuery := strings.ToLower(strings.TrimSpace(query))
	result := []Book{}
	for _, book := range books {
		statusMatches := status == "all" || book.Status == status
		if statusMatches && strings.Contains(strings.ToLower(book.Title), normalizedQuery) {
			result = append(result, book)
		}
	}
	return result
}

// ValidateBook returns an error message per invalid field ("title", "author").
// An empty map means valid.
func ValidateBook(title string, author string) map[string]string {
	errors := map[string]string{}
	if strings.TrimSpace(title) == "" {
		errors["title"] = "Title is required"
	}
	if strings.TrimSpace(author) == "" {
		errors["author"] = "Author is required"
	}
	return errors
}
