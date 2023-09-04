package main

import (
	"io"
	"log"

	pb "github.com/amit/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while streaming : %v", err)
		}
		log.Printf("Request received with %v", req.Name)
		
		res:= &pb.HelloResponse{
			Message : "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil{
			return err
		}

	}
}
