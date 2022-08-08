package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	help := flag.Bool("h", false, "usage help")
	flag.Parse()

	if *help || len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "[-d data] /tmp/app.sock /uri")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("Unix HTTP client")

	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", flag.Args()[0])
			},
		},
	}

	var response *http.Response
	var err error
	response, err = httpc.Get("http://unix" + flag.Args()[1])
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, response.Body)
}
