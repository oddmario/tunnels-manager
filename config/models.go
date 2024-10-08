package config

import "github.com/oddmario/tunnel-manager/tunnel"

type ConfigObject struct {
	Mode                             string
	ApplyKernelTuningTweaks          bool
	MainNetworkInterface             string
	Tunnels                          []*tunnel.Tunnel
	DynamicIPUpdaterAPIIsEnabled     bool
	DynamicIPUpdaterAPIListenAddress string
	DynamicIPUpdaterAPIListenPort    int
	PingTimeout                      int
	PingInterval                     int
	DynamicIPUpdateTimeout           int
	DynamicIPUpdateInterval          int
}
