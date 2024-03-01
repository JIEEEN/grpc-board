package group

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	pb.UnimplementedBoardGroupServiceServer
	db *sql.DB
}

func NewBoardServer(db *sql.DB) *userServer {
	return &userServer{
		db: db,
	}
}

func (s *userServer) GetAllGroupUsers(ctx context.Context, in *emptypb.Empty) (*pb.GroupUsersInfo, error) {
	var groupUsers []*pb.GroupUser

	rows, err := s.db.Query("select * from group_users")
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.GroupUsersInfo{}, err
	}
	defer rows.Close()

	var r_id, r_groupId, r_userId string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_groupId, &r_userId)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.GroupUsersInfo{}, err
		}
	}
}

func (s *userServer) GetGroupUserId(ctx context.Context, in *pb.GetGroupUserRequest) (*pb.GetGroupUserResponse, error) {

}

func (s *userServer) CreateGroupUser(ctx context.Context, in *pb.CreateBoardGroupRequest) (*pb.CreateBoardGroupResponse, error) {

}

func (s *userServer) UpdateGroupUser(ctx context.Context, in *pb.UpdateGroupUserRequest) (*pb.UpdateGroupUserResponse, error) {

}

func (s *userServer) DeleteGroupUser(ctx context.Context, in *pb.DeleteGroupUserRequest) (*pb.DeleteBoardGroupResponse, error) {

}
