package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DarpHome/gorcon"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	address := flag.String("address", "127.0.0.1:25575", "Address of RCON server")
	password := flag.String("password", "mypassw0rd", "RCON password")
	flag.Parse()
	client := gorcon.NewRCONClient(*address)
	PanicIf(client.Login(*password))
	defer client.Close()
	fmt.Printf("Logged with request ID %d!\n", client.RequestID)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmd, err := reader.ReadString('\n')
		PanicIf(err)
		cmd = strings.Trim(cmd, "\r\n")
		if cmd == "Q" {
			break
		}
		res, err := client.SendCommand(cmd)
		PanicIf(err)
		fmt.Print(res)
	}
}
