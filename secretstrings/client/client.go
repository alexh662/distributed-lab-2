package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"os"
	"time"

	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	filePath := flag.String("file", "../wordlist", "Path to file containing words")
	flag.Parse()

	clients := []*rpc.Client{}
	for i := 8030; true; i++ {
		addr := fmt.Sprintf("127.0.0.1:%d", i)
		client, err := rpc.Dial("tcp", addr)
		if err != nil {
			break
		}
		fmt.Println("connected to server: %s", addr)
		defer client.Close()
		clients = append(clients, client)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		word := scanner.Text()

		if word == "" {
			continue
		}

		request := stubs.Request{Message: word}
		response := new(stubs.Response)

		client := clients[i%len(clients)]
		i++

		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println(response.Message)
	}
}
