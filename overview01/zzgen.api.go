// Code generated by gozz:api github.com/go-zing/gozz. DO NOT EDIT.

package overview01

import (
	"context"
)

var _ = context.Context(nil)

type Apis struct {
	BookService BookService
}

func (s Apis) Iterate(fn func(interface{}, []map[string]interface{})) {
	for _, f := range []func() (interface{}, []map[string]interface{}){
		s._BookService,
	} {
		fn(f())
	}
}

func (s Apis) _BookService() (interface{}, []map[string]interface{}) {
	t := s.BookService
	return &t, []map[string]interface{}{
		{
			"name":     "Get",
			"resource": "get|{id}",
			"options": map[string]string{
				"prefix": "book",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryId
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Get(ctx, in)
			},
		},
		{
			"name":     "List",
			"resource": "get",
			"options": map[string]string{
				"prefix": "book",
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
			"name":     "Create",
			"resource": "post",
			"options": map[string]string{
				"prefix": "book",
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
				"prefix": "book",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in FormBook
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.Edit(ctx, in)
			},
		},
		{
			"name":     "Delete",
			"resource": "delete|{id}",
			"options": map[string]string{
				"prefix": "book",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in QueryId
				if err := dec(&in); err != nil {
					return nil, err
				}
				return nil, t.Delete(ctx, in)
			},
		},
	}
}
