package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	_ "github.com/abiosoft/caddy-exec"
	_ "github.com/caddy-dns/desec"
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/caddyserver/transform-encoder"
	_ "github.com/darkweak/souin/plugins/caddy"
	_ "github.com/dunglas/caddy-cbrotli"
	_ "github.com/dunglas/mercure/caddy"
	_ "github.com/dunglas/vulcain/caddy"
	_ "github.com/enum-gg/caddy-discord"
	_ "github.com/greenpau/caddy-git"
	_ "github.com/hairyhenderson/caddy-teapot-module"
	_ "github.com/mholt/caddy-l4"
	_ "github.com/mholt/caddy-ratelimit"
	_ "magnax.ca/caddy/gopkg"
)

func main() {
	caddycmd.Main()
}
