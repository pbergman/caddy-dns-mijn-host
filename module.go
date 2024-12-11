package template

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	mijn_host "github.com/pbergman/libdns-mijn-host"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *mijn_host.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.mijn-host",
		New: func() caddy.Module { return &Provider{mijn_host.NewProvider("")} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
func (p *Provider) Provision(ctx caddy.Context) error {
	//repl := caddy.NewReplacer()
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	providername [<api_key> <private_key_path>] {
//	    api_key <api_key>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.SetApiKey(d.Val())
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if p.Provider.GetApiKey() != "" {
					return d.Err("ApiKey already set")
				}
				if d.NextArg() {
					p.Provider.SetApiKey(d.Val())
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.GetApiKey() == "" {
		return d.Err("Missing ApiKey")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
