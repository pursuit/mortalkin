package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pursuit/mortalkin/internal/game"
	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
	"github.com/pursuit/mortalkin/internal/proto/out/api/portal"
	"github.com/pursuit/mortalkin/internal/proto/server"

	"google.golang.org/grpc"
)

func main() {
	defer log.Println("Shutdown the server success")

	go game.StartServer()
	defer game.Shutdown()

	portalConn, err := grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer portalConn.Close()
	userClient := portal_proto.NewUserClient(portalConn)

	lis, err := net.Listen("tcp", ":5004")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	mortalkin_proto.RegisterUserServer(grpcServer, server.UserServer{
		UserClient: userClient,
	})
	mortalkin_proto.RegisterGameServer(grpcServer, server.GameServer{
		UserClient: userClient,
	})

	go func() {
		log.Println("listen to 5004")
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

	log.Println("Server is ready")
	<-sigint

	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	gracefulChan := make(chan bool)

	go func() {
		grpcServer.GracefulStop()
		gracefulChan <- true
	}()

	select {
	case <-gracefulChan:
		break
	case <-ctx.Done():
		log.Println("Forcing shut down")
		grpcServer.Stop()
	}
}
