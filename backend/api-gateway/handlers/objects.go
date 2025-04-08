package handlers

import (
	user "messenger/user-service/user"
	"messenger/user-service/utils"
	post "messenger/wall-service/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ValidateClient    *utils.TAuthClient
	UserServiceClient user.UserProfileServiceClient
	WallServiceClient post.WallServiceClient
	connections       []*grpc.ClientConn
)

func InitClients(authConf utils.TAuthConfig, userServiceURL, wallServiceURL string) error {
	conn, err := grpc.NewClient(
		userServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	connections = append(connections, conn)

	UserServiceClient = user.NewUserProfileServiceClient(conn)

	conn, err = grpc.NewClient(
		wallServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	connections = append(connections, conn)

	WallServiceClient = post.NewWallServiceClient(conn)

	ValidateClient, err = utils.NewValidateClient(authConf.JwtPublicStr)
	if err != nil {
		return err
	}
	return nil
}

func CloseClients() {
	for _, conn := range connections {
		conn.Close()
	}
}
