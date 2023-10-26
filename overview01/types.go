package overview01

import (
	"context"
	"time"
)

//go:generate gozz run -p "doc" -p "api:prefix=book" -p "impl" -p "tag" ./

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
		// Entity Id
		Id int `json:"id" uri:"id"`
	}

	// Query book entities param
	// +zz:tag:query:{{ snake .FieldName }}
	QueryBook struct {
		// Query pagination offser no
		PageNo int `json:"page_no" query:"page_no"`
		// Query pagination size
		PageSize int `json:"page_size" query:"page_size"`
	}

	// Form to edit or create book entity
	FormBook struct {
		QueryId
		// Title of book entity
		Title string `json:"title"`
		// Type of book entity
		Type string `json:"type"`
	}

	// Book entity
	Book struct {
		FormBook
		// Book entity create time
		CreatedAt time.Time `json:"created_at"`
		// Book entity create user
		CreatedBy string `json:"created_by"`
		// Book entity last update time
		UpdatedAt time.Time `json:"updated_at"`
		// Book entity last update user
		UpdatedBy string `json:"updated_by"`
	}

	// Book entities list
	ListBook struct {
		// Query pagination total count in database
		Total int `json:"total"`
		// Entities list data
		List []Book `json:"list"`
	}
)

var (
	_ BookService = (*Implement)(nil)
)

type Implement struct{}

func (implement Implement) Get(ctx context.Context, query QueryId) (book Book, err error) {
	panic("not implemented")
}

func (implement Implement) List(ctx context.Context, query QueryBook) (list ListBook, err error) {
	panic("not implemented")
}

func (implement Implement) Create(ctx context.Context, form FormBook) (book Book, err error) {
	panic("not implemented")
}

func (implement Implement) Edit(ctx context.Context, form FormBook) (book Book, err error) {
	panic("not implemented")
}

func (implement Implement) Delete(ctx context.Context, query QueryId) (err error) {
	panic("not implemented")
}
