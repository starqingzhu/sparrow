package common

type Server struct {
	Addr string `json:"addr,default=127.0.0.1:8080"`
}

type ServerInterface interface {
	Connect()
	Close()
	Send()
	Recv()
}
