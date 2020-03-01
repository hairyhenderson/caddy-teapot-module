package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/hairyhenderson/caddy-teapot-module"
)

func main() {
	caddycmd.Main()
}
