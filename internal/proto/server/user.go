package server

import (
	"context"

	"github.com/pursuit/mortalkin/internal/game"
	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
	"github.com/pursuit/mortalkin/internal/proto/out/api/portal"

	"github.com/pursuit/portal/pkg"

	"github.com/dgrijalva/jwt-go"
)

type UserServer struct {
	mortalkin_proto.UnimplementedUserServer

	UserClient portal_proto.UserClient
}

func (this UserServer) Login(ctx context.Context, in *mortalkin_proto.LoginPayload) (*mortalkin_proto.LoginResponse, error) {
	userResp, err := this.UserClient.Login(ctx, &portal_proto.LoginPayload{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	var claims pkg.Jwt
	_, err = jwt.ParseWithClaims(userResp.GetToken(), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("zxcwqe"), nil
	})
	if err != nil {
		return nil, err
	}

	resp := mortalkin_proto.LoginResponse{
		Token: userResp.GetToken(),
	}

	gameCharacters := game.GetCharacters(claims.ID)
	resp.Characters = make([]*mortalkin_proto.Character, len(gameCharacters), len(gameCharacters))
	for i, char := range gameCharacters {
		resp.Characters[i] = &mortalkin_proto.Character{
			Name: char.GetName(),
		}
	}

	return &resp, nil
}

func (this UserServer) CreateCharacter(ctx context.Context, in *mortalkin_proto.CreateCharacterPayload) (*mortalkin_proto.Character, error) {
	var claims pkg.Jwt
	_, err := jwt.ParseWithClaims(in.GetToken(), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("zxcwqe"), nil
	})
	if err != nil {
		return nil, err
	}

	character, err := game.CreateCharacter(claims.ID, in.GetName())
	if err != nil {
		return nil, err
	}

	return &mortalkin_proto.Character{
		Id:   int64(character.GetID()),
		Name: character.GetName(),
		Position: &mortalkin_proto.Position{
			X: int32(character.GetPosition().GetX()),
			Y: int32(character.GetPosition().GetX()),
		},
	}, nil
}
