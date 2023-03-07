package book

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"

	pb "github.com/scylladb/scylla-cloud/gen/proto/srv/book/v1"

	"github.com/twitchtv/twirp"
)

// Service is a twirp book service.
type Service struct {
	bookIDCounter int
	booksByID     map[string]*pb.Book
	mu            sync.Mutex
}

var _ pb.BookService = (*Service)(nil)

// NewService returns a new book service.
func NewService() *Service {
	return &Service{
		booksByID: make(map[string]*pb.Book),
	}
}

// CreateBook creates a single new book.
func (s *Service) CreateBook(ctx context.Context, r *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bookIDCounter++
	id := fmt.Sprint(s.bookIDCounter)
	r.Book.Id = id

	if err := r.Book.Validate(); err != nil {
		s.bookIDCounter--
		return nil, err
	}

	s.booksByID[id] = r.Book
	return &pb.CreateBookResponse{
		Book: r.Book,
	}, nil
}

// GetBook gets a single book.
func (s *Service) GetBook(ctx context.Context, r *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	if r.Id == "123" {
		return nil, twirp.NewError(twirp.NotFound, "foo")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	return &pb.GetBookResponse{
		Book: s.booksByID[r.Id],
	}, nil
}

// ListBooks lists all requested books.
func (s *Service) ListBooks(ctx context.Context, r *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	allBooks := make([]*pb.Book, 0, len(s.booksByID))

	s.mu.Lock()
	if len(r.Ids) > 0 {
		for _, id := range r.Ids {
			b := s.booksByID[id]
			if b != nil {
				allBooks = append(allBooks, b)
			}
		}
	} else {
		for _, b := range s.booksByID {
			allBooks = append(allBooks, b)
		}
	}
	s.mu.Unlock()

	sort.Slice(allBooks, func(i, j int) bool {
		a, _ := strconv.ParseInt(allBooks[i].Id, 10, 64) // nolint: errcheck
		b, _ := strconv.ParseInt(allBooks[j].Id, 10, 64) // nolint: errcheck
		return a < b
	})

	// Pagination is not implemented

	return &pb.ListBooksResponse{
		Books: allBooks,
	}, nil
}

// DeleteBook deletes a single book.
func (s *Service) DeleteBook(ctx context.Context, r *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.booksByID, r.Id)
	return &pb.DeleteBookResponse{}, nil
}

// UpdateBook updates a single book.
func (s *Service) UpdateBook(ctx context.Context, r *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r.Book.Id = r.Id

	if err := r.Book.Validate(); err != nil {
		return nil, err
	}

	s.booksByID[r.Id] = r.Book
	return &pb.UpdateBookResponse{
		Book: r.Book,
	}, nil
}
