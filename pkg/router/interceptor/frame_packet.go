package interceptor

import (
	"context"
	"sparrow/pkg/net/link/middleware"
)

func FrameEncode() middleware.DispatchInterceptor {
	return func(ctx context.Context, req interface{}, next middleware.DispatchFunc) (interface{}, error) {
		//加密

		ret, err := next(ctx, req)
		return ret, err
	}
}

func FrameDecode() middleware.DispatchInterceptor {
	return func(ctx context.Context, req interface{}, next middleware.DispatchFunc) (interface{}, error) {
		//解密
		ret, err := next(ctx, req)
		return ret, err
	}
}
