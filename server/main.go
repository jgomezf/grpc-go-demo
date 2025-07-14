package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"grpc-go-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Unary: %v", req.Name)
	return &proto.HelloReply{Message: "Hola " + req.Name}, nil
}

func (s *server) GreetManyTimes(req *proto.HelloRequest, stream proto.Greeter_GreetManyTimesServer) error {
	log.Printf("Server Streaming: %v", req.Name)
	for i := 0; i < 5; i++ {
		resp := &proto.HelloReply{Message: fmt.Sprintf("Hola %s - #%d", req.Name, i)}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) LongGreet(stream proto.Greeter_LongGreetServer) error {
	log.Println("Client Streaming")
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloReply{Message: "Saludos a: " + fmt.Sprint(names)})
		}
		if err != nil {
			return err
		}
		log.Printf("Recibido: %s", req.Name)
		names = append(names, req.Name)
	}
}

func (s *server) GreetEveryone(stream proto.Greeter_GreetEveryoneServer) error {
	log.Println("Bidirectional Streaming")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Cliente dice: %s", req.Name)
		resp := &proto.HelloReply{Message: "Hola " + req.Name + "!"}
		stream.Send(resp)
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Println("Servidor gRPC escuchando en :50051")
	s.Serve(lis)
}
