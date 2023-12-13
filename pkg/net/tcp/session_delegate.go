package tcp

import (
	"context"
	"sparrow/pkg/net/link/ggnet"
	"sparrow/pkg/net/link/middleware"
)

type MiddlewareErrCallBack func(ggnet.Session, error)
type RequestHandler func(ctx context.Context, requestList []byte, dispatcher func() []byte) (resp interface{}, err error)

type DelegateFilter struct {
	Dispathcher middleware.Dispatcher
	ErrCallBack MiddlewareErrCallBack
	Handler     RequestHandler
}

//func BasicSessionDelegate(
//	delegate *DelegateFilter,
//	contextInit func(ctx context.Context, session ggnet.Session) context.Context,
//) ggnet.SessionDelegate {
//	chain := middleware.ChainDispatchInterceptor(
//		interceptor.FrameEncode(),
//		interceptor.FrameDecode(),
//	)
//	return ggnet.SessionDelegateFunc(func(ctx context.Context, session ggnet.Session, b interface{}) []byte {
//		ret, err := chain(ctx,, b.([]byte), func(ctx context.Context, packet interface{})(interface{}, error) {
//			return delegate.Handler(ctx, ctx, func() []byte){
//				return nil, nil
//			}
//		})
//	})
//}
