# Teapot module for Caddy v2

Package teapot implements a sample [Caddy v2 module](https://caddyserver.com/docs/extending-caddy).
Its only purpose is to respond with [`418 I'm a teapot`](https://en.wikipedia.org/wiki/Hyper_Text_Coffee_Pot_Control_Protocol)
to every request.

The simplest use could be in a Caddyfile like:

```
*:80

route {
    teapot
}
```

Then, when using a Caddy server with this module enabled:

```console
$ curl -I localhost/any/path
HTTP/1.1 418 I'm a teapot
Server: Caddy
Date: Sun, 01 Mar 2020 15:48:23 GMT
```

See also the [`example/`](./example) folder for an example of usage.
