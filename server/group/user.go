package group

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	pb.UnimplementedGroupUserServiceServer
	db *sql.DB
}

func NewUserServer(db *sql.DB) *userServer {
	return &userServer{
		db: db,
	}
}

func (s *userServer) GetAllGroupUsers(ctx context.Context, in *emptypb.Empty) (*pb.GroupUsersInfo, error) {
	groupUsers := []*pb.GroupUser{}

	rows, err := s.db.Query("select id, group_id, user_id from group_users")
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

		groupUser := &pb.GroupUser{
			Id:      r_id,
			GroupId: r_groupId,
			UserId:  r_userId,
		}

		groupUsers = append(groupUsers, groupUser)
	}

	return &pb.GroupUsersInfo{GroupUsers: groupUsers}, nil
}

func (s *userServer) GetGroupUserId(ctx context.Context, in *pb.GetGroupUserRequest) (*pb.GetGroupUserResponse, error) {
	id := in.GetId()

	var r_id, r_groupId, r_userId string
	err := s.db.QueryRow("select * from group_users where id = ?", id).Scan(&r_id, &r_groupId, &r_userId)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.GetGroupUserResponse{}, err
	}

	groupUser := &pb.GroupUser{
		Id:      r_id,
		GroupId: r_groupId,
		UserId:  r_userId,
	}

	return &pb.GetGroupUserResponse{GroupUser: groupUser}, nil
}

func (s *userServer) CreateGroupUser(ctx context.Context, in *pb.CreateGroupUserRequest) (*pb.CreateGroupUserResponse, error) {
	groupUser := in.GetGroupUserRequest()

	_, err := s.db.Exec("insert into group_users (group_id, user_id) values (?, ?)",
		groupUser.GetGroupId(), groupUser.GetUserId())
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.CreateGroupUserResponse{
			Result: false,
		}, err
	}

	return &pb.CreateGroupUserResponse{
		Result: true,
	}, nil
}

func (s *userServer) UpdateGroupUser(ctx context.Context, in *pb.UpdateGroupUserRequest) (*pb.UpdateGroupUserResponse, error) {
	groupUser := in.GetGroupUserRequest()

	_, err := s.db.Exec("update ignore boards set group_id = ?, user_id = ?", groupUser.GetGroupId(), groupUser.GetUserId())
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UpdateGroupUserResponse{
			Result: false,
		}, err
	}

	return &pb.UpdateGroupUserResponse{
		Result: true,
	}, nil
}

func (s *userServer) DeleteGroupUser(ctx context.Context, in *pb.DeleteGroupUserRequest) (*pb.DeleteGroupUserResponse, error) {
	id := in.GetId()

	_, err := s.db.Exec("delete from group_users where id = ?", id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.DeleteGroupUserResponse{
			Result: false,
		}, err
	}

	return &pb.DeleteGroupUserResponse{
		Result: true,
	}, nil
}
