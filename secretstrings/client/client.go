package main

import (
	"bufio"
	"flag"
	"log"
	"net/rpc"
	"os"

	//	"bufio"
	//	"os"
	//	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
	"fmt"

	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	filePath := flag.String("file", "../wordlist", "Path to file containing words")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		if word == "" {
			continue
		}

		request := stubs.Request{Message: word}
		response := new(stubs.Response)
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println(response.Message)
	}
}
