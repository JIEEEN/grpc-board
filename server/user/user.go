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

	var r_id, r_email, r_nickname string
	err := s.db.QueryRow("select id, email, nickname from users where id = ?", id).Scan(&r_id, &r_email, &r_nickname)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	user := pb.User{
		Id:       r_id,
		Email:    r_email,
		Nickname: r_nickname,
	}

	return &pb.UserResponse{User: &user}, nil
}

func (s *server) GetAllUser(ctx context.Context, in *emptypb.Empty) (*pb.UsersInfo, error) {
	users := []*pb.User{}

	rows, err := s.db.Query("select id, email, nickname from users")
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	var r_id, r_email, r_nickname string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_email, &r_nickname)
		if err != nil {
			log.Fatalf("Failed to get row data: %v", err)
		}

		user := &pb.User{
			Id:       r_id,
			Email:    r_email,
			Nickname: r_nickname,
		}

		users = append(users, user)
	}

	return &pb.UsersInfo{Users: users}, nil
}
