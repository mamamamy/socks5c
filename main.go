package main

import (
	"flag"
	"io"
	"os"
	"sync"

	"golang.org/x/net/proxy"
)

func main() {
	serverAddress := flag.String("server", "", "socks5 server address")
	targetAddress := flag.String("target", "", "target address")
	hasAuth := flag.Bool("auth", false, "socks5 has auth")
	user := flag.String("user", "", "socks5 auth user")
	password := flag.String("password", "", "socks5 auth password")

	flag.Parse()

	if len(*serverAddress) <= 0 || len(*targetAddress) <= 0 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	var auth *proxy.Auth
	if *hasAuth {
		if len(*user) <= 0 || len(*password) <= 0 {
			flag.PrintDefaults()
			os.Exit(-1)
		}
		auth = &proxy.Auth{
			User:     *user,
			Password: *password,
		}
	}

	pd, err := proxy.SOCKS5("tcp", *serverAddress, auth, nil)
	if err != nil {
		panic(err)
	}
	c, err := pd.Dial("tcp", *targetAddress)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		io.Copy(c, os.Stdin)
		wg.Done()
	}()
	go func() {
		io.Copy(os.Stdout, c)
		wg.Done()
	}()
	wg.Wait()
}
