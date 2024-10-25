package chat

import (
	"context"
	"log"
)

type Server struct{}

func (s Server) mustEmbedUnimplementedChatServiceServer() {}

func (s Server) SayHello(c context.Context, m *Message) (*Message, error) {

	log.Printf("Received message body from client: %s", m.Body)

	return &Message{Body: "Helo from server!"}, nil
}
