package book

import (
	"context"
	"testing"

	"github.com/scylladb/scylla-cloud/gen/proto/srv/book/v1"
)

var s *Service

func init() {
	s = NewService()
}

func Test_BookValidation(t *testing.T) {
	testCases := []struct {
		name  string
		book  *book.Book
		valid bool
		err   string
	}{
		{
			name: "Valid Book",
			book: &book.Book{
				Title:  "Some Book",
				Author: "Kevin Barbour",
			},
			valid: true,
		},
		{
			name: "No Title",
			book: &book.Book{
				Author: "Kevin Barbour",
			},
			valid: false,
			err:   `invalid Book.Title: value length must be between 1 and 512 bytes, inclusive`,
		},
		{
			name: "Invalid Author",
			book: &book.Book{
				Title:  "An Okay Book",
				Author: "kevin Barbour",
			},
			valid: false,
			err:   `invalid Book.Author: value does not match regex pattern "^[A-Z]+.*$"`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			req := &book.CreateBookRequest{
				Book: test.book,
			}

			_, err := s.CreateBook(context.Background(), req)

			if test.valid {
				if err != nil {
					t.Fatalf("Expected no error, got %s", err)
				}
			} else {
				if err == nil || err.Error() != test.err {
					t.Fatalf("Expected err (%s) got err (%s)", test.err, err)
				}
			}
		})
	}
}
