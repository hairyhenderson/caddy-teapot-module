module github.com/hairyhenderson/caddy-teapot-module/example

go 1.14

require (
	github.com/caddyserver/caddy/v2 v2.0.0-beta.15
	github.com/hairyhenderson/caddy-teapot-module v0.0.1
)

replace github.com/hairyhenderson/caddy-teapot-module => ../
