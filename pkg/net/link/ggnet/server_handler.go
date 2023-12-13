package ggnet

type Handler interface {
	HandleSession(Session)
}

type HandlerFunc func(Session)

func (f HandlerFunc) HandleSession(sess Session) {
	f(sess)
}
