package main

import (
	"fmt"
	"github.com/openziti/zrok/environment"
	"github.com/openziti/zrok/sdk"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("usage: pastefrom <shrToken>")
	}
	shrToken := os.Args[1]

	root, err := environment.LoadRoot()
	if err != nil {
		panic(err)
	}

	acc, err := sdk.CreateAccess(root, &sdk.AccessRequest{ShareToken: shrToken})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := sdk.DeleteAccess(root, acc); err != nil {
			panic(err)
		}
	}()

	conn, err := sdk.NewDialer(shrToken, root)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	buf := make([]byte, 10240)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
