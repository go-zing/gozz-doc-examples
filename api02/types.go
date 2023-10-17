package api02

import (
	"context"
)

//go:generate gozz run -p "api" ./

// +zz:api:./
type (
	T interface {
		// +zz:api:get:
		Empty()
		// +zz:api:get:
		Ret() (ret int)
		// +zz:api:get:
		Error() (err error)
		// +zz:api:get:
		RetError() (ret int, err error)
		// +zz:api:get:
		Context(ctx context.Context)
		// +zz:api:get:
		ContextRet(ctx context.Context) (ret int)
		// +zz:api:get:
		ContextError(ctx context.Context) (err error)
		// +zz:api:get:
		ContextRetError(ctx context.Context) (ret int, err error)
		// +zz:api:get:
		Param(param int)
		// +zz:api:get:
		ParamRet(param int) (ret error)
		// +zz:api:get:
		ParamError(param int) (err error)
		// +zz:api:get:
		ParamRetError(param int) (ret int, err error)
		// +zz:api:get:
		ContextParam(ctx context.Context, param int)
		// +zz:api:get:
		ContextParamRet(ctx context.Context, param int) (ret int)
		// +zz:api:get:
		ContextParamError(ctx context.Context, param int) (err error)
		// +zz:api:get:
		ContextParamRetError(ctx context.Context, param int) (ret int, err error)
		// +zz:api:get:
		ComplexParam(param map[context.Context][]struct {
			Field []func(context.Context) interface {
				context.Context
			}
		})
		// +zz:api:get:
		PtrParam(*int)
	}
)
