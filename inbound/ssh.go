package inbound

import (
	"context"
	"net"
	"os"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/uot"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/auth"
	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/sing/protocol/socks"
)

var (
	_ adapter.Inbound           = (*Socks)(nil)
	_ adapter.InjectableInbound = (*Socks)(nil)
)

type SSH struct {
	myInboundAdapter
}

func NewSSH(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.SSHInboundOptions) *SSH {
	inbound := &SSH{
		myInboundAdapter{
			protocol:      C.TypeSOCKS,
			network:       options.Network.Build(),
			ctx:           ctx,
			router:        uot.NewRouter(router, logger),
			logger:        logger,
			tag:           tag,
			listenOptions: options.ListenOptions,
		},
	}
	inbound.connHandler = inbound
	return inbound
}

func (h *SSH) NewConnection(ctx context.Context, conn net.Conn, metadata adapter.InboundContext) error {
	
	return socks.HandleConnection(ctx, conn, h.authenticator, h.upstreamUserHandler(metadata), adapter.UpstreamMetadata(metadata))
}

func (h *SSH) NewPacketConnection(ctx context.Context, conn N.PacketConn, metadata adapter.InboundContext) error {
	return os.ErrInvalid
}


