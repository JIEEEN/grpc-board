package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
)

type server struct {
	pb.UnimplementedAuthServiceServer
	db *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{
		db: db,
	}
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	email := in.GetEmail()
	password := in.GetPassword()

	var success bool
	var r_email, r_password string
	err := s.db.QueryRow("select email from users where email = ?", email).Scan(&r_email)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		success = false
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, err
	}
	if r_email == "" {
		err = errors.New("no Email data on db")
		log.Printf("No Email data on db: %v", err)
		success = false
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Printf("Error occured while hash password: %v", err)
		success = false
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, err
	}

	err = s.db.QueryRow("select password from where email = ?", email).Scan(&r_password)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		success = false
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, err
	}

	if ComparePassword(hashedPassword, r_password) {
		success = true
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, nil
	} else {
		log.Printf("Password Mismatch: %v", err)
		success = false
		return &pb.LoginResponse{
			Result:      success,
			AccessToken: "",
		}, err
	}
}
