package main

import (
	"context"
	"io"
	"log"
	"time"

	"grpc-go-demo/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// Unary
	doUnary(c)

	// Server Streaming
	doServerStreaming(c)

	// Client Streaming
	doClientStreaming(c)

	// Bidirectional Streaming
	doBidirectionalStreaming(c)
}

func doUnary(c proto.GreeterClient) {
	resp, _ := c.SayHello(context.Background(), &proto.HelloRequest{Name: "Juan"})
	log.Printf("[Unary] Respuesta: %s", resp.Message)
}

func doServerStreaming(c proto.GreeterClient) {
	stream, _ := c.GreetManyTimes(context.Background(), &proto.HelloRequest{Name: "Carlos"})
	log.Println("[Server Streaming] Respuestas:")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("  ->", msg.Message)
	}
}

func doClientStreaming(c proto.GreeterClient) {
	stream, _ := c.LongGreet(context.Background())
	names := []string{"Ana", "Luis", "Maria"}
	for _, name := range names {
		log.Println("Enviando:", name)
		stream.Send(&proto.HelloRequest{Name: name})
		time.Sleep(500 * time.Millisecond)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("[Client Streaming] Respuesta: %s", resp.Message)
}

func doBidirectionalStreaming(c proto.GreeterClient) {
	stream, _ := c.GreetEveryone(context.Background())
	done := make(chan struct{})

	// Envío (cliente -> servidor)
	go func() {
		for _, name := range []string{"Pepe", "Laura", "Andrés"} {
			log.Println("Enviando:", name)
			stream.Send(&proto.HelloRequest{Name: name})
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// Recepción (servidor -> cliente)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			log.Println("  <-", resp.Message)
		}
		close(done)
	}()

	<-done
	log.Println("[Bidirectional] Completado")
}
