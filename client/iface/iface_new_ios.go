//go:build ios

package iface

import (
	"github.com/netbirdio/netbird/client/iface/bind"
	"github.com/netbirdio/netbird/client/iface/device"
	"github.com/netbirdio/netbird/client/iface/netstack"
	"github.com/netbirdio/netbird/client/iface/wgproxy"
)

// NewWGIFace Creates a new WireGuard interface instance
func NewWGIFace(opts WGIFaceOpts) (*WGIface, error) {
	iceBind := bind.NewICEBind(opts.TransportNet, opts.Address, opts.MTU)
	if netstack.IsEnabled() {
		return &WGIface{
			tun:            device.NewNetstackDevice(opts.IFaceName, opts.Address, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, netstack.ListenAddr()),
			userspaceBind:  true,
			wgProxyFactory: wgproxy.NewUSPFactory(iceBind, opts.MTU),
		}, nil
	}

	wgIFace := &WGIface{
		tun:            device.NewTunDevice(opts.IFaceName, opts.Address, opts.WGPort, opts.WGPrivKey, opts.MTU, iceBind, opts.MobileArgs.TunFd),
		userspaceBind:  true,
		wgProxyFactory: wgproxy.NewUSPFactory(iceBind, opts.MTU),
	}
	return wgIFace, nil
}
