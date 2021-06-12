package server

import (
	"io"

	"github.com/pursuit/mortalkin/internal/game"
	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
	"github.com/pursuit/mortalkin/internal/proto/out/api/portal"

	"github.com/pursuit/portal/pkg"

	"github.com/dgrijalva/jwt-go"
)

type GameServer struct {
	mortalkin_proto.UnimplementedGameServer

	UserClient portal_proto.UserClient
}

func (this GameServer) Play(stream mortalkin_proto.Game_PlayServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		var claims pkg.Jwt
		_, err = jwt.ParseWithClaims(in.GetToken(), &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("zxcwqe"), nil
		})
		if err != nil {
			return err
		}

		return game.Connect(int(in.GetCharacterId()), claims.ID, stream)
	}
}
