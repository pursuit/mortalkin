// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mortalkin_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (Game_PlayClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Play(ctx context.Context, opts ...grpc.CallOption) (Game_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[0], "/pursuit.api.mortalkin.Game/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &gamePlayClient{stream}
	return x, nil
}

type Game_PlayClient interface {
	Send(*PlayGamePayload) error
	Recv() (*GameNotif, error)
	grpc.ClientStream
}

type gamePlayClient struct {
	grpc.ClientStream
}

func (x *gamePlayClient) Send(m *PlayGamePayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gamePlayClient) Recv() (*GameNotif, error) {
	m := new(GameNotif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	Play(Game_PlayServer) error
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) Play(Game_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Play(&gamePlayServer{stream})
}

type Game_PlayServer interface {
	Send(*GameNotif) error
	Recv() (*PlayGamePayload, error)
	grpc.ServerStream
}

type gamePlayServer struct {
	grpc.ServerStream
}

func (x *gamePlayServer) Send(m *GameNotif) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gamePlayServer) Recv() (*PlayGamePayload, error) {
	m := new(PlayGamePayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pursuit.api.mortalkin.Game",
	HandlerType: (*GameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _Game_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mortalkin/game.proto",
}