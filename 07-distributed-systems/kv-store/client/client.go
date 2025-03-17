package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const SOCK = "/tmp/kvs.sock"

func main() {
	conn, err := net.Dial("unix", SOCK)
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	// create chan to receive responses from
	responseChan := make(chan string)

	// start go routine to handle async responses
	go func() {
		for {
			// handle response from server
			response := make([]byte, 1024)
			n, err := conn.Read(response)
			if err != nil {
				log.Fatal("Failed to read from server:", err)
				close(responseChan)
				return
			}

			responseChan <- string(response[:n])
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		fmt.Printf("\nReceived signal %s, shutting down... ", sig)
		os.Exit(0)
	}()

	for {
		fmt.Print("> ")

		if scanner.Scan() {
			// read user input
			input := scanner.Text()

			if input == "clear" {
				clearScreen()
				continue
			}

			_, err := conn.Write([]byte(input + "\n"))
			if err != nil {
				log.Fatal("Failed to write to server:", err)
			}

			if msg, open := <-responseChan; open {
				fmt.Println(msg)
			} else {
				fmt.Println("Server connection closed.")
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from: stdin: %v", err)
	}

}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		fmt.Print("\033[H\033[02J")
	}
}
