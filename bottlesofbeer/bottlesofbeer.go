package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

var nextAddr string

type Beer int

func (b *Beer) SingVerse(n int, reply *bool) error {
	if n <= 0 {
		fmt.Println("no more bottles of beer on the wall")
		*reply = false
		return nil
	}

	fmt.Printf("%d bottles of beer on the wall, %d bottles of beer, take one down, pass it around...\n", n, n)

	client, err := rpc.Dial("tcp", nextAddr)
	if err != nil {
		fmt.Println(err)
		*reply = false
		return err
	}
	defer client.Close()

	var ok bool
	call := client.Go("Beer.SingVerse", n-1, &ok, nil)
	<-call.Done
	*reply = ok
	return nil
}

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	beer := new(Beer)
	rpc.Register(beer)
	listener, _ := net.Listen("tcp", ":"+*thisPort)

	go rpc.Accept(listener)

	if *bottles > 0 {
		var reply bool
		beer.SingVerse(*bottles, &reply)
	}

	select {}
}
