// Package teapot implements a sample Caddy v2 module. Its only purpose is to
// respond with "418 I'm a teapot" to every request.
//
// The simplest use could be in a Caddyfile like:
//
//    localhost
//
//    route {
//        teapot
//    }
//
package teapot

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Teapot{})
	httpcaddyfile.RegisterHandlerDirective("teapot", parseCaddyfile)
}

// CaddyModule returns the Caddy module information.
func (Teapot) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.teapot",
		New: func() caddy.Module { return new(Teapot) },
	}
}

// Teapot implements a static "418 I'm a teapot" response to all requests on the route
type Teapot struct{}

// UnmarshalCaddyfile - this is a no-op
func (s *Teapot) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	t := new(Teapot)
	err := t.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s Teapot) ServeHTTP(w http.ResponseWriter, r *http.Request, _ caddyhttp.Handler) error {
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("I'm a teapot\r\n"))

	return nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*Teapot)(nil)
	_ caddyfile.Unmarshaler       = (*Teapot)(nil)
)
