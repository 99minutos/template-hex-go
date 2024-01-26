package cmux

import (
	"example-service/internal/domain/core"
	"github.com/soheilhy/cmux"
	"net"
)

type CmuxContainer struct {
	listener net.Listener
	grpcL    net.Listener
	httpL    net.Listener
	cmux     cmux.CMux
}

func NewCmux(acx *core.AppContext) *CmuxContainer {
	l, err := net.Listen("tcp", ":"+acx.Envs.Port)
	if err != nil {
		acx.Fatalw("unable to listen", "error", err)
	}
	m := cmux.New(l)
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	anyL := m.Match(cmux.Any())
	mux := &CmuxContainer{}
	mux.SetGrpcListener(grpcL)
	mux.SetHttpListener(anyL)
	mux.SetCMux(m)
	return mux
}

func (c *CmuxContainer) GetGrpcListener() net.Listener {
	return c.grpcL
}
func (c *CmuxContainer) GetHttpListener() net.Listener {
	return c.httpL
}
func (c *CmuxContainer) SetGrpcListener(l net.Listener) {
	c.grpcL = l
}
func (c *CmuxContainer) SetHttpListener(l net.Listener) {
	c.httpL = l
}
func (c *CmuxContainer) SetCMux(m cmux.CMux) {
	c.cmux = m
}

func (c *CmuxContainer) Start() {
	err := c.cmux.Serve()
	if err != nil {
		panic(err)
	}
}
