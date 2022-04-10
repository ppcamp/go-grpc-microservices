package chat

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)

type ChatServer struct {
	UnimplementedChatServiceServer
}

func NewServer() ChatServiceServer {
	return &ChatServer{}
}

func (s *ChatServer) SayHello(ctx context.Context, msg *Message) (*Message, error) {
	if peer, ok := peer.FromContext(ctx); ok {
		logrus.Infof("Client %v", peer.Addr)
	}
	logrus.Info("Received: %v", msg.Body)
	return &Message{Body: "Hello from server"}, nil
}
