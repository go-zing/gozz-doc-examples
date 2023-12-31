// Code generated by gozz:api github.com/go-zing/gozz. DO NOT EDIT.

package api01

import (
	"context"
)

var _ = context.Context(nil)

type Apis struct {
	BookService BookService
	UserService UserService
}

func (s Apis) Iterate(fn func(interface{}, []map[string]interface{})) {
	for _, f := range []func() (interface{}, []map[string]interface{}){
		s._BookService,
		s._UserService,
	} {
		fn(f())
	}
}

func (s Apis) _BookService() (interface{}, []map[string]interface{}) {
	t := s.BookService
	return &t, []map[string]interface{}{
		{
			"name":     "List",
			"resource": "get",
			"options": map[string]string{
				"id":     "BookService.List",
				"prefix": "book_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryBook
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.List(ctx, in)
			},
		},
		{
			"name":     "Get",
			"resource": "get|{id}",
			"options": map[string]string{
				"id":     "BookService.Get",
				"prefix": "book_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryBook
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Get(ctx, in)
			},
		},
		{
			"name":     "Create",
			"resource": "post",
			"options": map[string]string{
				"id":     "BookService.Create",
				"prefix": "book_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in FormBook
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Create(ctx, in)
			},
		},
		{
			"name":     "Edit",
			"resource": "put|{id}",
			"options": map[string]string{
				"id":     "BookService.Edit",
				"prefix": "book_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in FormBook
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Edit(ctx, in)
			},
		},
	}
}

func (s Apis) _UserService() (interface{}, []map[string]interface{}) {
	t := s.UserService
	return &t, []map[string]interface{}{
		{
			"name":     "List",
			"resource": "get",
			"options": map[string]string{
				"id":     "UserService.List",
				"prefix": "user_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryUser
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.List(ctx, in)
			},
		},
		{
			"name":     "Get",
			"resource": "get|{id}",
			"options": map[string]string{
				"id":     "UserService.Get",
				"prefix": "user_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryUser
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Get(ctx, in)
			},
		},
		{
			"name":     "Create",
			"resource": "post",
			"options": map[string]string{
				"id":     "UserService.Create",
				"prefix": "user_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in FormUser
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Create(ctx, in)
			},
		},
		{
			"name":     "Edit",
			"resource": "put|{id}",
			"options": map[string]string{
				"id":     "UserService.Edit",
				"prefix": "user_service",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in FormUser
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Edit(ctx, in)
			},
		},
	}
}
