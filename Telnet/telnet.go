package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var host string
var port string

func init() {
	pflag.StringVarP(&port, "port", "p", "", "Port")
	pflag.StringVarP(&host, "host", "h", "", "Host or IP")
}

func ConnReader(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				log.Fatalf("Cannot scan from %v\n", conn.RemoteAddr())
			}
			text := scanner.Text()
			log.Printf(" | Receive from %v : %s", conn.RemoteAddr(), text)
		}
	}
}

func ConnWriter(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				log.Printf("Cannot scan from stdin\n")
				break
			}
			text := scanner.Text()
			log.Printf("| Send to %v : %s", conn.RemoteAddr(), text)

			_, err := conn.Write([]byte(fmt.Sprintf("%s\n", text)))
			if err != nil {
				log.Fatalf("Cannot write to %v: %v", conn.RemoteAddr(), err)
			}
		}

	}
}

func CloseSignalInterceptor() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Interrupt by Ctrl+C")
		os.Exit(0)
	}()
}

func main() {

	pflag.Parse()
	if host == "" {
		log.Fatalf("Host argument is missed\n")
	}
	if port == "" {
		log.Fatalf("Port argument is missed\n")
	}

	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10 * time.Second)

	conn, err := dialer.DialContext(ctx, "tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Fatalf("Cannot connect to %v:%v %v\n",host, port,  err)
	}

	defer conn.Close()

	log.Printf("Connected to %v:%v\n", host, port)

	var wg sync.WaitGroup

	// Ctrl+C Handler
	CloseSignalInterceptor()

	//Reader goroutine
	wg.Add(1)
	go func() {
		ConnReader(ctx, conn)
		wg.Done()
	}()

	//Writer goroutine
	wg.Add(1)
	go func() {
		ConnWriter(ctx, conn)
		wg.Done()
	}()

	wg.Wait()
}
