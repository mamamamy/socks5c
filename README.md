# socks5c

```sh
ssh -o ProxyCommand="socks5c -server=<proxy.net>:<port> -target=%h:%p -auth -user=<user> -password=<password>" user@server.net
```