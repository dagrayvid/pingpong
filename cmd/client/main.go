package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/dagrayvid/pingpong/grpc/pong"
)

var server string
var certfile, keyfile, cafile string

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(cafile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func run() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	log.Printf("Connecting to server %s", server)
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewPongServiceClient(conn)
	stream, err := client.PingPongRPC(context.Background())
	if err != nil {
		log.Fatalf("Cannot open gRPC stream: %v", err)
	}

	waitc := make(chan struct{})

	msg := &pb.PongData{
		Msg:  "Ping",
		Ball: 0,
	}

	go func() {
		for {
			log.Printf("%d Sending: %s\n", msg.Ball, msg.Msg)
			stream.Send(msg)
			time.Sleep(2 * time.Second)
			msg, err = stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%d Received: %s\n", msg.Ball, msg.Msg)
			msg.Ball++
			msg.Msg = "Ping"
		}
	}()
	<-waitc
	stream.CloseSend()
}

func main() {
	flag.StringVar(&server, "server", "localhost:12021", "Server to connect to")
	flag.StringVar(&cafile, "ca-file", "certs/ca-cert.pem", "Location of cafile")
	flag.StringVar(&keyfile, "key-file", "certs/client-key.pem", "Location of key-file")
	flag.StringVar(&certfile, "cert-file", "certs/client-cert.pem", "Location of cert-file")

	flag.Parse()
	run()
}
