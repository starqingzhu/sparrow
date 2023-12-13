package middleware

import "context"

type DispatchFunc func(context.Context, interface{}) (interface{}, error)

func (f DispatchFunc) Dispatch(ctx context.Context, in interface{}) (interface{}, error) {
	return f(ctx, in)
}

type Dispatcher interface {
	Dispatch(ctx context.Context, in interface{}) (interface{}, error)
}

type DispatchInterceptor func(context.Context, interface{}, DispatchFunc) (interface{}, error)

func ChainDispatchInterceptor(interceptors ...DispatchInterceptor) DispatchInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, req interface{}, handler DispatchFunc) (interface{}, error) {

		chain := func(currentInter DispatchInterceptor, currentHandler DispatchFunc) DispatchFunc {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chain(interceptors[i], chainedHandler)
		}
		return chainedHandler(ctx, req)
	}
}
