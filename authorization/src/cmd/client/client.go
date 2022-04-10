package main

import (
	"context"
	"streamer/configs"
	"streamer/controllers/chat"
	"streamer/utils"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn := utils.Must(grpc.Dial(*configs.APP_PORT, options))
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	ctx := context.Background()
	in := &chat.Message{Body: "teste"}

	msg := utils.Must(client.SayHello(ctx, in))

	logrus.Info(msg.GetBody())

}
