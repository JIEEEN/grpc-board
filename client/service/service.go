package service

import (
	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/grpc"
)

var (
	Auth  pb.AuthServiceClient
	User  pb.UserServiceClient
	Board pb.BoardServiceClient
	Reply pb.ReplyServiceClient
)

func InitService(conn *grpc.ClientConn) {
	Auth = pb.NewAuthServiceClient(conn)
	User = pb.NewUserServiceClient(conn)
	Board = pb.NewBoardServiceClient(conn)
	Reply = pb.NewReplyServiceClient(conn)
}
