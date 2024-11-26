package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	address, timeout := parseArgs()
	if address == "" {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(2)
	}

	conn, err := connectToServer(address, timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	stopChan := make(chan struct{})
	go handleInput(conn, stopChan)
	go handleOutput(conn, stopChan)

	waitForSignal(stopChan)
}

func parseArgs() (string, time.Duration) {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		return "", timeout
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	return net.JoinHostPort(host, port), timeout
}

func connectToServer(address string, timeout time.Duration) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func handleInput(conn net.Conn, stopChan chan struct{}) {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			fmt.Println("EOF received, closing connection")
			close(stopChan)
			return
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			close(stopChan)
			return
		}

		if _, err := conn.Write(buf[:n]); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
			close(stopChan)
			return
		}
	}
}

func handleOutput(conn net.Conn, done chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in handleOutput:", r)
		}
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Received:", scanner.Text())
	}

	select {
	case <-done:
		fmt.Println("Channel already closed")
	default:
		close(done)
	}
}

func waitForSignal(stopChan chan struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		fmt.Println("\nSignal received, exiting...")
		close(stopChan)
	case <-stopChan:
	}
}
