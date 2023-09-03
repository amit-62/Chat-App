package main

import (
	"io"
	"log"

	pb "github.com/amit/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			log.Fatalf("error while streaming : %v", err)
		}
		log.Printf("Request received with %v", req.Name)
		messages = append(messages, "Hello", req.Name)

	}
}
