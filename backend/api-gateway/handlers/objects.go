package handlers

import (
	user "messenger/user-service/server/user"
	"messenger/user-service/server/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ValidateClient    *utils.TAuthClient
	UserServiceClient user.UserProfileServiceClient
	connections       []*grpc.ClientConn
)

func InitClients(authConf utils.TAuthConfig, userServiceURL string) error {
	conn, err := grpc.NewClient(
		userServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	connections = append(connections, conn)

	UserServiceClient = user.NewUserProfileServiceClient(conn)

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
