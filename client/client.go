package main

import (
	"context"
	"learning-grpc/chat"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	//WithTransportCredentials and insecure.NewCredentials() instead
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("couldb't connect: ", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := &chat.Message{Body: "Hello From the client"}

	resp, err := c.SayHello(context.Background(), message)

	if err != nil {
		log.Fatal("Error when calling say hello: ", err)
	}

	log.Println("Response from server: ", resp.Body)
}
