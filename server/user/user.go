package user

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{
		db: db,
	}
}

func (s *server) GetUserId(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	id := in.GetId()

	var r *pb.User
	err := s.db.QueryRow("select * from users where id = ?", id).Scan(r)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{User: r}, nil
}

func (s *server) GetAllUsers(ctx context.Context, in *emptypb.Empty) (*pb.UsersInfo, error) {
	users := []*pb.User{}

	rows, err := s.db.Query("select * from users")
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UsersInfo{}, err
	}
	defer rows.Close()

	var r *pb.User
	for rows.Next() {
		err = rows.Scan(r)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.UsersInfo{}, err
		}

		users = append(users, r)
	}

	return &pb.UsersInfo{Users: users}, nil
}
