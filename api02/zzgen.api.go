// Code generated by gozz:api github.com/go-zing/gozz. DO NOT EDIT.

package api02

import (
	"context"
)

var _ = context.Context(nil)

type Apis struct {
	T T
}

func (s Apis) Range(fn func(interface{}, []map[string]interface{})) {
	for _, f := range []func() (interface{}, []map[string]interface{}){
		s._T,
	} {
		fn(f())
	}
}

func (s Apis) _T() (interface{}, []map[string]interface{}) {
	t := s.T
	return &t, []map[string]interface{}{
		{
			"name":     "Empty",
			"resource": "get",
			"options": map[string]string{
				"path": "empty",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				t.Empty()
				return nil, nil
			},
		},
		{
			"name":     "Ret",
			"resource": "get",
			"options": map[string]string{
				"path": "ret",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) { return t.Ret(), nil },
		},
		{
			"name":     "Error",
			"resource": "get",
			"options": map[string]string{
				"path": "error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) { return nil, t.Error() },
		},
		{
			"name":     "RetError",
			"resource": "get",
			"options": map[string]string{
				"path": "ret_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) { return t.RetError() },
		},
		{
			"name":     "Context",
			"resource": "get",
			"options": map[string]string{
				"path": "context",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				t.Context(ctx)
				return nil, nil
			},
		},
		{
			"name":     "ContextRet",
			"resource": "get",
			"options": map[string]string{
				"path": "context_ret",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				return t.ContextRet(ctx), nil
			},
		},
		{
			"name":     "ContextError",
			"resource": "get",
			"options": map[string]string{
				"path": "context_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				return nil, t.ContextError(ctx)
			},
		},
		{
			"name":     "ContextRetError",
			"resource": "get",
			"options": map[string]string{
				"path": "context_ret_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				return t.ContextRetError(ctx)
			},
		},
		{
			"name":     "Param",
			"resource": "get",
			"options": map[string]string{
				"path": "param",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				t.Param(in)
				return nil, nil
			},
		},
		{
			"name":     "ParamRet",
			"resource": "get",
			"options": map[string]string{
				"path": "param_ret",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return nil, t.ParamRet(in)
			},
		},
		{
			"name":     "ParamError",
			"resource": "get",
			"options": map[string]string{
				"path": "param_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return nil, t.ParamError(in)
			},
		},
		{
			"name":     "ParamRetError",
			"resource": "get",
			"options": map[string]string{
				"path": "param_ret_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.ParamRetError(in)
			},
		},
		{
			"name":     "ContextParam",
			"resource": "get",
			"options": map[string]string{
				"path": "context_param",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				t.ContextParam(ctx, in)
				return nil, nil
			},
		},
		{
			"name":     "ContextParamRet",
			"resource": "get",
			"options": map[string]string{
				"path": "context_param_ret",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.ContextParamRet(ctx, in), nil
			},
		},
		{
			"name":     "ContextParamError",
			"resource": "get",
			"options": map[string]string{
				"path": "context_param_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return nil, t.ContextParamError(ctx, in)
			},
		},
		{
			"name":     "ContextParamRetError",
			"resource": "get",
			"options": map[string]string{
				"path": "context_param_ret_error",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				return t.ContextParamRetError(ctx, in)
			},
		},
		{
			"name":     "ComplexParam",
			"resource": "get",
			"options": map[string]string{
				"path": "complex_param",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in map[context.Context][]struct {
					Field []func(context.Context) interface {
						context.Context
					}
				}
				if err := dec(&in); err != nil {
					return nil, err
				}
				t.ComplexParam(in)
				return nil, nil
			},
		},
		{
			"name":     "PtrParam",
			"resource": "get",
			"options": map[string]string{
				"path": "ptr_param",
			},
			"invoke": func(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
				var in int
				if err := dec(&in); err != nil {
					return nil, err
				}
				t.PtrParam(&in)
				return nil, nil
			},
		},
	}
}
