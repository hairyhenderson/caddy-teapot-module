module github.com/hairyhenderson/caddy-teapot-module/example

go 1.14

replace github.com/hairyhenderson/caddy-teapot-module => ../

require (
	github.com/caddyserver/caddy/v2 v2.0.0-beta.15
	github.com/hairyhenderson/caddy-teapot-module v0.0.0-00010101000000-000000000000
)
