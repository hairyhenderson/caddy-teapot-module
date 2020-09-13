module github.com/hairyhenderson/caddy-teapot-module/example

go 1.15

require (
	github.com/caddyserver/caddy/v2 v2.2.0-rc.1.0.20200911194521-4217217badf2
	github.com/hairyhenderson/caddy-teapot-module v0.0.3-0
)

replace github.com/hairyhenderson/caddy-teapot-module => ../
