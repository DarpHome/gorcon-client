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
	addressp := flag.String("address", "127.0.0.1:25575", "Address of RCON server")
	passwordp := flag.String("password", "mypassw0rd", "RCON password")
	flag.Parse()
	address := *addressp
	password := *passwordp
	/* rs := gorcon.NewRCONServer(nil)
	rs.Check(gorcon.ForPassword(password)).OnCommand(func(ctx *gorcon.RCONCommandContext) {
		ctx.Reply("hey, " + ctx.Command)
	}).OnLogged(func(ctx *gorcon.RCONContext) {
		fmt.Printf("Someone logged, IP: %s\n", ctx.Connection.RemoteAddr().String())
	}).OnExit(func(ctx *gorcon.RCONContext) {
		fmt.Printf("Someone quit: %s\n", ctx.Connection.RemoteAddr().String())
	}).OnError(func(cctx *gorcon.RCONCommandContext, ctx *gorcon.RCONContext, err error) {
		fmt.Fprintf(os.Stderr, "Cctx: %v\nCtx: %v\nError: %v", cctx, ctx, err)
	})
	go rs.Run(":25585")
	time.Sleep(1 * time.Second) */
	client := gorcon.NewRCONClient()
	PanicIf(client.Connect(address))
	PanicIf(client.Login(password))
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
		if !strings.HasSuffix(res, "\n") {
			res += "\n"
		}
		fmt.Print(res)
	}
	client.Close()
	// rs.Close()
}
