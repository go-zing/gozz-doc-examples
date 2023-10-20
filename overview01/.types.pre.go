package overview01

import (
	"context"
	"time"
)

//go:generate gozz run -p "doc" -p "api:prefix=book:foo=bar" -p "impl" -p "tag" ./

// +zz:doc
// +zz:api:./
// +zz:impl:./types.go:type=Implement
// Book manager service
type BookService interface {
	// Get book entity by id
	// +zz:api:get|{id}
	Get(ctx context.Context, query QueryId) (book Book, err error)
	// Get book entities list by query params
	// +zz:api:get
	List(ctx context.Context, query QueryBook) (list ListBook, err error)
	// Create book entity with form
	// +zz:api:post
	Create(ctx context.Context, form FormBook) (book Book, err error)
	// Edit book entity from form
	// +zz:api:put|{id}
	Edit(ctx context.Context, form FormBook) (book Book, err error)
	// Delete book entity by id
	// +zz:api:delete|{id}
	Delete(ctx context.Context, query QueryId) (err error)
}

// +zz:doc
// +zz:tag:json:{{ snake .FieldName }}
type (
	// Uri path id param
	// +zz:tag:uri:{{ snake .FieldName }}
	QueryId struct {
		Id int
	}

	// Query book entities param
	// +zz:tag:query:{{ snake .FieldName }}
	QueryBook struct {
		// Query pagination offser no
		PageNo int
		// Query pagination size
		PageSize int
	}

	// Form to edit or create book entity
	FormBook struct {
		QueryId
		// Title of book entity
		Title string
		// Type of book entity
		Type string
	}

	// Book entity
	Book struct {
		FormBook
		// Book entity create time
		CreatedAt time.Time
		// Book entity create user
		CreatedBy string
		// Book entity last update time
		UpdatedAt time.Time
		// Book entity last update user
		UpdatedBy string
	}

	// Book entities list
	ListBook struct {
		// Query pagination total count in database
		Total int
		// Entities list data
		List []Book
	}
)
