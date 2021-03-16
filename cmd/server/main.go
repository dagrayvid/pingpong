package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/ArangoGutierrez/pingpong/grpc/pong"
)

var certfile, keyfile, cafile string

// PongServer empty struct for gRPC interfaces
type PongServer struct {
}

// PingPongRPC stream gRPC func
func (ps *PongServer) PingPongRPC(stream pb.PongService_PingPongRPCServer) error {
	log.Println("Started stream")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("%d Received: %s\n", in.Ball, in.Msg)
		in.Msg = "Pong"
		log.Printf("%d Sending: %s\n", in.Ball, in.Msg)
		err = stream.Send(in)
		if err != nil {
			return err
		}
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {

	flag.StringVar(&keyfile, "key-file", "certs/server-key.pem", "Location of key-file")
	flag.StringVar(&certfile, "cert-file", "certs/server-cert.pem", "Location of cert-file")
	flag.StringVar(&cafile, "ca-file", "certs/ca-cert.pem", "Location of cafile")
	flag.Parse()

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	pb.RegisterPongServiceServer(grpcServer, &PongServer{})

	l, err := net.Listen("tcp", ":12021")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:12021")
	grpcServer.Serve(l)
}
