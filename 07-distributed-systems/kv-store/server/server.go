package main

import (
	"bufio"
	"fmt"
	"kvstore/kvs"
	"log"
	"net"
	"os"
	"strings"
)

const (
	SOCK         = "/tmp/kvs.sock"
	STORAGE_PATH = "../kvstore.json"
)

type Server struct {
	kvstore *kvs.KeyValStore
	l       net.Listener
}

func (s *Server) handleClientConn(conn net.Conn) {
	defer conn.Close()

	// read and process commands
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "get":
			if len(args) < 2 {
				conn.Write([]byte("Usage: get <key>"))
				continue
			}

			result := s.handleGet(args[1])
			conn.Write([]byte(result))
		case "set":
			if len(args) < 2 {
				conn.Write([]byte("Usage: set <key>=<val>"))
				continue
			}

			result := s.handleSet(args[1])
			conn.Write([]byte(result))
		case "exit":
			conn.Write([]byte("...exiting"))
			conn.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}
}

func (s *Server) handleGet(key string) string {
	key = strings.TrimSpace(key)

	val, exists := s.kvstore.Store[key]

	if exists {
		return val
	} else {
		return fmt.Sprintf("Key: %s not found \n", key)
	}
}

func (s *Server) handleSet(input string) string {
	parts := strings.SplitN(input, "=", 2)

	if len(parts) != 2 {
		return fmt.Sprintf("Usage: set <key>=<val>")
	}

	key := strings.TrimSpace(parts[0])
	if key == "" {
		return fmt.Sprint("Key cannot be empty")
	}

	val := strings.TrimSpace(parts[1])

	s.kvstore.Store[key] = val

	if err := s.kvstore.SaveToDisk(); err != nil {
		return fmt.Sprintf("Error saving data to disk: %s\n", err)
	}

	return fmt.Sprintf("Set: [%s] = %s", key, val)
}

func (s *Server) ListenAndServe() {
	os.Remove(SOCK)

	l, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatalf("Failed to listen on socket: %v", err)
	}

	// set server listener
	s.l = l
	defer s.l.Close()
	log.Printf("Server listening on: %s", SOCK)

	for {
		// accept new connections
		conn, err := s.l.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// spin up new go routine
		go s.handleClientConn(conn)
	}
}

func main() {
	kvstore, err := kvs.NewKeyValStore(STORAGE_PATH)
	if err != nil {
		fmt.Println("Error initializing KeyValue store: ", err)
		return
	}

	server := &Server{kvstore: kvstore}
	server.ListenAndServe()
}
