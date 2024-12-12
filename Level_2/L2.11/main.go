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
		fmt.Fprintf(os.Stderr, "Ошибка подключения: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Соединение установленно %s\n", address)

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
			fmt.Println("EOF получен, закрываем соединение")
			close(stopChan)
			return
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения при вводе: %v\n", err)
			close(stopChan)
			return
		}

		_, err = conn.Write(buf[:n])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка записи в соединении: %v\n", err)
			close(stopChan)
			return
		}
	}
}

func handleOutput(conn net.Conn, done chan struct{}) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Восстановление от паники в handleOutput:", r)
		}
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Полученно:", scanner.Text())
	}

	select {
	case <-done:
		fmt.Println("Канал уже закрыт")
	default:
		close(done)
	}
}

func waitForSignal(stopChan chan struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signalChan:
		fmt.Println("\nСигнал получен, выходим...")
		close(stopChan)
	case <-stopChan:
	}
}
