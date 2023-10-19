package api01

import (
	"context"
)

type (
	QueryBook struct{}
	ListBook  struct{}
	DataBook  struct{}
	FormBook  struct{}

	QueryUser struct{}
	ListUser  struct{}
	DataUser  struct{}
	FormUser  struct{}
)

//go:generate gozz run -p "api" ./

// +zz:api:./:prefix={{ snake .Name }}:id={{ .Name }}.{{ .FieldName }}
type (
	BookService interface {
		// +zz:api:get
		List(ctx context.Context, query QueryBook) (ret ListBook, err error)
		// +zz:api:get|{id}
		Get(ctx context.Context, query QueryBook) (data DataBook, err error)
		// +zz:api:post
		Create(ctx context.Context, form FormBook) (data DataBook, err error)
		// +zz:api:put|{id}
		Edit(ctx context.Context, form FormBook) (data DataBook, err error)
	}

	UserService interface {
		// +zz:api:get
		List(ctx context.Context, query QueryUser) (ret ListUser, err error)
		// +zz:api:get|{id}
		Get(ctx context.Context, query QueryUser) (data DataBook, err error)
		// +zz:api:post
		Create(ctx context.Context, query FormUser) (data DataBook, err error)
		// +zz:api:put|{id}
		Edit(ctx context.Context, form FormUser) (data DataBook, err error)
	}
)
