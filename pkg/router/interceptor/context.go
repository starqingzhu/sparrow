package interceptor

import (
	"context"
	"errors"
	"sparrow/pkg/net/link/ggnet"
	"sparrow/pkg/net/link/middleware"
)

func ContextInit(call func(ctx context.Context, session ggnet.Session) context.Context) middleware.DispatchInterceptor {
	return func(ctx context.Context, req interface{}, next middleware.DispatchFunc) (resp interface{}, err error) {
		sess, ok := ggnet.SessionFromContext(ctx)
		if !ok {
			return nil, errors.New("ContextInit session lost")
		}
		ctx = call(ctx, sess)
		return next(ctx, req)
	}
}
