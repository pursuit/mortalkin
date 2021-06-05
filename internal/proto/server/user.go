package server

import (
	"context"

	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
	"github.com/pursuit/mortalkin/internal/proto/out/api/portal"
)

type UserServer struct {
	mortalkin_proto.UnimplementedUserServer

	UserClient portal_proto.UserClient
}

func (this UserServer) Login(ctx context.Context, in *mortalkin_proto.LoginPayload) (*mortalkin_proto.LoginResponse, error) {
	resp, err := this.UserClient.Login(ctx, &portal_proto.LoginPayload{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &mortalkin_proto.LoginResponse{
		Token: resp.GetToken(),
	}, nil
}
